<template>
  <v-container fluid class="main" id="signUp">
    <v-row style="margin-top:10vh" align="center" justify="center">
      <p class="textTitle">SIGN UP</p>
    </v-row>
    <v-row align="center" justify="center">
      <v-form ref="form" lazy-validation>
        <v-row align="center" justify="start">
          <div class="text_detail">
            <p>I'm</p>
          </div>
          <v-radio-group v-model="row" row>
            <v-radio label="Student" value="radio-1" />
            <v-radio label="Teacher" value="radio-2" />
          </v-radio-group>
        </v-row>
        <v-row align="center" justify="start">
          <div class="text_detail">
            <p>Sex</p>
          </div>
          <v-radio-group v-model="row" row>
            <v-radio label="Male" value="radio-1" />
            <v-radio label="Female" value="radio-2" />
          </v-radio-group>
        </v-row>
        <!--name-->
        <v-row align="center" justify="start">
          <v-col>
            <label>Frist name:</label>
            <div class="input_button">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
          <v-col>
            <label>Family name:</label>
            <div class="input_button">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
        </v-row>
        <v-row align="center" justify="start">
          <v-col>
            <label>Birthday:</label>
            <div class="input_button">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
          <v-col>
            <label>Education:</label>
            <v-select
              :items="items"
              label="Choose education"
              solo
              rounded
              outlined
            ></v-select>
          </v-col>
        </v-row>
        <v-row align="center" justify="start">
          <v-col>
            <label>Email:</label>
            <div class="input_button">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
          <v-col>
            <label>Comfrim email:</label>
            <div class="input_button">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
        </v-row>
        <v-row align="center" justify="start">
          <v-col col="6">
            <div>
              <label class="">Password:</label>
            </div>
            <div class="input_button">
              <v-text-field
                color="primary"
                solo
                rounded
                outlined
                v-model="password"
                :type="showPassword ? 'text' : 'password'"
              />
            </div>
          </v-col>
          <v-col col="6">
            <label>Comfrim Password:</label>
            <div class="input_button">
              <v-text-field
                color="primary"
                solo
                rounded
                outlined
                v-model="password"
                :type="showPassword ? 'text' : 'password'"
              />
            </div>
          </v-col>
        </v-row>
        <router-link to="/signupsuccessfully">
          <button class="button1">Sign Up</button>
        </router-link>
      </v-form>
    </v-row>
    <div class="d-flex flex-column justify-bottom align-center">
      <v-img
        lazy-src="../assets/imgbit.png"
        max-height="328"
        max-width="1290"
        style="margin-top:20px"
        src="../assets/imgbit.png"
      ></v-img>
    </div>
  </v-container>
</template>

<script>
export default {
  name: "signUp",
  components: {},
  data: () => ({
    items: ["A", "B", "C", "D"]
  })
};
</script>

<style>
p.textTitle {
  font-weight: bold;
  font-size: 46px;
  color: #5c5c5c;
  font-family: "Average Sans", sans-serif;
}

.text_detail {
  margin-top: 15px;
  margin-left: 24vw;
  margin-right: 20px;
}

.input_button {
  width: 242px;
  height: 20px;
  margin: 20px;
}

button {
  background-color: #5cbbf6;
  border: none;
  color: white;
  padding: 15px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 15px;
  margin: 4px 2px;
  cursor: pointer;
  border-radius: 20px;
}

.button1 {
  background-color: #5cbbf6;
  border: none;
  color: white;
  padding: 15px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 15px;
  margin: 4px 2px;
  cursor: pointer;
  border-radius: 20px;
  margin-top: 15px;
  margin-bottom: 100px;
}

div {
  font-family: "CSChatThaiUI";
  text-align: center;
}

body {
  background-color: #ecede8;
}
</style>
