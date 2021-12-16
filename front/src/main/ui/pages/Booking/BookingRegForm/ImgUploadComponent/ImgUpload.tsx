/* eslint-disable */
import s from "./ImgUpload.module.css";
import { FileUploader } from "react-drag-drop-files";
import { changePhotoUrl, changeProgressStatus, ProgressType } from "../../../../../bll/reducers/BookingRegForm-reducer";
import { useAppDispatch } from "../../../../../bll/store/store";


const { uploadTablet, customTablet } = s;
const fileTypes = ["JPG", "PNG", "GIF"];

type ImgUploadPropsType = {
  progress: ProgressType
}


export const ImgUpload = ({ progress }: ImgUploadPropsType) => {

  const dispatch = useAppDispatch();

  const handleChange = (file:any) => {
    dispatch(changeProgressStatus({ newStatus: "uploaded" }));
    const correctFile = Object.assign(file, {
      preview: URL.createObjectURL(file)
    });
    dispatch(changePhotoUrl({newPhotoUrl: correctFile.preview }));
  };


  return (
    <div className={uploadTablet}>
      {progress === "getUpload" &&
      <FileUploader classes={customTablet} handleChange={handleChange} name="file" types={fileTypes} />}
    </div>
  );
};