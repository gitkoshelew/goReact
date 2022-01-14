import { AxiosResponse } from 'axios'
import { $api } from './API'

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

export const AuthAPI = {
  async logIn(user: UserRequestDataType): Promise<AxiosResponse<LogInResponseType>> {
    const res = await $api.post('api/login', user)
    debugger
    return res
  },
  async logOut(): Promise<AxiosResponse> {
    const res = await $api.post('api/logout')
    return res
  },
  async AuthMe(): Promise<AxiosResponse<LogInResponseType>> {
    const res = await $api.post('api/me')
    return res
  },
}
