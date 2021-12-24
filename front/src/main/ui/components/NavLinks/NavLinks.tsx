import s from './NavLinks.module.scss'
import { NavLink } from 'react-router-dom'

const { navLinkTitle, oneLink, oneLinkActive } = s

type NavLinksPropsType = {
  navNames: string[]
}

export const NavLinks = (props: NavLinksPropsType) => {
  const { navNames } = props

  const correctNavLinks = navNames.map((t, i) => (
    <NavLink
      key={i}
      className={({ isActive }) => (isActive ? oneLinkActive : oneLink)}
      to={`/${t.replace(/\s/g, '').toLowerCase()}`}
    >
      {t}
    </NavLink>
  ))

  return <div className={navLinkTitle}>{correctNavLinks}</div>
}
