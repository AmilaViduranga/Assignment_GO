import { Component, OnInit, Input, EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-input',
  templateUrl: './input-text.component.html',
  styleUrls: ['./input-text.component.css']
})
export class InputTextComponent implements OnInit {

  @Input() placeholder:string = "";
  @Input() label:string = "";
  @Input() type:string = "text";
  @Input() model:string;
  @Input() key:string;
  @Input() isLabelEnable:boolean = false;

  @Output() modelValue:EventEmitter<any> = new EventEmitter<any>();

  constructor() { }

  ngOnInit() {
    if(this.label == "") {
      this.label = this.placeholder;
    }
  }

  sendModelValue() {
    this.modelValue.emit({
      "key": this.key,
      "value": this.model
    })
  }

}
