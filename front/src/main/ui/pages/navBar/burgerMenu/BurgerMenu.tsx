import s from './BurgerMenu.module.css';
import {NavLinks} from '../../../components/NavLinks/NavLinks';
import {NavLinksBurger} from '../../../components/NavLinks/NavLinksBurger/NavLinksBurger';

const {burgerMenu, navLinksBurgerMenu} = s

type BurgerMenuPropsType = {
    setIsBurgerCollapse: (newStatus: boolean) => void
    isBurgerCollapse: boolean
}


export const BurgerMenu = ({isBurgerCollapse, setIsBurgerCollapse}: BurgerMenuPropsType) => {

    const setBurgerStatus = () => {
        setIsBurgerCollapse(!isBurgerCollapse)
    }


    return (
        <div>
            <div onClick={setBurgerStatus} className={burgerMenu}>
                <span>
                </span>
            </div>
            {isBurgerCollapse && <div onClick={setBurgerStatus} className={navLinksBurgerMenu}>
                <NavLinksBurger navNames={['Home', 'About Us', 'Room', 'Service', 'Blog', 'Gallery']}/>
            </div>}
        </div>
    )
}