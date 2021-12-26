import { isDev } from './env/env'
import bookingOrderDay from './mockData/BookingMockData'

export const BookingPageAPI = {
  getRoomList() {
    if (isDev) return Promise.resolve(bookingOrderDay)
    else return Promise.reject('error')
  },
}
