import s from './footer.module.css';
import {ContactInfo} from './ContactInfo/ContactInfo';
import {UsefulLinks} from './UsefulLinks/UsefulLinks';

const {footerTitle, contactInfo, usefulLinks, subscribeFieldsBlock} = s;


export const Footer = () => {
    return (
        <div className={footerTitle}>
            <div>
                <ContactInfo/>
            </div>
            <UsefulLinks/>
            <div className={subscribeFieldsBlock}></div>
        </div>
    )
}