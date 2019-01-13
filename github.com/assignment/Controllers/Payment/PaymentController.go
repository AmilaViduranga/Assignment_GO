package Payment

import (
	"encoding/json"
	"github.com/assignment/Connections"
	"github.com/assignment/Services"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type FareDetails struct {
	Fare_id     int    `json:"fare_id"`
	Bookings_id string `json:"bookings_id"`
	Price       string `json:"price"`
	Status      string `json:"status"`
}

func FareDetailsbyBookingid(w http.ResponseWriter, r *http.Request) {
	/*Cors Handling*/
	Services.HandleCors(&w, r)
	//Open Db Connction
	dbcon := Connections.CreateConnection()

	params := mux.Vars(r)
	var Booking_Id = params["booking_id"]

	FareObj := FareDetails{}
	responseObject := []FareDetails{}

	// Create db connection
	results, err := dbcon.Query("SELECT fare_id, bookings_id, price, status FROM fares WHERE bookings_id='" + Booking_Id + "'")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var fare_id int
		var price, status, bookings_id string
		//var rides GetAllRidesModel
		err = results.Scan(&fare_id, &bookings_id, &price, &status)
		if err != nil {
			panic(err.Error())
		} else {
			FareObj.Bookings_id = bookings_id
			FareObj.Fare_id = fare_id
			FareObj.Price = price
			FareObj.Status = status
			responseObject = append(responseObject, FareObj)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseObject)
	dbcon.Close()
}

func UpdateFairDetails(w http.ResponseWriter, r *http.Request) {
	Services.HandleCors(&w, r)
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		} else {
			var result FareDetails
			json.Unmarshal(body, &result)
			var Status = result.Status
			var Price = result.Price
			var Bookings_id = result.Bookings_id

			//Open Db Connction
			dbcon := Connections.CreateConnection()
			//update function
			insForm, err := dbcon.Prepare("UPDATE fares SET price =?, `status` = ? WHERE `bookings_id` =?")
			insForm.Exec(Price, Status, Bookings_id)
			if err != nil {
				panic(err.Error())
			}
			//load updated row
			FareObj := FareDetails{}
			responseObject := []FareDetails{}

			results, err := dbcon.Query("SELECT fare_id, bookings_id, price, status FROM fares WHERE bookings_id='" + Bookings_id + "'")
			if err != nil {
				panic(err.Error())
			}
			for results.Next() {
				var fare_id int
				var price, status, bookings_id string
				//var rides GetAllRidesModel
				err = results.Scan(&fare_id, &bookings_id, &price, &status)
				if err != nil {
					panic(err.Error())
				} else {
					FareObj.Bookings_id = bookings_id
					FareObj.Fare_id = fare_id
					FareObj.Price = price
					FareObj.Status = status
					responseObject = append(responseObject, FareObj)
				}
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseObject)
			dbcon.Close()
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
