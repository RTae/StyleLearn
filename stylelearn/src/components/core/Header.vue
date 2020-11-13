<template>
  <div id="header">
    <v-app-bar absolute color="secondary" hide-on-scroll height="70">
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="../../assets/logo.png"
          transition="scale-transition"
          width="200"
        />
      </div>
      <v-spacer></v-spacer>
      <v-layout row wrap justify-end align-start>
        <v-flex xs2 md1
          ><router-link to="/"><p class="text">Home</p></router-link></v-flex
        >
        <v-flex xs2 md1
          ><router-link to="/basicuse"
            ><p class="text">Basic Use</p></router-link
          ></v-flex
        >
        <v-flex xs2 md1
          ><router-link to="/about"
            >
              <div class="textContainer">
                <p class="text">About Us</p>
              </div>
          </router-link></v-flex
        >
        <v-flex xs2 md1.5
          ><router-link to="/signup"
            ><button class="btnHeader">Sign Up</button></router-link
          ></v-flex
        >
        <v-flex xs2 md1.5
          ><router-link to="/login"
            >
            <button class="btnHeader">Login In</button></router-link
          ></v-flex
        >
      </v-layout>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  name: "Header"
};
</script>

<style>
.btnHeader {
  background-color: #5cbbf6;
  font-family: "Average Sans", sans-serif;
  border-radius: 100px;
  width: 130px;
  height: 45px;
  opacity: 1;
  transition: 0.3s;
  font-size: 12px;
  text-transform: uppercase;
  color: white;
  box-shadow: 0 0 4px #999;
  cursor: pointer;
  outline: none;
}

.btnHeader:hover {
  background: #47a7f5 radial-gradient(circle, transparent 1%, #47a7f5 1%)
    center/15000%;
}

.btnHeader:active {
  background-color: #6eb9f7;
  background-size: 100%;
  transition: background 0s;
}

.text {
  color: black;
  margin-top: 20px;
  font-family: "Average Sans", sans-serif;
  font-size: 20px;
}
</style>
