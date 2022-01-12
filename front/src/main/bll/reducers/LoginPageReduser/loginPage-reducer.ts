import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { LogInResponseType } from '../../../dal/api_client/AuthService'

const initialState: InitialStateLoginPageType = {
  loadingStatus: 'onWaiting',
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
    reqLoginSuccess(state, action: PayloadAction<{ user: LogInResponseType }>) {
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
  },
})

export const LoginPageReducer = loginPageSlice.reducer
export const { reqLoginLogoutStart, reqLoginSuccess, reqLoginError, reqLogOutSuccess, reqLogOutError } =
  loginPageSlice.actions

//types

export type InitialStateLoginPageType = {
  loadingStatus: LoginPageLoadingStatusType
  errorMsg: string
  user: UserType
}

export type LoginPageLoadingStatusType = 'onWaiting' | 'loading' | 'success' | 'error'
export type UserType = LogInResponseType | null
