/* eslint-disable */
import React, { useState } from 'react'
import s from './BookingRegForm.module.css'
import { ImgUpload } from './ImgUploadComponent/ImgUpload'
import { Button } from '../../../components/Button/Button'
import { useSelector } from 'react-redux'
import { AppRootStateType } from '../../../../bll/store/store'
import { ProgressType } from '../../../../bll/reducers/BookingRegFormReducer/BookingRegForm-reducer'
import Preloader from '../../../components/preloader/preloader'

const { bookingForm, clientDescription, inputInfo, uploadedImg } = s

export const BookingRegForm = () => {
  const progress = useSelector<AppRootStateType, ProgressType>((state) => state.BookingRegForm.progress)
  const photoUrl = useSelector<AppRootStateType, string | null>((state) => state.BookingRegForm.photoUrl)
  const errorMSG = useSelector<AppRootStateType, string>((state) => state.BookingRegForm.errorMSG)

  const isActiveBtn = progress === 'uploaded'

  const actualContent = () => {
    switch (progress) {
      case 'getUpload':
        return (
          <>
            Upload your pet photo:
            <ImgUpload progress={progress} />
          </>
        )
      case 'uploading':
        return <Preloader />
      case 'uploaded':
        return (
          <div className={uploadedImg}>
            Your pet :
            <img src={photoUrl ? photoUrl : ''} alt="uploadedImg" />
          </div>
        )
      case 'uploadError':
        return (
          <>
            <div>{errorMSG}</div>
            <div>Upload an image please</div>
          </>
        )
    }
  }

  return (
    <div className={bookingForm}>
      <div className={clientDescription}>
        <div className={inputInfo}>
          Your name:
          <input type="text" />
        </div>
        <div className={inputInfo}>
          Your email:
          <input type="text" />
        </div>
        <div className={inputInfo}>{actualContent()}</div>
        <Button type={'Upload'} isActive={isActiveBtn} />
      </div>
    </div>
  )
}
