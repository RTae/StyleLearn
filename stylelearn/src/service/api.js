import httpClient from "../service/httpClient";
import { server } from "../service/constants";
import * as courseAPI from "../service/api_couse";
import * as subjectAPI from "../service/api_subject";
import * as lessonAPI from "../service/api_lesson";
import * as progressLessonAPI from "../service/api_progresslesson";
import * as invoiceAPI from "../service/api_invoice";
import * as uploadAPI from "../service/api_upload";
import * as paymentAPI from "../service/api_payment";
import * as videoAPI from "../service/api_video";

const isLoggedIn = () => {
  const token = localStorage.getItem(server.TOKEN_KEY);
  return token != null;
};

const logoff = () => {
  localStorage.removeItem(server.TOKEN_KEY);
  localStorage.removeItem(server.USER_TYPE);
  localStorage.removeItem(server.USERNAME);
};

const login = async values => {
  var bodyFormData = new FormData();
  bodyFormData.append("email", values.email);
  bodyFormData.append("password", values.password);

  const result = await httpClient.post(server.LOGIN_URL, bodyFormData);
  if (result.data.status === "1") {
    localStorage.setItem(server.USERNAME, result.data.result.userID);
    localStorage.setItem(server.USER_TYPE, result.data.result.userType);
    localStorage.setItem(server.TOKEN_KEY, "TOKEN123qweasd");
    return result.data;
  } else {
    return result.data;
  }
};

const register = async values => {
  var bodyFormData = new FormData();
  bodyFormData.append("firstname", values.firstname);
  bodyFormData.append("familyname", values.familyname);
  bodyFormData.append("brithday", values.birthday);
  bodyFormData.append("sex", values.sex);
  bodyFormData.append("email", values.email);
  bodyFormData.append("password", values.password);
  bodyFormData.append("userType", values.role);
  bodyFormData.append("educationType", values.edu);

  const result = await httpClient.post(server.REGISTER_URL, bodyFormData);
  if (result.data.status === "1") {
    return result.data
  } else {
    return result.data
  }
};

export const getUser = (id) => {
  return httpClient.get(server.USER + `?id=${id}&mode=2`);
};

export const updateProfile = async values => {
  var bodyFormData = new FormData();
  bodyFormData.append("id", values.id);
  bodyFormData.append("firstname", values.firstname);
  bodyFormData.append("familyname", values.familyname);
  bodyFormData.append("birthday", values.birthday);
  bodyFormData.append("linkPic", values.linkPic);
  bodyFormData.append("sex", values.sex);
  bodyFormData.append("edu", values.edu);
  bodyFormData.append("bio", values.bio);
  const result = await httpClient.put(server.USER, bodyFormData);
  if (result.data.status === "1") {
    return result.data
  } else {
    return result.data
  }
}

export const changePassword = async values => {
  var bodyFormData = new FormData();
  bodyFormData.append("id", values.id);
  bodyFormData.append("oldPassword", values.oldPassword);
  bodyFormData.append("newPassword", values.newPassword);
  const result = await httpClient.post(server.CHAGNGE_PASSWORD, bodyFormData);
  if (result.data.status === "1") {
    return result.data
  } else {
    return result.data
  }
}

export default {
  login,
  register,
  updateProfile,
  changePassword,
  isLoggedIn,
  logoff,
  getUser,
  ...progressLessonAPI,
  ...courseAPI,
  ...subjectAPI,
  ...lessonAPI,
  ...invoiceAPI,
  ...uploadAPI,
  ...paymentAPI,
  ...videoAPI
};
