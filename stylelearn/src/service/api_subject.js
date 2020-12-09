import httpClient from "@/service/httpClient";
import { server } from "@/service/constants";

export const getAllSubject = () => {
  return httpClient.get(server.SUBJECT);
};
