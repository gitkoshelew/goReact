import React, { useEffect, useRef, useState } from 'react'
import './App.scss'
import { NavBar } from './navBar/navBar'
import { RoutesInfo } from '../Routes/RoutesInfo'
import { Footer } from './footer/footer'
import { Notification } from '../components/Notification/Notification'
import { io, Socket } from 'socket.io-client'
import { useAppDispatch } from '../../bll/store/store'
import { showNotificationRequest } from '../../bll/reducers/NotificationReducer/notification-saga'

function App() {
  const [isBurgerCollapse, setIsBurgerCollapse] = useState(false)

  const socket = useRef<Socket | null>(null)
  const dispatch = useAppDispatch()

  useEffect(() => {
    socket.current = io('ws://localhost:5001')

    return () => {
      socket.current?.disconnect()
    }
  }, [])

  useEffect(() => {
    socket.current?.on('BROKER_RECEIVED_NOTIFICATION', (notification: string) => {
      dispatch(showNotificationRequest(JSON.parse(notification)))
    })
  }, [])

  /*
                  *TODO:-routes system for faster navigation by application
                  * all links are located in /Routes folder
                  * isBurgerCollapse created for burger menu correct work

     */

  return (
    <div className={'app'}>
      <Notification />
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
