import { call, put } from 'redux-saga/effects'
import {
  reqLoginError,
  reqLoginLogoutStart,
  reqLoginSuccess,
  reqLogOutError,
  reqLogOutSuccess,
} from './loginPage-reducer'
import { AuthAPI, LogInResponseType, UserRequestDataType } from '../../../dal/api_client/API'
import { AxiosResponse } from 'axios'

export function* LoginPageLoginSagaWorker(action: LoginRequestType) {
  try {
    yield put(reqLoginLogoutStart())
    const { data }: AxiosResponse<LogInResponseType> = yield call(AuthAPI.logIn, action.user)
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
