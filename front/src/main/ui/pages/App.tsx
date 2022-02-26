import React, { useEffect, useState } from 'react'
import './App.scss'
import { NavBar } from './navBar/navBar'
import { RoutesInfo } from '../Routes/RoutesInfo'
import { Footer } from './footer/footer'
import { Notification } from '../components/Notification/Notification'
import { AppRootState, useAppDispatch } from '../../bll/store/store'
import { closeNotificationChannelRequest } from '../../bll/reducers/NotificationReducer/notification-saga'
import { openNotificationChannelRequest } from '../../bll/reducers/NotificationReducer/socketChannel'
import { useSelector } from 'react-redux'

function App() {
  const [isBurgerCollapse, setIsBurgerCollapse] = useState(false)

  const clientId = useSelector((state: AppRootState) => state.LoginPage.user?.userId)
  const notificationSocketChannel = useSelector((state: AppRootState) => state.Notification.socketChannel)
  const notifications = useSelector((state: AppRootState) => state.Notification.notifications)

  const dispatch = useAppDispatch()

  useEffect(() => {
    if (clientId) {
      dispatch(openNotificationChannelRequest(clientId))
    }

    return () => {
      if (notificationSocketChannel) {
        dispatch(closeNotificationChannelRequest())
      }
    }
  }, [clientId])

  /*
                  *TODO:-routes system for faster navigation by application
                  * all links are located in /Routes folder
                  * isBurgerCollapse created for burger menu correct work

     */

  return (
    <div className={'app'}>
      <div className={'notificationContainer'}>
        {notifications.map((notification) => (
          <Notification
            key={notification.content.id}
            id={notification.content.id}
            type={notification.content.type}
            reason={notification.content.reason}
            description={notification.content.description}
          />
        ))}
      </div>
      <NavBar setIsBurgerCollapse={setIsBurgerCollapse} isBurgerCollapse={isBurgerCollapse} />
      {!isBurgerCollapse && (
        <div>
          <RoutesInfo />
          <Footer />
        </div>
      )}
    </div>
  )
}

export default App
