import axios from "axios";
import {
  server,
  apiUrl,
  NOT_CONNECT_NETWORK,
  NETWORK_CONNECTION_MESSAGE
} from "./constants";
import router from "../router";

const isAbsoluteURLRegex = /^(?:\w+:)\/\//;

axios.interceptors.request.use(async config => {
  if (!isAbsoluteURLRegex.test(config.url)) {
    const userToken = localStorage.getItem(server.TOKEN_KEY);
    if (userToken) {
      config.headers = { "x-access-token": userToken };
    }
    config.url = apiUrl + "/" + config.url;
  }
  config.timeout = 600000; // 600 Second
  return config;
});

axios.interceptors.response.use(
  response => {
    return response;
  },
  error => {
    console.log(JSON.stringify(error, undefined, 2));
    router.push("/");
    if (axios.isCancel(error)) {
      return Promise.reject(error);
    } else if (!error.response) {
      return Promise.reject({
        code: NOT_CONNECT_NETWORK,
        message: NETWORK_CONNECTION_MESSAGE
      });
    }
    return Promise.reject(error);
  }
);

export default axios;
