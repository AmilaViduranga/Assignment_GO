package Passengers

import (
	"encoding/json"
	"github.com/assignment/Connections"
	"github.com/assignment/Services"
	"net/http"
)

/*Create Model Passenger*/
type PassengerModel struct {
	Passenger_id int    `json:"Passenger_id"`
	Name         string `json:"Name"`
	Contact_no   string `json:"Contact_no"`
	Email        string `json:"Email"`
	Password     string `json:"Password"`
}

/*Get All Passengers*/
func GetAllPassengers(w http.ResponseWriter, r *http.Request) {

	Services.HandleCors(&w, r)

	dbcon := Connections.CreateConnection()

	PassengerObj := PassengerModel{}
	responseObject := []PassengerModel{}

	results, err := dbcon.Query("SELECT passenger_id, namep, contact_no, email, password FROM passengers")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var passenger_id int
		var passenger_namep, passenger_contact_no, passenger_email, passenger_password string
		//var rides GetAllRidesModel
		err = results.Scan(&passenger_id, &passenger_namep, &passenger_contact_no, &passenger_email, &passenger_password)
		if err != nil {
			panic(err.Error())
		} else {
			PassengerObj.Contact_no = passenger_contact_no
			PassengerObj.Email = passenger_email
			PassengerObj.Passenger_id = passenger_id
			PassengerObj.Name = passenger_namep
			PassengerObj.Password = passenger_contact_no
			responseObject = append(responseObject, PassengerObj)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseObject)
	dbcon.Close()
}
