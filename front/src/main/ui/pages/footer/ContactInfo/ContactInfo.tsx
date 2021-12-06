import s from './ConstactInfo.module.css'
import { ft } from '../../../svgWrapper/footerSvgWrapper'
import { home } from '../../../svgWrapper/footerSvgWrapper'
import { mobile } from '../../../svgWrapper/footerSvgWrapper'
import { gmail } from '../../../svgWrapper/footerSvgWrapper'
import { ContactInfoElement } from './ContactInfoElement/ContactInfoElement'

const { contactInfo, linkInMap, ftLogo } = s

export const ContactInfo = () => {
  return (
    <div className={contactInfo}>
      <div className={ftLogo}>
        <img src={ft} alt="ftLogo" />
      </div>
      <ContactInfoElement img={home} link={'Hotel.np002@gmai.com'} />
      <ContactInfoElement img={mobile} link={'+84 0934 425 031'} />
      <ContactInfoElement img={gmail} link={'497 Evergeen Rd. Roseville, CA 98823'} />
      <div className={linkInMap}>Check map</div>
    </div>
  )
}
