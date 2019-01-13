import { Component, OnInit, Input, EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.css']
})
export class MapComponent implements OnInit {
  @Input() title:string = "";
  @Input() key:string;
  @Output() modelValue:EventEmitter<any> = new EventEmitter<any>();
  
  lat: number = 51.678418;
  lng: number = 7.809007;
  constructor() { }

  ngOnInit() {
  }

  getClickedInformation(event) {
    this.lat = event.coords.lat;
    this.lng = event.coords.lng;
    this.sendModelValue();
  }

  sendModelValue() {
    this.modelValue.emit({
      "key": this.key,
      "value": {
        "lat": this.lat,
        "lng": this.lng
      }
    })
  }


}
