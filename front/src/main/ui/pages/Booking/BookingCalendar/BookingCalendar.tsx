import React, { SetStateAction, useCallback, useEffect, useState } from 'react'
import Calendar, { CalendarTileProperties } from 'react-calendar'
import 'react-calendar/dist/Calendar.css'
import s from './BookingCalendar.module.scss'
import moment, { MomentInput } from 'moment'
import { useSelector } from 'react-redux'
import { AppRootState, useAppDispatch } from '../../../../bll/store/store'
import { changeActualDay } from '../../../../bll/reducers/BookingRoomsPickReducer/BookingRoomPick-reducer'
import { CalendarTooltipAnchor } from '../../../components/CalendarTooltip/CalendarTooltipAnchor'
import { CalendarTooltip } from '../../../components/CalendarTooltip/CalendarTooltip'
import { fetchSeatsRequest } from '../../../../bll/reducers/SeatsReducer/seats-saga'

export const BookingCalendar = () => {
  const isRentArr = useSelector((state: AppRootState) => state.BookingRoomPick.isRent)
  const startDate = useSelector((state: AppRootState) => state.Seats.seatsSearch[0])
  const dispatch = useAppDispatch()

  const [dateState, setDateState] = useState<Date>(new Date(startDate?.rentFrom ? startDate.rentFrom : new Date()))

  const [tooltipVisible, setTooltipVisible] = useState<boolean>(false)
  const [tooltipDate, setTooltipDate] = useState<Date | null>(null)

  useEffect(() => {
    dispatch(fetchSeatsRequest())
  }, [])

  const showTooltip = useCallback((isVisible: boolean) => {
    setTooltipVisible(isVisible)
  }, [])
  const handleTooltipDate = useCallback((date: Date | null) => {
    setTooltipDate(date)
  }, [])

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
    <div className={s.calendarContainer}>
      <Calendar
        tileContent={(props) => (
          <CalendarTooltipAnchor date={props.date} showTooltip={showTooltip} handleTooltipDate={handleTooltipDate} />
        )}
        tileClassName={s.calendarCell}
        tileDisabled={searchInRentArr}
        minDate={new Date(2022, 2, 30)}
        defaultActiveStartDate={dateState}
        onChange={changeDate}
      />
      {tooltipVisible && <CalendarTooltip tooltipDate={tooltipDate} />}
    </div>
  )
}
