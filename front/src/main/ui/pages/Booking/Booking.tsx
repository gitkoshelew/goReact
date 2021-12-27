import React, { useEffect, useMemo } from 'react'
import s from './Booking.module.css'
import { TitlePageTextBlock } from '../../components/TitlePageTextBlock/TitlePageTextBlock'
import { BookingRegForm } from './BookingRegForm/BookingRegForm'
import { BookingCalendar } from './BookingCalendar/BookingCalendar'
import { BookingRoom } from './BookingRoom/BookingRoom'
import { useSelector } from 'react-redux'
import { AppRootStateType, useAppDispatch } from '../../../bll/store/store'
import {
  IsRentType,
  LoadingStatusBookingPickType,
  OrderedRoomsType,
} from '../../../bll/reducers/BookingRoomsPickReducer/BookingRoomPick-reducer'
import { Button } from '../../components/Button/Button'
import { ProgressType } from '../../../bll/reducers/BookingRegFormReducer/BookingRegForm-reducer'
import { SelectedToOrderRoom } from './SelectedToOrderRom/SelectedToOrderRoom'
import { FormikErrors, useFormik } from 'formik'
import Preloader from '../../components/preloader/preloader'
import { ErrorMsg } from '../../components/ErrorMsg/ErrorMsg'
import { BookingRoomPickSaga } from '../../../bll/reducers/BookingRoomsPickReducer/BookingRoomPick-saga'

const { bookingPage, bookingForm, bookingProcess, bookingCalendar, uploadOrderedRoomsBlock } = s

type FormValues = {
  firstName: string
  lastName: string
  email: string
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
    return errors
  }

  const formik = useFormik({
    initialValues: {
      firstName: '',
      lastName: '',
      email: '',
    },
    validate,
    onSubmit: (values) => {
      alert(JSON.stringify(values, null, 2))
    },
  })

  const loadingStatus = useSelector<AppRootStateType, LoadingStatusBookingPickType>(
    (state) => state.BookingRoomPick.loadingStatus
  )

  const isError = loadingStatus === 'error' ? <ErrorMsg /> : <BookingCalendar />
  const correctView = loadingStatus === 'loading' ? <Preloader /> : isError

  const dispatch = useAppDispatch()

  useEffect(() => {
    dispatch(BookingRoomPickSaga())
  }, [])

  const progress = useSelector<AppRootStateType, ProgressType>((state) => state.BookingRegForm.progress)
  const actualDay = useSelector<AppRootStateType, string | Date>((state) => state.BookingRoomPick.actualDay)
  const isRentArr = useSelector<AppRootStateType, IsRentType[]>((state) => state.BookingRoomPick.isRent)
  const orderedRoomBasket = useSelector<AppRootStateType, OrderedRoomsType[]>(
    (state) => state.BookingRoomPick.orderedRoomBasket
  )

  const isActiveBtn = progress === 'uploaded' && orderedRoomBasket.length !== 0

  const roomIndicate = useMemo(() => {
    const newActualDay = isRentArr && isRentArr.find((t) => t.id === actualDay)
    return newActualDay ? newActualDay : null
  }, [actualDay, isRentArr])

  return (
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
          <Button type={'upload'} isActive={isActiveBtn} />
        </div>
      </div>
    </form>
  )
}
