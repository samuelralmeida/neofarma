package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/samuelralmeida/neofarma/external/repository/firestore"
	"github.com/samuelralmeida/neofarma/external/web/handlers"
	customMiddlewares "github.com/samuelralmeida/neofarma/external/web/middlewares"
	"github.com/samuelralmeida/neofarma/internal/patient"
	responsibilityUC "github.com/samuelralmeida/neofarma/internal/responsibility/usecases"
	"github.com/samuelralmeida/neofarma/internal/user"

	_ "github.com/samuelralmeida/neofarma/external/web/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample Neofarma server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		http://localhost:3000
// @BasePath	/v2
func main() {
	ctx := context.Background()
	firestoreClient, err := firestore.NewFirestoreClient(ctx)
	if err != nil {
		log.Fatalln("error to start firestore client:", err)
	}
	defer firestoreClient.Close()

	firestoreRepository := firestore.NewFirestoreRepository(firestoreClient)

	userUseCases := user.NewUserUseCases(firestoreRepository)
	patientUseCases := patient.NewPatientUseCases(firestoreRepository, userUseCases)
	responsibilityUseCases := responsibilityUC.NewResponsibilityUseCases(firestoreRepository, userUseCases, patientUseCases)

	webHandler := handlers.NewWebHandler(patientUseCases, userUseCases, responsibilityUseCases)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(customMiddlewares.SetUser(userUseCases))
	r.Use(middleware.Timeout(60 * time.Second))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/patients", func(r chi.Router) {
		r.Post("/save", webHandler.SavePatient)
		r.Get("/{id}", webHandler.GetPatientById)
	})

	r.Route("/admin", func(r chi.Router) {
		r.Post("/create", webHandler.SignUp)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/signin", webHandler.SignIn)
		r.Post("/signout", webHandler.SignOut)
	})

	r.Route("/responsibilities", func(r chi.Router) {
		r.Post("/create", webHandler.CreateRelationship)
		r.Post("/remove", webHandler.RemoveRelationship)
	})

	log.Println("listening...", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)
}
