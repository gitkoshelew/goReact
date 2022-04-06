import { changePhotoUrl, changeProgressStatus, setBookingPhoto } from './BookingRegForm-reducer'
import { put, call } from 'redux-saga/effects'
import { AxiosResponse } from 'axios'
import {
  BookingInfoPhoto,
  BookingRegFormAPI,
  FetchBookingPhotoResponse,
} from '../../../dal/api_bookingPhoto/BookingService'

const dataGivers = async (data: Blob | MediaSource) => {
  const result: Blob | MediaSource = await new Promise<Blob | MediaSource>((res) => {
    setTimeout(() => {
      res(data)
    }, 1000)
  })
  const correctImg = Object.assign(result, {
    preview: URL.createObjectURL(result),
  })

  return correctImg.preview
}

export function* BookingUploadPetImgSagaWorker(action: BookingUploadPetImgType) {
  yield put(changeProgressStatus({ newStatus: 'uploading' }))
  const file: string = yield call(dataGivers, action.file)
  yield put(changePhotoUrl({ newPhotoUrl: file }))
  // const bookingPhotoResponse: AxiosResponse<FetchBookingPhotoResponse> = yield call(
  //   BookingRegFormAPI.createBookingPhoto,
  //   action.newPhotoInfo,
  //   action.file
  // )
  // yield put(
  //   setBookingPhoto({
  //     bookingPhoto: bookingPhotoResponse.data,
  //     bookingInfo: {
  //       type: 'user',
  //       ownerId: 1,
  //     },
  //   })
  // )
  yield put(changeProgressStatus({ newStatus: 'uploaded' }))
}

export const BookingUploadPetImg = (file: File, newPhotoInfo: BookingInfoPhoto) => ({
  type: 'BOOKING_REG_FORM/BOOKING_PET_IMG_UPLOAD',
  file,
  newPhotoInfo,
})
type BookingUploadPetImgType = ReturnType<typeof BookingUploadPetImg>
