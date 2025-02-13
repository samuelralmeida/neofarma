install-binaries:
	go install github.com/swaggo/swag/cmd/swag@latest

gcloud-login:
	gcloud auth login

gcloud-set-project:
	gcloud config set project neofarma-project

firestore-start-emulator:
	firebase emulators:start

swag-generate:
	swag init -o ./external/web/docs
