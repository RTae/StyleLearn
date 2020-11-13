<template>
  <v-container fluid class="main" id="signUp">
    <v-row style="margin-top: 10vh" align="center" justify="center">
      <p class="textTitle">SIGN UP</p>
    </v-row>
    <v-row align="center" justify="center">
      <v-form ref="form" lazy-validation>
        <v-row align="center" justify="center">
          <div class="radioContainer">
            <div class="radioLineContainer">
              <div class="text_detail">
                <p>I'm</p>
              </div>
              <v-radio-group v-model="row" row>
                <v-radio label="Student" value="radio-1" />
                <v-radio label="Teacher" value="radio-2" />
              </v-radio-group>
            </div>
            <div class="radioLineContainer">
              <div class="text_detail">
                <p>Sex</p>
              </div>
              <v-radio-group v-model="row" row>
                <v-radio label="Male" value="radio-1" />
                <v-radio label="Female" value="radio-2" />
              </v-radio-group>
            </div>
          </div>
        </v-row>
        <!--name-->
        <v-row align="center" justify="center">
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Frist name:</label>
            </v-row>
            <div class="inputFiled">
                <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Family name:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
        </v-row>
        <!--birthday and education-->
        <v-row align="center" justify="start">
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Birthday:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Education:</label>
            </v-row>
            <div class="inputFiled">
              <v-select
                :items="items"
                label="Choose education"
                solo
                rounded
                outlined
              />
            </div>
          </v-col>
        </v-row>
        <v-row align="center" justify="start">
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Email:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
          <v-col>
            <!--email-->
            <v-row class="ml-8" justify="start">
              <label>Comfrim email:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
        </v-row>
        <!--password-->
        <v-row align="center" justify="start">
          <v-col col="6">
            <v-row class="ml-8" justify="start">
              <label class="">Password:</label>
            </v-row>
            <div>
            </div>
            <div class="inputFiled">
              <v-text-field
                color="primary"
                solo
                rounded
                outlined
                :type="'password'"
              />
            </div>
          </v-col>
          <v-col col="6">
            <v-row class="ml-8" justify="start">
              <label>Comfrim Password:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field
                solo
                rounded
                outlined
                :type="'password'"
              />
            </div>
          </v-col>
        </v-row>
        <v-row justify='center'>
          <router-link to="/signupsuccessfully">
            <button class="signUpBtn">SIGN UP</button>
          </router-link>
        </v-row>
      </v-form>
    </v-row>
    <div class="d-flex flex-column justify-bottom align-center">
      <v-img
          alt="bitButton"
          contain
          style="margin-top: 50px"
          src="../assets/imgbit.png"
          width="1290"
      />
    </div>
  </v-container>
</template>

<script scope>
export default {
  name: "signUp",
  components: {},
  data: () => ({
    items: ["A", "B", "C", "D"]
  })
};
</script>

<style scope>
.radioContainer {
  display: flex;
  flex-wrap: wrap;
  align-content: center;
  flex-direction: column;
}

.radioLineContainer{
  display: flex;
  flex-wrap: wrap;
  align-content: flex-start;
  flex-direction: row;
}

p.textTitle {
  font-weight: bold;
  font-size: 20px;
  color: #5c5c5c;
  font-family: "Average Sans", sans-serif;
}

label {
  font-weight: bold;
  font-size: 15px;
  color: #5c5c5c;
  font-family: "Average Sans", sans-serif;
}

.text_detail {
  margin-top: 20px;
  margin-right: 20px;
}

.inputFiled {
  width: 375px;
  height: 32px;
  margin-top: 20px;
  margin-left: 20px;
  margin-right: 20px;
  margin-bottom: 35px;
}

.signUpBtn {
  background-color: #5cbbf6;
  background-position: center;
  font-family: "Average Sans", sans-serif;
  border-radius: 100px;
  margin-right: 20px;
  width: 130px;
  height: 45px;
  opacity: 1;
  transition: 0.3s;
  font-size: 13px;
  text-transform: uppercase;
  color: white;
  box-shadow: 0 0 4px #999;
  cursor: pointer;
  outline: none;
  margin-top: 50px;
}

.signUpBtn:hover {
  background: #47a7f5 radial-gradient(circle, transparent 1%, #47a7f5 1%)
    center/15000%;
}
.signUpBtn:active {
  background-color: #6eb9f7;
  background-size: 100%;
  transition: background 0s;
}
</style>
