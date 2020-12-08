import httpClient from "@/service/httpClient";
import { server } from "@/service/constants";

export const createVideo = async values => {
  const timeElapsed = Date.now()
  const today = new Date(timeElapsed)
  var bodyForm = new FormData();
  bodyForm.append("userID", values.userID);
  bodyForm.append("lessonID", values.lessonID);
  bodyForm.append("description", values.description);
  bodyForm.append("videoFile", values.videoURL);
  bodyForm.append("uploadDate", today.toISOString().slice(0, 10));
  bodyForm.append("status", values.status);

  const result = await httpClient.post(
    server.VIDEO,
    bodyForm
  );
  if (result.data.status === "1") {
    return result.data;
  } else {
    return result.data;
  }
};
