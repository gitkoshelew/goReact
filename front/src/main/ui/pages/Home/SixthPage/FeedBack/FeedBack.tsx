import s from './FeedBack.module.css';


const {feedBackText, feedBack, feedBackUser, userFeedbackName, userFeedBackNationality, infoUserBlock, userPhoto} = s;


type FeedBackPropsType = {
    photo: string
    feedBackMess: string
    userName: string
    nationality: string
}


export const FeedBack = ({photo, feedBackMess, userName, nationality}: FeedBackPropsType) => {
    return (
        <div className={feedBack}>
            <div className={feedBackText}>
                {`"${feedBackMess}"`}
            </div>
            <div className={feedBackUser}>
                <div className={userPhoto}>
                    <img src={photo} alt="userPhoto"/>
                </div>
                <div className={infoUserBlock}>
                    <div className={userFeedbackName}>{userName}</div>
                    <div className={userFeedBackNationality}>{nationality}</div>
                </div>
            </div>

        </div>
    )
}