import s from './ImgLinks.module.css';
import { facebook } from '../../../../svgWrapper/footerSvgWrapper'
import { twitter } from '../../../../svgWrapper/footerSvgWrapper'
import { instagram } from '../../../../svgWrapper/footerSvgWrapper'
import { youtube } from '../../../../svgWrapper/footerSvgWrapper'

const { imgLinksBlock, imgLinkElement } = s;


export const ImgLinks = () => {
    return (
        <div className={imgLinksBlock}>

            <div className={imgLinkElement}>
                <img src={facebook} alt="facebook"/>
            </div>
            <div className={imgLinkElement}>
                <img src={twitter} alt="twitter"/>
            </div>
            <div className={imgLinkElement}>
                <img src={instagram} alt="instagram"/>
            </div>
            <div className={imgLinkElement}>
                <img src={youtube} alt="youtube"/>
            </div>

        </div>
    )
}