import s from './Button.module.css'
import buy from '../../../../assets/img/navBar/Buy.svg'
import { NavLink } from 'react-router-dom'

const { loginBtn, buyBtn,uploadPetPhoto,uploadPetPhotoActive } = s

type BtnPropsType = {
  type: string
  isActive?:boolean
}

export const Button = ({ type,isActive }: BtnPropsType) => {
  return (
    <div>
      {type === 'login' && (
        <NavLink className={loginBtn} to={'/login'}>
          <div>Login</div>
        </NavLink>
      )}
      {type === 'buy' && (
        <div className={buyBtn}>
          <img src={buy} alt="buyContainer" />
        </div>
      )}
      {type === 'Upload' && (
          <button disabled={isActive} className={ isActive? uploadPetPhotoActive:uploadPetPhoto}>Upload</button>
      )}
    </div>
  )
}
