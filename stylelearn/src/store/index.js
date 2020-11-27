import Vue from "vue";
import Vuex from "vuex";
import api from "../service/api";
import router from "@/router";
import { server } from "../service/constants";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    coreHeader: true,
    logoHeader: false,
    loginHeaderStudent: false,
    loginHeaderTutor: false,
    dialogState: false,
    dialogMessage: "",
    dialogLoading: false,
    username: null,
    userType: null
  },
  getters: {
    getCoreHeader (state) {
      return state.coreHeader
    },
    getLogoHeader (state) {
      return state.logoHeader
    },
    getLoginHeaderStudent (state) {
      return state.loginHeaderStudent
    },
    getLoginHeaderTutor (state) {
      return state.loginHeaderTutor
    },
    getDialogState (state) {
      return state.dialogState
    },
    getDialogMsg (state) {
      return state.dialogMessage
    },
    getDialogLoading (state) {
      return state.dialogLoading
    },
    getUserName (state) {
      return state.username
    },
    getUserType (state) {
      return state.userType
    }
  },
  mutations: {
    SET_CORE_HEADER (state, value) {
      state.coreHeader = value
    },
    SET_LOGO_HEADER (state, value) {
      state.logoHeader = value
    },
    SET_LOGIN_HEADER_STUDENT (state, value) {
      state.loginHeaderStudent = value
    },
    SET_LOGIN_HEADER_TUTOR (state, value) {
      state.loginHeaderTutor = value
    },
    SET_DIALOG_STATE (state, value) {
      state.dialogState = value
    },
    SET_DIALOG_MESSAGE (state, msg) {
      state.dialogMessage = msg
    },
    SET_DIALOG_LOADING (state, value) {
      state.dialogLoading = value
    },
    SET_USERNAME (state, value) {
      state.username = value
    },
    SET_USER_TYPE (state, value) {
      state.userType = value
    }
  },
  actions: {
    enterDefault ({ commit }) {
      commit("SET_CORE_HEADER", true)
      commit("SET_LOGO_HEADER", false)
      commit("SET_LOGIN_HEADER_STUDENT", false)
      commit("SET_LOGIN_HEADER_TUTOR", false)
    },
    enterDialog ({ commit }) {
      commit("SET_CORE_HEADER", false)
      commit("SET_LOGO_HEADER", true)
      commit("SET_LOGIN_HEADER_STUDENT", false)
      commit("SET_LOGIN_HEADER_TUTOR", false)
    },
    dialogPopup ({ commit }, { value, msg }) {
      commit("SET_DIALOG_STATE", value)
      commit("SET_DIALOG_MESSAGE", msg)
    },
    restoreLogin ({ commit }) {
      if (api.isLoggedIn() === true) {
        const userType = localStorage.getItem(server.USER_TYPE);
        const username = localStorage.getItem(server.USERNAME);
        if (userType === "Student") {
          commit("SET_DIALOG_LOADING", false)
          commit("SET_USERNAME", username);
          commit("SET_USER_TYPE", userType);
          commit("SET_CORE_HEADER", false)
          commit("SET_LOGO_HEADER", false)
          commit("SET_LOGIN_HEADER_STUDENT", true)
          commit("SET_LOGIN_HEADER_TUTOR", false)
        } else if (userType === "Tutor") {
          commit("SET_DIALOG_LOADING", false)
          commit("SET_USERNAME", username);
          commit("SET_USER_TYPE", userType);
          commit("SET_CORE_HEADER", false)
          commit("SET_LOGO_HEADER", false)
          commit("SET_LOGIN_HEADER_STUDENT", false)
          commit("SET_LOGIN_HEADER_TUTOR", true)
        }
      }
    },
    doLogout ({ commit }) {
      api.logoff()
      commit("SET_USERNAME", null);
      commit("SET_CORE_HEADER", true)
      commit("SET_LOGO_HEADER", false)
      commit("SET_LOGIN_HEADER_STUDENT", false)
      commit("SET_LOGIN_HEADER_TUTOR", false)
      router.push({ name: "Home" })
    },
    async doLogin ({ commit, dispatch }, { email, password }) {
      commit("SET_DIALOG_LOADING", true)
      const result = await api.login({ email, password });
      if (result.status === "1") {
        const userType = localStorage.getItem(server.USER_TYPE);
        const username = localStorage.getItem(server.USERNAME);
        if (userType === "Student") {
          commit("SET_DIALOG_LOADING", false)
          commit("SET_USERNAME", username);
          commit("SET_USER_TYPE", userType);
          commit("SET_CORE_HEADER", false)
          commit("SET_LOGO_HEADER", false)
          commit("SET_LOGIN_HEADER_STUDENT", true)
          commit("SET_LOGIN_HEADER_TUTOR", false)
          router.push({ name: "Home" })
        } else if (userType === "Tutor") {
          commit("SET_DIALOG_LOADING", false)
          commit("SET_USERNAME", username);
          commit("SET_USER_TYPE", userType);
          commit("SET_CORE_HEADER", false)
          commit("SET_LOGO_HEADER", false)
          commit("SET_LOGIN_HEADER_STUDENT", false)
          commit("SET_LOGIN_HEADER_TUTOR", true)
          router.push({ name: "HomeTutor" })
        }
      } else {
        commit("SET_DIALOG_LOADING", false)
        dispatch({ type: "dialogPopup", value: true, msg: result.msg })
        commit("SET_USERNAME", null);
        commit("SET_CORE_HEADER", true)
        commit("SET_LOGIN_HEADER_STUDENT", false)
        commit("SET_LOGIN_HEADER_TUTOR", false)
      }
    },
    async doRegister ({ commit, dispatch }, { firtname, familyname, birthday, sex, email, password, role, edu }) {
      commit("SET_DIALOG_LOADING", true)
      const result = await api.register({ firtname, familyname, birthday, sex, email, password, role, edu });
      if (result.status === "1") {
        commit("SET_DIALOG_LOADING", false)
        router.push({ name: "SignUpSec" })
      } else {
        commit("SET_DIALOG_LOADING", false)
        dispatch({ type: "dialogPopup", value: true, msg: result.msg })
      }
    }
  },
  modules: {}
});
