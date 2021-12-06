import s from './AfterPanel.module.css'
import { planet } from '../../../../svgWrapper/footerSvgWrapper'

const { afterPanelBlock, imgBlock, textBlock, addressText } = s

export const AfterPanel = () => {
  return (
    <div className={afterPanelBlock}>
      <div className={addressText}>© 2021 Hotel FT. Designed by Ovatheme</div>
      <div>
        <div className={imgBlock}>
          <img src={planet} alt="planet" />
          <div className={textBlock}>English</div>
        </div>
      </div>
    </div>
  )
}
