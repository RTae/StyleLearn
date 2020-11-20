import httpClient from "../service/httpClient";
import { server } from "../service/constants";

const login = async values => {
  var bodyFormData = new FormData();
  bodyFormData.append("email", values.email);
  bodyFormData.append("password", values.password);

  const result = await httpClient.post(server.LOGIN_URL, bodyFormData);
  if (result.data.status === "1") {
    localStorage.setItem(server.USERNAME, values.email);
    return result.data;
  } else {
    return result.data;
  }
};

const register = async values => {
  var bodyFormData = new FormData();
  bodyFormData.append("firstname", values.firtname);
  bodyFormData.append("familyname", values.familyname);
  bodyFormData.append("brithday", values.birthday);
  bodyFormData.append("sex", values.sex);
  bodyFormData.append("email", values.email);
  bodyFormData.append("password", values.password);
  bodyFormData.append("userType", values.role);
  bodyFormData.append("educationType", values.edu);
  console.log(bodyFormData)

  const result = await httpClient.post(server.REGISTER_URL, bodyFormData);
  if (result.data.status === "1") {
    return result.data
  } else {
    return result.data
  }
};

export default {
  login,
  register
};
