import s from './TitlePageTextBlock.module.css';
import {arrow3} from '../../svgWrapper/HomeSvgWrapper';


const {TextBlockTitle, mainText, linkText} = s;

type TitleTextBlockPropsType = {
    mainTextMess: string
    linkTextMess?: string
    isWithLink: boolean
}


export const TitlePageTextBlock = ({mainTextMess, linkTextMess, isWithLink}: TitleTextBlockPropsType) => {
    return (
        <>
            {
                isWithLink ? <div className={TextBlockTitle}>
                    <div className={mainText}>
                        {mainTextMess.toUpperCase()}
                    </div>
                    <div className={linkText}>
                        {linkTextMess}
                        <img src={arrow3} alt="greenArrow"/>
                    </div>
                </div> : <div className={TextBlockTitle}>
                    <div className={mainText}>
                        {mainTextMess.toUpperCase()}
                    </div>
                </div>
            }
        </>
    )
}