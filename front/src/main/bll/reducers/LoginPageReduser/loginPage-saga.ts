import { call, put } from 'redux-saga/effects'
import {
  reqLoginError,
  reqLoginLogoutStart,
  reqLoginSuccess,
  reqLogOutError,
  reqLogOutSuccess,
} from './loginPage-reducer'
import { AxiosResponse } from 'axios'
import { AuthAPI, LogInResponseType, UserRequestDataType } from '../../../dal/api_client/AuthService'

export function* LoginPageLoginSagaWorker(action: LoginRequestType) {
  try {
    yield put(reqLoginLogoutStart())
    const { data, headers }: AxiosResponse<LogInResponseType> = yield call(AuthAPI.logIn, action.user)
    yield call(storeToken, headers['access-token'])
    yield put(reqLoginSuccess({ user: data }))
  } catch (err) {
    if (err instanceof Error) {
      yield put(reqLoginError({ errorMsg: err.message }))
    }
  }
}

export const LoginRequest = (user: UserRequestDataType) => ({
  type: 'LOGIN_PAGE/LOGIN_SAGA',
  user,
})

type LoginRequestType = ReturnType<typeof LoginRequest>

export function* LoginPageLogoutSagaWorker() {
  try {
    yield put(reqLoginLogoutStart())
    yield call(AuthAPI.logOut)
    localStorage.removeItem('token')
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

async function storeToken(token: string) {
  try {
    await localStorage.setItem('token', token)
  } catch (error) {
    console.log('Localstorage error during token store:', error)
  }
}
