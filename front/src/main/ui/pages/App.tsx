import React, { useState } from 'react'
import './App.css'
import { NavBar } from './navBar/navBar'
import { RoutesInfo } from '../Routes/RoutesInfo'
import { Footer } from './footer/footer'

function App() {
  const [isBurgerCollapse, setIsBurgerCollapse] = useState(false)

  /*
                  *TODO:-routes system for faster navigation by application
                  * all links are located in /Routes folder
                  * isBurgerCollapse created for burger menu correct work

     */

  const a = 10

  return (
    <div className={'app'}>
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
