import Vue from "vue";
import Vuex from "vuex";
import api from "../service/api"
import router from "@/router";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    coreHeader: true,
    logoHeader: false,
    loginHeader: false,
    dialogState: false,
    dialogMessage: "",
    dialogLoading: false
  },
  getters: {
    getCoreHeader (state) {
      return state.coreHeader
    },
    getLogoHeader (state) {
      return state.logoHeader
    },
    getLoginHeader (state) {
      return state.loginHeader
    },
    getDialogState (state) {
      return state.dialogState
    },
    getDialogMsg (state) {
      return state.dialogMessage
    },
    getDialogLoading (state) {
      return state.dialogLoading
    }
  },
  mutations: {
    SET_CORE_HEADER (state, value) {
      state.coreHeader = value
    },
    SET_LOGO_HEADER (state, value) {
      state.logoHeader = value
    },
    SET_LOGIN_HEADER (state, value) {
      state.loginHeader = value
    },
    SET_DIALOG_STATE (state, value) {
      state.dialogState = value
    },
    SET_DIALOG_MESSAGE (state, msg) {
      state.dialogMessage = msg
    },
    SET_DIALOG_LOADING (state, value) {
      state.dialogLoading = value
    }
  },
  actions: {
    enterSignUp ({ commit }) {
      commit("SET_CORE_HEADER", true)
      commit("SET_LOGO_HEADER", false)
      commit("SET_LOGIN_HEADER", false)
    },
    enterHome ({ commit }) {
      commit("SET_CORE_HEADER", true)
      commit("SET_LOGO_HEADER", false)
      commit("SET_LOGIN_HEADER", false)
    },
    enterDialog ({ commit }) {
      commit("SET_CORE_HEADER", false)
      commit("SET_LOGO_HEADER", true)
      commit("SET_LOGIN_HEADER", false)
    },
    enterLogin ({ commit }) {
      commit("SET_CORE_HEADER", false)
      commit("SET_LOGO_HEADER", false)
      commit("SET_WITHOUT_SEARCH_HEADER", false)
      commit("SET_LOGIN_HEADER", true)
    },
    dialogPopup ({ commit }, { value, msg }) {
      commit("SET_DIALOG_STATE", value)
      commit("SET_DIALOG_MESSAGE", msg)
    },
    async doLogin ({ commit, dispatch }, { email, password }) {
      commit("SET_DIALOG_LOADING", true)
      const result = await api.login({ email, password });
      if (result.status === "1") {
        commit("SET_DIALOG_LOADING", false)
        dispatch({ type: "enterLogin" })
        router.push({ name: "Home" })
      } else {
        commit("SET_DIALOG_LOADING", false)
        dispatch({ type: "dialogPopup", value: true, msg: result.msg })
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
