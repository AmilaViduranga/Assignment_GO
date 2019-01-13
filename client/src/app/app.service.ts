import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
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
  urls:UrlsList = new UrlsList();

  headers:HttpHeaders = new HttpHeaders({
    'authorization': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6ImFkbWluIiwidXNlcm5hbWUiOiJhZG1pbiJ9.4PWUk3ApGtV6sexF4lOYeWgozLKKdh0iOudZ0PV4krQ',
    'content-type': 'application/json',
    'Access-Control-Allow-Origin': '*',
    "Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
  });
  
  constructor(private http: HttpClient) { 
    
  }

  getHardCodedValues(): Observable<any> {
    return this.http.get('assets/hardcoded.json');
  }  

  getBookingInformations(): Observable<any> {
    return this.http.get(this.urls.GET_BOOKING_INFO, {headers: this.headers});
  }

  getSerchedData(searchedData):Observable<any> {
    return this.http.post(this.urls.SEARCH_RIDES_INFO,searchedData,{headers: this.headers});
  }

  getPassengerDetails():Observable<any>{
    return this.http.get(this.urls.PASSENGERS_INFO,{headers: this.headers});
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

