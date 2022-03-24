import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { BookingPaymentFormType } from '../../../dal/api_bookingPayment/BookingService'

enum BookingPaymentLoadingStatuses {
  IDLE = 'IDLE',
  LOADING = 'LOADING',
  SUCCESS = 'SUCCESS',
  ERROR = 'ERROR',
}

const initialState: InitialStateSeats = {
  loadingStatus: BookingPaymentLoadingStatuses.IDLE,
  bookingPayment: [],
  successMsg: '',
  errorMsg: '',
}
const seatsSlice = createSlice({
  name: 'bookingPayment',
  initialState,
  reducers: {
    bookingPaymentStart(state) {
      state.loadingStatus = BookingPaymentLoadingStatuses.LOADING
      state.successMsg = ''
      state.errorMsg = ''
    },
    setBookingPayment(state, action: PayloadAction<{ bookingPayment: BookingPaymentFormType[] }>) {
      state.loadingStatus = BookingPaymentLoadingStatuses.SUCCESS
      state.bookingPayment = action.payload.bookingPayment
    },
    bookingPaymentMessage(state, action: PayloadAction<{ successMsg: string }>) {
      state.loadingStatus = BookingPaymentLoadingStatuses.SUCCESS
      state.successMsg = action.payload.successMsg
    },
    bookingPaymentError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = BookingPaymentLoadingStatuses.ERROR
      state.errorMsg = action.payload.errorMsg
    },
  },
})

export const BookingPaymentFormReducer = seatsSlice.reducer
export const { setBookingPayment, bookingPaymentMessage, bookingPaymentError, bookingPaymentStart } = seatsSlice.actions

//types

type InitialStateSeats = {
  loadingStatus: BookingPaymentLoadingStatus
  bookingPayment: BookingPaymentFormType[]
  successMsg: string
  errorMsg: string
}

export type BookingPaymentLoadingStatus =
  | BookingPaymentLoadingStatuses.IDLE
  | BookingPaymentLoadingStatuses.LOADING
  | BookingPaymentLoadingStatuses.SUCCESS
  | BookingPaymentLoadingStatuses.ERROR
