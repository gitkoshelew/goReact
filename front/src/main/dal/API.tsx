import bookingOrderDay from './mockData/BookingMockData';
import {isDev} from './env/env';


export const BookingPageAPI = {
    getRoomList() {
        return isDev ? Promise.resolve(bookingOrderDay) : Promise.reject('error')

    },
}
