<template>
  <v-container fluid class="main" id="upload">
      <v-row style="margin-top:0vh" justify="center">
        <router-link to="/uploadvideotutor">
          <button class="UploadVideoBtn">+ Upload My Video</button>
        </router-link>
      </v-row>
      <v-row align="center" justify="center">
        <p class="homeTitle">My Video</p>
      </v-row>
      <v-row align="center" justify="center">
        <router-link to="/myvideotutor">
          <p> All</p>
        </router-link>
      </v-row>
  </v-container>
</template>

<script>
export default {

}
</script>

<style>
.UploadVideoBtn {
  background-color: #5cbbf6;
  font-family: "Average Sans", sans-serif;
  border-radius: 10px;
  margin-right: 1300px;
  margin-top: 100px;
  width: 140px;
  height: 45px;
  transition: 0.3s;
  font-size: 13px;
  color: white;
  box-shadow: 0 0 4px #999;
  cursor: pointer;
}

.homeTitle {
  font-size: 30px;
  font-weight: bold;
  font-family: "Average Sans", sans-serif;
  margin-right: 1300px;
  margin-top: 60px;
}

.text {
  font-size: 100px;
}

</style>
