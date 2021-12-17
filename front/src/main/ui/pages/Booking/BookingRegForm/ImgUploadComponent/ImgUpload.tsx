/* eslint-disable */
import s from './ImgUpload.module.css'
import { FileUploader } from 'react-drag-drop-files'
import {
  changePhotoUrl,
  changeProgressStatus,
  ProgressType,
} from '../../../../../bll/reducers/BookingRegFormReducer/BookingRegForm-reducer'
import { useAppDispatch } from '../../../../../bll/store/store'
import { BookingUploadPetImg } from '../../../../../bll/reducers/BookingRegFormReducer/BookindRegForm-saga'

const { uploadTablet, customTablet } = s
const fileTypes = ['JPG', 'PNG', 'GIF', 'JPEG']

type ImgUploadPropsType = {
  progress: ProgressType
}

export const ImgUpload = ({ progress }: ImgUploadPropsType) => {
  const dispatch = useAppDispatch()

  const handleChange = (file: any) => {
    dispatch(BookingUploadPetImg(file))
  }

  return (
    <div className={uploadTablet}>
      {progress === 'getUpload' && (
        <FileUploader maxSize={20} classes={customTablet} handleChange={handleChange} name="file" types={fileTypes} />
      )}
    </div>
  )
}
