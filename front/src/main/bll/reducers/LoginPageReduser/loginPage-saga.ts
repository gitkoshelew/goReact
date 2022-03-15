import { call, put } from 'redux-saga/effects'
import {
  reqLoginError,
  reqLoginLogoutStart,
  reqLoginSuccess,
  reqLogOutError,
  reqLogOutSuccess,
  reqMeError,
  reqMeSuccess,
} from './loginPage-reducer'
import { AxiosResponse } from 'axios'
import { AuthAPI, LogInResponse, UserRequestData } from '../../../dal/api_client/AuthService'
import { LoginResponse } from '../../../dal/mockData/LoginUserMockData'
import { isDev } from '../../../dal/env/env'

export function* LoginPageLoginSagaWorker(action: LoginRequestType) {
  try {
    yield put(reqLoginLogoutStart())
    const { data, headers }: AxiosResponse<LogInResponse> = yield call(AuthAPI.logIn, action.user)
    if (!isDev) {
      yield call(storeToken, headers['access-token'])
    } else if (isDev) {
      yield call(storeToken, LoginResponse.accessToken)
    }
    yield put(reqLoginSuccess({ user: data }))
  } catch (err) {
    if (err instanceof Error) {
      yield put(reqLoginError({ errorMsg: err.message }))
    }
  }
}

export const LoginRequest = (user: UserRequestData) => ({
  type: 'LOGIN_PAGE/LOGIN_SAGA',
  user,
})

type LoginRequestType = ReturnType<typeof LoginRequest>

export function* LoginPageLogoutSagaWorker() {
  try {
    yield put(reqLoginLogoutStart())
    yield call(AuthAPI.logOut)
    if (!isDev) {
      localStorage.removeItem('token')
    } else {
      localStorage.removeItem('MockToken')
    }
    yield put(reqLogOutSuccess())
  } catch (err) {
    if (err instanceof Error) {
      yield put(reqLogOutError({ errorMsg: err.message }))
    }
  }
}

export const LogOutRequest = () => ({
  type: 'LOGIN_PAGE/LOGOUT_SAGA',
})

export function* LoginPageMeRequestSagaWorker() {
  try {
    yield put(reqLoginLogoutStart())
    const { data }: AxiosResponse<LogInResponse> = yield call(AuthAPI.AuthMe)
    yield put(reqMeSuccess({ user: data }))
  } catch (err) {
    if (err instanceof Error) {
      yield put(reqMeError({ errorMsg: err.message }))
    }
  }
}

export const MeRequest = () => ({
  type: 'LOGIN_PAGE/ME_SAGA',
})

async function storeToken(token: string) {
  if (isDev) {
    try {
      if (token) {
        await localStorage.setItem('MockToken', token)
      }
    } catch (error) {
      console.log('Localstorage error during token store:', error)
    }
  } else {
    try {
      if (token) {
        await localStorage.setItem('token', token)
      }
    } catch (error) {
      console.log('Localstorage error during token store:', error)
    }
  }
}
