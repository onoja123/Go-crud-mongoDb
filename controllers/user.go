package controllers

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/onoja123/mongo-golang/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(sin *mgo.Session) *UserController {
	return &UserController{sin}
}

// Get user controller
func (userController UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//User id
	id := p.ByName("id")

	//checking for user else 404
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	//Get users model
	u := models.User{}

	if err := userController.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	// Send status 200
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

// Create a new new user
func (userController UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	userController.session.DB("mongo-golang").C("users").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	// Send status 200
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "%s\n", uj)
}

// Delete a user by id
func (userController UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//User id
	id := p.ByName("id")

	//checking for user else 404
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := userController.session.DB("mongo-golan").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Send status 200
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
