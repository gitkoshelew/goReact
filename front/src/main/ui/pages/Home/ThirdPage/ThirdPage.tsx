import s from './ThirdPage.module.css';
import {ThirdPageTextBlock} from './thirdPageTextBlock/ThirdPageTextBlock';
import {BigRoomBlock} from './BigRoomBlock/BigRoomBlock';
import {MiniRoomBlock} from './miniRoomBlock/MiniRoomBlock';


const {thirdPage, photoBlock} = s;

export const ThirdPage = () => {
    return (
        <div className={thirdPage}>
            <ThirdPageTextBlock/>
            <div className={photoBlock}>
                <BigRoomBlock roomName={'Luxure'} price={100} adultNum={2} childrenNum={2} squareNum={100}/>
                <MiniRoomBlock/>
            </div>

        </div>
    )
}