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

export const getPopCourse = () => {
  return httpClient.get(server.COURSE + "?mode=4");
};

export const getNewCourse = () => {
  return httpClient.get(server.COURSE + "?mode=5");
};

export const getPopCourseLimit = () => {
  return httpClient.get(server.COURSE + "?mode=4&limit=1");
};

export const getNewCourseLimit = () => {
  return httpClient.get(server.COURSE + "?mode=5&limit=1");
};
