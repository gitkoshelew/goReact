import React, { SetStateAction, useCallback, useState } from 'react'
import Calendar, { CalendarTileProperties } from 'react-calendar'
import 'react-calendar/dist/Calendar.css'
import moment, { MomentInput } from 'moment'
import { useSelector } from 'react-redux'
import { AppRootState, useAppDispatch } from '../../../../bll/store/store'
import { changeActualDay } from '../../../../bll/reducers/BookingRoomsPickReducer/BookingRoomPick-reducer'

export const BookingCalendar = () => {
  const isRentArr = useSelector((state: AppRootState) => state.BookingRoomPick.isRent)

  const dispatch = useAppDispatch()

  const [dateState, setDateState] = useState<Date>(new Date())

  const changeDate = useCallback(
    (e: MomentInput & SetStateAction<Date>) => {
      setDateState(e)
      dispatch(changeActualDay({ newActualDay: moment(e).format('MMDDYY') }))
    },
    [setDateState, dispatch]
  )

  const searchInRentArr = (props: CalendarTileProperties) => {
    const singleDay = isRentArr?.length && isRentArr.find((t) => t.id === moment(props.date).format('MMDDYY'))
    if (singleDay) {
      return !singleDay.secondRoom && !singleDay.firstRoom
    }
    return false
  }

  return (
    <>
      <Calendar
        tileDisabled={searchInRentArr}
        minDate={new Date()}
        defaultActiveStartDate={dateState}
        onChange={changeDate}
      />
    </>
  )
}
