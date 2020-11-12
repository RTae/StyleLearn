<template>
  <v-container fluid class="main" id="signUpSec">
    <v-row align="center" justify="center" style="margin-top:20vh">
      <v-card class="bgcard" width="461" height="184"  >
        <v-row><label class="text">Successfully Register</label></v-row>
        <v-row><label class="textdetail">Hope you will find your style</label></v-row>
      </v-card>
    </v-row>
    <router-link to="/login"
          ><button class="button">Go to Login</button></router-link
          >
    <div class="d-flex flex-column justify-bottom align-center">
      <v-img
        lazy-src="../assets/imgbit.png"
        max-height="328"
        max-width="1290"
        style="margin-top:80px"
        src="../assets/imgbit.png"
      ></v-img>
    </div>
  </v-container>
</template>

<script>
export default {
  name: "signUpSec"
};
</script>

<style scoped>
.bgcard {
  color: white;
  border-radius:20px;
}
.text {
  font-weight: bold;
  font-size: 30px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 30px;
  margin-left: 110px;
  margin-bottom: 20px;
}
.textdetail {
  font-size: 18px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 10px;
  margin-left: 140px;
}

.button {
  font-family: "Average Sans", sans-serif;
  border-radius: 100px;
  margin-right: 20px;
  width: 130.5px;
  height: 45px;
  opacity: 1;
  transition: 0.3s;
  font-size: 11px;
  text-transform: uppercase;
  color: white;
  box-shadow: 0 0 4px #999;
  cursor: pointer;
  outline: none;
  margin-top: 40px;
}
body {
  background-color: #ecede8;
}
</style>
