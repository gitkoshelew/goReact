import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { PhotoUrl } from '../../../dal/api_File/FileService'

const initialState: InitialStateType = {
  progress: 'getUpload',
  photoUrl: null,
  errorMSG: '',
  checkedOnlinePayment: false,
}

const BookingRegFormSlice = createSlice({
  name: 'BookingRegForm',
  initialState: initialState,
  reducers: {
    changeProgressStatus(state, action: PayloadAction<{ newStatus: ProgressType }>) {
      state.progress = action.payload.newStatus
    },
    changePhotoUrl(state, action: PayloadAction<{ newPhotoUrl: PhotoUrl }>) {
      state.photoUrl = action.payload.newPhotoUrl
    },
    changeErrorMSG(state, action: PayloadAction<{ newErrorMsg: string }>) {
      state.errorMSG = action.payload.newErrorMsg
    },
    changeCheckedOnlinePayment(state, action: PayloadAction<{ checkedOnlinePayment: boolean }>) {
      state.checkedOnlinePayment = action.payload.checkedOnlinePayment
    },
  },
})

export const BookingRegFormReducer = BookingRegFormSlice.reducer
export const { changePhotoUrl, changeProgressStatus, changeErrorMSG, changeCheckedOnlinePayment } =
  BookingRegFormSlice.actions

//types

export type ProgressType = 'getUpload' | 'uploading' | 'uploaded' | 'uploadError'

export type InitialStateType = {
  progress: ProgressType
  photoUrl: PhotoUrl | null
  errorMSG: string
  checkedOnlinePayment: boolean
}
