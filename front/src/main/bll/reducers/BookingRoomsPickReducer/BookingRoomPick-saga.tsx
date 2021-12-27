import { IsRentType, setIsRent, setLoadingStatus } from './BookingRoomPick-reducer'
import { call, put } from 'redux-saga/effects'
import { BookingPageAPI } from '../../../dal/API'

export function* BookingRoomPickSagaWorker() {
  try {
    yield put(setLoadingStatus({ newStatus: 'loading' }))
    const newIsRent: { bookingRoom: IsRentType[] } = yield call(BookingPageAPI.getRoomList)
    yield put(setIsRent({ newIsRent: newIsRent.bookingRoom }))
    yield put(setLoadingStatus({ newStatus: 'success' }))
  } catch (err) {
    yield put(setLoadingStatus({ newStatus: 'error' }))
  }
}

export const BookingRoomPickSaga = () => ({
  type: 'BOOKING_ROOM_PICK/NEW_IS_RENT_ROOMS_FOR_CALENDAR',
})
