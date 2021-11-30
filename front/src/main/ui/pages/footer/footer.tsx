import s from './footer.module.css';
import {ContactInfo} from './ContactInfo/ContactInfo';
import {UsefulLinks} from './UsefulLinks/UsefulLinks';
import {SubscribeFieldBlock} from './SubscribeFieldBlock/SubscribeFieldBlock';
import {AfterPanel} from './SubscribeFieldBlock/AfterPanel/AfterPanel';

const {footerTitle, footerBlock, afterPanel} = s;


export const Footer = () => {
    return (
        <div className={footerBlock}>
            <div className={footerTitle}>
                <ContactInfo/>
                <UsefulLinks/>
                <SubscribeFieldBlock/>
            </div>
            <div>
                <AfterPanel/>
            </div>
        </div>
    )
}