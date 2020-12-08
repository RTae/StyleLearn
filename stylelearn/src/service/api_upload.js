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
