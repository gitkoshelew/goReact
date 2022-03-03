import { ProgressType } from '../../../../../bll/reducers/BookingRegFormReducer/BookingRegForm-reducer'
import { DragDropFiles } from './DragDropFiles'
import s from './ImgUpload.module.scss'

const { uploadTablet } = s

type ImgUploadPropsType = {
  progress: ProgressType
}

export const ImgUpload = ({ progress }: ImgUploadPropsType) => {
  return <div className={uploadTablet}>{progress === 'getUpload' && <DragDropFiles />}</div>
}
