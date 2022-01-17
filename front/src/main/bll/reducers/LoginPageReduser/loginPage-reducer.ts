import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { LogInResponse } from '../../../dal/api_client/AuthService'

const initialState: InitialStateLoginPage = {
  loadingStatus: 'idle',
  errorMsg: '',
  user: null,
}

const loginPageSlice = createSlice({
  name: 'loginPage',
  initialState,
  reducers: {
    reqLoginLogoutStart(state) {
      state.loadingStatus = 'loading'
      state.errorMsg = ''
    },
    reqLoginSuccess(state, action: PayloadAction<{ user: LogInResponse }>) {
      state.loadingStatus = 'success'
      state.user = action.payload.user
    },
    reqLoginError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = 'error'
      state.errorMsg = action.payload.errorMsg
    },
    reqLogOutSuccess(state) {
      state.loadingStatus = 'success'
      state.user = null
    },
    reqLogOutError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = 'error'
      state.errorMsg = action.payload.errorMsg
    },
    reqMeSuccess(state, action: PayloadAction<{ user: LogInResponse }>) {
      state.loadingStatus = 'success'
      state.user = action.payload.user
    },
    reqMeError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = 'idle'
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

export type LoginPageLoadingStatus = 'idle' | 'loading' | 'success' | 'error'
export type User = LogInResponse | null
