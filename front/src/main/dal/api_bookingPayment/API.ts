import axios from 'axios'

// export const API_BOOKING_URL = process.env.NEST_APP_API_PAYMENT_LINK
export const API_BOOKING_URL = 'http://localhost:5003/'

export const apiBookingPayment = axios.create({
  withCredentials: true,
  baseURL: API_BOOKING_URL,
})
