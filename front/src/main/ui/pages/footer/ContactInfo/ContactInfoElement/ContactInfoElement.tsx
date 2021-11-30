import s from './ConstactInfoElement.module.css';


const {contactInfoElement, contactInfoElementLinks, contactInfoElementImg} = s;

type ContactInfoElementPropsType = {
    img: any
    link: string
}


export const ContactInfoElement = ({img, link}: ContactInfoElementPropsType) => {


    return (
        <div className={contactInfoElement}>
            <div className={contactInfoElementImg}><img src={img} alt={'contactElementMsg'}/></div>
            <div className={contactInfoElementLinks}>{link}</div>
        </div>
    )
}