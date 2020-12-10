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

export const getAllVideoByLessonID = async LessonID => {
  return httpClient.get(server.VIDEO + `?lessonID=${LessonID}&mode=2`);
}

export const getAllVideoByUserID = async userID => {
  return httpClient.get(server.VIDEO + `?userID=${userID}&mode=3`);
}

export const getShowVideoByVideoID = async videoID => {
  return httpClient.get(server.VIDEO + `?videoID=${videoID}&mode=4`);
}
