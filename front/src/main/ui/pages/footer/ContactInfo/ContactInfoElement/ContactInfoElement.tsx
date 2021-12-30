import s from './ConstactInfoElement.module.scss'

const { contactInfoElement, contactInfoElementLinks, contactInfoElementImg } = s

type ContactInfoElementPropsType = {
  img: string
  link: string
}

export const ContactInfoElement = ({ img, link }: ContactInfoElementPropsType) => {
  return (
    <div className={contactInfoElement}>
      <div className={contactInfoElementImg}>
        <img src={img} alt={'contactElementMsg'} />
      </div>
      <div className={contactInfoElementLinks}>{link}</div>
    </div>
  )
}
