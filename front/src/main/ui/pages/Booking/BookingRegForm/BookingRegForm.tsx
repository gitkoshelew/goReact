import React from 'react'
import s from './BookingRegForm.module.scss'
import { ImgUpload } from './ImgUploadComponent/ImgUpload'
import { useDispatch, useSelector } from 'react-redux'
import { AppRootState } from '../../../../bll/store/store'
import Preloader from '../../../components/preloader/preloader'
import { FormikValues } from 'formik/dist/types'
import { Checkbox, FormControl, FormControlLabel, Radio, RadioGroup } from '@mui/material'
import { americanExpress, mastercard, visa } from '../../../svgWrapper/BookingRoomSvgWrapper'
import { changeCheckedOnlinePayment } from '../../../../bll/reducers/BookingRegFormReducer/BookingRegForm-reducer'

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

  const dispatch = useDispatch()

  const handleChangeCheckbox = (event: React.ChangeEvent<HTMLInputElement>) => {
    dispatch(changeCheckedOnlinePayment({ checkedOnlinePayment: event.target.checked }))
  }
  const checked = useSelector((state: AppRootState) => state.BookingRegForm.checkedOnlinePayment)

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
            placeholder={'Ivan'}
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
            placeholder={'Ivanov'}
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
            placeholder={'example@gmail.com'}
          />
        </div>
        {formik.errors.email && formik.touched.email ? <div className={errorMsg}>{formik.errors.email}</div> : null}
        <div className={inputInfo}>{actualContent()}</div>
        <div>
          <FormControlLabel
            control={<Checkbox checked={checked} onChange={handleChangeCheckbox} />}
            label="Online payment"
          />
        </div>

        {checked && (
          <>
            <div className={formik.errors.company && formik.touched.company ? errorInputInfo : inputInfo}>
              <FormControl>
                <RadioGroup row value={formik.values.company} onChange={formik.handleChange}>
                  <FormControlLabel
                    id={'mastercard'}
                    value={'mastercard'}
                    control={<Radio />}
                    label={<img src={mastercard} alt={'mastercard'} className={cardCompanyImages} />}
                    labelPlacement={'top'}
                    name={'company'}
                  />
                  <FormControlLabel
                    id={'visa'}
                    value={'visa'}
                    control={<Radio />}
                    label={<img src={visa} alt={'visa'} className={cardCompanyImages} />}
                    labelPlacement={'top'}
                    name={'company'}
                  />
                  <FormControlLabel
                    id={'americanExpress'}
                    value={'americanExpress'}
                    control={<Radio />}
                    label={<img src={americanExpress} alt={'americanExpress'} className={cardCompanyImages} />}
                    labelPlacement={'top'}
                    name={'company'}
                  />
                </RadioGroup>
              </FormControl>
            </div>
            {formik.errors.company && formik.touched.company ? (
              <div className={errorMsg}>{formik.errors.company}</div>
            ) : null}

            <div className={formik.errors.cardNumber && formik.touched.cardNumber ? errorInputInfo : inputInfo}>
              Card number:
              <input
                className={formik.errors.cardNumber && formik.touched.cardNumber ? errorInputInfoInput : inputInfoInput}
                id={'cardNumber'}
                onBlur={formik.handleBlur}
                value={formik.values.cardNumber}
                onChange={formik.handleChange}
                type="text"
                placeholder={'0000000000000000'}
              />
            </div>
            {formik.errors.cardNumber && formik.touched.cardNumber ? (
              <div className={errorMsg}>{formik.errors.cardNumber}</div>
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
                placeholder={'03'}
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
                placeholder={'2022'}
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
                type="password"
                placeholder={'xxx'}
              />
            </div>
            {formik.errors.cvv && formik.touched.cvv ? <div className={errorMsg}>{formik.errors.cvv}</div> : null}
          </>
        )}
      </div>
    </div>
  )
}
