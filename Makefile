# setup

install-binaries:
	go install github.com/swaggo/swag/cmd/swag@latest

# gcloud

gcloud-login:
	gcloud auth login

gcloud-set-project:
	gcloud config set project neofarma-project

gcloud-current-project:
	gcloud config get-value project

gcloud-deploy:
	gcloud run deploy neofarma-app --region southamerica-east1 --source .

firestore-start-emulator:
	firebase emulators:start

# swagger

swag-generate:
	swag init -o ./external/web/docs
