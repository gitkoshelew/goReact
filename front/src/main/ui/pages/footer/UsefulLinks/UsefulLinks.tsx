import s from './UseFulLinks.module.css'
import { UsefulLinksElement } from './UseFulLinksElement/UsefulLinksElement'

const { usefulLinksBlock, linksTitle, linksList, groupOfLinks } = s

export const UsefulLinks = () => {
  return (
    <div className={usefulLinksBlock}>
      <div className={linksTitle}>USEFUL LINKS</div>
      <div className={linksList}>
        <div className={groupOfLinks}>
          <UsefulLinksElement linkName={'About Us'} />
          <UsefulLinksElement linkName={'Rooms'} />
          <UsefulLinksElement linkName={'Service'} />
        </div>
        <div>
          <UsefulLinksElement linkName={'Home'} />
          <UsefulLinksElement linkName={'Gallery'} />
          <UsefulLinksElement linkName={'Booking'} />
        </div>
      </div>
    </div>
  )
}
