package main

import (
	"github.com/assignment/Constants"
	"github.com/assignment/Controllers/Bookings"
	"github.com/assignment/Controllers/Drivers"
	"github.com/assignment/Controllers/Passengers"
	"github.com/assignment/Controllers/Payment"
	"github.com/assignment/Controllers/Rides"
	"github.com/assignment/Routes"
	"github.com/assignment/Services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	/*Create The Router*/
	log.Print("Router Created" + Constants.APP_VERSION)
	router := mux.NewRouter();

	/*Set SSL Certification*/
	//log.Fatal(http.ListenAndServeTLS(":8447", "server.crt", "muve-interviews.pem", router))
	log.Print("SSL Secured" + Constants.APP_VERSION)

	/*Start The Application*/
	log.Print("Create The Application " + Constants.APP_VERSION)

	log.Print("Create Routes " + Constants.APP_VERSION)
	router.HandleFunc("/", Routes.HomePage).Methods("GET")

	/*Authenticate Routes¶*/
	router.HandleFunc("/authenticate", Services.CreateTokenEndpoint).Methods("POST")
	router.HandleFunc("/protected", Services.ProtectedEndpoint).Methods("GET")
	router.HandleFunc("/login", Services.ValidateMiddleware(Services.AuthenticateTest)).Methods("GET")

	/*Payment Routes¶*/
	router.HandleFunc("/payment/search/fare/{booking_id}", Services.ValidateMiddleware(Payment.FareDetailsbyBookingid)).Methods("GET")
	router.HandleFunc("/payment/fare/update", Services.ValidateMiddleware(Payment.UpdateFairDetails)).Methods("POST")

	/*Rider Routes¶*/
	router.HandleFunc("/rides", Services.ValidateMiddleware(Rides.GetAllRides)).Methods("GET")
	router.HandleFunc("/rides/search/", Services.ValidateMiddleware(Rides.Dashboardsearch)).Methods("POST")

	/*Passenger Routes¶*/
	router.HandleFunc("/passengers", Services.ValidateMiddleware(Passengers.GetAllPassengers)).Methods("GET")

	/*Drivers Routes¶*/
	router.HandleFunc("/drivers", Services.ValidateMiddleware(Drivers.GetAllDrivers)).Methods("GET")

	/*Bookings Routes¶*/
	router.HandleFunc("/bookings", Services.ValidateMiddleware(Bookings.CreateBooking)).Methods("POST")
	router.HandleFunc("/bookings", Services.ValidateMiddleware(Bookings.UpdateBooking)).Methods("PUT")
	router.HandleFunc("/all-bookings", Services.ValidateMiddleware(Bookings.GetBookingInfo)).Methods("GET")
	router.HandleFunc("/search-bookings", Services.ValidateMiddleware(Bookings.GetSearchedBookingInfo)).Methods("POST")

	/*Start The Applicaiton*/
	log.Print("Start The Application " + Constants.APP_VERSION)
	log.Fatal(http.ListenAndServe(":8001", router))

}
