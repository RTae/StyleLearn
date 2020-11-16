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
      <v-row align="center" justify="center">
        <v-toolbar-items >
          <div class="input_serach">
            <v-text-field
              label="Serach"
              append-icon="mdi-magnify"
              solo
              dense
              outlined
              flat
            ></v-text-field>
          </div>
          <v-btn text><router-link to="/"><p class="text">Home</p></router-link></v-btn>
          <v-btn text><router-link to="/basicuse"><p class="text">Basic User</p></router-link></v-btn>
          <v-btn text><router-link to="/about"><p class="text">About Us</p></router-link></v-btn>
          <router-link to="/signup"><button class="btnHeader" >Sign Up</button></router-link>
          <router-link to="/login"><button class="btnHeader">Login In</button></router-link>
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
.input_serach{
  width: 600px;
  height: 10px;
  margin-top: 10px;
  margin-right: 20px;
}
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
  margin-top: 10px;
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
