import React, { useMemo } from 'react'
import s from './CalendarTooltip.module.scss'
import { useSelector } from 'react-redux'
import { AppRootState } from '../../../bll/store/store'
import moment from 'moment'

type PropsType = {
  tooltipDate: Date | null
}

export const CalendarTooltip = React.memo(({ tooltipDate }: PropsType) => {
  const rooms = useSelector((state: AppRootState) => state.Seats.rooms)
  const seats = useSelector((state: AppRootState) => state.Seats.seats)

  const availableSeats = useMemo(() => {
    return seats.filter((seat) => {
      if (tooltipDate) {
        if (!moment(tooltipDate).isBetween(moment(seat.rentFrom), moment(seat.rentTo), 'date', '[]')) return seat
      }
    })
  }, [seats, tooltipDate])

  const searchRoomData = (seatRoomId: number) => {
    const petType = rooms.find((room) => room.roomId === seatRoomId)?.petType
    const roomNum = rooms.find((room) => room.roomId === seatRoomId)?.roomNum
    return { petType, roomNum }
  }

  const availableSeatsInRoom = useMemo(() => {
    return availableSeats.map((seat) => {
      return searchRoomData(seat.roomId)
    })
  }, [availableSeats])

  return (
    <div className={s.tooltipContainer}>
      <div className={s.tooltipData}>
        <h3 className={s.tooltipTitle}>Available seats on {moment(tooltipDate).format('MM/DD/YY')}:</h3>
        {availableSeatsInRoom.map((seat, index) => (
          <p key={index} className={s.tooltipText}>
            We have available seat in room <span className={s.tooltipTextMark}>№{seat.roomNum}</span> for your{' '}
            <span className={s.tooltipTextMark}>{seat.petType}</span>
          </p>
        ))}
      </div>
    </div>
  )
})
