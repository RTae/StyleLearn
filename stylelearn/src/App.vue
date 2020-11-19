<template>
  <div id="app">
    <v-app>
      <Header v-if="$store.getters.getCoreHeader"/>
      <HeaderNoSearchBar v-if="$store.getters.getWithOutSearchHeader"/>
      <HeaderLogo v-if="$store.getters.getLogoHeader"/>
      <HeaderWithLogin v-if="$store.getters.getWithLogin"/>
      <Content />
      <Footer />
    </v-app>
    <!--
    <h1>{{ $store.getters.getCount }}</h1>
    <button @click="add">Add</button>
    <button @click="del">Del</button>
    -->
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
