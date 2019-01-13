package Drivers

import (
	"encoding/json"
	"github.com/assignment/Connections"
	"github.com/assignment/Services"
	"net/http"
)

/*Create driver Model*/
type DriverModel struct {
	Driver_id  int    `json:"Driver_id"`
	Name       string `json:"Name"`
	Contact_no string `json:"Contact_no"`
	Email      string `json:"Email"`
	Password   string `json:"Password"`
}

type DriverVehichleModel struct {
	Driver_Vehicle_id	int	`json:"driver_vehicle_id"`
	No_Plate			string `json:"no_plate"`
}

func GetAllDrivers(w http.ResponseWriter, r *http.Request) {
	/*Cors Handling*/
	Services.HandleCors(&w, r)
	//Open Db Connction
	dbcon := Connections.CreateConnection()

	DriverObj := DriverModel{}
	responseObject := []DriverModel{}
	results, err := dbcon.Query("SELECT driver_id, namec,contact, email, password FROM drivers")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var driver_id int
		var driver_name, driver_contact_no, driver_email, driver_password string
		//var rides GetAllRidesModel
		err = results.Scan(&driver_id, &driver_name, &driver_contact_no, &driver_email, &driver_password)
		if err != nil {
			panic(err.Error())
		} else {
			DriverObj.Contact_no = driver_contact_no
			DriverObj.Email = driver_email
			DriverObj.Driver_id = driver_id
			DriverObj.Name = driver_name
			DriverObj.Password = driver_password
			responseObject = append(responseObject, DriverObj)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseObject)
	dbcon.Close()
}

func GetVehicleDetails(w http.ResponseWriter, r *http.Request) {
	/*Cors Handling*/
	Services.HandleCors(&w, r)
	//Open Db Connction
	dbcon := Connections.CreateConnection()

	DriverVehicleObj := DriverVehichleModel{}
	responseObject := []DriverVehichleModel{}
	results, err := dbcon.Query("SELECT d.driver_vehicles_id, d.no_plate from driver_vehicles d")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var driver_vehicle_id int
		var no_plate string
		//var rides GetAllRidesModel
		err = results.Scan(&driver_vehicle_id, &no_plate)
		if err != nil {
			panic(err.Error())
		} else {
			DriverVehicleObj.Driver_Vehicle_id = driver_vehicle_id
			DriverVehicleObj.No_Plate = no_plate
			responseObject = append(responseObject, DriverVehicleObj)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseObject)
	dbcon.Close()
}


