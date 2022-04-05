import s from './SelectedToOrderRoom.module.scss'
import {
  deleteOrderedRoom,
  OrderedRoomsType,
} from '../../../../bll/reducers/BookingRoomsPickReducer/BookingRoomPick-reducer'
import moment from 'moment'
import { deleteIcon } from '../../../svgWrapper/BookingRoomSvgWrapper'
import { useAppDispatch } from '../../../../bll/store/store'
import { useCallback } from 'react'

const {
  selectedRoomsTablet,
  selectedRoomsTabletElem,
  selectedRoomsTabletTitle,
  orderedRoomList,
  onDeleteOrderRoomIcon,
} = s

type SelectedToOrderRoomPropsType = {
  orderedRoomBasket: OrderedRoomsType[]
}

export const SelectedToOrderRoom = ({ orderedRoomBasket }: SelectedToOrderRoomPropsType) => {
  const dispatch = useAppDispatch()

  const onDeleteOrderRoomHandler = useCallback(
    (roomToDelete: OrderedRoomsType) => {
      dispatch(deleteOrderedRoom({ newOrderedRooms: roomToDelete }))
    },
    [dispatch]
  )

  const orderedRoomsView = orderedRoomBasket.map((t, i) => (
    <div className={orderedRoomList} key={i}>
      <div className={selectedRoomsTabletElem}>{moment(t.id, 'MMDDYY').format('DD MMMM YYYY')}</div>
      <div className={selectedRoomsTabletElem}>
        {t.orderedRoomType}
        <img
          className={onDeleteOrderRoomIcon}
          onClick={() => onDeleteOrderRoomHandler(t)}
          src={deleteIcon}
          alt="deleteHandler"
        />
      </div>
    </div>
  ))

  return (
    <div className={selectedRoomsTablet}>
      <div className={selectedRoomsTabletTitle}>
        <div className={selectedRoomsTabletElem}>Date</div>
        <div className={selectedRoomsTabletElem}>Room Name</div>
      </div>
      {orderedRoomsView}
    </div>
  )
}
