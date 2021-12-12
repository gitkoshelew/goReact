import s from './Button.module.css'
import buy from '../../../../assets/img/navBar/Buy.svg'
import { NavLink } from 'react-router-dom'

const { loginBtn, buyBtn } = s

type BtnPropsType = {
  type: string
}

export const Button = ({ type }: BtnPropsType) => {
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
    </div>
  )
}
