import s from './MiniRoomElement.module.css';
import {PriceWindow} from '../../../../../components/priceWindow/PriceWindow';
import {RoomInfo} from '../../../../../components/RoomInfo/RoomInfo';

const {oneRoom, rePositionInfo, roomNames} = s;

type MiniRoomElement = {
    childrenNum: number
    adultNum: number
    squareNum: number
    price: number
    roomName: string
}

export const MiniRoomElement = ({childrenNum, adultNum, squareNum, price, roomName}: MiniRoomElement) => {
    return (
        <div className={oneRoom}>
            <PriceWindow price={price}/>
            <span>300x300</span>
            <div className={roomNames}>{roomName} room</div>
            <div className={rePositionInfo}>
                <RoomInfo childrenNum={childrenNum} adultNum={adultNum} squareNum={squareNum}/>
            </div>
        </div>

    )
}