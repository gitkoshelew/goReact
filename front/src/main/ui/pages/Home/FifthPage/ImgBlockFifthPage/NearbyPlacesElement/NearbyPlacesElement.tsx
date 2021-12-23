import s from './NearbyPlacesElement.module.css'

type NearbyPlacesBlockType = 'main' | 'secondary'

type NearbyPlacesElementPropsType = {
  type: NearbyPlacesBlockType
  infoNameMess: string
  infoDistanceMess: number
}

const { oneElement, mainBlock, secondaryBlock, informBlock, infoName, infoDistance, informBlockSecondary } = s

export const NearbyPlacesElement = ({ type, infoNameMess, infoDistanceMess }: NearbyPlacesElementPropsType) => {
  return (
    <div className={oneElement}>
      {type === 'main' && (
        <div>
          <div className={mainBlock}>400x300</div>
          <div className={informBlock}>
            <div className={infoName}>{infoNameMess}</div>
            <div className={infoDistance}>{`${infoDistanceMess} m`}</div>
          </div>
        </div>
      )}
      {type === 'secondary' && (
        <div>
          <div className={secondaryBlock}>300x300</div>
          <div className={informBlockSecondary}>
            <div className={infoName}>{infoNameMess}</div>
            <div className={infoDistance}>{`${infoDistanceMess} m`}</div>
          </div>
        </div>
      )}
    </div>
  )
}
