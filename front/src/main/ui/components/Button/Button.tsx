import s from './Button.module.css'
import buy from '../../../../assets/img/navBar/Buy.svg'
import { NavLink } from 'react-router-dom'

const { loginBtn, buyBtn, uploadPetPhoto, uploadPetPhotoActive, orderBtn } = s

type BtnPropsType = {
  type: string
  isActive?: boolean
  onClick?: () => void
}

export const Button = ({ onClick, type, isActive }: BtnPropsType) => {
  return (
    <div>
      {type === 'login' && (
        <NavLink className={loginBtn} to={'/login'}>
          <div>login</div>
        </NavLink>
      )}
      {type === 'buy' && (
        <div className={buyBtn}>
          <img src={buy} alt="buyContainer" />
        </div>
      )}
      {type === 'order' && (
        <button type={'button'} onClick={onClick} className={orderBtn}>
          order
        </button>
      )}
      {type === 'upload' && (
        <button type={'submit'} disabled={!isActive} className={isActive ? uploadPetPhotoActive : uploadPetPhoto}>
          Confirm
        </button>
      )}
    </div>
  )
}
