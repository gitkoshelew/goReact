import s from './RoomInfo.module.css';
import { children } from '../../svgWrapper/HomeSvgWrapper';
import { adult } from '../../svgWrapper/HomeSvgWrapper';
import { square } from '../../svgWrapper/HomeSvgWrapper';

const { roomInfo, roomInfoElement } = s
type RoomInfoPropsTypes = {
    childrenNum: number
    adultNum: number
    squareNum: number
}


export const RoomInfo = ({ childrenNum, adultNum, squareNum }: RoomInfoPropsTypes) => {
    return (
        <div className={roomInfo}>
            <div className={roomInfoElement}>
                <img src={children} alt="children"/>
                <p>{childrenNum} Children</p>
            </div>
            <div className={roomInfoElement}>
                <img src={adult} alt="adult"/>
                <p>{adultNum} Adult</p>
            </div>
            <div className={roomInfoElement}>
                <img src={square} alt="square"/>
                <p>{squareNum} ft²</p>
            </div>
        </div>
    )

}