import s from './BurgerMenu.module.css'
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

  return (
    <div>
      <div onClick={setBurgerStatus} className={isBurgerCollapse ? burgerClose : burgerMenu}>
        <span></span>
      </div>
      {isBurgerCollapse && (
        <div onClick={setBurgerStatus} className={navLinksBurgerMenu}>
          <NavLinksBurger
            navNames={['Home', 'About Us', 'Room', 'Service', 'Blog', 'Gallery', 'Login', 'Basket']}
          />
        </div>
      )}
    </div>
  )
}
