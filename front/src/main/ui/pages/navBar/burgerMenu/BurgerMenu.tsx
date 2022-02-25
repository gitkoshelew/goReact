import s from './BurgerMenu.module.scss'
import { NavLinksBurger } from '../../../components/NavLinks/NavLinksBurger/NavLinksBurger'

const { burgerMenu, navLinksBurgerMenu, burgerClose } = s

type BurgerMenuPropsType = {
  setIsBurgerCollapse: (newStatus: boolean) => void
  isBurgerCollapse: boolean
}

export const BurgerMenu = ({ isBurgerCollapse, setIsBurgerCollapse }: BurgerMenuPropsType) => {
  const setBurgerStatus = () => {
    setIsBurgerCollapse(!isBurgerCollapse)
  }
  const online = navigator.onLine;

  return (
    <div>
      <div data-testid="test" onClick={setBurgerStatus} className={isBurgerCollapse ? burgerClose : burgerMenu}>
        <span> </span>
      </div>
      {isBurgerCollapse && (
        <div onClick={setBurgerStatus} className={navLinksBurgerMenu}>
          <NavLinksBurger
            navNames={online? ['Home', 'About Us', 'Room', 'Service', 'Booking', 'Gallery', 'Login', 'Basket', 'Chat'] : ['Home', 'About Us', 'Login', 'Basket']}
          />
        </div>
      )}
    </div>
  )
}
