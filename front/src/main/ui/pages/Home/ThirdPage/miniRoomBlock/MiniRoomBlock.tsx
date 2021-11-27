import s from './MiniRoomBlock.module.css';
import {MiniRoomElement} from './MiniRoomElement/MiniRoomElement';
const{miniRoom,oneRoomTop,oneRoomBottom}=s;

export const MiniRoomBlock=()=>{
    return(
        <div className={miniRoom}>
           <div className={oneRoomTop}>
               <div>
                   <MiniRoomElement roomName={'Modern'} childrenNum={2} adultNum={2} squareNum={100} price={100}/>
               </div>
               <div>
                   <MiniRoomElement roomName={'Luxe'} childrenNum={2} adultNum={2} squareNum={100} price={100}/>
               </div>
           </div>
            <div className={oneRoomBottom}>
                <div>
                    <MiniRoomElement roomName={'Premium'} childrenNum={2} adultNum={2} squareNum={100} price={100}/>
                </div>
                <div>
                    <MiniRoomElement roomName={'Single'} childrenNum={2} adultNum={2} squareNum={100} price={100}/>
                </div>
           </div>

        </div>
    )
}