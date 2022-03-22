import { call, put } from 'redux-saga/effects'
import { AxiosResponse } from 'axios'
import {
  BookingPageAPI,
  BookingPaymentFormType,
  FetchBookingPaymentResponse,
} from '../../../dal/api_bookingPayment/BookingService'
import {
  bookingPaymentError,
  bookingPaymentMessage,
  bookingPaymentStart,
  setBookingPayment,
} from './bookingPaymentForm-reducer'

export function* fetchBookingPaymentSagaWorker(action: ReturnType<typeof fetchBookingPaymentRequest>) {
  try {
    yield put(bookingPaymentStart())
    const bookingPaymentResponse: AxiosResponse<FetchBookingPaymentResponse> = yield call(
      BookingPageAPI.createBookingPayment,
      action.newPaymentCard
    )
    yield put(setBookingPayment({ bookingPayment: bookingPaymentResponse.data }))
    yield put(bookingPaymentMessage({ successMsg: bookingPaymentResponse.statusText }))
  } catch (err) {
    if (err instanceof Error) {
      yield put(bookingPaymentError({ errorMsg: err.name }))
    }
  }
}

export const fetchBookingPaymentRequest = (newPaymentCard: BookingPaymentFormType) => ({
  type: 'BOOKING_PAYMENT/FETCH_PAYMENT_SAGA',
  newPaymentCard,
})
