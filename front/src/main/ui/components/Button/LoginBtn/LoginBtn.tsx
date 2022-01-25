import s from './LoginBtn.module.scss'
import { NavLink } from 'react-router-dom'

const { loginBtnTitle } = s

export const LoginButton = () => {
  return (
    <NavLink className={loginBtnTitle} to={'/login'}>
      <div>login</div>
    </NavLink>
  )
}
