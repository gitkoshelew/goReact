import React, { FormEvent, useEffect, useMemo, useState } from 'react'
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
import { fetchBookingPaymentRequest } from '../../../bll/reducers/BookingPaymentFormReducer/bookingPaymentForm-saga'
import { Modal } from '../../components/Modal/Modal'
import { BookingSearchForm } from './BookingSearchForm/BookingSearchForm'
import { searchSeatsRequest } from '../../../bll/reducers/SeatsReducer/seats-saga'
import moment from 'moment'

const {
  bookingPage,
  bookingForm,
  bookingProcess,
  bookingCalendar,
  uploadOrderedRoomsBlock,
  searchPage,
  searchForm,
  searchProcess,
  updateCalendarInfo,
} = s

type SearchFormValues = {
  hotelId: string
  petType: string
  rentFrom: Date | string
  rentTo: Date | string
}

type BookingFormValues = {
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
  const [modalActiveBookingPayment, setModalActiveBookingPayment] = useState(false)
  const [modalActiveSeats, setModalActiveSeats] = useState(false)

  const dispatch = useAppDispatch()
  const checked = useSelector((state: AppRootState) => state.BookingRegForm.checkedOnlinePayment)

  const successMessageBookingPayment = useSelector((state: AppRootState) => state.BookingPaymentForm.successMsg)
  const errorMessageBookingPayment = useSelector((state: AppRootState) => state.BookingPaymentForm.errorMsg)

  const successMessageSeats = useSelector((state: AppRootState) => state.Seats.successMsg)
  const errorMessageSeats = useSelector((state: AppRootState) => state.Seats.errorMsg)

  const modalActiveBookingPaymentHandler = () => {
    setModalActiveBookingPayment(true)
  }

  const modalActiveSeatsHandler = () => {
    setModalActiveSeats(true)
  }

  const customSubmitBooking = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    formikBooking.handleSubmit()
    if (
      formikBooking.isValid &&
      formikBooking.values.firstName &&
      formikBooking.values.lastName &&
      formikBooking.values.email
    ) {
      modalActiveBookingPaymentHandler()
    }
  }

  const customSubmitSearch = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    formikSearch.handleSubmit()
    if (
      formikSearch.isValid &&
      formikSearch.values.hotelId &&
      formikSearch.values.petType &&
      formikSearch.values.rentTo &&
      formikSearch.values.rentFrom
    ) {
      modalActiveSeatsHandler()
    }
  }
  //Formik Search
  const initialDate: Date = new Date()
  const rentFromInitialDate = new Date(moment().format())
  const rentToInitialDate = new Date(moment(initialDate.setMonth(initialDate.getMonth() + 1)).format())

  const validateSearch = (values: SearchFormValues) => {
    const errors: FormikErrors<SearchFormValues> = {}
    if (!values.hotelId) {
      errors.hotelId = 'Required field'
    } else if (values.hotelId.toString().length > 1) {
      errors.hotelId = 'Must be 1 character'
    }
    if (!values.petType) {
      errors.petType = 'Required field'
    } else if (values.petType !== 'cat' && values.petType !== 'dog') {
      errors.petType = 'Please write cat or dog'
    }
    return errors
  }

  const formikSearch = useFormik({
    initialValues: {
      hotelId: '',
      petType: '',
      rentTo: rentToInitialDate,
      rentFrom: rentFromInitialDate,
    },
    validate: validateSearch,
    onSubmit: (values) => {
      dispatch(
        searchSeatsRequest({
          hotelId: Number(values.hotelId),
          petType: values.petType,
          rentTo: new Date(moment(values.rentTo).format()),
          rentFrom: new Date(moment(values.rentFrom).format()),
        })
      )
      if (formikSearch.isValid) {
        formikSearch.resetForm()
      }
    },
  })

  //Formik Booking
  useEffect(() => {
    if (!checked) {
      formikBooking.resetForm({
        values: {
          firstName: formikBooking.values.firstName,
          lastName: formikBooking.values.lastName,
          email: formikBooking.values.email,
          cardNumber: '',
          company: '',
          mm: '',
          yy: '',
          cvv: '',
        },
      })
    }
  }, [checked])

  const validateBooking = (values: BookingFormValues) => {
    const errors: FormikErrors<BookingFormValues> = {}
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

    if (checked && !values.cardNumber) {
      errors.cardNumber = 'Required field'
    } else if (
      checked &&
      !/^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$/i.test(
        values.cardNumber
      )
    ) {
      errors.cardNumber = 'Invalid card number'
    }

    if (checked && !values.company) {
      errors.company = 'Required field'
    }

    if (checked && !values.mm) {
      errors.mm = 'Required field'
    } else if (checked && !/(0[1-9]|1[012])$/i.test(values.mm)) {
      errors.mm = 'Invalid month'
    }

    if (checked && !values.yy) {
      errors.yy = 'Required field'
    } else if (checked && !/^20(2[2-9]|[3-9][0-9])$/i.test(values.yy)) {
      errors.yy = 'Invalid year'
    }

    if (checked && !values.cvv) {
      errors.cvv = 'Required field'
    } else if (checked && !/^(\d{3})$/i.test(values.cvv)) {
      errors.cvv = 'Invalid cvv'
    }
    return errors
  }

  const formikBooking = useFormik({
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
    validate: validateBooking,
    onSubmit: (values) => {
      dispatch(fetchBookingPaymentRequest(values))
      if (formikBooking.isValid) {
        formikBooking.resetForm()
      }
    },
  })

  const loadingStatus = useSelector((state: AppRootState) => state.BookingRoomPick.loadingStatus)

  const ErrorView = loadingStatus === 'error' ? <div>error</div> : <BookingCalendar />
  const correctView = loadingStatus === 'loading' ? <Preloader /> : ErrorView

  useEffect(() => {
    dispatch(BookingRoomPickSaga())
  }, [])

  const progress = useSelector((state: AppRootState) => state.BookingRegForm.progress)
  const actualDay = useSelector((state: AppRootState) => state.BookingRoomPick.actualDay)
  const isRentArr = useSelector((state: AppRootState) => state.BookingRoomPick.isRent)
  const orderedRoomBasket = useSelector((state: AppRootState) => state.BookingRoomPick.orderedRoomBasket)

  const isActiveBtnBooking = progress === 'uploaded' && formikBooking.isValid
  const isActiveBtnSearch = formikSearch.isValid

  const roomIndicate = useMemo(() => {
    const newActualDay = isRentArr && isRentArr.find((t) => t.id === actualDay)
    return newActualDay ? newActualDay : null
  }, [actualDay, isRentArr])

  const bookingPaymentLoadingStatus = useSelector((state: AppRootState) => state.BookingPaymentForm.loadingStatus)
  const seatsLoadingStatus = useSelector((state: AppRootState) => state.Seats.loadingStatus)

  if (bookingPaymentLoadingStatus === 'LOADING') {
    return <Preloader />
  }
  if (seatsLoadingStatus === 'LOADING') {
    return <Preloader />
  }

  return (
    <div className="bookingContainer">
      <form onSubmit={customSubmitSearch}>
        <div className={searchPage}>
          <TitlePageTextBlock mainTextMess={'Write pet and hotel'} isWithLink={false} />
          <div className={searchProcess}>
            <div className={searchForm}>
              <BookingSearchForm formik={formikSearch} />
              <div className={updateCalendarInfo}>
                <Button view={'upload'} disabled={!isActiveBtnSearch} />
              </div>
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

          <Modal active={modalActiveSeats} setActive={setModalActiveSeats}>
            <p>{successMessageSeats ? 'Please, look at the calendar. Choose an available room!' : errorMessageSeats}</p>
          </Modal>
        </div>
      </form>

      <form onSubmit={customSubmitBooking}>
        <div className={bookingPage}>
          <TitlePageTextBlock mainTextMess={'Book room for pet'} isWithLink={false} />
          <div className={bookingProcess}>
            <div className={bookingForm}>
              <BookingRegForm formik={formikBooking} />
            </div>
          </div>
          <div className={uploadOrderedRoomsBlock}>
            {orderedRoomBasket.length !== 0 && <SelectedToOrderRoom orderedRoomBasket={orderedRoomBasket} />}
            <Button view={'upload'} disabled={!isActiveBtnBooking} />
            <Modal active={modalActiveBookingPayment} setActive={setModalActiveBookingPayment}>
              <p>
                {successMessageBookingPayment
                  ? 'Congratulations! You have successfully made a payment!'
                  : errorMessageBookingPayment}
              </p>
            </Modal>
          </div>
        </div>
      </form>
    </div>
  )
}
