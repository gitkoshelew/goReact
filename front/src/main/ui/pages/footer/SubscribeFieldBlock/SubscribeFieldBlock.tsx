import s from './SubscribeFieldBlock.module.css';
import { EmailField } from './EmailField/EmailField';
import { ImgLinks } from './ImgLinks/ImgLinks';


const { subscribeField, linksTitle } = s;


export const SubscribeFieldBlock = () => {
    return (
        <div className={subscribeField}>
            <div className={linksTitle}>
                STAY IN TOUCH
            </div>
            <div>
                <EmailField/>
                <ImgLinks/>
            </div>

        </div>
    )
}