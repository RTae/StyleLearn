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
    userType: null,
    bukectListLesson: [],
    unPaidState: false,
    stateLoginDialog: false
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
    },
    getBukectList (state) {
      return state.bukectListLesson
    },
    getUnPaidSate (state) {
      return state.unPaidState
    },
    getStateLoginDialog (state) {
      return state.stateLoginDialog
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
    },
    ADD_ITEM_TO_BUKECT (state, values) {
      state.bukectListLesson.push(values)
    },
    REMOVE_ITEM_FROM_BUKECT (state, idx) {
      if (idx > -1) {
        state.bukectListLesson.splice(idx, 1);
      }
    },
    SET_UNPAID_STATE (state, value) {
      state.unPaidState = value
    },
    CLEAR_ITEM_IN_BUKECT (state) {
      state.bukectListLesson = []
    },
    SET_STATE_LOGIN_DIALOG (state, value) {
      state.stateLoginDialog = value
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
    restoreLogin ({ commit, dispatch }) {
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
          dispatch({ type: "restoreBukect" });
          dispatch({ type: "checkUnPaidInvoice" });
          commit("SET_STATE_LOGIN_DIALOG", true)
          setInterval(() => {
            commit("SET_STATE_LOGIN_DIALOG", false)
          }, 3000);
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
          dispatch({ type: "restoreBukect" });
          dispatch({ type: "checkUnPaidInvoice" });
          router.push({ name: "Home" })
          commit("SET_STATE_LOGIN_DIALOG", true)
          setInterval(() => {
            commit("SET_STATE_LOGIN_DIALOG", false)
          }, 3000);
        } else if (userType === "Tutor") {
          commit("SET_DIALOG_LOADING", false)
          commit("SET_USERNAME", username);
          commit("SET_USER_TYPE", userType);
          commit("SET_CORE_HEADER", false)
          commit("SET_LOGO_HEADER", false)
          commit("SET_LOGIN_HEADER_STUDENT", false)
          commit("SET_LOGIN_HEADER_TUTOR", true)
          commit("SET_DIALOG_LOADING", false)
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
    async doRegister ({ commit, dispatch }, { firstname, familyname, birthday, sex, email, password, role, edu }) {
      commit("SET_DIALOG_LOADING", true)
      const result = await api.register({ firstname, familyname, birthday, sex, email, password, role, edu });
      if (result.status === "1") {
        commit("SET_DIALOG_LOADING", false)
        router.push({ name: "SignUpSec" })
      } else {
        commit("SET_DIALOG_LOADING", false)
        dispatch({ type: "dialogPopup", value: true, msg: result.msg })
      }
    },
    async editProfile ({ commit, dispatch }, { id, firstname, familyname, birthday, sex, edu, bio, newImage, image }) {
      commit("SET_DIALOG_LOADING", true)
      if (newImage !== null) {
        const result = await api.uploadImage({ newImage })
        var linkPic = result.data.url
        if (result.status === 200) {
          const result = await api.updateProfile({ id, firstname, familyname, birthday, linkPic, sex, edu, bio })
          if (result.status === "1") {
            commit("SET_DIALOG_LOADING", false)
            dispatch({
              type: "dialogPopup",
              value: true,
              msg: "Update done"
            });
          } else {
            commit("SET_DIALOG_LOADING", false)
            dispatch({
              type: "dialogPopup",
              value: true,
              msg: result.msg
            });
          }
        } else {
          commit("SET_DIALOG_LOADING", false)
          dispatch({
            type: "dialogPopup",
            value: true,
            msg: "Upload picture fail" + " : " + result.status.toString()
          });
        }
      } else {
        linkPic = image
        const result = await api.updateProfile({ id, firstname, familyname, birthday, linkPic, sex, edu, bio })
        if (result.status === "1") {
          commit("SET_DIALOG_LOADING", false)
          dispatch({
            type: "dialogPopup",
            value: true,
            msg: "Update done"
          });
        } else {
          commit("SET_DIALOG_LOADING", false)
          dispatch({
            type: "dialogPopup",
            value: true,
            msg: result.msg
          });
        }
      }
    },
    async changePassword ({ commit, dispatch }, { id, oldPassword, newPassword }) {
      commit("SET_DIALOG_LOADING", true)
      const result = await api.changePassword({ id, oldPassword, newPassword })
      if (result.status === "1") {
        commit("SET_DIALOG_LOADING", false)
        dispatch({
          type: "dialogPopup",
          value: true,
          msg: "Update done"
        });
      } else {
        commit("SET_DIALOG_LOADING", false)
        dispatch({
          type: "dialogPopup",
          value: true,
          msg: result.msg
        });
      }
    },
    addItemToBukect ({ commit }, { id, name }) {
      var list = this.getters.getBukectList
      var c = 0
      for (var i = 0; i < list.length; i++) {
        if (list[i].id !== id) {
          c++
        }
      }
      if (c === list.length) {
        commit("ADD_ITEM_TO_BUKECT", { id: id, name: name })
      }
      localStorage.setItem(server.BUKECT, JSON.stringify(this.getters.getBukectList));
    },
    removeItemFromBuekect ({ commit }, { idx }) {
      commit("REMOVE_ITEM_FROM_BUKECT", idx)
      localStorage.setItem(server.BUKECT, JSON.stringify(this.getters.getBukectList));
    },
    restoreBukect ({ dispatch }) {
      if (api.isLoggedIn()) {
        const userType = localStorage.getItem(server.USER_TYPE);
        var bukect = JSON.parse(localStorage.getItem(server.BUKECT))
        if (bukect === null) {
          bukect = []
        }
        if (userType === "Student") {
          for (var i = 0; i < bukect.length; i++) {
            dispatch({
              type: "addItemToBukect",
              id: bukect[i].id,
              name: bukect[i].name
            });
          }
        }
      }
    },
    async confirmInvoice ({ commit, dispatch }, { userId, total, lesson }) {
      commit("SET_DIALOG_LOADING", true)
      const result = await api.createInvoice({ userId, total, lesson });
      if (result.status === "1") {
        commit("SET_DIALOG_LOADING", false)
        commit("SET_UNPAID_STATE", true)
        router.push({ name: "DetailPayment" })
        localStorage.removeItem(server.BUKECT);
      } else {
        commit("SET_DIALOG_LOADING", false)
        commit("SET_UNPAID_STATE", false)
        dispatch({ type: "dialogPopup", value: true, msg: result.msg })
      }
    },
    async checkUnPaidInvoice ({ commit }) {
      commit("SET_DIALOG_LOADING", true)
      var result = await api.getUnPaidInvoice(this.state.username)
      if (result[1]) {
        commit("SET_UNPAID_STATE", true)
        commit("SET_DIALOG_LOADING", false)
      } else {
        commit("SET_DIALOG_LOADING", false)
        commit("SET_UNPAID_STATE", false)
      }
    },
    async confirmPayment ({ commit, dispatch }, { invoiceID, userID, tranferDate, amountTranfer, total, tranferForm, tranferTo, image }) {
      commit("SET_DIALOG_LOADING", true)
      var result = await api.createPayment({ invoiceID, userID, tranferDate, amountTranfer, total, tranferForm, tranferTo })
      if (result.status === "1") {
        var newImage = image
        var resultImage = await api.uploadImage({ newImage })
        if (resultImage.status === 200) {
          var status = "1"
          var resultInvoice = await api.updateStatusInvoice({ invoiceID, status })
          if (resultInvoice.status === "1") {
            var resultInvoiceLine = await api.getLineItemInvoice(invoiceID)
            if (resultInvoiceLine.status === "1") {
              for (var idx = 0; idx < resultInvoiceLine.result.length; idx++) {
                var lessonID = resultInvoiceLine.result[idx].LessonID
                var quantityDay = resultInvoiceLine.result[idx].QuantityDay
                var resutlProgressLesson = await api.createProgressLesson({ userID, lessonID, quantityDay })
                if (resutlProgressLesson.status === "1") {
                } else {
                  commit("SET_DIALOG_LOADING", false)
                  dispatch({ type: "dialogPopup", value: true, msg: resutlProgressLesson.msg })
                }
              }
              commit("SET_DIALOG_LOADING", false)
              commit("SET_UNPAID_STATE", false)
              commit("CLEAR_ITEM_IN_BUKECT")
              router.push({ name: "Home" })
            } else {
              commit("SET_DIALOG_LOADING", false)
              dispatch({ type: "dialogPopup", value: true, msg: resultInvoiceLine.msg })
            }
          } else {
            commit("SET_DIALOG_LOADING", false)
            dispatch({ type: "dialogPopup", value: true, msg: resultInvoice.msg })
          }
        } else {
          commit("SET_DIALOG_LOADING", false)
          dispatch({
            type: "dialogPopup",
            value: true,
            msg: "Upload fail" + " : " + resultImage.status.toString()
          })
        }
      } else {
        commit("SET_DIALOG_LOADING", false)
        dispatch({ type: "dialogPopup", value: true, msg: result.msg })
      }
    },
    async uploadVideo ({ commit, dispatch }, { userID, lessonID, description, videoFile, status }) {
      commit("SET_DIALOG_LOADING", true)
      var resultVideo = await api.uploadVideo({ videoFile })
      if (resultVideo.status === 200) {
        var videoURL = resultVideo.data.url
        var result = await api.createVideo({ userID, lessonID, description, videoURL, status })
        if (result.status === "1") {
          commit("SET_DIALOG_LOADING", false)
          dispatch({
            type: "dialogPopup",
            value: true,
            msg: "Upload Done !!"
          })
        } else {
          commit("SET_DIALOG_LOADING", false)
          dispatch({
            type: "dialogPopup",
            value: true,
            msg: result.msg
          })
        }
      } else {
        commit("SET_DIALOG_LOADING", false)
        dispatch({
          type: "dialogPopup",
          value: true,
          msg: "Upload faile" + " : " + resultVideo.status.toString()
        })
      }
    }
  }
});
