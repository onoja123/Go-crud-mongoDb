package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/onoja123/mongo-golang/controllers"

	"gopkg.in/mgo.v2"
)

func main() {

	// Hanlde Http request
	router := httprouter.New()

	userController := controllers.NewUserController(getSession())

	// Route to get a user
	router.GET("/user:/:id", userController.GetUser)

	// Route to create a user
	router.POST("/user", userController.CreateUser)

	// Route to delete a user
	router.DELETE("/user/:id", userController.DeleteUser)

	//Listen to port
	http.ListenAndServe("localhost:3000", router)
}

func getSession() *mgo.Session {
	url := "put your mongodb url"
	sin, err := mgo.Dial(url)

	if err != nil {
		panic(err)

	}
	return sin

}
