import bookingOrderDay from '../mockData/BookingMockData'
import axios from 'axios'

export type IsRent = { id: string; firstRoom: boolean; secondRoom: boolean }

export const API_URL = process.env.REACT_APP_API_LINK

export const $api = axios.create({
  withCredentials: true,
  baseURL: API_URL,
})

$api.interceptors.request.use((config) => {
  config.headers = config.headers ?? {}
  config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`
  return config
})

$api.interceptors.response.use(
  (config) => {
    return config
  },
  async (error) => {
    if (localStorage.getItem('token') || localStorage.getItem('MockToken')) {
      let refreshStopper = 0
      const originalRequest = error.config
      if (error.response.status === 401 && refreshStopper === 0) {
        const response = await $api.post('/api/refresh')
        localStorage.setItem('token', response.headers['Access-Token'])
        refreshStopper = refreshStopper + 1
        return $api.request(originalRequest)
      }
    }
  }
)

export const BookingPageAPI = {
  getCalendarData(): Promise<IsRent[]> {
    return Promise.resolve(bookingOrderDay.bookingRoomPageMockData.calendarData).then((res) => res)
  },
}
