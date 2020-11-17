<template>
  <v-container fluid class="main" id="login">
    <v-row style="margin-top: 10vh" align="center" justify="center">
      <p class="textTitle">Upload My Video</p>
    </v-row>
    <v-row justify="center">
      <v-form ref="form" lazy-validation>
        <!--video-->
        <v-row align="center" justify="center">
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Video</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Cover Video</label>
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
              <label>Subject</label>
            </v-row>
            <div class="inputFiled">
              <v-select
                :items="items"
                label="Choose subject"
                solo
                rounded
                outlined
              />
            </div>
          </v-col>
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Course</label>
            </v-row>
            <div class="inputFiled">
              <v-select
                :items="items"
                label="Choose course"
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
              <label>Lesson</label>
            </v-row>
            <div class="inputFiled">
              <v-select
                :items="items"
                label="Choose lesson"
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
              <label>Description</label>
            </v-row>
            <div class="inputFiled2">
             <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
        </v-row>
        <v-row justify="center">
          <router-link to="/hometutor">
            <button class="signUpBtn">UPLOAD</button>
          </router-link>
        </v-row>
        <div div align="center" justify="center" style="margin-top:10px">
          <p class="txtuploaddetail">Send to a check specialist</p>
          <p class="txtuploaddetail2">*Please wait for a response within 2 days.</p>
        </div>
      </v-form>
    </v-row>
  </v-container>
</template>

<script scope>
export default {
  name: "Course",
  components: {},
  data: () => ({
    items: ["A", "B", "C", "D"]
  })
};
</script>

<style>
.radioContainer {
  display: flex;
  flex-wrap: wrap;
  align-content: center;
  flex-direction: column;
}

.radioLineContainer {
  display: flex;
  flex-wrap: wrap;
  align-content: flex-start;
  flex-direction: row;
}

.inputFiled {
  width: 375px;
  height: 32px;
  margin-top: 20px;
  margin-left: 20px;
  margin-right: 20px;
  margin-bottom: 35px;
}

.inputFiled2 {
  width: 375px;
  height: 32px;
  margin-top: 20px;
  margin-left: 20px;
  margin-right: 20px;
  margin-bottom: 35px;
}

.uploaddetail{
  font-size: 8px;
  color: #5c5c5c;
  font-family: "Average Sans", sans-serif;
  margin-top: 20px;
  margin-right: 3px;
}

.txtuploaddetail{
  font-size: 13px;
}

.txtuploaddetail2{
  font-size: 13px;
  color: red;
}

</style>
