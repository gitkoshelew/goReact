import React, {useState} from 'react';
import './App.css';
import {NavBar} from './navBar/navBar';
import {RoutesInfo} from '../Routes/RoutesInfo';


function App() {

    const [isBurgerCollapse, setIsBurgerCollapse] = useState(false);


/*
              *TODO:-routes system for faster navigation by application
              * all links are located in /Routes folder
              * isBurgerCollapse created for burger menu correct work

 */

  return (
      <div>
          <NavBar setIsBurgerCollapse={setIsBurgerCollapse} isBurgerCollapse={isBurgerCollapse}/>
          {!isBurgerCollapse && <div>
              <RoutesInfo/>
          </div>}
      </div>
  );
}

export default App;
