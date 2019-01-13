import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-table',
  templateUrl: './table.component.html',
  styleUrls: ['./table.component.css']
})
export class TableComponent implements OnInit {
  @Input() data:any[] = [];
  @Input() rowsOnPage:number = 10;
  @Input() driverName:string = "";
  @Input() driverPhone:string = "";
  constructor() { }

  ngOnInit() {
  }
}
