import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { BookingPaymentFormType } from '../../../dal/api_bookingPayment/BookingService'

const initialState: InitialStateSeats = {
  bookingPayment: [],
}

const seatsSlice = createSlice({
  name: 'bookingPayment',
  initialState,
  reducers: {
    setBookingPayment(state, action: PayloadAction<{ bookingPayment: BookingPaymentFormType[] }>) {
      state.bookingPayment = action.payload.bookingPayment
    },
  },
})

export const BookingPaymentFormReducer = seatsSlice.reducer
export const { setBookingPayment } = seatsSlice.actions

//types

type InitialStateSeats = {
  bookingPayment: BookingPaymentFormType[]
}
