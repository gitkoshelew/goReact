import axios from 'axios'

export const API_BOOKING_REG_FORM_URL = process.env.REACT_APP_API_PAYMENT_LINK

export const apiBookingRegForm = axios.create({
  withCredentials: true,
  baseURL: API_BOOKING_REG_FORM_URL,
})
