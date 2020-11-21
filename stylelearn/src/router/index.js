import Vue from "vue";
import VueRouter from "vue-router";
import Login from "../views/Login";
import Signup from "../views/Signup";
import Home from "../views/Home";
import HomeTutor from "../views/Tutor/HomeTutor";
import TestTutor from "../views/Tutor/TestTutor";
import MyVideoTutor from "../views/Tutor/MyVideoTutor";
import ProfileTutor from "../views/Tutor/ProfileTutor";
import EditProfileTutor from "../views/Tutor/EditProfileTutor";
import EditVideoTutor from "../views/Tutor/EditVideoTutor";
import UploadVideoTutor from "../views/Tutor/UploadVideoTutor";
import BasicUse from "../views/BasicUse";
import About from "../views/About";
import ForgetPassword from "../views/ForgetPassword";
import ForgetPasswordPop from "../views/ForgetPasswordPop";
import SignUpSec from "../views/SignUpSuc";
import NewPassword from "../views/NewPassword";
import CoursesPage from "../views/CoursesPage";
import LessonPage from "../views/LessonPage";
import AdminHome from "../views/Admin/Adminhome";
import AdminEditvideo from "../views/Admin/Admineditvideo";
import AdminclickvideoPop from "../views/Admin/AdminclickvideoPop";
import Adminclickvideo from "../views/Admin/Adminclickvideo";
import Adminapprovevideo from "../views/Admin/Adminapprovevideo";
import AdminapprovevideoPop1 from "../views/Admin/AdminapprovevideoPop1";
import AdminapprovevideoPop2 from "../views/Admin/AdminapprovevideoPop2";
import Admincustomer from "../views/Admin/Admincustomer";
import Adminapprovepayment from "../views/Admin/Adminapprovepayment";
import AdminapprovepaymentPop1 from "../views/Admin/AdminapprovepaymentPop1.vue"
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
    path: "/editprofiletutor",
    name: "EditProfileTutor",
    component: EditProfileTutor
  },
  {
    path: "/editvideotutor",
    name: "EditVideoTutor",
    component: EditVideoTutor
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
  },
  {
    path: "/mycourse",
    name: "MyCourse",
    component: MyCourse
  },
  {
    path: "/learncourse",
    name: "LearnCourse",
    component: LearnCourse
  },
  {
    path: "/learncoursetutorpage",
    name: "LearnCourseTutorPage",
    component: LearnCourseTutorPage
  },
  {
    path: "/adminhome",
    name: "AdminHome",
    component: AdminHome
  },
  {
    path: "/admineditvideo",
    name: "Admineditvideo",
    component: AdminEditvideo
  },
  {
    path: "/adminclickvideopop",
    name: "adminclickvideoPop",
    component: AdminclickvideoPop
  },
  {
    path: "/adminclickvideo",
    name: "adminclickvideo",
    component: Adminclickvideo
  },
  {
<<<<<<< HEAD
    path: "/video",
=======
    path: "/adminapprovevideo",
    name: "Adminapprovevideo",
    component: Adminapprovevideo
  },
  {
    path: "/adminapprovevideoPop1",
    name: "AdminapprovevideoPop1",
    component: AdminapprovevideoPop1
  },
  {
    path: "/adminapprovevideoPop2",
    name: "AdminapprovevideoPop2",
    component: AdminapprovevideoPop2
  },
  {
    path: "/admincustomer",
    name: "Admincustomer",
    component: Admincustomer
  },
  {
    path: "/adminapprovepayment",
    name: "Adminapprovepayment",
    component: Adminapprovepayment
  },
  {
    path: "/adminapprovepaymentPop1",
    name: "AdminapprovepaymentPop1",
    component: AdminapprovepaymentPop1
  },
  {
    path: "/videostudent",
>>>>>>> c8bfd0eb3812ec13a122cf869b3bdc9ba6d16a9e
    name: "VideoStudent",
    component: VideoStudent
  },
  {
    path: "/invoice",
    name: "SelectedItemInvoice",
    component: SelectedItemInvoice
  },
  {
    path: "/selectitem",
    name: "SelectItem",
    component: SelectItem
  }

];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
