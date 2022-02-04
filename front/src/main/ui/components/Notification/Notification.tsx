import React from 'react'
import s from './Notification.module.scss'
import { Alert, AlertTitle } from '@mui/material'
import { useSelector } from 'react-redux'
import { AppRootState } from '../../../bll/store/store'

const { notification, hideNotification } = s

export const Notification = () => {
  const isNotificationOpened = useSelector((state: AppRootState) => state.Notification.isOpened)
  const notificationData = useSelector((state: AppRootState) => state.Notification.data)

  return (
    <div className={`${notification} ${!isNotificationOpened ? hideNotification : ''}`}>
      <Alert variant="filled" severity={notificationData.type}>
        <AlertTitle>{notificationData.reason}</AlertTitle>
        {notificationData.description}
      </Alert>
    </div>
  )
}
