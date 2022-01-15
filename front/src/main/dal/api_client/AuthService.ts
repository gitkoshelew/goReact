import { AxiosResponse } from 'axios'
import { $api } from './API'
import { LoginResponse } from '../mockData/LoginUserMockData'
import { userPhotoBoilerPlate } from '../../ui/svgWrapper/navBarSvgWrapper'
import { isDev } from '../env/env'

export type UserRequestData = { email: string; password: string }
export type LogInResponse = {
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
  async logIn(user: UserRequestData): Promise<AxiosResponse<LogInResponse> | { data: LogInResponse }> {
    if (isDev) {
      const res = await Promise.resolve(LoginResponse)
      return mockDataPhotoFieldChecker(res)
    }
    const res = await $api.post('api/login', user)
    return photoFieldChecker(res)
  },
  async logOut(): Promise<AxiosResponse | string> {
    if (isDev) {
      return Promise.resolve('ok')
    }
    const res = await $api.post('api/logout')
    return res
  },
  async AuthMe(): Promise<AxiosResponse<LogInResponse> | { data: LogInResponse }> {
    if (isDev) {
      const res = await Promise.resolve(LoginResponse)
      return mockDataPhotoFieldChecker(res)
    }
    const res = await $api.post('api/me')
    return photoFieldChecker(res)
  },
}

const photoFieldChecker = (response: AxiosResponse<LogInResponse>): AxiosResponse<LogInResponse> => {
  return response.data.photo === 'PhotoURL...'
    ? { ...response, data: { ...response.data, photo: userPhotoBoilerPlate } }
    : response
}

const mockDataPhotoFieldChecker = (response: { data: LogInResponse }): { data: LogInResponse } => {
  return response.data.photo === 'PhotoURL...'
    ? { ...response, data: { ...response.data, photo: userPhotoBoilerPlate } }
    : response
}
