/* eslint-disable */
import React, { useState } from "react";
import s from "./BookingRegForm.module.css";
import { ImgUpload } from "./ImgUploadComponent/ImgUpload";
import { Button } from "../../../components/Button/Button";

const { bookingForm, clientDescription, inputInfo, uploadedImg } = s;

export type ProgressType = "getUpload" | "uploading" | "uploaded" | "uploadError"


export const BookingRegForm = () => {

  const [progress, setProgress] = useState<ProgressType>("getUpload");
  const [photoUrl, setPhotoUrl] = useState<any>(undefined);
  const [errorMSG, setErrorMSG] = useState<string | null>(null);
  const [isActiveBtn, setIsActiveBtn] = useState<boolean>(true);


  const actualContent = () => {
    switch (progress) {
      case ("getUpload"):
        return <>
          Upload your pet photo:
          <ImgUpload setPhotoUrl={setPhotoUrl} setProgress={setProgress} progress={progress} />
        </>;
      case("uploading"):
        return <div>uploading...</div>;
      case ("uploaded"):
        return <div className={uploadedImg}>
          <img src={photoUrl} alt="uploadedImg" />
        </div>;
      case("uploadError"):
        return <>
          <div>{errorMSG}</div>
          <div>Upload an image please</div>
        </>;

    }
  };


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
      <div className={inputInfo}>
        {actualContent()}
      </div>
      <Button type={"Upload"} isActive={isActiveBtn} />
    </div>
  </div>;
};
