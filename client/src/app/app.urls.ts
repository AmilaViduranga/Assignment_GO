export class UrlsList {
    BASE_URL:string = "http://localhost:4200/";
    SERVICE_URL:string =  "http://localhost:8001/";
    GET_BOOKING_INFO:string;
    SEARCH_RIDES_INFO:string;
    PASSENGERS_INFO:string;

    constructor() {
        this.GET_BOOKING_INFO = this.SERVICE_URL + "all-bookings";
        this.SEARCH_RIDES_INFO = this.SERVICE_URL + "search-bookings"
        this.PASSENGERS_INFO = this.SERVICE_URL + "passengers";
    }
}