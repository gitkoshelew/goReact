import bookingOrderDay from '../mockData/BookingMockData'

export type IsRentType = { id: string; firstRoom: boolean; secondRoom: boolean }

export const BookingPageAPI = {
  getCalendarData(): Promise<IsRentType[]> {
    return Promise.resolve(bookingOrderDay.bookingRoomPageMockData.calendarData).then((res) => res)
  },
}
