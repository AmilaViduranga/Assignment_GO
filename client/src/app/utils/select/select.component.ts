import { Component, OnInit, Input, EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-select',
  templateUrl: './select.component.html',
  styleUrls: ['./select.component.css']
})
export class SelectComponent implements OnInit {
  @Input() labelValue:string = "";
  @Input() itemList:any[] = [];
  @Input() defaultOption:string = "";
  @Input() model:string;
  @Input() enableLabel:boolean = false;
  @Input() key:string;

  @Output() modelValue:EventEmitter<any> = new EventEmitter<any>();

  items:Array<Item> = new Array<Item>();

  constructor() { }

  ngOnInit() {
    this.prepareItems();
    if(this.model == undefined) {
      this.model = this.defaultOption;
    }
  }

  prepareItems() {
    if(this.itemList.length != 0) {
      console.log(this.itemList);
      this.itemList.forEach(item => {
        this.items.push(new Item(item));
      })
    }
  }

  sendModelValue() {
    this.modelValue.emit({
      "key": this.key,
      "value": this.model
    })
  }

}

export class Item {
  label:string;
  value:string;
  constructor(item) {
    this.label = item.value;
    this.value = item.key;
  }
}
