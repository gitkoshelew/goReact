/* eslint-disable */
import s from "./ImgUpload.module.css";
import { FileUploader } from "react-drag-drop-files";
import { useState } from "react";
import { ProgressType } from "../BookingRegForm";

const {uploadTablet,customTablet} = s;
const fileTypes = ["JPG", "PNG", "GIF"];

type ImgUploadPropsType = {
  setProgress: (newStatus: ProgressType) => void
  progress: ProgressType
  setPhotoUrl: (newPhoto: any) => void
}


export const ImgUpload = ({ progress, setProgress, setPhotoUrl }: ImgUploadPropsType) => {

  const [file, setFile] = useState<null | Blob>(null);


  const handleChange = (file: any) => {
    setFile(file);
    setProgress("uploaded");
    const correctFile = Object.assign(file, {
      preview: URL.createObjectURL(file)
    })
    setPhotoUrl(correctFile.preview)
  };


  return (
    <div className={uploadTablet}>
      {progress === "getUpload" && <FileUploader classes={customTablet} handleChange={handleChange} name="file" types={fileTypes} />}
    </div>
  );
};