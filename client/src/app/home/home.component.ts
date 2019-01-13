import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { AppService } from '../app.service';
import { SelectComponent } from '../utils/select/select.component';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  statics:any;
  bookingData:any[] = [];
  dataSendToTable:any[] = [];
  passengerList:any[] = [];
  vehicleList:any[] = [];
  recordsOnTable:number = 10;
  totalPages:number = 0;
  bookingTypes:number = 0;
  searchByStatus:number = 0;
  searchedRider:string = "";
  searchedByRiderPhone:string = "";
  searchByCustomerName:string = "";
  searchByCustomerPhone:string = "";
  searchByDate:string = "";
  isMessageStateEnable:boolean = false;
  responseType:string = "";
  respondMessage:string = "";

  startLatitude:number = 0;
  startLongtitude:number = 0;
  endLatitude:number = 0;
  endLongtitude:number = 0;

  @ViewChild("customerList")  customerListElement: SelectComponent;
  @ViewChild("vehicleListElement") vehicleListLoader: SelectComponent;
  
  //filtering keys
  isTypeFilterDirty:boolean = false;
  isStatusDirty:boolean = false;
  isRiderFieldDirty:boolean = false;
  isRiderPhoneDirty:boolean = false;
  isCustomerNameDirty:boolean = false;
  isCustomerPhoneDirty:boolean = false;
  isSearchedByDateDirty:boolean =  false;

  newBookingInstance:BookingModel = new BookingModel();
  constructor(private service: AppService) { }

  ngOnInit() {
    this.getAllBookings();
    this.getVehicleInfo();
    this.getPassengerInfo();
    this.loadHardCodedValues();
  }

  getDataFromComponents(value) {
    if(value.key == "recordsOnTable") {
      this.recordsOnTable = value.value;
    }
    if(value.key == "bookingTypes") {
      this.bookingTypes = value.value;
      this.isTypeFilterDirty = true;
      this.filtering();
    }
    if(value.key == "searchByStatus") {
      this.searchByStatus = value.value;
      this.isStatusDirty = true;
      this.filtering();
    }
    if(value.key == "riderName") {
      this.searchedRider = value.value;
      this.isRiderFieldDirty = true;
      this.filtering();
      if(this.searchedRider.length == 0) {
        this.dataSendToTable = this.bookingData;
      }
    }

    if(value.key == "riderPhone") {
      this.searchedByRiderPhone = value.value;
      this.isRiderPhoneDirty = true;
      this.filtering();
      if(this.searchedByRiderPhone.length == 0) {
        this.dataSendToTable = this.bookingData;
      }
    }

    if(value.key == "customerName") {
      this.searchByCustomerName = value.value;
      this.isCustomerNameDirty = true;
      this.filtering();
      if(this.searchByCustomerName.length == 0) {
        this.dataSendToTable = this.bookingData;
      }
    }

    if(value.key == "customerPhone") {
      this.searchByCustomerPhone = value.value;
      this.isCustomerPhoneDirty = true;
      this.filtering();
      if(this.searchByCustomerPhone.length == 0) {
        this.dataSendToTable = this.bookingData;
      }
    }

    if(value.key == "searchFromDate") {
      this.searchByDate = value.value;
      this.isSearchedByDateDirty = true;
      this.filtering();
      if(this.searchByDate.length == 0) {
        this.dataSendToTable = this.bookingData;
      }
    }

    if(value.key == "create-startLocation") {
      this.startLatitude = value.value.lat;
      this.startLongtitude = value.value.lng;
      this.newBookingInstance.start_latitude = this.startLatitude.toString();
      this.newBookingInstance.start_logitude = this.startLongtitude.toString();
    }

    if(value.key == "create-endLocation") {
      this.endLatitude = value.value.lat;
      this.endLongtitude = value.value.lng;
      this.newBookingInstance.end_latitud = this.endLatitude.toString();
      this.newBookingInstance.end_logitude = this.endLongtitude.toString();
    }

    if(value.key == "create-bookingDate") {
      this.newBookingInstance.datebooking = value.value;
    }

    if(value.key == "create-bookingType") {
      this.newBookingInstance.booking_type = parseInt(value.value);
    }

    if(value.key == "create-passenger") {
      this.newBookingInstance.passenger_id = parseInt(value.value);
    }

    if(value.key == "create-vehicle") {
      this.newBookingInstance.driver_vehicles_id = parseInt(value.value);
    }

    this.totalPages = Math.round(this.dataSendToTable.length / this.recordsOnTable);
  }

  ReSetAllData() {
    this.dataSendToTable = this.bookingData;
    this.getDataFromComponents({key:"All",value:"All"});
  }

  loadHardCodedValues() {
    this.service.getHardCodedValues().subscribe(data => {
      this.statics = data;
    })
  }

  getAllBookings() {
    this.service.getBookingInformations().subscribe(data => {
      this.bookingData = data;
      this.dataSendToTable = this.bookingData;
      this.totalPages = Math.round(this.bookingData.length / 10);
    })
  }

  filtering() {
    if(this.isTypeFilterDirty == true) {
      this.dataSendToTable = [];
      this.bookingData.forEach(item => {
        if(item.booking_type == this.bookingTypes) {
          this.dataSendToTable.push(item);
        }
      })
      this.isTypeFilterDirty = false;
    }

    if(this.isStatusDirty == true) {
      this.dataSendToTable = [];
      this.bookingData.forEach(item => {
        if(item.status == this.searchByStatus) {
          this.dataSendToTable.push(item);
        }
      })
      this.isStatusDirty = false;
    }

    if(this.isRiderFieldDirty == true) {
      this.dataSendToTable = [];
      this.bookingData.forEach(item => {
        if(item.driver_name == this.searchedRider) {
          this.dataSendToTable.push(item);
        }
      })
      this.isRiderFieldDirty = false;
    } 

    if(this.isRiderPhoneDirty == true) {
      this.dataSendToTable = [];
      this.bookingData.forEach(item => {
        if(item.driver_phone == this.searchedByRiderPhone) {
          this.dataSendToTable.push(item);
        }
      })
      this.isRiderPhoneDirty = false;
    }

    if(this.isCustomerNameDirty == true) {
      this.dataSendToTable = [];
      this.bookingData.forEach(item => {
        if(item.customer_name == this.searchByCustomerName) {
          this.dataSendToTable.push(item);
        }
      })
      this.isCustomerNameDirty = false;
    }

    if(this.isCustomerPhoneDirty == true) {
      this.dataSendToTable = [];
      this.bookingData.forEach(item => {
        if(item.customer_phone == this.searchByCustomerPhone) {
          this.dataSendToTable.push(item);
        }
      })
      this.isCustomerPhoneDirty = false;
    }

    if(this.isSearchedByDateDirty == true) {
      this.dataSendToTable = [];
      this.bookingData.forEach(item => {
        if(item.CreatedDate == this.searchByDate) {
          this.dataSendToTable.push(item);
        }
      })
      this.isSearchedByDateDirty = false;
    }
  }

  searchInfo() {
    this.dataSendToTable = [];
    let searchData = {
      "booking_type" : this.bookingTypes,
      "status" : this.searchByStatus,
      "driver_phone" : this.searchedByRiderPhone,
      "driver_name" : this.searchedRider,
      "customer_name" : this.searchByCustomerName,
      "customer_phone" : this.searchByCustomerPhone,
      "booked_date" : this.searchByDate
    }

    this.service.getSerchedData(searchData).subscribe(result => {
      this.dataSendToTable = result;
    })
  }

  exportAsXLSX() {
    this.service.exportAsExcelFile(this.dataSendToTable, 'sample');
  }

  getPassengerInfo() {
    this.service.getPassengerDetails().subscribe(data => {
      data.forEach(passenger => {
        this.passengerList.push({
          "key": passenger.Passenger_id,
          "value": passenger.Email
        })
      })
      this.customerListElement.prepareItems();
    })
  }

  getVehicleInfo() {
    this.service.getVehicleDetails().subscribe(data => {
      data.forEach(vehicle => {
        this.vehicleList.push({
          "key": vehicle.driver_vehicle_id,
          "value": vehicle.no_plate
        })
      })
      this.vehicleListLoader.prepareItems();
    })
  }

  createNewBooking() {
    this.service.createNewBooking(this.newBookingInstance).subscribe(data => {
      this.isMessageStateEnable = true;
      this.responseType = "success";
      this.respondMessage = "Successfylly submitted your booking";
      setTimeout(() => {
        this.isMessageStateEnable = false;
      }, 5000);
    }, err => {
      this.isMessageStateEnable = true;
      this.responseType = "danger";
      this.respondMessage = "Booking is failed, please try again";
      setTimeout(() => {
        this.isMessageStateEnable = false;
      }, 5000);
    })
  }
  
}

export class BookingModel {
  passenger_id:number;
  driver_vehicles_id: number;
  datebooking:any;
  timebooking:any;
  start_logitude:string;
  start_latitude:string;
  end_logitude:string;
  end_latitud:string;
  booking_type:number;
}