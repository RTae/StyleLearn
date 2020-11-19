import Vue from "vue";
import Vuex from "vuex";
import api from "../service/api"

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    coreHeader: false,
    logoHeader: false,
    withOutSearchHeader: false,
    loginHeader: false,
    dialogLogin: false,
    dialogMessage: ""
  },
  getters: {
    getCoreHeader (state) {
      return state.coreHeader
    },
    getLogoHeader (state) {
      return state.logoHeader
    },
    getWithOutSearchHeader (state) {
      return state.withOutSearchHeader
    },
    getLoginHeader (state) {
      return state.loginHeader
    },
    getDialogState (state) {
      return state.dialogLogin
    },
    getDialogMsg (state) {
      return state.dialogMessage
    }
  },
  mutations: {
    SET_CORE_HEADER (state, value) {
      state.coreHeader = value
    },
    SET_LOGO_HEADER (state, value) {
      state.logoHeader = value
    },
    SET_WITHOUT_SEARCH_HEADER (state, value) {
      state.withOutSearchHeader = value
    },
    SET_LOGIN_HEADER (state, value) {
      state.loginHeader = value
    },
    SET_DIALOG_LOGIN (state, value) {
      state.dialogLogin = value
    },
    SET_DIALOG_MESSAGE (state, msg) {
      state.dialogMessage = msg
    }
  },
  actions: {
    enterSignUp ({ commit }) {
      commit("SET_CORE_HEADER", false)
      commit("SET_LOGO_HEADER", false)
      commit("SET_WITHOUT_SEARCH_HEADER", true)
      commit("SET_LOGIN_HEADER", false)
    },
    enterHome ({ commit }) {
      commit("SET_CORE_HEADER", true)
      commit("SET_LOGO_HEADER", false)
      commit("SET_WITHOUT_SEARCH_HEADER", false)
      commit("SET_LOGIN_HEADER", false)
    },
    enterDialog ({ commit }) {
      commit("SET_CORE_HEADER", false)
      commit("SET_LOGO_HEADER", true)
      commit("SET_WITHOUT_SEARCH_HEADER", false)
      commit("SET_LOGIN_HEADER", false)
    },
    dialogPopup ({ commit }, { value, msg }) {
      commit("SET_DIALOG_LOGIN", value)
      commit("SET_DIALOG_MESSAGE", msg)
    },
    async doLogin ({ commit, dispatch }, { email, password }) {
      const result = await api.login({ email, password });
      if (result.status === "1") {
        console.log("Success")
      } else {
        dispatch({ type: "dialogPopup", value: true, msg: result.msg })
      }
    }
  },
  modules: {}
});
