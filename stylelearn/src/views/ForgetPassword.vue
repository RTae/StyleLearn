<template>
  <v-container fluid class="main" id="forgetPassword">
        <v-row align="center" justify="center" style="margin-top:20vh"
          ><label class="text"
            >Please enter the registered email address to reset your
            password</label
          ></v-row
        >
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
          </v-form>
        </v-row>
        <router-link to="/forgetpasswordpop"
          ><button class="mybtn">Submit</button></router-link
        >
  </v-container>
</template>

<script>
export default {
  name: "forgetPassword"
};
</script>

<style scoped>
.input_button {
  width: 417px;
  height: 50px;
  margin: 20px;
}
.text {
  font-weight: bold;
  font-size: 20px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 30px;
  margin-bottom: 20px;
}
.textdetail {
  font-size: 12px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 10px;
}

.mybtn {
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
  margin-top: 30px;
}
body {
  background-color: #ecede8;
}
</style>
