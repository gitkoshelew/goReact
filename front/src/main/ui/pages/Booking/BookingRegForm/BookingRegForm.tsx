/* eslint-disable */
import React, { SyntheticEvent, useState } from 'react'
import s from './BookingRegForm.module.css'
import { ImgUpload } from './ImgUploadComponent/ImgUpload'
import { useSelector } from 'react-redux'
import { AppRootStateType, useAppDispatch } from '../../../../bll/store/store'
import { ProgressType } from '../../../../bll/reducers/BookingRegFormReducer/BookingRegForm-reducer'
import Preloader from '../../../components/preloader/preloader'
import { changeUserParams } from '../../../../bll/reducers/BookingRoomsPickReducer/BookingRoomPick-reducer'

const { bookingForm, clientDescription, inputInfo, uploadedImg } = s

type InputValueSetType = (newValue: string) => void

export const BookingRegForm = () => {
  const dispatch = useAppDispatch()

  const progress = useSelector<AppRootStateType, ProgressType>((state) => state.BookingRegForm.progress)
  const photoUrl = useSelector<AppRootStateType, string | null>((state) => state.BookingRegForm.photoUrl)
  const errorMSG = useSelector<AppRootStateType, string>((state) => state.BookingRegForm.errorMSG)

  const userName = useSelector<AppRootStateType, string>((state) => state.BookingRoomPick.userName)
  const userEmail = useSelector<AppRootStateType, string>((state) => state.BookingRoomPick.userEmail)

  const inputHandler = (e: SyntheticEvent<HTMLInputElement>, type: string) => {
    dispatch(changeUserParams({ newTextParams: e.currentTarget.value, params: type }))
  }

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
          <input value={userName} onChange={(e) => inputHandler(e, 'name')} type="text" />
        </div>
        <div className={inputInfo}>
          Your email:
          <input value={userEmail} onChange={(e) => inputHandler(e, 'email')} type="text" />
        </div>
        <div className={inputInfo}>{actualContent()}</div>
      </div>
    </div>
  )
}
