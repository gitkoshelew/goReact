import { createSlice, PayloadAction } from '@reduxjs/toolkit'

const initialState: InitialRegisterStateType = {
  loadingStatus: 'waitingForUser',
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
  },
})

export const RegisterPageReducer = registerPageSlice.reducer

export const { reqRegisterStart, reqRegisterSuccess, reqRegisterError } = registerPageSlice.actions

type InitialRegisterStateType = {
  loadingStatus: RegisterPageLoadingStatusType
  errorMsg: string
}

export type RegisterPageLoadingStatusType = 'waitingForUser' | 'loading' | 'success' | 'error'
