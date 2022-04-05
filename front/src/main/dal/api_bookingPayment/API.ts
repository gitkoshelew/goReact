import axios from 'axios'

export const API_BOOKING_URL = process.env.REACT_APP_API_PAYMENT_LINK

export const apiBookingPayment = axios.create({
  withCredentials: true,
  baseURL: API_BOOKING_URL,
})
