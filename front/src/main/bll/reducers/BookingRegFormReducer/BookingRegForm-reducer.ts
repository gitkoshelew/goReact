import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { BookingInfoPhoto, BookingPhotoType } from '../../../dal/api_bookingPhoto/BookingService'

const initialState: InitialStateType = {
  progress: 'getUpload',
  photoUrl: null,
  errorMSG: '',
  checkedOnlinePayment: false,
  bookingPhoto: [],
  bookingInfo: [],
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
    changeCheckedOnlinePayment(state, action: PayloadAction<{ checkedOnlinePayment: boolean }>) {
      state.checkedOnlinePayment = action.payload.checkedOnlinePayment
    },
    setBookingPhoto(state, action: PayloadAction<{ bookingInfo: BookingInfoPhoto; bookingPhoto: BookingPhotoType }>) {
      state.bookingInfo = action.payload.bookingInfo
      state.bookingPhoto = action.payload.bookingPhoto
    },
  },
})

export const BookingRegFormReducer = BookingRegFormSlice.reducer
export const { changePhotoUrl, changeProgressStatus, changeErrorMSG, changeCheckedOnlinePayment, setBookingPhoto } =
  BookingRegFormSlice.actions

//types

export type ProgressType = 'getUpload' | 'uploading' | 'uploaded' | 'uploadError'

export type InitialStateType = {
  progress: ProgressType
  photoUrl: null | string
  errorMSG: string
  checkedOnlinePayment: boolean
  bookingPhoto: any
  bookingInfo: any
}
