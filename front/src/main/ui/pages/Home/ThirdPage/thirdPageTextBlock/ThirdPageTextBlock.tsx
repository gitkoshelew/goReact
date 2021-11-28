import s from './ThirdPageTextBlock.module.css';
import {arrow3} from '../../../../svgWrapper/HomeSvgWrapper';


const {thirdPageTextBlockTitle, mainText, linkText} = s;

export const ThirdPageTextBlock = () => {
    return (
        <div className={thirdPageTextBlockTitle}>
            <div className={mainText}>
                <p>OUR FAVORITE ROOMS</p>
            </div>
            <div className={linkText}>
                <p>See more</p>
                <img src={arrow3} alt="greenArrow"/>
            </div>
        </div>
    )
}