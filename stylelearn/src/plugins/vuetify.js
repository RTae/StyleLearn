import Vue from "vue";
import Vuetify from "vuetify/lib";

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: "#70CCFF",
        background: "#EFEFEF",
        secondary: "#FFFFFF",
        font_b: "#000000",
        font_w: "#FFFFFF",
        error: "#EB5757"
      }
    }
  }
});
