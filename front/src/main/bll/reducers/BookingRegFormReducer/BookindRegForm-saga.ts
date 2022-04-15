import { changePhotoUrl, changeProgressStatus } from './BookingRegForm-reducer'
import { put, call } from 'redux-saga/effects'
import { AxiosResponse } from 'axios'
import { FilesAPI } from '../../../dal/api_File/FileService'

export function* BookingUploadPetImgSagaWorker(action: BookingUploadPetImgType) {
  yield put(changeProgressStatus({ newStatus: 'uploading' }))
  const bookingPhotoResponse: AxiosResponse = yield call(FilesAPI.uploadFile, action.file)
  yield put(changePhotoUrl({ newPhotoUrl: bookingPhotoResponse.data }))

  yield put(changeProgressStatus({ newStatus: 'uploaded' }))
}

export const BookingUploadPetImg = (file: File) => ({
  type: 'BOOKING_REG_FORM/BOOKING_PET_IMG_UPLOAD',
  file,
})
type BookingUploadPetImgType = ReturnType<typeof BookingUploadPetImg>
