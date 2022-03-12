import React, { useEffect, useMemo } from 'react'
import s from './Booking.module.scss'
import { TitlePageTextBlock } from '../../components/TitlePageTextBlock/TitlePageTextBlock'
import { BookingRegForm } from './BookingRegForm/BookingRegForm'
import { BookingCalendar } from './BookingCalendar/BookingCalendar'
import { BookingRoom } from './BookingRoom/BookingRoom'
import { useSelector } from 'react-redux'
import { AppRootState, useAppDispatch } from '../../../bll/store/store'
import { Button } from '../../components/Button/Button'
import { SelectedToOrderRoom } from './SelectedToOrderRom/SelectedToOrderRoom'
import { FormikErrors, useFormik } from 'formik'
import Preloader from '../../components/preloader/preloader'
import { BookingRoomPickSaga } from '../../../bll/reducers/BookingRoomsPickReducer/BookingRoomPick-saga'

const { bookingPage, bookingForm, bookingProcess, bookingCalendar, uploadOrderedRoomsBlock } = s

type FormValues = {
  firstName: string
  lastName: string
  email: string
  cardNumber: string
  company: string
  mm: string
  yy: string
  cvv: string
}

export const Booking = () => {
  const validate = (values: FormValues) => {
    const errors: FormikErrors<FormValues> = {}
    if (!values.firstName) {
      errors.firstName = 'Required field'
    } else if (values.firstName.length > 15) {
      errors.firstName = 'Must be 15 characters or less'
    }
    if (!values.lastName) {
      errors.lastName = 'Required field'
    } else if (values.lastName.length > 15) {
      errors.lastName = 'Must be 15 characters or less'
    }
    if (!values.email) {
      errors.email = 'Required field'
    } else if (!/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(values.email)) {
      errors.email = 'Invalid email address'
    }

    if (!values.cardNumber) {
      errors.cardNumber = 'Required field'
    } else if (
      !/^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$/i.test(
        values.cardNumber
      )
    ) {
      errors.cardNumber = 'Invalid card number'
    }

    if (!values.company) {
      errors.company = 'Required field'
    } else if (values.company.length > 15) {
      errors.company = 'Invalid company name'
    }

    if (!values.mm) {
      errors.mm = 'Required field'
    } else if (!/(0[1-9]|1[012])$/i.test(values.mm)) {
      errors.mm = 'Invalid month'
    }

    if (!values.yy) {
      errors.yy = 'Required field'
    } else if (!/^20(2[2-9]|[2-9][0-9])$/i.test(values.yy)) {
      errors.yy = 'Invalid year'
    }

    if (!values.cvv) {
      errors.cvv = 'Required field'
    } else if (!/(\d{3})$/i.test(values.cvv)) {
      errors.cvv = 'Invalid cvv'
    }
    return errors
  }

  const formik = useFormik({
    initialValues: {
      firstName: '',
      lastName: '',
      email: '',
      cardNumber: '',
      company: '',
      mm: '',
      yy: '',
      cvv: '',
    },
    validate,
    onSubmit: (values) => {
      alert(JSON.stringify(values, null, 2))
    },
  })

  const loadingStatus = useSelector((state: AppRootState) => state.BookingRoomPick.loadingStatus)

  const ErrorView = loadingStatus === 'error' ? <div>error</div> : <BookingCalendar />
  const correctView = loadingStatus === 'loading' ? <Preloader /> : ErrorView
  const dispatch = useAppDispatch()

  useEffect(() => {
    dispatch(BookingRoomPickSaga())
  }, [])

  const progress = useSelector((state: AppRootState) => state.BookingRegForm.progress)
  const actualDay = useSelector((state: AppRootState) => state.BookingRoomPick.actualDay)
  const isRentArr = useSelector((state: AppRootState) => state.BookingRoomPick.isRent)
  const orderedRoomBasket = useSelector((state: AppRootState) => state.BookingRoomPick.orderedRoomBasket)

  const isActiveBtn = progress === 'uploaded' && orderedRoomBasket.length !== 0

  const roomIndicate = useMemo(() => {
    const newActualDay = isRentArr && isRentArr.find((t) => t.id === actualDay)
    return newActualDay ? newActualDay : null
  }, [actualDay, isRentArr])

  return (
    <div className="bookingContainer">
      <form onSubmit={formik.handleSubmit}>
        <div className={bookingPage}>
          <TitlePageTextBlock mainTextMess={'Book room for pet'} isWithLink={false} />
          <div className={bookingProcess}>
            <div className={bookingForm}>
              <BookingRegForm formik={formik} />
            </div>
            <div className={bookingCalendar}>
              {correctView}
              {roomIndicate && (
                <BookingRoom
                  dayId={roomIndicate.id}
                  firstRoom={roomIndicate.firstRoom}
                  secondRoom={roomIndicate.secondRoom}
                />
              )}
            </div>
          </div>
          <div className={uploadOrderedRoomsBlock}>
            {orderedRoomBasket.length !== 0 && <SelectedToOrderRoom orderedRoomBasket={orderedRoomBasket} />}
            <Button view={'upload'} disabled={!isActiveBtn} />
          </div>
        </div>
      </form>
    </div>
  )
}
