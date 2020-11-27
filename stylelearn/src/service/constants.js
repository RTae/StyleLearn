export const NETWORK_CONNECTION_MESSAGE =
  "Cannot connect to server, Please try again.";
export const NETWORK_TIMEOUT_MESSAGE =
  "A network timeout has occurred, Please try again.";
export const UPLOAD_PHOTO_FAIL_MESSAGE =
  "An error has occurred. The photo was unable to upload.";
export const NOT_CONNECT_NETWORK = "NOT_CONNECT_NETWORK";

export const apiUrl = "https://stylelearn-backend.vercel.app/api";

export const server = {
  LOGIN_URL: apiUrl + "/login/index",
  REGISTER_URL: apiUrl + "/register/index",
  COURSE: apiUrl + "/course/index",
  SUBJECT: apiUrl + "/subject/index",
  LESSON: apiUrl + "/lesson/index",
  USERNAME: "username",
  USER_TYPE: "userType",
  TOKEN_KEY: "TOKEN"
};
