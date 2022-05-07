import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { BookingPaymentFormType } from '../../../dal/api_bookingPayment/BookingService'
import { LoadingStatuses } from '../types/enum'

const initialState: InitialStateSeats = {
  loadingStatus: LoadingStatuses.IDLE,
  bookingPayment: [],
  successMsg: '',
  errorMsg: '',
}
const seatsSlice = createSlice({
  name: 'bookingPayment',
  initialState,
  reducers: {
    bookingPaymentStart(state) {
      state.loadingStatus = LoadingStatuses.LOADING
      state.successMsg = ''
      state.errorMsg = ''
    },
    setBookingPayment(state, action: PayloadAction<{ bookingPayment: BookingPaymentFormType[] }>) {
      state.loadingStatus = LoadingStatuses.SUCCESS
      state.bookingPayment = action.payload.bookingPayment
    },
    bookingPaymentMessage(state, action: PayloadAction<{ successMsg: string }>) {
      state.loadingStatus = LoadingStatuses.SUCCESS
      state.successMsg = action.payload.successMsg
    },
    bookingPaymentError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = LoadingStatuses.ERROR
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
  | LoadingStatuses.IDLE
  | LoadingStatuses.LOADING
  | LoadingStatuses.SUCCESS
  | LoadingStatuses.ERROR
