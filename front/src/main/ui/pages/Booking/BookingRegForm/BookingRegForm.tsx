import React from 'react'
import s from './BookingRegForm.module.css'
import { ImgUpload } from './ImgUploadComponent/ImgUpload'
import { useSelector } from 'react-redux'
import { AppRootStateType } from '../../../../bll/store/store'
import { ProgressType } from '../../../../bll/reducers/BookingRegFormReducer/BookingRegForm-reducer'
import Preloader from '../../../components/preloader/preloader'
import { FormikValues } from 'formik/dist/types'

const {
  bookingForm,
  clientDescription,
  inputInfo,
  inputInfoInput,
  uploadedImgContent,
  uploadedImg,
  errorMsg,
  errorInputInfo,
  errorInputInfoInput,
} = s

type BookingRegFormType = {
  formik: FormikValues
}

export const BookingRegForm = ({ formik }: BookingRegFormType) => {
  const progress = useSelector<AppRootStateType, ProgressType>((state) => state.BookingRegForm.progress)
  const photoUrl = useSelector<AppRootStateType, string | null>((state) => state.BookingRegForm.photoUrl)
  const errorMSG = useSelector<AppRootStateType, string>((state) => state.BookingRegForm.errorMSG)

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
            <img className={uploadedImgContent} src={photoUrl ? photoUrl : ''} alt="uploadedImg" />
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
        <div className={formik.errors.firstName && formik.touched.firstName ? errorInputInfo : inputInfo}>
          First name:
          <input
            className={formik.errors.firstName && formik.touched.firstName ? errorInputInfoInput : inputInfoInput}
            id={'firstName'}
            onBlur={formik.handleBlur}
            value={formik.values.firstName}
            onChange={formik.handleChange}
            type="text"
          />
        </div>
        {formik.errors.firstName && formik.touched.firstName ? (
          <div className={errorMsg}>{formik.errors.firstName}</div>
        ) : null}

        <div className={formik.errors.lastName && formik.touched.lastName ? errorInputInfo : inputInfo}>
          Last name:
          <input
            className={formik.errors.lastName && formik.touched.lastName ? errorInputInfoInput : inputInfoInput}
            id={'lastName'}
            onBlur={formik.handleBlur}
            value={formik.values.lastName}
            onChange={formik.handleChange}
            type="text"
          />
        </div>
        {formik.errors.lastName && formik.touched.lastName ? (
          <div className={errorMsg}>{formik.errors.lastName}</div>
        ) : null}

        <div className={formik.errors.email && formik.touched.email ? errorInputInfo : inputInfo}>
          Email:
          <input
            className={formik.errors.email && formik.touched.email ? errorInputInfoInput : inputInfoInput}
            id={'email'}
            onBlur={formik.handleBlur}
            value={formik.values.email}
            onChange={formik.handleChange}
            type="text"
          />
        </div>
        {formik.errors.email && formik.touched.email ? <div className={errorMsg}>{formik.errors.email}</div> : null}
        <div className={inputInfo}>{actualContent()}</div>
      </div>
    </div>
  )
}
