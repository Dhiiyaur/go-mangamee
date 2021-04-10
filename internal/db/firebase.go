package db

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

type Config struct {
	FirebaseApp       *firebase.App
	FirebaseDB        *db.Client
	FbRootRef         *db.Ref
	FbUsersRef        *db.Ref
	FbUsersHistoryRef *db.Ref
}

var config Config

func InitDb() {
	_, firebaseDb := initFirebaseApp()
	fbRootRef := firebaseDb.NewRef("")
	fbUsersRef := fbRootRef.Child("usersData")
	fbUsersHistoryRef := fbRootRef.Child("usersDataHistory")

	config = Config{
		FbRootRef:         fbRootRef,
		FbUsersRef:        fbUsersRef,
		FbUsersHistoryRef: fbUsersHistoryRef,
	}

}

func ConDB() Config {
	return config
}

func initFirebaseApp() (*firebase.App, *db.Client) {

	conf := &firebase.Config{
		DatabaseURL: "https://mangamee-c7e7f-default-rtdb.firebaseio.com",
	}
	opt := option.WithCredentialsFile(os.Getenv("KEY"))

	firebaseApp, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatalln("Error initializing firebase app:", err)
	}

	firebaseDB, err := firebaseApp.Database(context.Background())
	if err != nil {
		log.Fatalln("Error initializing firebase db:", err)
	}

	return firebaseApp, firebaseDB
}
