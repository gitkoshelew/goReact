import s from './HomeTabletElement.module.css';
import calendar from '../../../../../assets/img/home/Calendar.svg'
import usersShadow from '../../../../../assets/img/home/usersShadow.svg'

const {
    tabletComponent,
    tabletComponentTitle,
    tabletComponentText,
    tabletComponentBody,
    tabletComponentBodyBtn
} = s;

type HomeTabletElementPropsType = {
    type: string
}


export const HomeTabletElement = (props: HomeTabletElementPropsType) => {
    const {type} = props;
    return (
        <>
            {type === 'checkIn' &&
            <div className={tabletComponent}>
                <div className={tabletComponentTitle}>
                    <div><img src={calendar} alt="calendar"/></div>
                    <div className={tabletComponentText}>Check in</div>
                </div>
                <div className={tabletComponentBody}>
                    <input required={true} type="date" value="2021-07-22"/>
                </div>
            </div>}
            {type === 'checkOut' &&
            <div className={tabletComponent}>
                <div className={tabletComponentTitle}>
                    <div><img src={calendar} alt="calendar"/></div>
                    <div className={tabletComponentText}>Check out</div>
                </div>
                <div className={tabletComponentBody}>
                    <input type="date" value="2021-07-22"/>
                </div>
            </div>}
            {type === 'calendar' &&
            <div className={tabletComponent}>
                <div className={tabletComponentTitle}>
                    <div><img src={usersShadow} alt="calendar"/></div>
                    <div className={tabletComponentText}>Guest</div>
                </div>
                <div className={tabletComponentBody}>
                    <select name="howMuchPerson">
                        <option>4 Persons</option>
                        <option>3 Persons</option>
                        <option>2 Persons</option>
                        <option>1 Persons</option>
                        <option>more</option>
                    </select>
                </div>
            </div>}
            {type === 'btnAvailability' &&
            <div className={tabletComponentBodyBtn}>
                <p>Check Availability</p>
            </div>}

        </>
    )
}