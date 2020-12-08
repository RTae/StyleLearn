import httpClient from "@/service/httpClient";
import { server } from "@/service/constants";

export const getCourseFromProgressLesson = uid => {
  return httpClient.get(server.PROGRESS_LESSON + `?uid=${uid}&mode=1`);
};

export const getProgressLesson = (uid, cid) => {
  return httpClient.get(
    server.PROGRESS_LESSON + `?uid=${uid}&lid=${cid}&mode=2`
  );
};

export const createProgressLesson = async values => {
  var bodyForm = new FormData();
  bodyForm.append("uid", values.userID);
  bodyForm.append("lid", values.lessonID);
  bodyForm.append("quantityDay", values.quantityDay);
  const result = await httpClient.post(
    server.PROGRESS_LESSON,
    bodyForm
  );
  if (result.data.status === "1") {
    return result.data;
  } else {
    return result.data;
  }
};
