import React from 'react';
import s from './navBar.module.css';
import {Logo} from '../../components/Logo/Logo';
import {SelectUI} from '../../components/Select/Select';
import {NavLinks} from '../../components/NavLinks/NavLinks';
import {Button} from '../../components/Button/Button';
import {BurgerMenu} from './burgerMenu/BurgerMenu';

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
    logoSelectContainer
} = s;

type NavBarPropsType = {
    setIsBurgerCollapse: (newStatus: boolean) => void
    isBurgerCollapse: boolean
}


export const NavBar = ({isBurgerCollapse, setIsBurgerCollapse}: NavBarPropsType) => {


    return (
        <div className={headerContainer}>
            <div className={headerNavBar}>
                <div className={contentContainer}>
                    <div className={logoSelectContainer}>
                    <div className={logoContainer}>
                        <Logo/>
                    </div>
                    <div className={selectContainer}>
                        <SelectUI/>
                    </div>
                    </div>
                        <div>
                    <div className={burgerMenu}>
                        <BurgerMenu setIsBurgerCollapse={setIsBurgerCollapse} isBurgerCollapse={isBurgerCollapse}/>
                    </div>
                    <div className={commonNavBar}>
                        <div className={navLinksBlock}>
                            <NavLinks navNames={['Home', 'About Us', 'Room', 'Service', 'Blog', 'Gallery']}/>
                        </div>
                        <div className={btnBlock}>
                            <Button type={'buy'}/>
                            <Button type={'login'}/>
                        </div>
                    </div>
                    </div>
                </div>
            </div>
        </div>
    );
}
