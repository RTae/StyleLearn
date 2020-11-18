import Vue from "vue";
import VueRouter from "vue-router";
import Login from "../views/Login.vue";
import Signup from "../views/Signup.vue";
import Home from "../views/Home.vue";
import HomeTutor from "../views/Tutor/HomeTutor.vue";
import TestTutor from "../views/Tutor/TestTutor.vue";
import MyVideoTutor from "../views/Tutor/MyVideoTutor.vue";
import ProfileTutor from "../views/Tutor/ProfileTutor.vue";
import UploadVideoTutor from "../views/Tutor/UploadVideoTutor.vue";
import BasicUse from "../views/BasicUse.vue";
import About from "../views/About.vue";
import ForgetPassword from "../views/ForgetPassword.vue";
import ForgetPasswordPop from "../views/ForgetPasswordPop.vue";
import SignUpSec from "../views/SignUpSuc.vue";
import NewPassword from "../views/NewPassword.vue";
import CoursesPage from "../views/CoursesPage.vue";
import LessonPage from "../views/LessonPage.vue";
import Adminhome from "../views/Admin/Adminhome.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/login",
    name: "Login",
    component: Login
  },
  {
    path: "/signup",
    name: "Signup",
    component: Signup
  },
  {
    path: "/hometutor",
    name: "HomeTutor",
    component: HomeTutor
  },
  {
    path: "/testtutor",
    name: "TestTutor",
    component: TestTutor
  },
  {
    path: "/myvideotutor",
    name: "MyVideoTutor",
    component: MyVideoTutor
  },
  {
    path: "/profiletutor",
    name: "ProfileTutor",
    component: ProfileTutor
  },
  {
    path: "/uploadvideotutor",
    name: "UploadVideoTutor",
    component: UploadVideoTutor
  },
  {
    path: "/basicuse",
    name: "BasicUse",
    component: BasicUse
  },
  {
    path: "/about",
    name: "About",
    component: About
  },
  {
    path: "/forgetpassword",
    name: "ForgetPassword",
    component: ForgetPassword
  },
  {
    path: "/forgetpasswordpop",
    name: "ForgetPasswordPop",
    component: ForgetPasswordPop
  },
  {
    path: "/signupsuccessfully",
    name: "SignUpSec",
    component: SignUpSec
  },
  {
    path: "/newpassword",
    name: "NewPassword",
    component: NewPassword
  },
  {
    path: "/coursespage",
    name: "CoursesPage",
    component: CoursesPage
  },
  {
    path: "/lessonpage",
    name: "LessonPage",
    component: LessonPage
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
