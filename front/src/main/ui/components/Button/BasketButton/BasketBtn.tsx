import s from './BasketBtn.module.scss'
import buy from '../../../../../assets/img/navBar/Buy.svg'

const { buyBtnTitle } = s

export const BasketButton = () => {
  return (
    <div className={buyBtnTitle}>
      <img src={buy} alt="buyContainer" />
    </div>
  )
}
