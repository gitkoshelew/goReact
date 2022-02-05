import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import styles from './Room.module.scss'
import { useDispatch, useSelector } from "react-redux";
import { AppRootState } from "../../../bll/store/store";
import { ChangeEvent, useEffect, useState } from "react";
import { fetchRoomRequest } from "../../../bll/reducers/RoomPageReducer/roomPage-saga";
import { Pagination } from "@mui/material";

const { roomPage } = styles

export const Room = () => {
    const [page, setPage] = useState(1)
    const dispatch = useDispatch();
    const rooms = useSelector((state: AppRootState) => state.RoomPage.rooms)
    const totalRoomsCount = useSelector((state: AppRootState) => state.RoomPage.totalRoomsCount)
    const pageSize = useSelector((state: AppRootState) => state.RoomPage.pageSize)
    const currentPage = useSelector((state: AppRootState) => state.RoomPage.currentPage)

    useEffect(() => {
        dispatch(fetchRoomRequest(currentPage, pageSize))
    }, [])

    const onPageChange = (e: ChangeEvent<unknown>, pageChange: number) => {
        setPage(pageChange)
    }
    const fetchRoomRequestHandler = () => {
        dispatch(fetchRoomRequest(page, pageSize))

    }

    return (
        <>
            <Pagination count={totalRoomsCount} size="small" onChange={onPageChange} onClick={fetchRoomRequestHandler}
                        showFirstButton
                        showLastButton
                        sx={{ marginY: 3, marginX: "auto" }}/>
            {
                rooms.length === 0
                    ? (<span
                        className={styles.textNameWithoutCards}>Sorry, but there are no available rooms at the moment</span>)
                    : (rooms.map((room) => {
                                const { roomNum, hotelId, roomPhotoUrl, petType } = room
                                return (
                                    <div key={room.roomId} className={roomPage}>
                                        <Card sx={{ minWidth: 275, minHeight: 250, maxWidth: 350, maxHeight: 300, padding: 2 }}>
                                            <CardContent className='Wrapper'>
                                                <Typography sx={{ fontSize: 24 }} color="text.secondary" gutterBottom>
                                                    The room number: {roomNum}
                                                </Typography>
                                                <Typography sx={{ fontSize: 24 }} color="text.secondary" gutterBottom>
                                                    The hotel number: {hotelId}
                                                </Typography>
                                                <Typography sx={{ fontSize: 24 }} color="text.secondary" gutterBottom>
                                                    Photo: {roomPhotoUrl}
                                                </Typography>
                                                <Typography sx={{ fontSize: 24 }} color="text.secondary" gutterBottom>
                                                    Pet type: {petType}
                                                </Typography>
                                            </CardContent>
                                            <CardActions style={{ justifyContent: 'center' }}>
                                                <Button variant="contained" size="small" color={"success"}
                                                        onClick={() => {
                                                            alert('Booking!!')
                                                        }}>Booking</Button>
                                            </CardActions>
                                        </Card>
                                    </div>

                                );

                            }
                        )
                    )
            }
        </>
    )
}
