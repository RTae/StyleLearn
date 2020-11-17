<template>
  <v-container fluid class="main" id="upload">
    <!--upload video button-->
      <router-link to="/uploadvideotutor">
        <button class="UploadVideoBtn">+ Upload My Video</button>
      </router-link>
    <!--link-->
    <v-row justify="center">
      <div style="margin-right:20px;"><router-link to="/hometutor">home</router-link></div>
      <div style="margin-right:20px;"><router-link to="/uploadvideotutor">upload</router-link></div>
      <div style="margin-right:20px;"><router-link to="/myvideotutor">my video</router-link></div>
      <router-link to="/profiletutor">
        <div>profile</div></router-link>
    </v-row>
    <!--title-->
    <v-row align="center" justify="center">
      <p class="homeTitle">My Video</p>
    </v-row>
    <v-row align="center" justify="center">
      <router-link to="/myvideotutor">
        <p class="text2">All</p>
      </router-link>
    </v-row>
  <!--course picture 1-->
    <v-row>
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:220px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://sv1.picz.in.th/images/2020/11/16/bHaFJ2.png"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            16,215 Views
          </v-btn>
        </v-card>
      </v-col>
      <!--course picture 2-->
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:220px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://sv1.picz.in.th/images/2020/11/16/bHaFJ2.png"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            14,1756 Views
          </v-btn>
        </v-card>
      </v-col>
      <!--course picture 3-->
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:220px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://sv1.picz.in.th/images/2020/11/16/bHaFJ2.png"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            19,125 Views
          </v-btn>
        </v-card>
      </v-col>
    </v-row>
    <v-row align="center" justify="center">
      <p class="text3">Lesson Name: Differential I</p>
    </v-row>
    <v-row>
      <v-card
        justify="center"
        width="400px"
        height="250px"
        style="center"
        border-radius="40px"
        color="#70CCFF"
      >
      </v-card>
    </v-row>
  </v-container>
</template>

<script>
export default {};
</script>

<style scope>
.UploadVideoBtn {
  background-color: #5cbbf6;
  font-family: "Average Sans", sans-serif;
  border-radius: 10px;
  margin-right: 1300px;
  margin-top: 100px;
  width: 150px;
  height: 55px;
  transition: 0.3s;
  font-size: 14px;
  color: white;
  box-shadow: 0 0 4px #999;
  cursor: pointer;
}

.homeTitle {
  font-size: 40px;
  font-weight: bold;
  font-family: "Average Sans", sans-serif;
  margin-right: 1200px;
  margin-top: 60px;
}

.text {
  font-size: 13px;
}

.text2 {
  font-size: 25px;
  margin-left: 1200px;
}

.text3 {
  font-size: 13px;
  margin-right: 780px;
}

.botton {
  display: block;
  border-style: none;
  position: absolute;
  box-shadow: none;
  background-repeat: no-repeat;
}

li {
  margin-right: 10px;
}
</style>
