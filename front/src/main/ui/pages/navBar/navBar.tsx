import React from 'react'
import s from './navBar.module.scss'
import { Logo } from '../../components/Logo/Logo'
import { SelectUI } from '../../components/Select/Select'
import { NavLinks } from '../../components/NavLinks/NavLinks'
import { Button } from '../../components/Button/Button'
import { BurgerMenu } from './burgerMenu/BurgerMenu'
import { useSelector } from 'react-redux'
import { AppRootState } from '../../../bll/store/store'
import { UserNavBarView } from './UserView/userNavBarView'

const {
  headerNavBar,
  contentContainer,
  logoContainer,
  headerContainer,
  selectContainer,
  burgerMenu,
  navLinksBlock,
  btnBlock,
  commonNavBar,
  logoSelectContainer,
} = s

type NavBarPropsType = {
  setIsBurgerCollapse: (newStatus: boolean) => void
  isBurgerCollapse: boolean
}

export const NavBar = ({ isBurgerCollapse, setIsBurgerCollapse }: NavBarPropsType) => {
  const userLogin = useSelector((state: AppRootState) => state.LoginPage.user)

  const userView = userLogin ? <UserNavBarView user={userLogin} /> : <Button type={'login'} />

  return (
    <div className={headerContainer}>
      <div className={headerNavBar}>
        <div className={contentContainer}>
          <div className={logoSelectContainer}>
            <div className={logoContainer}>
              <Logo />
            </div>
            <div className={selectContainer}>
              <SelectUI />
            </div>
          </div>
          <div>
            <div className={burgerMenu}>
              <BurgerMenu setIsBurgerCollapse={setIsBurgerCollapse} isBurgerCollapse={isBurgerCollapse} />
            </div>
            <div className={commonNavBar}>
              <div className={navLinksBlock}>
                <NavLinks navNames={['Home', 'About Us', 'Room', 'Service', 'Booking', 'Gallery', 'Chat']} />
              </div>
              <div className={btnBlock}>
                <Button type={'buy'} />
                {userView}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
