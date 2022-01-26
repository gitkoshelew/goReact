import s from './NearbyPlaces.module.scss'
import { NearbyPlacesElement } from './NearbyPlacesElement/NearbyPlacesElement'

const { secondaryNearby, nearbyPlaces } = s

export const NearbyPlaces = () => {
  return (
    <div className={nearbyPlaces}>
      <div>
        <NearbyPlacesElement infoDistanceMess={500} infoNameMess={'Nile river'} type={'main'} />
      </div>
      <div className={secondaryNearby}>
        <NearbyPlacesElement type={'secondary'} infoNameMess={'Lakeside Asia'} infoDistanceMess={500} />
        <NearbyPlacesElement type={'secondary'} infoNameMess={'Beach resort'} infoDistanceMess={500} />
      </div>
    </div>
  )
}
