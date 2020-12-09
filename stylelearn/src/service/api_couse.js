import httpClient from "@/service/httpClient";
import { server } from "@/service/constants";

export const getAllCouse = () => {
  return httpClient.get(server.COURSE);
};

export const getCourseBySubject = (id) => {
  return httpClient.get(server.COURSE + `?id=${id}&mode=2`);
};

export const getCourseBySubjectName = (name) => {
  return httpClient.get(server.COURSE + `?id=${name}&mode=3`);
};
