import s from './BigRoomBlock.module.css';
import {PriceWindow} from '../../../../components/priceWindow/PriceWindow';
import {RoomInfo} from '../../../../components/RoomInfo/RoomInfo';

const {bigRoomBlock, roomNames} = s;

type BigRoomBlockType = {
    childrenNum: number
    adultNum: number
    squareNum: number
    price: number
    roomName: string
}


export const BigRoomBlock = ({childrenNum, adultNum, squareNum, price, roomName}: BigRoomBlockType) => {
    return (
        <div className={bigRoomBlock}>
            <span>630x630</span>
            <PriceWindow price={price}/>
            <RoomInfo childrenNum={childrenNum} adultNum={adultNum} squareNum={squareNum}/>
            <div className={roomNames}>{roomName} room</div>
        </div>
    )
}