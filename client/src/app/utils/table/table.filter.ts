import * as _ from "lodash";
import {Pipe, PipeTransform, Injectable} from "@angular/core";

@Pipe({
  name: "filter",
  pure: false
})

@Injectable()
export class FilterPipe implements PipeTransform {

  transform(array: any[], query: string): any {
    if (query) {
      return array.filter(function(row) {
        let word = query.toLowerCase();
        let driver_name = row.driver_name.toLowerCase();
        console.log(driver_name);
        if(driver_name.indexOf(word) > -1) {
          return row;
        }
      })
    }
    return array;
  }
}