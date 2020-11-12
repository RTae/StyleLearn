<template>
  <v-container fluid class="main" id="login">
    <v-row style="margin-top:10vh" align="center" justify="center">
      <p class="textTitle">LOGIN</p>
    </v-row>
    <v-row align="center" justify="center">
      <v-form ref="form" lazy-validation>
        <div class="input_button">
          <v-text-field
            color="primary"
            solo
            rounded
            outlined
            placeholder="Email"
          />
        </div>

        <div class="input_button">
          <v-text-field
            color="primary"
            solo
            rounded
            outlined
            placeholder="Password"
            v-model="password"
            :type="showPassword ? 'text' : 'password'"
            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            @click:append="showPassword = !showPassword"
          />
        </div>
        <v-row align="center" justify="center">
          <p style="margin-right:10px;" class="text">Don't have an account?</p>
          <router-link to="/signup"
            ><p style="margin-right:165px;" class="text">
              Sign up
            </p></router-link
          >
          <router-link to="/forgetpassword"
            ><p class="text">Forget password?</p></router-link
          >
        </v-row>
        <button class="button">Sign In</button>
      </v-form>
    </v-row>
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
  name: "login"
};
</script>

<style>
.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
}

.input_button {
  width: 423px;
  height: 60px;
  margin: 20px;
}

.textTitle {
  font-weight: bold;
  font-size: 46px;
  color: #5c5c5c;
  font-family: "Average Sans", sans-serif;
  margin-top: 20px;
}

.text {
  font-size: 12px;
  font-family: "Average Sans", sans-serif;
}

.button {
  background-position: center;
  font-family: "Average Sans", sans-serif;
  border-radius: 100px;
  margin-right: 20px;
  width: 130px;
  height: 45px;
  opacity: 1;
  transition: 0.3s;
  font-size: 10px;
  text-transform: uppercase;
  color: white;
  box-shadow: 0 0 4px #999;
  cursor: pointer;
  outline: none;
  margin-top: 50px;
}

.button:hover {
  background: #47a7f5 radial-gradient(circle, transparent 1%, #47a7f5 1%)
    center/15000%;
}
.button:active {
  background-color: #6eb9f7;
  background-size: 100%;
  transition: background 0s;
}
</style>
