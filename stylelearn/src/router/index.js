import Vue from "vue";
import VueRouter from "vue-router";
import { server } from "../service/constants";
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
import DetailPayment from "../views/DetailPayment.vue";
import ConfirmPayment from "../views/ConfirmPayment.vue";
import ProfileStudent from "../views/ProfileStudent.vue";

Vue.use(VueRouter);
const routes = [
  // General
  {
    path: "/home",
    name: "Home",
    component: Home,
    meta: {
      isAllowGuest: true
    }
  },
  {
    path: "/login",
    name: "Login",
    meta: {
      requiresVisitor: true
    },
    component: Login
  },
  {
    path: "/signup",
    name: "Signup",
    meta: {
      requiresVisitor: true
    },
    component: Signup
  },
  {
    path: "/signupsuccessfully",
    name: "SignUpSec",
    meta: {
      requiresVisitor: true
    },
    component: SignUpSec
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
    path: "/courses",
    name: "CoursesPage",
    component: CoursesPage,
    meta: {
      isAllowGuest: true
    }
  },
  {
    path: "/lessonpage",
    name: "LessonPage",
    component: LessonPage,
    meta: {
      isAllowGuest: true
    }
  },
  // Tutor
  {
    path: "/hometutor",
    name: "HomeTutor",
    meta: {
      isSecured: true,
      isTutor: true,
      isStudent: false
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
      isSecured: true,
      isTutor: true,
      isStudent: false
    },
    component: MyVideoTutor
  },
  {
    path: "/profiletutor",
    name: "ProfileTutor",
    meta: {
      isSecured: true,
      isTutor: true,
      isStudent: false
    },
    component: ProfileTutor
  },
  {
    path: "/editprofiletutor",
    name: "EditProfileTutor",
    meta: {
      isSecured: true,
      isTutor: true,
      isStudent: false
    },
    component: EditProfileTutor
  },
  {
    path: "/editvideotutor",
    name: "EditVideoTutor",
    meta: {
      isSecured: true,
      isTutor: true,
      isStudent: false
    },
    component: EditVideoTutor
  },
  {
    path: "/uploadvideotutor",
    name: "UploadVideoTutor",
    meta: {
      isSecured: true,
      isTutor: true,
      isStudent: false
    },
    component: UploadVideoTutor
  },
  // Student
  {
    path: "/mycourse",
    name: "MyCourse",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true
    },
    component: MyCourse
  },
  {
    path: "/learn",
    name: "LearnCourse",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true
    },
    component: LearnCourse
  },
  {
    path: "/tutorpage",
    name: "TutorPage",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true
    },
    component: LearnCourseTutorPage
  },
  {
    path: "/video",
    name: "Video",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true
    },
    component: VideoStudent
  },
  {
    path: "/invoice",
    name: "SelectedItemInvoice",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true
    },
    component: SelectedItemInvoice
  },
  {
    path: "/selectitem",
    name: "SelectItem",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true

    },
    component: SelectItem
  },
  {
    path: "/detailpayment",
    name: "DetailPayment",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true
    },
    component: DetailPayment
  },
  {
    path: "/confirmpayment",
    name: "ConfirmPayment",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true
    },
    component: ConfirmPayment
  },
  {
    path: "/",
    redirect: "/home" // Home
  },
  {
    path: "*",
    redirect: "/home" // page not found
  },
  {
    path: "/profilestudent",
    name: "ProfileStudent",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true
    },
    component: ProfileStudent
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

router.beforeEach((to, from, next) => {
  // is geust
  if (to.matched.some(record => record.meta.isAllowGuest)) {
    if (localStorage.getItem(server.USER_TYPE) === "Tutor") {
      next("/hometutor")
    } else {
      next()
    }
  } else if (to.matched.some(record => record.meta.isSecured)) {
    // secure route
    if (api.isLoggedIn()) {
      // is tutor
      if (to.matched.some(record => record.meta.isTutor && !record.meta.isStudent)) {
        if (localStorage.getItem(server.USER_TYPE) === "Tutor") {
          next()
        } else {
          next("/")
        }
      // is student
      } else if (to.matched.some(record => !record.meta.isTutor && record.meta.isStudent)) {
        if (localStorage.getItem(server.USER_TYPE) === "Student") {
          next()
        } else {
          next("/hometutor")
        }
      } else {
        next("/")
      }
    } else {
      next("/");
    }
  } else if (to.matched.some(record => record.meta.requiresVisitor)) {
    // when login can't not enter login and signup
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
