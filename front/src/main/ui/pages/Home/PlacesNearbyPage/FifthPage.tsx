import s from './FifthPage.module.scss'
import { TitlePageTextBlock } from '../../../components/TitlePageTextBlock/TitlePageTextBlock'
import { NearbyPlaces } from './ImgBlockFifthPage/NearbyPlaces'

const { fifthPage, titleText } = s

export const FifthPage = () => {
  return (
    <div className={fifthPage}>
      <div className={titleText}>
        <TitlePageTextBlock isWithLink={true} linkTextMess={'See more'} mainTextMess={'Places nearby'} />
      </div>
      <NearbyPlaces />
    </div>
  )
}
