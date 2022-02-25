import React from 'react'
import s from './navBar.module.scss'
import { Logo } from '../../components/Logo/Logo'
import { SelectUI } from '../../components/Select/Select'
import { NavLinks } from '../../components/NavLinks/NavLinks'
import { BurgerMenu } from './burgerMenu/BurgerMenu'
import { useSelector } from 'react-redux'
import { AppRootState } from '../../../bll/store/store'
import { UserNavBarView } from './UserView/userNavBarView'
import { LoginButton } from '../../components/Button/LoginBtn/LoginBtn'
import { BasketButton } from '../../components/Button/BasketButton/BasketBtn'

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
  offlineMode,
  disableMode,
} = s

type NavBarPropsType = {
  setIsBurgerCollapse: (newStatus: boolean) => void
  isBurgerCollapse: boolean
}

export const NavBar = ({ isBurgerCollapse, setIsBurgerCollapse }: NavBarPropsType) => {
  const userLogin = useSelector((state: AppRootState) => state.LoginPage.user)

  const userView = userLogin ? <UserNavBarView user={userLogin} /> : <LoginButton />
  const online = navigator.onLine;
  const offlineMessage = `You're offline`

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
          <div className={online? disableMode : offlineMode}>
            {offlineMessage}
          </div>
          <div>
            <div className={burgerMenu}>
              <BurgerMenu setIsBurgerCollapse={setIsBurgerCollapse} isBurgerCollapse={isBurgerCollapse} />
            </div>
            <div className={commonNavBar}>
              <div className={navLinksBlock}>
                <NavLinks navNames={online ? ['Home', 'About Us', 'Room', 'Service', 'Booking', 'Gallery', 'Chat'] :['Home', 'About Us'] } />
              </div>
              <div className={btnBlock}>
                <BasketButton />
                {userView}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
