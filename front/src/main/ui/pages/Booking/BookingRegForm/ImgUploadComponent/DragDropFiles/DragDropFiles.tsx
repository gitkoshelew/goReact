import React from 'react'
import { FileUploader } from 'react-drag-drop-files'
import { useAppDispatch } from '../../../../../../bll/store/store'
import { BookingUploadPetImg } from '../../../../../../bll/reducers/BookingRegFormReducer/BookindRegForm-saga'
import s from './DragDropFiles.module.scss'

const { customTablet, dragDropFilesContainer } = s
const fileTypes = ['JPG', 'PNG', 'GIF', 'JPEG']

export const DragDropFiles = () => {
  const dispatch = useAppDispatch()

  const handleChange = (file: File) => {
    dispatch(BookingUploadPetImg(file))
  }

  return (
    <div className={dragDropFilesContainer}>
      <FileUploader maxSize={3} classes={customTablet} handleChange={handleChange} name="file" types={fileTypes} />
    </div>
  )
}
