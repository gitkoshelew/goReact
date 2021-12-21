import s from "./SelectedToOrderRoom.module.css";
import {
  deleteOrderedRoom,
  OrderedRoomsType
} from "../../../../bll/reducers/BookingRoomsPickReducer/BookingRoomPick-reducer";
import moment from "moment";
import { deleteIcon } from "../../../svgWrapper/BookingRoomSvgWrapper";
import { useAppDispatch } from "../../../../bll/store/store";
import { useCallback } from "react";

const { selectedRoomsTablet, roomType } = s;

type SelectedToOrderRoomPropsType = {
  orderedRoomBasket: OrderedRoomsType[]
}

export const SelectedToOrderRoom = ({ orderedRoomBasket }: SelectedToOrderRoomPropsType) => {
  const dispatch = useAppDispatch();

  const onDeleteOrderRoomHandler = useCallback((roomToDelete: OrderedRoomsType) => {
    dispatch(deleteOrderedRoom({ newOrderedRooms: roomToDelete }));
  }, []);

  const orderedRoomsView = orderedRoomBasket.map((t, i) => (
    <div key={i}>
      <div>{moment(t.id, "MMDDYY").format("DD MMMM YYYY")}</div>
      <div className={roomType}>
        {t.orderedRoomType}
        <img onClick={() => onDeleteOrderRoomHandler(t)} src={deleteIcon} alt="deleteHandler" />
      </div>
    </div>
  ));

  return (
    <div className={selectedRoomsTablet}>
      <div>
        <div>Date</div>
        <div>Room Name</div>
      </div>
      {orderedRoomsView}
    </div>
  );
};
