import bookingOrderDay from '../mockData/BookingMockData'
import axios from 'axios';

export type IsRentType = { id: string; firstRoom: boolean; secondRoom: boolean }
export type UserType = { email: string, password: string }
export type LogInResponseType = {
    userId: number,
    'email': string,
    'role': string,
    'verified': true,

    'name': string
    'sName': string
    'mName': string
    'sex': string,
    'birthDate': string,
    'address': string,
    'phone': string,
    'photo': string
}


const API_URL = 'http://localhost:8080/';

const settings = {
    withCredentials: true,
    baseURL: API_URL,
}

export const $api = axios.create({
    ...settings,
})


export const AuthAPI = {
    async logIn(user: UserType):Promise<LogInResponseType> {
        const res = await $api.post('api/login', user)
        return res.data
    }


}

export const BookingPageAPI = {
    getCalendarData(): Promise<IsRentType[]> {
        return Promise.resolve(bookingOrderDay.bookingRoomPageMockData.calendarData).then((res) => res)
    },
}
