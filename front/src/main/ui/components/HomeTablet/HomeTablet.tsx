import s from './HomeTablet.module.scss'
import { HomeTabletElement } from './HomeTabletElement/HomeTabletElement'

const { tabletTitle } = s

export const HomeTablet = () => {
  return (
    <div className={tabletTitle}>
      <HomeTabletElement type={'checkIn'} />
      <HomeTabletElement type={'checkOut'} />
      <HomeTabletElement type={'calendar'} />
      <HomeTabletElement type={'btnAvailability'} />
    </div>
  )
}
