import s from './HomeTabletSmall.module.css';
import {HomeTabletSmallElement} from './HomeTabletSmallElement/HomeTabletSmallElement';

const {tabletTitle} = s

export const HomeTabletSmall = () => {
    return (
        <div className={tabletTitle}>
            <HomeTabletSmallElement type={'checkIn'}/>
            <HomeTabletSmallElement type={'checkOut'}/>
            <HomeTabletSmallElement type={'calendar'}/>
            <HomeTabletSmallElement type={'btnAvailability'}/>
        </div>
    )
}