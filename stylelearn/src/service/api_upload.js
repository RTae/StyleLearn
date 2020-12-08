import httpClient from "@/service/httpClient";
import { server } from "@/service/constants";

export const uploadImage = async values => {
  var bodyFormData = new FormData();
  bodyFormData.append("file", values.newImage);
  bodyFormData.append("upload_preset", "image_upload");
  const result = await httpClient.post(server.UPLOAD_IMAGE, bodyFormData);
  return {
    data: result.data,
    status: result.status
  }
}

export const uploadVideo = async values => {
  var bodyFormData = new FormData();
  bodyFormData.append("file", values.videoFile);
  bodyFormData.append("upload_preset", "video_upload");
  const result = await httpClient.post(server.UPLOAD_VIDEO, bodyFormData);
  return {
    data: result.data,
    status: result.status
  }
}
