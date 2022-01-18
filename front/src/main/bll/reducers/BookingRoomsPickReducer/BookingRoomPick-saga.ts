import { reqCalendarDataFailed, reqCalendarDataStarted, reqCalendarDataSucceeded } from './BookingRoomPick-reducer'
import { call, put } from 'redux-saga/effects'
import { BookingPageAPI, IsRent } from '../../../dal/api_client/API'

export function* BookingRoomPickSagaWorker() {
  try {
    yield put(reqCalendarDataStarted())
    const newIsRent: IsRent[] = yield call(BookingPageAPI.getCalendarData)
    yield put(reqCalendarDataSucceeded({ newCalendarData: newIsRent }))
  } catch (err) {
    yield put(reqCalendarDataFailed())
  }
}

export const BookingRoomPickSaga = () => ({
  type: 'BOOKING_ROOM_PICK/NEW_IS_RENT_ROOMS_FOR_CALENDAR',
})
