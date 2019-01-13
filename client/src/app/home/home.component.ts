import { Component, OnInit, AfterContentInit } from '@angular/core';
import { AppService } from '../app.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit, AfterContentInit {

  statics:any;
  bookingData:any[] = [];
  dataSendToTable:any[] = [];
  passengerList:any[] = [];
  recordsOnTable:number = 10;
  totalPages:number = 0;
  bookingTypes:number = 0;
  searchByStatus:number = 0;
  searchedRider:string = "";
  searchedByRiderPhone:string = "";
  searchByCustomerName:string = "";
  searchByCustomerPhone:string = "";
  searchByDate:string = "";
  
  //filtering keys
  isTypeFilterDirty:boolean = false;
  isStatusDirty:boolean = false;
  isRiderFieldDirty:boolean = false;
  isRiderPhoneDirty:boolean = false;
  isCustomerNameDirty:boolean = false;
  isCustomerPhoneDirty:boolean = false;
  isSearchedByDateDirty:boolean =  false;
  constructor(private service: AppService) { }

  ngOnInit() {
    this.getAllBookings();
  }

  ngAfterContentInit() {
    this.loadHardCodedValues();
    this.getPassengerDetails();
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

  getPassengerDetails() {
    this.passengerList = [];
    this.service.getPassengerDetails().subscribe(data => {
      data.forEach(passenger => {
        this.passengerList.push({
          "key": passenger.Passenger_id,
          "value": passenger.Email
        })
      })
    })
    console.log(this.passengerList);
  }
}