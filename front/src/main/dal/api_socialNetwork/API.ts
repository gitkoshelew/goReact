import axios from 'axios'

export const API_BOOKING_URL = process.env.REACT_APP_API_SOCIAL_AUTH_LINK

export const apiAuthSocialNetwork = axios.create({
  withCredentials: true,
  baseURL: API_BOOKING_URL,
})
