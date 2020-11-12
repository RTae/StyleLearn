<template>
  <v-container fluid class="main" id="forgetPassword">
    <v-row align="center" justify="start" style="margin-top:20vh">
      <label class="text_password">New password</label>
    </v-row>
    <v-row align="center" justify="center">
      <div class="input_button">
        <v-text-field
          solo
          rounded
          outlined
          placeholder="Password"
          v-model="password"
          :type="showPassword ? 'text' : 'password'"
        />
      </div>
    </v-row>
    <v-row align="center" justify="start" style="margin-top:10vh">
      <label class="text_password">Confirm new Password</label>
    </v-row>
    <v-row align="center" justify="center">
      <div class="input_button">
        <v-text-field
          solo
          rounded
          outlined
          placeholder="Password"
          v-model="password"
          :type="showPassword ? 'text' : 'password'"
        />
      </div>
    </v-row>
    <router-link to="/forgetpasswordpop"><button class="button">Confirm</button></router-link>
  </v-container>
</template>

<script>
export default {
  name: "forgetPassword"
};
</script>

<style scoped>
.text_password {
  font-weight: bold;
  font-size: 20px;
  color: #5c5c5c;
  font-family: "TH Sarabun New";
  margin-left: 80vh;
  margin-top: 10px;
}

.input_button {
  width: 320px;
  height: 12px;
  margin: 10px;
  font-size: 12px;
}

.button {
  font-family: "Average Sans", sans-serif;
  border-radius: 100px;
  margin-right: 20px;
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
  margin-top: 90px;
}
body {
  background-color: #ecede8;
}
</style>
