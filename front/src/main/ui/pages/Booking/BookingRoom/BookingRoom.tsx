import { MiniRoomElement } from '../../Home/ThirdPage/miniRoomBlock/MiniRoomElement/MiniRoomElement'
import s from './BookingRoom.module.css'
import { rented } from '../../../svgWrapper/BookingRoomSvgWrapper'

const { roomBookingTitle, oneBookRoom, rentedImg } = s

type BookingRoomPropsTypes = {
  firstRoom: boolean
  secondRoom: boolean
}

export const BookingRoom = (props: BookingRoomPropsTypes) => {
  const { secondRoom, firstRoom } = props

  return (
    <div className={roomBookingTitle}>
      <div className={oneBookRoom}>
        {firstRoom ? (
          <MiniRoomElement childrenNum={1} adultNum={1} squareNum={100} price={100} roomName={'First Room'} />
        ) : (
          <div className={rentedImg}>
            <img src={rented} alt="rentedRoom" />
          </div>
        )}
      </div>
      <div>
        {secondRoom ? (
          <MiniRoomElement childrenNum={1} adultNum={1} squareNum={100} price={100} roomName={'Second Room'} />
        ) : (
          <div className={rentedImg}>
            <img src={rented} alt="rentedRoom" />
          </div>
        )}
      </div>
    </div>
  )
}
