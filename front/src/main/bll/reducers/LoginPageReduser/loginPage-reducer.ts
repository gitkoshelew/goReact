import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { LogInResponse } from '../../../dal/api_client/AuthService'
import { LoadingStatuses } from '../types/enum'

const initialState: InitialStateLoginPage = {
  loadingStatus: LoadingStatuses.IDLE,
  errorMsg: '',
  user: null,
}

const loginPageSlice = createSlice({
  name: 'loginPage',
  initialState,
  reducers: {
    reqLoginLogoutStart(state) {
      state.loadingStatus = LoadingStatuses.LOADING
      state.errorMsg = ''
    },
    reqLoginSuccess(state, action: PayloadAction<{ user: LogInResponse }>) {
      state.loadingStatus = LoadingStatuses.SUCCESS
      state.user = action.payload.user
    },
    reqLoginError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = LoadingStatuses.ERROR
      state.errorMsg = action.payload.errorMsg
    },
    reqLogOutSuccess(state) {
      state.loadingStatus = LoadingStatuses.SUCCESS
      state.user = null
    },
    reqLogOutError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = LoadingStatuses.ERROR
      state.errorMsg = action.payload.errorMsg
    },
    reqMeSuccess(state, action: PayloadAction<{ user: LogInResponse }>) {
      state.loadingStatus = LoadingStatuses.SUCCESS
      state.user = action.payload.user
    },
    reqMeError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = LoadingStatuses.IDLE
      state.errorMsg = action.payload.errorMsg
    },
  },
})

export const LoginPageReducer = loginPageSlice.reducer
export const {
  reqLoginLogoutStart,
  reqLoginSuccess,
  reqLoginError,
  reqLogOutSuccess,
  reqLogOutError,
  reqMeSuccess,
  reqMeError,
} = loginPageSlice.actions

//types

export type InitialStateLoginPage = {
  loadingStatus: LoginPageLoadingStatus
  errorMsg: string
  user: User
}

export type LoginPageLoadingStatus =
  | LoadingStatuses.IDLE
  | LoadingStatuses.LOADING
  | LoadingStatuses.SUCCESS
  | LoadingStatuses.ERROR
export type User = LogInResponse | null
