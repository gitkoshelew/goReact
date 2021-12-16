import React from "react";
import s from "./BookingRegForm.module.css";


const { bookingForm, clientDescription,inputInfo } = s;

export const BookingRegForm = () => {
  return <div className={bookingForm}>
    <div className={clientDescription}>
      <div className={inputInfo}>
        Your name:
        <input type="text" />
      </div>
      <div className={inputInfo}>
        Your email:
        <input type="text" />
      </div>
    </div>
  </div>;
};
