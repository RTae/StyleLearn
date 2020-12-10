import httpClient from "@/service/httpClient";
import { server } from "@/service/constants";

export const getAllLeson = () => {
  return httpClient.get(server.LESSON);
};

export const getLessonByCourse = (id) => {
  return httpClient.get(server.LESSON + `?id=${id}&mode=2`);
};

export const getLessonByCourseID = (id) => {
  return httpClient.get(server.LESSON + `?courseName=${id}&mode=3`);
};
