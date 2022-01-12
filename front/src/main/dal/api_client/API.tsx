import bookingOrderDay from '../mockData/BookingMockData'
import axios from 'axios'

export type IsRentType = { id: string; firstRoom: boolean; secondRoom: boolean }

const API_URL = 'http://localhost:8080/'

const settings = {
  withCredentials: true,
  baseURL: API_URL,
}

export const $api = axios.create({
  ...settings,
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
    const originalRequest = error.config
    if (error.response.status === 401) {
      const response = await $api.post('/api/refresh')
      localStorage.setItem('token', response.headers['Access-Token'])
      return $api.request(originalRequest)
    }
  }
)

export const BookingPageAPI = {
  getCalendarData(): Promise<IsRentType[]> {
    return Promise.resolve(bookingOrderDay.bookingRoomPageMockData.calendarData).then((res) => res)
  },
}
