<template>
  <div id="app">
    <v-app>
      <Header v-if="$store.getters.getCoreHeader"/>
      <HeaderLogo v-if="$store.getters.getLogoHeader"/>
      <HeaderStudent v-if="$store.getters.getLoginHeaderStudent"/>
      <HeaderTutor v-if="$store.getters.getLoginHeaderTutor"/>
      <Content />
      <Footer />
      <!-- Popup overlay -->
      <v-overlay :value="$store.getters.getDialogLoading">
        <v-progress-circular
          indeterminate
          size="64"
        ></v-progress-circular>
      </v-overlay>
    </v-app>
  </div>
</template>

<script>
import Header from "@/components/core/Header";
import HeaderLogo from "@/components/core/HeaderLogo";
import HeaderStudent from "@/components/core/HeaderLoginStudent";
import HeaderTutor from "@/components/core/HeaderLoginTutor";
import Content from "@/components/core/Content";
import Footer from "@/components/core/Footer";

export default {
  name: "App",
  components: {
    Header,
    HeaderLogo,
    HeaderStudent,
    HeaderTutor,
    Content,
    Footer
  },
  mounted () {
    this.$store.dispatch({ type: "restoreLogin" });
  }
};
</script>

<style>
@font-face {
  font-family: "Delius";
  src: local("Delius"),
   url(./assets/font/Delius-Regular.ttf) format("truetype");
}
@font-face {
  font-family: "THSarabunNew";
  src: local("THSarabunNew"),
   url(./assets/font/THSarabunNew.ttf) format("truetype");
}
</style>
