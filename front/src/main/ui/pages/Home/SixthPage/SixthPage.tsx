import {TitlePageTextBlock} from '../../../components/TitlePageTextBlock/TitlePageTextBlock';
import s from './SixthPage.module.css';
import {FeedBack} from './FeedBack/FeedBack';
import {flower} from '../../../svgWrapper/HomeSvgWrapper';

const {sixthPage} = s;

export const SixthPage = () => {
    return (
        <div className={sixthPage}>
            <TitlePageTextBlock mainTextMess={'our guests love us'} isWithLink={false}/>
            <FeedBack feedBackMess={`The staff have been amazing and extremely helpful.
                They respond in a very friendly manner to all questions of us and we will comeback if we have a chance`}
                      nationality={'American'} userName={'Ralph Edwards'} photo={flower}/>
        </div>
    )
}