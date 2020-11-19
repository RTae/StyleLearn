<template>
  <div id="app">
    <v-app>
      <Header v-if="$store.getters.getCoreHeader"/>
      <HeaderNoSearchBar v-if="$store.getters.getWithOutSearchHeader"/>
      <HeaderLogo v-if="$store.getters.getLogoHeader"/>
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
import HeaderLogo from "@/components/core/HeaderLogo"
import Content from "@/components/core/Content";
import Footer from "@/components/core/Footer";

export default {
  name: "App",
  components: {
    Header,
    HeaderNoSearchBar,
    HeaderLogo,
    Content,
    Footer
  }
};
</script>
