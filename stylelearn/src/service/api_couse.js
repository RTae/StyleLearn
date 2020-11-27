import httpClient from "@/service/httpClient";
import { server } from "@/service/constants";

export const getCourseBySubject = (id) => {
  return httpClient.get(server.COURSE + `?id=${id}&mode=2`);
};
