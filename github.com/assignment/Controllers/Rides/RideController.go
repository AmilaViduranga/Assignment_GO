package Rides

import (
	"encoding/json"
	"fmt"
	"github.com/assignment/Connections"
	"github.com/assignment/Services"
	"io/ioutil"
	"net/http"
)

type GetAllRidesModel struct {
	Idbookings     int    `json:"idbookings"`
	Start_logitude string `json:"start_logitude"`
	Start_latitude string `json:"start_latitude"`
	End_logitude   string `json:"end_logitude"`
	End_latitude   string `json:"end_latitude"`
	Driver_id      int    `json:"driver_id"`
	Drivername     string `json:"drivername"`
	Passenger_id   int    `json:"passenger_id"`
	Passengername  string `json:"passengername"`
	Contact_no     string `json:"contact_no"`
	Date           string `json:"date"`
	Time           string `json:"time"`
}

type SearchModel struct {
	Passenger_name    string `json:"passengername"`
	Passenger_contact string `json:"contact_no"`
	Driver_name       string `json:"drivername"`
	Datebooking       string `json:"date"`
}

func GetAllRides(w http.ResponseWriter, r *http.Request) {
	/*Cors Handling*/
	Services.HandleCors(&w, r)
	//Open Db Connction
	dbcon := Connections.CreateConnection();
	ridesObj := GetAllRidesModel{}
	responseObject := []GetAllRidesModel{}
	// Invoke the sql
	results, err := dbcon.Query("SELECT idbookings, start_logitude, start_latitude, end_logitude, end_latitude, driver_id, namec, passenger_id, namep, contact_no,datebooking,timebooking FROM getallrides")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var idbookings, driver_id, passenger_id int
		var start_logitude, drivername, passengername, contact_no, start_latitude, end_logitude, end_latitude, date, time string
		//var rides GetAllRidesModel
		err = results.Scan(&idbookings, &start_logitude, &start_latitude, &end_logitude, &end_latitude, &driver_id, &drivername, &passenger_id, &passengername, &contact_no, &date, &time)
		if err != nil {
			panic(err.Error())
		} else {
			ridesObj.Idbookings = idbookings
			ridesObj.Start_logitude = start_logitude
			ridesObj.Start_latitude = start_latitude
			ridesObj.End_logitude = end_logitude
			ridesObj.End_latitude = end_latitude
			ridesObj.Driver_id = driver_id
			ridesObj.Drivername = drivername
			ridesObj.Passenger_id = passenger_id
			ridesObj.Passengername = passengername
			ridesObj.Contact_no = contact_no
			ridesObj.Date = date
			ridesObj.Time = time
			responseObject = append(responseObject, ridesObj)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseObject)
	dbcon.Close()

}

func Dashboardsearch(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	var result SearchModel
	json.Unmarshal(body, &result)
	fmt.Println(string(body))
	/*Cors Handling*/
	Services.HandleCors(&w, r)
	//Open Db Connction
	dbcon := Connections.CreateConnection();
	var passenger_name = result.Passenger_name
	var passenger_contact = result.Passenger_contact
	var driver_name = result.Driver_name
	var datebooking = result.Datebooking

	ridesObj := GetAllRidesModel{}
	responseObject := []GetAllRidesModel{}

	fmt.Println("SELECT idbookings, start_logitude, start_latitude, end_logitude, end_latitude, driver_id, namec, passenger_id, namep, contact_no, datebooking, timebooking FROM getallrides WHERE namep='" + passenger_name + "' OR contact_no='" + passenger_contact + "' OR namec='" + driver_name + "' OR datebooking='" + datebooking + "'")

	// Invoke the sql
	results, err := dbcon.Query("SELECT idbookings, start_logitude, start_latitude, end_logitude, end_latitude, driver_id, namec, passenger_id, namep, contact_no, datebooking, timebooking FROM getallrides WHERE namep='" + passenger_name + "' OR contact_no='" + passenger_contact + "' OR namec='" + driver_name + "' OR datebooking='" + datebooking + "'")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var idbookings, driver_id, passenger_id int
		var start_logitude, drivername, passengername, contact_no, start_latitude, end_logitude, end_latitude, date, time string
		//var rides GetAllRidesModel
		err = results.Scan(&idbookings, &start_logitude, &start_latitude, &end_logitude, &end_latitude, &driver_id, &drivername, &passenger_id, &passengername, &contact_no, &date, &time)
		if err != nil {
			panic(err.Error())
		} else {
			ridesObj.Idbookings = idbookings
			ridesObj.Start_logitude = start_logitude
			ridesObj.Start_latitude = start_latitude
			ridesObj.End_logitude = end_logitude
			ridesObj.End_latitude = end_latitude
			ridesObj.Driver_id = driver_id
			ridesObj.Drivername = drivername
			ridesObj.Passenger_id = passenger_id
			ridesObj.Passengername = passengername
			ridesObj.Contact_no = contact_no
			ridesObj.Date = date
			ridesObj.Time = time
			responseObject = append(responseObject, ridesObj)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(responseObject)
	json.NewEncoder(w).Encode(responseObject)

	dbcon.Close()
}
