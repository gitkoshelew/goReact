import React from "react";
import s from "./Booking.module.css";
import { TitlePageTextBlock } from "../../components/TitlePageTextBlock/TitlePageTextBlock";
import { BookingRegForm } from "./BookingRegForm/BookingRegForm";


const { bookingPage, bookingForm } = s;

export const Booking = () => {
  return <div className={bookingPage}>
      <TitlePageTextBlock mainTextMess={"Book room for pet"} isWithLink={false} />
    <div className={bookingForm}>
      <BookingRegForm />
    </div>
  </div>;
};
