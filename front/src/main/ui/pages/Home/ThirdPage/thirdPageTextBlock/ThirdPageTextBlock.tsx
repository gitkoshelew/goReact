import s from './ThirdPageTextBlock.module.css';
import arrow from '../../../../../../assets/img/home/thirdPage/greenArrow.svg';


const {thirdPageTextBlockTitle,bigBlackText,smallGreenText} = s;

export const ThirdPageTextBlock = () => {
    return (
        <div className={thirdPageTextBlockTitle}>
            <div className={bigBlackText}>
                <p>OUR FAVORITE ROOMS</p>
            </div>
            <div className={smallGreenText}>
                <p>See more</p>
                <img src={arrow} alt="greenArrow"/>
            </div>
        </div>
    )
}