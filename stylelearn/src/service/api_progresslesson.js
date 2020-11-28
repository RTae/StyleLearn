import httpClient from "@/service/httpClient";
import { server } from "@/service/constants";

export const getCourseFromProgressLesson = (uid) => {
  return httpClient.get(server.PROGRESS_LESSON + `?uid=${uid}&mode=1`);
};

export const getProgressLesson = (uid, cid) => {
  return httpClient.get(server.PROGRESS_LESSON + `?uid=${uid}&cid=${cid}&mode=2`);
};
