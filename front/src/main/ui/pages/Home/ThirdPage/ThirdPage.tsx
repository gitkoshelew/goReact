import s from './ThirdPage.module.css'
import { TitlePageTextBlock } from '../../../components/TitlePageTextBlock/TitlePageTextBlock'
import { BigRoomBlock } from './BigRoomBlock/BigRoomBlock'
import { MiniRoomBlock } from './miniRoomBlock/MiniRoomBlock'

const { thirdPage, photoBlock } = s

export const ThirdPage = () => {
  return (
    <div className={thirdPage}>
      <TitlePageTextBlock isWithLink={false} mainTextMess={'OUR FAVORITE ROOMS'} />
      <div className={photoBlock}>
        <BigRoomBlock
          roomName={'Luxure'}
          price={100}
          adultNum={2}
          childrenNum={2}
          squareNum={100}
        />
        <MiniRoomBlock />
      </div>
    </div>
  )
}
