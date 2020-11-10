import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Signup from '../views/Signup.vue'
import Home from '../views/Home.vue'
import HomeTutor from '../views/Tutor/HomeTutor.vue'
import TestTutor from '../views/Tutor/TestTutor.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/signup',
    name: 'Signup',
    component: Signup
  },
  {
    path: '/homeTutor',
    name: 'HomeTutor',
    component: HomeTutor
  },
  {
    path: '/testTutor',
    name: 'TestTutor',
    component: TestTutor
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
