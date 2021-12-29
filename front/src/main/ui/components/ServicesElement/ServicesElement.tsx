import s from './ServicesElement.module.scss'
import { Case, flower, forkAndSpoon, swimmer } from '../../svgWrapper/HomeSvgWrapper'
import { IsActiveServiceElementType } from '../../pages/Home/FourthPage/ServiceBlock/ServiceBlock'

type ServicesType = 'forkAndSpoon' | 'flower' | 'swimmer' | 'case'

type ServicesElementPropsType = {
  mainTextMess: string
  secondaryTextMess: string
  isActive: IsActiveServiceElementType
  type: ServicesType
  setIsActive: (newIsActiveObj: IsActiveServiceElementType) => void
}

const { oneElement, mainText, secondaryText, textBlock, imgBlock, oneElementActive } = s

export const ServicesElement = ({
  secondaryTextMess,
  mainTextMess,
  type,
  isActive,
  setIsActive,
}: ServicesElementPropsType) => {
  const onActiveClassHandler = (propertyName: string) => {
    const newIsActive: IsActiveServiceElementType = { ...isActive }
    for (const key in newIsActive) {
      key === propertyName ? (newIsActive[key] = true) : (newIsActive[key] = false)
    }
    return setIsActive(newIsActive)
  }

  return (
    <>
      {type === 'forkAndSpoon' && (
        <div onClick={() => onActiveClassHandler('elem1')} className={isActive.elem1 ? oneElementActive : oneElement}>
          <div className={imgBlock}>
            <img src={forkAndSpoon} alt="forkAndSpoon" />
          </div>
          <div className={textBlock}>
            <div className={mainText}>{mainTextMess}</div>
            <div className={secondaryText}>{secondaryTextMess}</div>
          </div>
        </div>
      )}
      {type === 'swimmer' && (
        <div onClick={() => onActiveClassHandler('elem4')} className={isActive.elem4 ? oneElementActive : oneElement}>
          <div className={imgBlock}>
            <img src={swimmer} alt="swimmer" />
          </div>
          <div className={textBlock}>
            <div className={mainText}>{mainTextMess}</div>
            <div className={secondaryText}>{secondaryTextMess}</div>
          </div>
        </div>
      )}
      {type === 'flower' && (
        <div onClick={() => onActiveClassHandler('elem3')} className={isActive.elem3 ? oneElementActive : oneElement}>
          <div className={imgBlock}>
            <img src={flower} alt="flower" />
          </div>
          <div className={textBlock}>
            <div className={mainText}>{mainTextMess}</div>
            <div className={secondaryText}>{secondaryTextMess}</div>
          </div>
        </div>
      )}
      {type === 'case' && (
        <div onClick={() => onActiveClassHandler('elem2')} className={isActive.elem2 ? oneElementActive : oneElement}>
          <div className={imgBlock}>
            <img src={Case} alt="case" />
          </div>
          <div className={textBlock}>
            <div className={mainText}>{mainTextMess}</div>
            <div className={secondaryText}>{secondaryTextMess}</div>
          </div>
        </div>
      )}
    </>
  )
}
