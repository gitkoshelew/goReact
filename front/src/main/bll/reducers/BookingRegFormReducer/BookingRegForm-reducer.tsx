import { createSlice, PayloadAction } from '@reduxjs/toolkit'

const initialState: InitialStateType = {
  progress: 'getUpload',
  photoUrl: null,
  errorMSG: '',
}

const BookingRegFormSlice = createSlice({
  name: 'BookingRegForm',
  initialState: initialState,
  reducers: {
    changeProgressStatus(state, action: PayloadAction<{ newStatus: ProgressType }>) {
      state.progress = action.payload.newStatus
    },
    changePhotoUrl(state, action: PayloadAction<{ newPhotoUrl: string }>) {
      state.photoUrl = action.payload.newPhotoUrl
    },
    changeErrorMSG(state, action: PayloadAction<{ newErrorMsg: string }>) {
      state.errorMSG = action.payload.newErrorMsg
    },
  },
})

export const BookingRegFormReducer = BookingRegFormSlice.reducer
export const { changePhotoUrl, changeProgressStatus, changeErrorMSG } = BookingRegFormSlice.actions

//types

export type ProgressType = 'getUpload' | 'uploading' | 'uploaded' | 'uploadError'

export type InitialStateType = {
  progress: ProgressType
  photoUrl: null | string
  errorMSG: string
}
