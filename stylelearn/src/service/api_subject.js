import httpClient from "@/service/httpClient";
import { server } from "@/service/constants";

export const getSubject = () => {
  return httpClient.get(server.SUBJECT);
};
