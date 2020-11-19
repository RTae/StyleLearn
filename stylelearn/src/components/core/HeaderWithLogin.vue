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
      <v-row class="hidden-sm-and-down" align="center" justify="end">
        <v-toolbar-items>
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
          <v-btn text
            ><router-link to="/"><p class="text">Home</p></router-link></v-btn
          >
          <v-btn text
            ><router-link to="/basicuse"
              ><p class="text">Basic User</p></router-link
            ></v-btn
          >
          <v-btn text
            ><router-link to="/about"
              ><p class="text">About Us</p></router-link
            ></v-btn
          >
          <router-link to="/signup"
            ><v-img class="iconBar" src="../../assets/shopping.png"
          /></router-link>
          <router-link to="/signup"
            ><v-img class="iconBar" contain src="../../assets/bell.png"
          /></router-link>
          <v-img class="iconBarLine" src="../../assets/line.png"></v-img>
          <router-link to="/signup"
            ><v-img class="iconBarPro" contain src="../../assets/profile.png"
          /></router-link>
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

<style scope>
.input_serach {
  width: 600px;
  height: 10px;
  margin-top: 10px;
  margin-right: 20px;
}
.iconBar {
  width: 34px;
  height: 34px;
  margin-left: 20px;
  margin-top: 8px;
}
.iconBarPro {
  width: 40px;
  height: 40px;
  margin-left: 20px;
  margin-right: 20px;
  margin-top: 8px;
}
.iconBarLine {
  width: 10px;
  height: 34px;
  margin-left: 20px;
  margin-top: 8px;
}

.text {
  color: black;
  margin-top: 20px;
  font-family: "Average Sans", sans-serif;
  font-size: 20px;
}
</style>
