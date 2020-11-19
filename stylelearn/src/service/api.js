import httpClient from "../service/httpClient";
import { server } from "../service/constants";
import router from "@/router";

const login = async values => {
  var bodyFormData = new FormData();
  bodyFormData.append("email", values.email);
  bodyFormData.append("password", values.password);

  const result = await httpClient.post(server.LOGIN_URL, bodyFormData);
  if (result.data.status === "1") {
    localStorage.setItem(server.USERNAME, values.email);
    router.push({ name: "Home" })
    return result.data;
  } else {
    return result.data;
  }
};

const register = async values => {
  console.log(values)
  /*
  const result = await httpClient.post(server.REGISTER_URL, values);
  if (result.data.status === "1") {
    router.go(-1);
  } else {
    alert(JSON.stringify(result));
  }
  */
};

export default {
  login,
  register
};
