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
	"github.com/samuelralmeida/neofarma/internal/user"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	ctx := context.Background()
	firestoreClient, err := firestore.NewFirestoreClient(ctx)
	if err != nil {
		log.Fatalln("error to start firestore client:", err)
	}
	defer firestoreClient.Close()

	firestoreRepository := firestore.NewFirestoreRepository(firestoreClient)

	patientUseCases := patient.NewPatientUseCases(firestoreRepository)
	userUseCases := user.NewUserUseCases(firestoreRepository)

	webHandler := handlers.NewWebHandler(patientUseCases, userUseCases)

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

	log.Println("listening...", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)
}
