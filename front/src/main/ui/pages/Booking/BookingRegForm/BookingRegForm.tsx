import React from 'react'
import s from './BookingRegForm.module.scss'
import { ImgUpload } from './ImgUploadComponent/ImgUpload'
import { useSelector } from 'react-redux'
import { AppRootState } from '../../../../bll/store/store'
import Preloader from '../../../components/preloader/preloader'
import { FormikValues } from 'formik/dist/types'
import { Checkbox, FormControl, FormControlLabel, Radio, RadioGroup } from '@mui/material'
import { mastercard, visa } from '../../../svgWrapper/BookingRoomSvgWrapper'

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
  cardCompanyImages,
} = s

type BookingRegFormType = {
  formik: FormikValues
}

export const BookingRegForm = ({ formik }: BookingRegFormType) => {
  const progress = useSelector((state: AppRootState) => state.BookingRegForm.progress)
  const photoUrl = useSelector((state: AppRootState) => state.BookingRegForm.photoUrl)
  const errorMSG = useSelector((state: AppRootState) => state.BookingRegForm.errorMSG)

  const [checked, setChecked] = React.useState(false)

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setChecked(event.target.checked)
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
        <div>
          <FormControlLabel control={<Checkbox checked={checked} onChange={handleChange} />} label="Online payment" />
        </div>

        {checked && (
          <>
            <FormControl>
              <RadioGroup row>
                <FormControlLabel
                  value="mastercard"
                  control={<Radio />}
                  label={<img src={mastercard} alt="mastercard" className={cardCompanyImages} />}
                  labelPlacement="top"
                />
                <FormControlLabel
                  value="visa"
                  control={<Radio />}
                  label={<img src={visa} alt="visa" className={cardCompanyImages} />}
                  labelPlacement="top"
                />
                <FormControlLabel value="top" control={<Radio />} label="Top" labelPlacement="top" />
                <FormControlLabel value="top" control={<Radio />} label="Top" labelPlacement="top" />
              </RadioGroup>
            </FormControl>
            <div className={formik.errors.cardNumber && formik.touched.cardNumber ? errorInputInfo : inputInfo}>
              Card number:
              <input
                className={formik.errors.cardNumber && formik.touched.cardNumber ? errorInputInfoInput : inputInfoInput}
                id={'cardNumber'}
                onBlur={formik.handleBlur}
                value={formik.values.cardNumber}
                onChange={formik.handleChange}
                type="text"
              />
            </div>
            {formik.errors.cardNumber && formik.touched.cardNumber ? (
              <div className={errorMsg}>{formik.errors.cardNumber}</div>
            ) : null}

            <div className={formik.errors.company && formik.touched.company ? errorInputInfo : inputInfo}>
              Company:
              <input
                className={formik.errors.company && formik.touched.company ? errorInputInfoInput : inputInfoInput}
                id={'company'}
                onBlur={formik.handleBlur}
                value={formik.values.company}
                onChange={formik.handleChange}
                type="text"
              />
            </div>
            {formik.errors.company && formik.touched.company ? (
              <div className={errorMsg}>{formik.errors.company}</div>
            ) : null}

            <div className={formik.errors.mm && formik.touched.mm ? errorInputInfo : inputInfo}>
              Month:
              <input
                className={formik.errors.mm && formik.touched.mm ? errorInputInfoInput : inputInfoInput}
                id={'mm'}
                onBlur={formik.handleBlur}
                value={formik.values.mm}
                onChange={formik.handleChange}
                type="text"
              />
            </div>
            {formik.errors.mm && formik.touched.mm ? <div className={errorMsg}>{formik.errors.mm}</div> : null}

            <div className={formik.errors.yy && formik.touched.yy ? errorInputInfo : inputInfo}>
              Year:
              <input
                className={formik.errors.yy && formik.touched.yy ? errorInputInfoInput : inputInfoInput}
                id={'yy'}
                onBlur={formik.handleBlur}
                value={formik.values.yy}
                onChange={formik.handleChange}
                type="text"
              />
            </div>
            {formik.errors.yy && formik.touched.yy ? <div className={errorMsg}>{formik.errors.yy}</div> : null}

            <div className={formik.errors.cvv && formik.touched.cvv ? errorInputInfo : inputInfo}>
              Cvv:
              <input
                className={formik.errors.cvv && formik.touched.cvv ? errorInputInfoInput : inputInfoInput}
                id={'cvv'}
                onBlur={formik.handleBlur}
                value={formik.values.cvv}
                onChange={formik.handleChange}
                type="text"
              />
            </div>
            {formik.errors.cvv && formik.touched.cvv ? <div className={errorMsg}>{formik.errors.cvv}</div> : null}
          </>
        )}
      </div>
    </div>
  )
}
