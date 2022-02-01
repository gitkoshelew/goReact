import React, { useState } from 'react'
import './App.scss'
import { NavBar } from './navBar/navBar'
import { RoutesInfo } from '../Routes/RoutesInfo'
import { Footer } from './footer/footer'
import { Notification } from '../components/Notification/Notification'

function App() {
  const [isBurgerCollapse, setIsBurgerCollapse] = useState(false)

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
