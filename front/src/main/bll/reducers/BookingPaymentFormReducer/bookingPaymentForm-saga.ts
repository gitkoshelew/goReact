import { call, put } from 'redux-saga/effects'
import { AxiosResponse } from 'axios'
import {
  BookingPageAPI,
  BookingPaymentFormType,
  FetchBookingPaymentResponse,
} from '../../../dal/api_bookingPayment/BookingService'
import { setBookingPayment } from './bookingPaymentForm-reducer'

export function* fetchBookingPaymentSagaWorker(action: ReturnType<typeof fetchBookingPaymentRequest>) {
  const bookingPaymentResponse: AxiosResponse<FetchBookingPaymentResponse> = yield call(
    BookingPageAPI.createBookingPayment,
    action.newPaymentCard
  )
  yield put(setBookingPayment({ bookingPayment: bookingPaymentResponse.data }))
}

export const fetchBookingPaymentRequest = (newPaymentCard: BookingPaymentFormType) => ({
  type: 'BOOKING_PAYMENT/FETCH_PAYMENT_SAGA',
  newPaymentCard,
})
