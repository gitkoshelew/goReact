import React, { useEffect } from 'react'
import s from './CalendarTooltip.module.scss'
import { useSelector } from 'react-redux'
import { AppRootState, useAppDispatch } from '../../../bll/store/store'
import moment from 'moment'
import { seatsSearchDay } from '../../../bll/reducers/SeatsReducer/seats-reducer'

type PropsType = {
  tooltipDate: Date | null
}

export const CalendarTooltip = React.memo(({ tooltipDate }: PropsType) => {
  const dispatch = useAppDispatch()

  const strDate = tooltipDate ? tooltipDate.toString().split(' ') : ''

  const day = Number(strDate[2])
  useEffect(() => {
    dispatch(seatsSearchDay({ searchDay: day }))
  }, [day])

  console.log(day)
  const currentPickDay = useSelector((state: AppRootState) => state.Seats.searchDay)
  const seatSearchData = useSelector((state: AppRootState) =>
    state.Seats.seatsSearch[currentPickDay].seatIds.toString().split(' ')
  )

  console.log(seatSearchData)
  console.log(currentPickDay)

  return (
    <div className={s.tooltipData}>
      <h3 className={s.tooltipTitle}>Available seats on {moment(tooltipDate).format('MM/DD/YY')}:</h3>

      <p className={s.tooltipText}>
        We have available seat in room <span className={s.tooltipTextMark}>{seatSearchData}</span> for your{' '}
        <span className={s.tooltipTextMark}>pet</span>
      </p>
    </div>
  )
})
