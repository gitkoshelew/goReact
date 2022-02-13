import React from 'react'
import s from './Notification.module.scss'
import { Alert, AlertTitle } from '@mui/material'
import { useSelector } from 'react-redux'
import { AppRootState } from '../../../bll/store/store'

const { notification, hideNotification } = s

export const Notification = () => {
  const isNotificationOpened = useSelector((state: AppRootState) => state.Notification.isOpened)
  const currentNotification = useSelector((state: AppRootState) => state.Notification.currentNotification)

  return (
    <div className={`${notification} ${!isNotificationOpened ? hideNotification : ''}`}>
      {currentNotification && (
        <Alert variant="filled" severity={currentNotification.type}>
          <AlertTitle>{currentNotification.reason}</AlertTitle>
          {currentNotification.description}
        </Alert>
      )}
    </div>
  )
}
