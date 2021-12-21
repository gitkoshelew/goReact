import React, { useEffect, useState } from "react";
import s from "./Booking.module.css";
import { TitlePageTextBlock } from "../../components/TitlePageTextBlock/TitlePageTextBlock";
import { BookingRegForm } from "./BookingRegForm/BookingRegForm";
import { BookingCalendar } from "./BookingCalendar/BookingCalendar";
import { BookingRoom } from "./BookingRoom/BookingRoom";
import { useSelector } from "react-redux";
import { AppRootStateType } from "../../../bll/store/store";
import { IsRentType, OrderedRoomsType } from "../../../bll/reducers/BookingRoomsPickReducer/BookingRoomPick-reducer";
import { Button } from "../../components/Button/Button";
import { ProgressType } from "../../../bll/reducers/BookingRegFormReducer/BookingRegForm-reducer";
import { SelectedToOrderRoom } from "./SelectedToOrderRom/SelectedToOrderRoom";
import { useFormik, FormikErrors } from "formik";


const { bookingPage, bookingForm, bookingProcess, bookingCalendar, uploadOrderedRoomsBlock } = s;


type FormValues = {
  firstName: string
  lastName: string
  email: string
}


export const Booking = () => {

  const validate = (values: FormValues) => {
    const errors: FormikErrors<FormValues> = {};
    if (!values.firstName) {
      errors.firstName = "Required field";
    } else if (values.firstName.length > 15) {
      errors.firstName = "Must be 20 characters or less";
    }
    if (!values.lastName) {
      errors.lastName = "Required field";
    } else if (values.lastName.length > 15) {
      errors.lastName = "Must be 20 characters or less";
    }

    if (!values.email) {
      errors.email = "Required field";
    } else if (!/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(values.email)) {
      errors.email = "Invalid email address";
    }
    return errors;
  };


  const formik = useFormik({
    initialValues: {
      firstName: "",
      lastName: "",
      email: ""
    },
    validate,
    onSubmit: values => {
      alert(JSON.stringify(values, null, 2));
    }
  });


  const progress = useSelector<AppRootStateType, ProgressType>((state) => state.BookingRegForm.progress);
  const actualDay = useSelector<AppRootStateType, string | Date>((state) => state.BookingRoomPick.actualDay);
  const isRentArr = useSelector<AppRootStateType, IsRentType[]>((state) => state.BookingRoomPick.isRent);
  const orderedRoomBasket = useSelector<AppRootStateType, OrderedRoomsType[]>(
    (state) => state.BookingRoomPick.orderedRoomBasket
  );

  const [roomIndicate, setRoomIndicate] = useState<null | IsRentType>(null);

  const isActiveBtn = progress === "uploaded" && orderedRoomBasket.length !== 0;

  useEffect(() => {
    const newActualDay = isRentArr.find((t) => t.id === actualDay);
    newActualDay ? setRoomIndicate(newActualDay) : setRoomIndicate(null);
  }, [actualDay, isRentArr]);


  return (
    <form onSubmit={formik.handleSubmit}>
      <div className={bookingPage}>
        <TitlePageTextBlock mainTextMess={"Book room for pet"} isWithLink={false} />
        <div className={bookingProcess}>
          <div className={bookingForm}>
            <BookingRegForm formik={formik} />
          </div>
          <div className={bookingCalendar}>
            <BookingCalendar />
            {roomIndicate && (
              <BookingRoom
                dayId={roomIndicate.id}
                firstRoom={roomIndicate.firstRoom}
                secondRoom={roomIndicate.secondRoom}
              />
            )}
          </div>
        </div>
        <div className={uploadOrderedRoomsBlock}>
          {orderedRoomBasket.length != 0 && <SelectedToOrderRoom orderedRoomBasket={orderedRoomBasket} />}
          <Button type={"upload"} isActive={isActiveBtn} />
        </div>
      </div>
    </form>
  );
};
