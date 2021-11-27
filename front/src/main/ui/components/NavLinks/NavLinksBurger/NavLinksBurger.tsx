import s from './NavLinksBurger.module.css';
import {NavLink} from 'react-router-dom';


const {navLinkTitle, oneLink, oneLinkActive} = s;

type NavLinksPropsType = {
    navNames: string[]
}


export const NavLinksBurger = (props: NavLinksPropsType) => {

    const {navNames} = props;


    const correctNavLinks = navNames.map(t => <NavLink className={({isActive}) =>
        isActive ? oneLinkActive : oneLink} to={`/${t.replace(/\s/g, '').toLowerCase()}`}>{t}</NavLink>)


    return (
        <div className={navLinkTitle}>
            {correctNavLinks}
        </div>
    )
}