import React from 'react'
import s from './Notification.module.scss'
import { Alert, AlertTitle } from '@mui/material'
import { useAppDispatch } from '../../../bll/store/store'
import {
  confirmNotificationRequest,
  removeNotificationRequest,
} from '../../../bll/reducers/NotificationReducer/notification-saga'
import { NotificationData } from '../../../bll/reducers/NotificationReducer/notification-reducer'

const { notification } = s

export const Notification = ({ id, type, reason, description }: Omit<NotificationData, 'toUser'>) => {
  const dispatch = useAppDispatch()

  const handleClose = () => {
    dispatch(confirmNotificationRequest(id))
    dispatch(removeNotificationRequest(id))
  }

  return (
    <div className={`${notification}`}>
      <Alert variant="filled" severity={type} onClose={handleClose}>
        <AlertTitle>{reason}</AlertTitle>
        {description}
      </Alert>
    </div>
  )
}
