<template>
  <div id="header">
    <v-app-bar app color="secondary" height="70">
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="../../assets/image/main/logo.png"
          transition="scale-transition"
          width="200"
        />
      </div>
      <v-spacer></v-spacer>
      <v-row align="center" justify="end">
        <v-toolbar-items >
          <v-btn text name=home><router-link to="/"><p class="text">Home</p></router-link></v-btn>
          <v-btn text name=basicuse><router-link to="/basicuse"><p class="text">Basic Use</p></router-link></v-btn>
          <v-btn text name=about><router-link to="/about"><p class="text">About Us</p></router-link></v-btn>
          <router-link to="/signup"><button class="btnHeader" name=signup>Sign Up</button></router-link>
          <router-link to="/login"><button class="btnHeader" name=signin>LogIn</button></router-link>
        </v-toolbar-items>
      </v-row>
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
  font-size: 14px;
  text-transform: uppercase;
  color: white;
  box-shadow: 0 0 4px #999;
  cursor: pointer;
  outline: none;
  margin-right: 10px;
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
