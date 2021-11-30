import s from './NearbyPlaces.module.css';
import {NearbyPlacesElement} from './NearbyPlacesElement/NearbyPlacesElement';


const{}=s;

export const NearbyPlaces =()=>{
    return(
        <div>
            <NearbyPlacesElement infoDistanceMess={500} infoNameMess={'Nile river'} type={'main'}/>
            <NearbyPlacesElement type={'secondary'} infoNameMess={'Lakeside Asia'} infoDistanceMess={500}/>
            <NearbyPlacesElement type={'secondary'} infoNameMess={'Beach resort'} infoDistanceMess={500}/>
        </div>
    )
}