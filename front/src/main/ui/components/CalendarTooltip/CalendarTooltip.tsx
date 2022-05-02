import React from 'react'
import s from './CalendarTooltip.module.scss'
import { useSelector } from 'react-redux'
import { AppRootState } from '../../../bll/store/store'
import moment from 'moment'
import { SeatResponseWithNewKey } from '../../../dal/api_client/SeatsService'

type PropsType = {
  tooltipDate: Date | null
}

export const CalendarTooltip = React.memo(({ tooltipDate }: PropsType) => {
  const day = moment(tooltipDate).format('MM DD YYYY')
  let filterDay = useSelector<AppRootState, SeatResponseWithNewKey>((state: AppRootState) => state.Seats.seatsSearch)[
    day
  ]?.seatIds.join(',')

  return (
    <div className={s.tooltipData}>
      <h3 className={s.tooltipTitle}>Available seats on {moment(tooltipDate).format('MM/DD/YY')}:</h3>
      {filterDay ? (
        <p className={s.tooltipText}>
          We have available seat in room <span className={s.tooltipTextMark}>{filterDay}</span> for your{' '}
          <span className={s.tooltipTextMark}>pet</span>
        </p>
      ) : (
        <p className={s.tooltipText}>Sorry, but there are no available rooms for this date</p>
      )}
    </div>
  )
})
