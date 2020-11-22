import Vue from "vue";
import VueRouter from "vue-router";
import api from "../service/api";
import Login from "../views/Login";
import Signup from "../views/Signup";
import SignUpSec from "../views/SignUpSuc";
import Home from "../views/Home";
import BasicUse from "../views/BasicUse";
import About from "../views/About";
import CoursesPage from "../views/CoursesPage";
import LessonPage from "../views/LessonPage";
import HomeTutor from "../views/Tutor/HomeTutor";
import TestTutor from "../views/Tutor/TestTutor";
import MyVideoTutor from "../views/Tutor/MyVideoTutor";
import ProfileTutor from "../views/Tutor/ProfileTutor";
import EditProfileTutor from "../views/Tutor/EditProfileTutor";
import EditVideoTutor from "../views/Tutor/EditVideoTutor";
import UploadVideoTutor from "../views/Tutor/UploadVideoTutor";
import MyCourse from "../views/MyCourse.vue";
import LearnCourse from "../views/LearnCourse.vue";
import LearnCourseTutorPage from "../views/LearnCourseTutorPage.vue";
import VideoStudent from "../views/VideoStudent.vue";
import SelectedItemInvoice from "../views/SelectedItemInvoice.vue";
import SelectItem from "../views/SelectItem.vue";

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
    meta: {
      requiresVisitor: true,
      isSecured: false
    },
    component: Signup
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
    path: "/hometutor",
    name: "HomeTutor",
    meta: {
      isSecured: true
    },
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
    meta: {
      isSecured: true
    },
    component: MyVideoTutor
  },
  {
    path: "/profiletutor",
    name: "ProfileTutor",
    meta: {
      isSecured: true
    },
    component: ProfileTutor
  },
  {
    path: "/editprofiletutor",
    name: "EditProfileTutor",
    meta: {
      isSecured: true
    },
    component: EditProfileTutor
  },
  {
    path: "/editvideotutor",
    name: "EditVideoTutor",
    meta: {
      isSecured: true
    },
    component: EditVideoTutor
  },
  {
    path: "/uploadvideotutor",
    name: "UploadVideoTutor",
    meta: {
      isSecured: true
    },
    component: UploadVideoTutor
  },
  {
    path: "/signupsuccessfully",
    name: "SignUpSec",
    component: SignUpSec
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
  },
  {
    path: "/mycourse",
    name: "MyCourse",
    meta: {
      isSecured: true
    },
    component: MyCourse
  },
  {
    path: "/learncourse",
    name: "LearnCourse",
    meta: {
      isSecured: true
    },
    component: LearnCourse
  },
  {
    path: "/learncoursetutorpage",
    name: "LearnCourseTutorPage",
    meta: {
      isSecured: true
    },
    component: LearnCourseTutorPage
  },
  {
    path: "/video",
    name: "VideoStudent",
    meta: {
      isSecured: true
    },
    component: VideoStudent
  },
  {
    path: "/invoice",
    name: "SelectedItemInvoice",
    meta: {
      isSecured: true
    },
    component: SelectedItemInvoice
  },
  {
    path: "/selectitem",
    name: "SelectItem",
    meta: {
      isSecured: true
    },
    component: SelectItem
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => (record.meta.isSecured))) {
    // secure route
    if (api.isLoggedIn()) {
      next();
    } else {
      next("/");
    }
  } else if (to.matched.some(record => record.meta.requiresVisitor)) {
    // when login
    if (api.isLoggedIn()) {
      next("/");
    } else {
      next();
    }
  } else {
    // unsecure route
    next();
  }
});

export default router;
