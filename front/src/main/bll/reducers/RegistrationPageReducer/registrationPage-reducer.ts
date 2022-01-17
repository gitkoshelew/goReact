import { createSlice, PayloadAction } from '@reduxjs/toolkit'

const initialState: InitialRegisterState = {
  loadingStatus: 'idle',
  errorMsg: '',
}

const registerPageSlice = createSlice({
  name: 'registerPage',
  initialState,
  reducers: {
    reqRegisterStart(state) {
      state.loadingStatus = 'loading'
      state.errorMsg = ''
    },
    reqRegisterSuccess(state) {
      state.loadingStatus = 'success'
    },
    reqRegisterError(state, action: PayloadAction<{ errorMsg: string }>) {
      state.loadingStatus = 'error'
      state.errorMsg = action.payload.errorMsg
    },
    RegisterEnd(state) {
      state.loadingStatus = 'idle'
    },
  },
})

export const RegisterPageReducer = registerPageSlice.reducer

export const { reqRegisterStart, reqRegisterSuccess, reqRegisterError, RegisterEnd } = registerPageSlice.actions

type InitialRegisterState = {
  loadingStatus: RegisterPageLoadingStatus
  errorMsg: string
}

export type RegisterPageLoadingStatus = 'idle' | 'loading' | 'success' | 'error'
