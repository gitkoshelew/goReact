import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { LoadingStatuses } from '../types/enum'

const initialState: InitialRegisterState = {
  loadingStatus: LoadingStatuses.IDLE,
  errorMsg: '',
}

const registerPageSlice = createSlice({
  name: 'registerPage',
  initialState,
  reducers: {
    reqRegisterStart(state) {
      state.loadingStatus = LoadingStatuses.LOADING
      state.errorMsg = ''
    },
    reqRegisterSuccess(state) {
      state.loadingStatus = LoadingStatuses.SUCCESS
      state.loadingStatus = LoadingStatuses.IDLE
    },
    reqRegisterError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = LoadingStatuses.ERROR
      state.errorMsg = action.payload.errorMsg
    },
  },
})

export const RegisterPageReducer = registerPageSlice.reducer

export const { reqRegisterStart, reqRegisterSuccess, reqRegisterError } = registerPageSlice.actions

type InitialRegisterState = {
  loadingStatus: RegisterPageLoadingStatus
  errorMsg: string
}

export type RegisterPageLoadingStatus =
  | LoadingStatuses.IDLE
  | LoadingStatuses.LOADING
  | LoadingStatuses.SUCCESS
  | LoadingStatuses.ERROR
