import Vue from "vue"
import App from "./App.vue"
import router from "./router"
import store from "./store"
import vuetify from "./plugins/vuetify"
import "material-design-icons-iconfont/dist/material-design-icons.css"
import VueCoreVideoPlayer from "vue-core-video-player"

Vue.config.productionTip = false;

Vue.use(VueCoreVideoPlayer)

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount("#app");
