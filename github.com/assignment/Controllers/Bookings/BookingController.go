package Bookings

import (
	"encoding/json"
	"fmt"
	"github.com/assignment/Connections"
	"github.com/assignment/Services"
	"io/ioutil"
	"strconv"
	_ "strconv"
	"net/http"
)

/*CReate Model*/
type BookingModel struct {
	Idbookings         int    `json:"idbookings"`
	Passenger_id       int    `json:"passenger_id"`
	Driver_vehicles_id int    `json:"driver_vehicles_id"`
	Datebooking        string `json:"datebooking"`
	Timebooking        string `json:"timebooking"`
	Start_logitude     string `json:"start_logitude"`
	Start_latitude     string `json:"start_latitude"`
	End_logitude       string `json:"end_logitude"`
	End_latitud        string `json:"end_latitud"`
	booking_type       int    `json:"booking_type"`
}

type SearchBookingModel struct {
	Idbookings         int    `json:"idbookings"`
	Start_logitude     string `json:"start_logitude"`
	Start_latitude     string `json:"start_latitude"`
	End_logitude       string `json:"end_logitude"`
	End_latitud        string `json:"end_latitud"`
	DriverName         string `json:"driver_name"`
	CustomerName       string `json:"customer_name"`
	BillAmount         int    `json:"bill_amount"`
	Status             int 	  `json:"status"`
	TotalDistance      int `json:"total_distance"`
	CreatedDate		   string `json: "booked_date"`
	Booking_Type       int `json:"booking_type"`
	Driver_Phone       string `json:"driver_phone"`
	Customer_Phone     string `json:"customer_phone"`
}

type SearchModel struct {
	Booking_Type       int `json:"booking_type"`
	Status             int 	  `json:"status"`
	Driver_Phone       string `json:"driver_phone"`
	DriverName         string `json:"driver_name"`
	CustomerName       string `json:"customer_name"`
	Customer_Phone     string `json:"customer_phone"`
	CreatedDate		   string `json: "booked_date"`
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {

	Services.HandleCors(&w, r)

	dbcon := Connections.CreateConnection()
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		} else {
			var result BookingModel
			json.Unmarshal(body, &result)
			//var bookingId=result.Idbookings
			/*Assign to variables*/
			var passengerId = result.Passenger_id
			var driverid = result.Driver_vehicles_id
			var bookingDate = result.Datebooking
			var bookingTime = result.Timebooking
			var startlogi = result.Start_logitude
			var startlati = result.Start_latitude
			var endlati = result.End_latitud
			var endlogi = result.End_logitude
			var bookingType = result.booking_type
			/*Function to update*/
			fmt.Println(string(body))
			insForm, err := dbcon.Prepare("INSERT INTO bookings ( passenger_id, driver_vehicles_id, datebooking, timebooking, start_logitude, start_latitude, end_logitude, end_latitude, booking_type ) VALUES (?,?,?,?,?,?,?,?,?)")
			insForm.Exec(passengerId, driverid, bookingDate, bookingTime, startlogi, startlati, endlogi, endlati, bookingType)
			if err != nil {
				json.NewEncoder(w).Encode(err.Error())
				panic(err.Error())

			} else {
				fmt.Println(insForm)
				json.NewEncoder(w).Encode("created booking")
			}

		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	dbcon.Close();
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {

	Services.HandleCors(&w, r)

	dbcon := Connections.CreateConnection()
	if r.Method == "PUT" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		} else {

			var result BookingModel
			json.Unmarshal(body, &result)
			fmt.Println(string(body))

			/*Assign to variables*/
			var bookingId = result.Idbookings
			var passengerId = result.Passenger_id
			var driverid = result.Driver_vehicles_id
			var bookingDate = result.Datebooking
			var bookingTime = result.Timebooking
			var startlogi = result.Start_logitude
			var startlati = result.Start_latitude
			var endlati = result.End_latitud
			var endlogi = result.End_logitude
			var bookingType = result.booking_type
			/*function to update*/
			insForm, err := dbcon.Prepare("UPDATE bookings SET passenger_id = ?, driver_vehicles_id = ?, datebooking= ?, timebooking = ?, start_logitude = ?, start_latitude= ?, end_logitude = ?, end_latitude = ?, booking_type = ? WHERE idbookings = ?")
			res, err := insForm.Exec(passengerId, driverid, bookingDate, bookingTime, startlogi, startlati, endlogi, endlati, bookingId, bookingType)
			if res != nil {
				fmt.Println(res)
			}
			if err != nil {
				json.NewEncoder(w).Encode(err.Error())
				panic(err.Error())
			} else {
				json.NewEncoder(w).Encode("Updated Booking")
			}
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	dbcon.Close();
}

func GetBookingInfo(w http.ResponseWriter, r *http.Request) {
	Services.HandleCors(&w, r)
	dbcon := Connections.CreateConnection()
	serchResult := SearchBookingModel{}
	responseObject :=[] SearchBookingModel{}
	results, err := dbcon.Query("SELECT b.idbookings, b.start_logitude, b.start_latitude, b.end_logitude, b.end_latitude, d.namec,p.namep, f.price, f.status, b.datebooking, b.booking_type, d.contact, p.contact_no FROM bookings b, drivers d, passengers p, fares f, driver_vehicles v WHERE b.idbookings = f.bookings_id AND b.driver_vehicles_id = v.driver_vehicles_id AND b.passenger_id = p.passenger_id AND v.driver_id = d.driver_id AND b.passenger_id = p.passenger_id")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var booking_id,price, status, bookingType  int
		var startLogtitute, startLatitude, endLogtitude, endLatitude, driverName, customerName, bookedDate, contactNo, customer_phone  string
		//var rides GetAllRidesModel
		err = results.Scan(&booking_id, &startLogtitute, &startLatitude, &endLogtitude, &endLatitude, &driverName, &customerName, &price, &status, &bookedDate, &bookingType, &contactNo, &customer_phone)
		if err != nil {
			panic(err.Error())
		} else {
			serchResult.Idbookings = booking_id
			serchResult.Start_latitude = startLatitude
			serchResult.Start_logitude = startLogtitute
			serchResult.End_latitud = endLatitude
			serchResult.End_logitude = endLogtitude
			serchResult.BillAmount = price
			serchResult.CustomerName = customerName
			serchResult.DriverName = driverName
			serchResult.Status = status
			serchResult.TotalDistance = serchResult.BillAmount/100
			serchResult.CreatedDate = bookedDate
			serchResult.Booking_Type = bookingType
			serchResult.Driver_Phone = contactNo
			serchResult.Customer_Phone = customer_phone
			responseObject = append(responseObject, serchResult)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseObject)
	dbcon.Close()
}

func GetSearchedBookingInfo(w http.ResponseWriter, r *http.Request) {
	Services.HandleCors(&w, r)
	body, err := ioutil.ReadAll(r.Body)
	var resultItem SearchModel
	json.Unmarshal(body, &resultItem)
	var booking_type int = resultItem.Booking_Type
	var status int = resultItem.Status
	var driver_phone string = resultItem.Driver_Phone
	var driver_name string = resultItem.DriverName
	var customer_name string = resultItem.CustomerName
	var customer_phone string = resultItem.Customer_Phone
	var created_date string = resultItem.CreatedDate

	dbcon := Connections.CreateConnection()
	serchResult := SearchBookingModel{}
	responseObject :=[] SearchBookingModel{}
	results, err := dbcon.Query("SELECT b.idbookings, b.start_logitude, b.start_latitude, b.end_logitude, b.end_latitude, d.namec,p.namep, f.price, f.status, b.datebooking, b.booking_type, d.contact, p.contact_no FROM bookings b, drivers d, passengers p, fares f, driver_vehicles v WHERE b.idbookings = f.bookings_id AND b.driver_vehicles_id = v.driver_vehicles_id AND b.passenger_id = p.passenger_id AND v.driver_id = d.driver_id AND b.passenger_id = p.passenger_id AND (d.namec='"+driver_name+"' OR d.contact = '"+driver_phone+"' OR b.booking_type='"+strconv.Itoa(booking_type)+"' OR f.status='"+strconv.Itoa(status)+"' OR p.namep = '"+customer_name+"' OR p.contact_no = '"+customer_phone+"' OR b.datebooking = '"+created_date+"')")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var booking_id,price, status, bookingType  int
		var startLogtitute, startLatitude, endLogtitude, endLatitude, driverName, customerName, bookedDate, contactNo, customer_phone  string
		//var rides GetAllRidesModel
		err = results.Scan(&booking_id, &startLogtitute, &startLatitude, &endLogtitude, &endLatitude, &driverName, &customerName, &price, &status, &bookedDate, &bookingType, &contactNo, &customer_phone)
		if err != nil {
			panic(err.Error())
		} else {
			serchResult.Idbookings = booking_id
			serchResult.Start_latitude = startLatitude
			serchResult.Start_logitude = startLogtitute
			serchResult.End_latitud = endLatitude
			serchResult.End_logitude = endLogtitude
			serchResult.BillAmount = price
			serchResult.CustomerName = customerName
			serchResult.DriverName = driverName
			serchResult.Status = status
			serchResult.TotalDistance = serchResult.BillAmount/100
			serchResult.CreatedDate = bookedDate
			serchResult.Booking_Type = bookingType
			serchResult.Driver_Phone = contactNo
			serchResult.Customer_Phone = customer_phone
			responseObject = append(responseObject, serchResult)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseObject)
	dbcon.Close()
}