import React from 'react'
import s from './BookingSearchForm.module.scss'
import { FormikValues } from 'formik/dist/types'

const { bookingForm, clientDescription, inputInfo, inputInfoInput, errorMsg, errorInputInfo, errorInputInfoInput } = s

type BookingSearchFormType = {
  formik: FormikValues
}

export const BookingSearchForm = ({ formik }: BookingSearchFormType) => {
  return (
    <div className={bookingForm}>
      <div className={clientDescription}>
        <div className={formik.errors.hotelId && formik.touched.hotelId ? errorInputInfo : inputInfo}>
          Hotel Id:
          <input
            className={formik.errors.hotelId && formik.touched.hotelId ? errorInputInfoInput : inputInfoInput}
            id={'hotelId'}
            onBlur={formik.handleBlur}
            value={formik.values.hotelId}
            onChange={formik.handleChange}
            type="number"
            placeholder={'Id'}
          />
        </div>
        {formik.errors.hotelId && formik.touched.hotelId ? (
          <div className={errorMsg}>{formik.errors.hotelId}</div>
        ) : null}

        <div className={formik.errors.petType && formik.touched.petType ? errorInputInfo : inputInfo}>
          Pet Type:
          <input
            className={formik.errors.petType && formik.touched.petType ? errorInputInfoInput : inputInfoInput}
            id={'petType'}
            onBlur={formik.handleBlur}
            value={formik.values.petType}
            onChange={formik.handleChange}
            type="text"
            placeholder={'cat or dog'}
          />
        </div>
        {formik.errors.petType && formik.touched.petType ? (
          <div className={errorMsg}>{formik.errors.petType}</div>
        ) : null}

        <div className={formik.errors.rentFrom && formik.touched.rentFrom ? errorInputInfo : inputInfo}>
          Rent form
          <input
            className={formik.errors.rentFrom && formik.touched.rentFrom ? errorInputInfoInput : inputInfoInput}
            id={'rentFrom'}
            onBlur={formik.handleBlur}
            value={formik.values.rentFrom}
            onChange={formik.handleChange}
            type="date"
            placeholder={'dd/mm/yyyy'}
          />
        </div>
        {formik.errors.rentFrom && formik.touched.rentFrom ? (
          <div className={errorMsg}>{formik.errors.rentFrom}</div>
        ) : null}

        <div className={formik.errors.rentTo && formik.touched.rentTo ? errorInputInfo : inputInfo}>
          Rent To:
          <input
            className={formik.errors.rentTo && formik.touched.rentTo ? errorInputInfoInput : inputInfoInput}
            id={'rentTo'}
            onBlur={formik.handleBlur}
            value={formik.values.rentTo}
            onChange={formik.handleChange}
            type="date"
            placeholder={'dd/mm/yyyy'}
          />
        </div>
        {formik.errors.rentTo && formik.touched.rentTo ? <div className={errorMsg}>{formik.errors.rentTo}</div> : null}
      </div>
    </div>
  )
}
