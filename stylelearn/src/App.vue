<template>
  <div id="app">
    <v-app>
      <Header v-if="false"/>
      <HeaderNoSearchBar v-if="$store.getters.getWithOutSearchHeader"/>
      <HeaderLogo v-if="$store.getters.getLogoHeader"/>
      <HeaderWithLogin v-if="true"/>
      <Content />
      <Footer />
    </v-app>
  </div>
</template>

<script>
import Header from "@/components/core/Header";
import HeaderNoSearchBar from "@/components/core/HeaderWithOutSearchBar";
import HeaderLogo from "@/components/core/HeaderLogo";
import HeaderWithLogin from "@/components/core/HeaderWithLogin";
import Content from "@/components/core/Content";
import Footer from "@/components/core/Footer";

export default {
  name: "App",
  components: {
    Header,
    HeaderNoSearchBar,
    HeaderLogo,
    HeaderWithLogin,
    Content,
    Footer
  }
};
</script>
