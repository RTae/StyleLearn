import httpClient from "@/service/httpClient";
import { server } from "@/service/constants";

export const getLessonByCourse = (id) => {
  return httpClient.get(server.LESSON + `?id=${id}&mode=2`);
};
