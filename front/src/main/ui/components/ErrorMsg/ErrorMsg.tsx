import error from './error.gif';
import s from './ErrorMsg.module.scss';

const {ErrorImg} = s;


export const ErrorMsg = () => {
    return <img className={ErrorImg} src={error} alt={'errorMsg'}/>
}