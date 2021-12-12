import s from './UseFulLinksElement.module.css'
import { NavLink } from 'react-router-dom'

const { link } = s

type UsefulLinksElementPropsType = {
  linkName: string
}

export const UsefulLinksElement = ({ linkName }: UsefulLinksElementPropsType) => {
  return (
    <div className={link}>
      <NavLink className={link} to={`/${linkName.replace(/\s/g, '').toLowerCase()}`}>
        {linkName}
      </NavLink>
    </div>
  )
}
