import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { UrlsList } from './app.urls';
import * as FileSaver from 'file-saver';
import * as XLSX from 'xlsx';


const EXCEL_TYPE = 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;charset=UTF-8';
const EXCEL_EXTENSION = '.xlsx';

@Injectable({
  providedIn: 'root'
})
export class AppService {
  static TOKEN:string;
  urls:UrlsList = new UrlsList();
  headers:HttpHeaders;
  
  constructor(private http: HttpClient, private route: Router) { 
    
  }

  setHeaders() {
    if(AppService.TOKEN) {
      this.headers = new HttpHeaders({
        'authorization': AppService.TOKEN,
        'content-type': 'application/json',
        'Access-Control-Allow-Origin': '*',
        "Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
      });
    } else {
      this.route.navigateByUrl('/login');
    }
  }

  getHardCodedValues(): Observable<any> {
    return this.http.get('assets/hardcoded.json');
  }  

  getBookingInformations(): Observable<any> {
    this.setHeaders();
    return this.http.get(this.urls.GET_BOOKING_INFO, {headers: this.headers});
  }

  getSerchedData(searchedData):Observable<any> {
    this.setHeaders();
    return this.http.post(this.urls.SEARCH_RIDES_INFO,searchedData,{headers: this.headers});
  }

  getPassengerDetails():Observable<any>{
    this.setHeaders();
    return this.http.get(this.urls.PASSENGERS_INFO,{headers: this.headers});
  }

  getVehicleDetails():Observable<any> {
    this.setHeaders();
    return this.http.get(this.urls.VEHICLE_INFO, {headers: this.headers})
  }

  createNewBooking(data):Observable<any> {
    this.setHeaders();
    return this.http.post(this.urls.CREATE_BOOKING, data, {headers: this.headers});
  }

  loginToSystem(credentials):Observable<any> {
    return this.http.post(this.urls.LOGIN, credentials);
  }

  exportAsExcelFile(json: any[], excelFileName: string): void {
    const worksheet: XLSX.WorkSheet = XLSX.utils.json_to_sheet(json);
    const workbook: XLSX.WorkBook = { Sheets: { 'data': worksheet }, SheetNames: ['data'] };
    const excelBuffer: any = XLSX.write(workbook, { bookType: 'xlsx', type: 'array' });
    this.saveAsExcelFile(excelBuffer, excelFileName);
  }

  private saveAsExcelFile(buffer: any, fileName: string): void {
    const data: Blob = new Blob([buffer], {
      type: EXCEL_TYPE
    });
    FileSaver.saveAs(data, fileName + '_export_' + new Date().getTime() + EXCEL_EXTENSION);
  }
}

