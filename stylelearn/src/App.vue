<template>
  <div id="app">
    <v-app>
      <Header v-if="$store.getters.getCoreHeader"/>
      <HeaderLogo v-if="$store.getters.getLogoHeader"/>
      <HeaderWithLogin v-if="$store.getters.getLoginHeader"/>
      <Content />
      <Footer />
    </v-app>
  </div>
</template>

<script>
import Header from "@/components/core/Header";
import HeaderLogo from "@/components/core/HeaderLogo";
import HeaderWithLogin from "@/components/core/HeaderWithLogin";
import Content from "@/components/core/Content";
import Footer from "@/components/core/Footer";

export default {
  name: "App",
  components: {
    Header,
    HeaderLogo,
    HeaderWithLogin,
    Content,
    Footer
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
