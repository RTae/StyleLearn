import httpClient from "../service/httpClient";
import { server } from "../service/constants";
import * as courseAPI from "../service/api_couse";
import * as subjectAPI from "../service/api_subject";
import * as lessonAPI from "../service/api_lesson";
import * as progressLessonAPI from "../service/api_progresslesson";

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
  bodyFormData.append("firstname", values.firtname);
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
  bodyFormData.append("firstName", values.firstName);
  bodyFormData.append("familyname", values.familyname);
  bodyFormData.append("birthday", values.birthday);
  bodyFormData.append("sex", values.sex);
  bodyFormData.append("edu", values.edu);
  const result = await httpClient.post(server.USER, bodyFormData);
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

export const uploadFile = async values => {
  var bodyFormData = new FormData();
  bodyFormData.append("file", values.pic);
  bodyFormData.append("name", values.fileName);
  bodyFormData.append("fileType", values.fileType);
  const result = await httpClient.post(server.UPLOAD_FILE, bodyFormData);
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
  uploadFile,
  ...progressLessonAPI,
  ...courseAPI,
  ...subjectAPI,
  ...lessonAPI
};
