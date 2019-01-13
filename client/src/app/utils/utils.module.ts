import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AgmCoreModule } from '@agm/core';
import { FormsModule }   from '@angular/forms';
import { DataTableModule } from "angular-6-datatable";
import { SelectComponent } from './select/select.component';
import { InputTextComponent } from './input-text/input-text.component';
import { TableComponent } from './table/table.component';
import { FilterPipe } from './table/table.filter';
import { MapComponent } from './map/map.component';

@NgModule({
  declarations: [
    SelectComponent,
    InputTextComponent,
    TableComponent,
    FilterPipe,
    MapComponent
  ],
  imports: [
    CommonModule,
    DataTableModule,
    FormsModule,
    AgmCoreModule.forRoot({
      apiKey: 'AIzaSyAuHPL1nfHb_IqBAs2_Al8T681A-narWk4'
    })
  ],
  exports: [
    SelectComponent,
    InputTextComponent,
    TableComponent,
    MapComponent
  ]
})
export class UtilsModule { }
