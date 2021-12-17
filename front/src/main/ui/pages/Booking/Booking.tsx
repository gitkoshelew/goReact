import React, { useEffect, useState } from 'react'
import s from './Booking.module.css'
import { TitlePageTextBlock } from '../../components/TitlePageTextBlock/TitlePageTextBlock'
import { BookingRegForm } from './BookingRegForm/BookingRegForm'
import { BookingCalendar } from './BookingCalendar/BookingCalendar'
import { BookingRoom } from './BookingRoom/BookingRoom'
import { useSelector } from 'react-redux'
import { AppRootStateType } from '../../../bll/store/store'
import { IsRentType } from '../../../bll/reducers/BookingRoomsPickReducer/BookingRoomPick-reducer'
import { Button } from '../../components/Button/Button'
import { ProgressType } from '../../../bll/reducers/BookingRegFormReducer/BookingRegForm-reducer'

const { bookingPage, bookingForm, bookingProcess, bookingCalendar } = s

export const Booking = () => {
  const progress = useSelector<AppRootStateType, ProgressType>((state) => state.BookingRegForm.progress)
  const actualDay = useSelector<AppRootStateType, string | Date>((state) => state.BookingRoomPick.actualDay)
  const isRentArr = useSelector<AppRootStateType, IsRentType[]>((state) => state.BookingRoomPick.isRent)
  const userName = useSelector<AppRootStateType, string>((state) => state.BookingRoomPick.userName)
  const userEmail = useSelector<AppRootStateType, string>((state) => state.BookingRoomPick.userEmail)
  const [roomIndicate, setRoomIndicate] = useState<null | IsRentType>(null)

  const isActiveBtn = progress === 'uploaded' && !!userName && !!userEmail

  useEffect(() => {
    const newActualDay = isRentArr.find((t) => t.id === actualDay)
    newActualDay ? setRoomIndicate(newActualDay) : setRoomIndicate(null)
  }, [actualDay])

  return (
    <div className={bookingPage}>
      <TitlePageTextBlock mainTextMess={'Book room for pet'} isWithLink={false} />
      <div className={bookingProcess}>
        <div className={bookingForm}>
          <BookingRegForm />
        </div>
        <div className={bookingCalendar}>
          <BookingCalendar />
          {roomIndicate && <BookingRoom firstRoom={roomIndicate.firstRoom} secondRoom={roomIndicate.secondRoom} />}
        </div>
      </div>
      <Button type={'Upload'} isActive={isActiveBtn} />
    </div>
  )
}
