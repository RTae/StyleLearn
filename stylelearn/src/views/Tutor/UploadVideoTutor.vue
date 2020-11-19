<template>
  <v-container fluid class="main" id="login">
    <v-row style="margin-top: 10vh;margin-top:100px;margin-bottom:50px" align="center" justify="center">
      <v-card class="cardContainer" color=#70aaff>
        <p class="texttitle">Upload My Video</p>
      </v-card>
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
              <v-btn> + choose file </v-btn>
            </div>
          </v-col>
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Cover Video</label>
            </v-row>
            <div class="inputFiled">
              <v-btn> + choose file </v-btn>
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
        <v-row class="ml-8" justify="start">
          <label>Description</label>
        </v-row>
      </v-form>
    </v-row>
    <v-row justify="center">
      <div class="inputFiled2">
        <v-textarea color="primary" solo rounded outlined></v-textarea>
      </div>
    </v-row>
    <div div align="center" justify="center" style="margin-top:150px">
      <v-row justify="center">
        <router-link to="/myvideotutor">
          <button class="btnuploadvideo">Upload</button>
        </router-link>
      </v-row>
      <p class="txtuploaddetail">Send to a check specialist</p>
      <p class="txtuploaddetail2">
        *Please wait for a response within 2 days.
      </p>
    </div>
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
.texttitle {
  font-weight: bold;
  font-size: 46px;
  color: #f1f1f1;
  font-family: "Average Sans", sans-serif;
  margin-top: 20px;
}

.btnuploadvideo {
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
}
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
  width: 800px;
  height: 63px;
  margin-top: 20px;
}

.uploaddetail {
  font-size: 8px;
  color: #5c5c5c;
  font-family: "Average Sans", sans-serif;
  margin-top: 20px;
  margin-right: 3px;
}

.txtuploaddetail {
  font-size: 13px;
  margin-top: 20px;
}

.txtuploaddetail2 {
  font-size: 13px;
  color: red;
}

.cardContainer {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  height: 100px;
  width: 490px;
  background-color: #70aaff;
  border-radius: 50px;
}

.tutorCard {
  border-radius: 10px;
  background-color: #70CCFF;
  display: flex;
  flex-direction: row;
  text-align: start;
  align-items: center;
  margin-top: 20px;
  width: 383px;
  height: 238px;
}
</style>
