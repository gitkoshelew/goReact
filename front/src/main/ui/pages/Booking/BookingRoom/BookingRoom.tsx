import { MiniRoomElement } from '../../Home/FavoriteRoomsPage/miniRoomBlock/MiniRoomElement/MiniRoomElement'
import s from './BookingRoom.module.scss'
import { rented } from '../../../svgWrapper/BookingRoomSvgWrapper'
import { Button } from '../../../components/Button/Button'
import { useAppDispatch } from '../../../../bll/store/store'
import {
  addOrderedRoom,
  changeRoomStatus,
  RentRoomType,
} from '../../../../bll/reducers/BookingRoomsPickReducer/BookingRoomPick-reducer'

const { roomBookingTitle, oneBookRoom, rentedImg } = s

type BookingRoomPropsTypes = {
  firstRoom: boolean
  secondRoom: boolean
  dayId: string
}

export const BookingRoom = (props: BookingRoomPropsTypes) => {
  const { secondRoom, firstRoom, dayId } = props
  const dispatch = useAppDispatch()

  const onOrderClickHandler = (type: RentRoomType) => {
    dispatch(changeRoomStatus({ roomType: type, dayId }))
    dispatch(addOrderedRoom({ newOrderedRooms: { id: dayId, orderedRoomType: type } }))
  }

  return (
    <div className={roomBookingTitle}>
      <div className={oneBookRoom}>
        {firstRoom ? (
          <>
            <MiniRoomElement childrenNum={1} adultNum={1} squareNum={100} price={100} roomName={'First Room'} />
            <Button view={'order'} onClick={() => onOrderClickHandler('firstRoom')} />
          </>
        ) : (
          <div className={rentedImg}>
            <img src={rented} alt="rentedRoom" />
          </div>
        )}
      </div>
      <div className={oneBookRoom}>
        {secondRoom ? (
          <>
            <MiniRoomElement childrenNum={1} adultNum={1} squareNum={100} price={100} roomName={'Second Room'} />
            <Button view={'order'} onClick={() => onOrderClickHandler('secondRoom')} />
          </>
        ) : (
          <div className={rentedImg}>
            <img src={rented} alt="rentedRoom" />
          </div>
        )}
      </div>
    </div>
  )
}
