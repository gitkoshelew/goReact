import React from 'react'
import s from './UserChatItem.module.scss'
import { userPhotoBoilerPlate } from '../../svgWrapper/navBarSvgWrapper'
import { NavLink } from 'react-router-dom'

const { userChatItem, userPhoto, userDescriptions, userName, userRole, active } = s

type PropsType = {
  userId: number
  name: string
  sName: string
  role: string
  photo: string
}

export const UserChatItem = React.memo(({ userId, name, sName, role, photo }: PropsType) => {
  return (
    <NavLink to={`/chat/${userId}`} className={(link) => (link.isActive ? `${active} ${userChatItem}` : userChatItem)}>
      <div className={userPhoto}>
        <img src={photo !== 'PhotoURL...' ? photo : userPhotoBoilerPlate} alt="user photo" />
      </div>
      <div className={userDescriptions}>
        <div className={userName}>{`${name} ${sName}`}</div>
        <div className={userRole}>{role}</div>
      </div>
    </NavLink>
  )
})
