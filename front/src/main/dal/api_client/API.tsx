import bookingOrderDay from '../mockData/BookingMockData'
import axios, { AxiosResponse } from 'axios'

export type IsRentType = { id: string; firstRoom: boolean; secondRoom: boolean }
export type UserRequestDataType = { email: string; password: string }
export type LogInResponseType = {
  userId: number
  email: string
  role: string
  verified: boolean
  name: string
  sName: string
  mName: string
  sex: string
  birthDate: string
  address: string
  phone: string
  photo: string
}

const API_URL = 'http://localhost:8080/'

const settings = {
  withCredentials: true,
  baseURL: API_URL,
}

export const $api = axios.create({
  ...settings,
})

export const AuthAPI = {
  async logIn(user: UserRequestDataType): Promise<AxiosResponse<LogInResponseType>> {
    const res = await $api.post('api/login', user)
    return res
  },
  async logOut(): Promise<AxiosResponse> {
    const res = await $api.post('api/logout')
    return res
  },
}

export const BookingPageAPI = {
  getCalendarData(): Promise<IsRentType[]> {
    return Promise.resolve(bookingOrderDay.bookingRoomPageMockData.calendarData).then((res) => res)
  },
}
