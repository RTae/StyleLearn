import Vue from "vue";
import VueRouter from "vue-router";
import { server } from "../service/constants";
import api from "../service/api";
// General
import Login from "../views/Login";
import Signup from "../views/Signup";
import SignUpSec from "../views/SignUpSuc";
import Home from "../views/Home";
import BasicUse from "../views/BasicUse";
import About from "../views/About";
import CoursesPage from "../views/CoursesPage";
import LessonPage from "../views/LessonPage";
import CoursePopAndNew from "../views/CoursePopAndNewest.vue"
// Tutor
import HomeTutor from "../views/Tutor/HomeTutor";
import MyVideoTutor from "../views/Tutor/MyVideoTutor";
import ProfileTutor from "../views/Tutor/ProfileTutor";
import EditProfileTutor from "../views/Tutor/EditProfileTutor";
import changePasswordTutor from "../views/Tutor/changePasswordTutor";
import UploadVideoTutor from "../views/Tutor/UploadVideoTutor";
// Student
import MyCourse from "../views/Student/MyCourse";
import LearnCourse from "../views/Student/LearnCourse";
import LearnCourseTutorPage from "../views/Student/LearnCourseTutorPage";
import VideoStudent from "../views/Student/VideoStudent";
import SelectedItemInvoice from "../views/Student/SelectedItemInvoice";
import SelectItem from "../views/Student/SelectItem";
import DetailPayment from "../views/Student/DetailPayment";
import ConfirmPayment from "../views/Student/ConfirmPayment.vue";
import ProfileStudent from "../views/Student/ProfileStudent.vue";
import EditProfileStudent from "../views/Student/EditProfileStudent.vue";
import ChangePasswordStudent from "../views/Student/ChangePasswordStudent.vue";

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
    path: "/lesson",
    name: "LessonPage",
    component: LessonPage,
    meta: {
      isAllowGuest: true
    }
  },
  {
    path: "/coursePN",
    name: "CoursePandN",
    component: CoursePopAndNew,
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
    path: "/accountTutor",
    name: "ChangePasswordTutor",
    meta: {
      isSecured: true,
      isTutor: true,
      isStudent: false
    },
    component: changePasswordTutor
  },
  {
    path: "/upload",
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
    path: "/mycourse/lesson",
    name: "LearnCourse",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true
    },
    component: LearnCourse
  },
  {
    path: "/mycourse/lesson/tutor",
    name: "TutorPage",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true
    },
    component: LearnCourseTutorPage
  },
  {
    path: "/mycourse/lesson/tutor/video",
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
      isStudent: true,
      isPayment: true
    },
    component: SelectedItemInvoice
  },
  {
    path: "/selectitem",
    name: "SelectItem",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true,
      isPayment: true
    },
    component: SelectItem
  },
  {
    path: "/detailpayment",
    name: "DetailPayment",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true,
      isPayment: true
    },
    component: DetailPayment
  },
  {
    path: "/confirmpayment",
    name: "ConfirmPayment",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true,
      isPayment: true
    },
    component: ConfirmPayment
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
  },
  {
    path: "/editprofilestudent",
    name: "EditProfileStudent",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true
    },
    component: EditProfileStudent
  },
  {
    path: "/accountstudent",
    name: "ChangePasswordStudent",
    meta: {
      isSecured: true,
      isTutor: false,
      isStudent: true
    },
    component: ChangePasswordStudent
  },
  // ETC
  {
    path: "/",
    redirect: "/home" // Home
  },
  {
    path: "*",
    redirect: "/home" // page not found
  }
]

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

router.beforeEach(async (to, from, next) => {
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
    next();
  }
});

export default router;
