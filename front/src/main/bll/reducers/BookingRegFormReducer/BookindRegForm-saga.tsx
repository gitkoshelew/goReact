import { changePhotoUrl, changeProgressStatus } from './BookingRegForm-reducer'
import { put, call } from 'redux-saga/effects'

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
  yield put(changeProgressStatus({ newStatus: 'uploaded' }))
}

export const BookingUploadPetImg = (file: Blob | MediaSource) => ({
  type: 'BOOKING_REG_FORM/BOOKING_PET_IMG_UPLOAD',
  file,
})
type BookingUploadPetImgType = ReturnType<typeof BookingUploadPetImg>
