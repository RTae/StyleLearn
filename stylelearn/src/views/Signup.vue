<template>
  <v-container fluid class="main" id="signUp">
    <v-row style="margin-top:10vh" align="center" justify="center">
      <p class="textTitle">SIGN UP</p>
    </v-row>
    <v-row align="center" justify="center">
      <v-form ref="form" lazy-validation>
        <v-row class="d-flex flex-row align-center" flat height="100">
          <div style="margin-top:10px">
            <p>I am:</p>
          </div>
          <v-radio-group v-model="row" mandatory>
            <v-radio label="Student" value="student" />
            <v-radio label="Tutor" value="tutor" />
          </v-radio-group>
        </v-row>
        <v-row class="d-flex flex-row align-center" flat height="100">
          <div style="margin-top:10px">
            <p>Sex:</p>
          </div>
          <v-radio-group v-model="row" mandatory>
            <v-radio label="Female" value="female" />
            <v-radio label="Male" value="male" />
          </v-radio-group>
        </v-row>
        <!--name-->
        <v-row class="d-flex flex-row align-center">
          <v-col>
            <div class="input_button">
              First name:<v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
          <v-col>
            <div class="input_button">
              Family name:<v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
        </v-row>
        <v-row class="d-flex flex-row align-center">
          <v-col>
            <div class="input_button">
              Birthday:<v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
          <v-col>
            Education:<v-select
              :items="items"
              label="Solo field"
              solo
              rounded
              outlined
            ></v-select>
          </v-col>
        </v-row>
        <v-row class="d-flex flex-row align-center">
          <v-col>
            <div class="input_button">
              Email:<v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
          <v-col>
            <div class="input_button">
              Comfrim Email:<v-text-field
                color="primary"
                solo
                rounded
                outlined
              />
            </div>
          </v-col>
        </v-row>
        <v-row class="d-flex flex-row align-center">
          <v-col col="6">
            <div class="input_button">
              Password:<v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
          <v-col col="6">
            <div class="input_button">
              Confrim Password:<v-text-field
                color="primary"
                solo
                rounded
                outlined
              />
            </div>
          </v-col>
        </v-row>
        <button class="button">Sign Up</button>
      </v-form>
    </v-row>
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
div {
  font-family: "CSChatThaiUI";
  text-align: center;
}
h1 {
  text-align: center;
  font-weight: bold;
  font-size: 50px;
}
h2 {
  font-weight: normal;
  font-size: 15px;
}
li {
  margin: 0;
  padding: 0;
  color: #5cbbf6;
  display: inline;
  margin-right: 10px;
}
body {
  background-color: #ecede8;
}
</style>
