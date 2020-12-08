<template>
  <v-container fluid class="main" id="Home">
    <!-- Upload button -->
    <v-row class="uploadbtn">
      <v-btn
        color="white"
        class="text-none"
        style="height: 85px; border-radius: 30px; font-size: 24px;"
        elevation="5"
        rounded
        depressed
        @click="onClickUpload"
        name=upload
      >
      <v-icon left> add </v-icon>
        Upload video
      </v-btn>
    </v-row>
    <!-- My video -->
    <v-row>
      <v-col cols="6">
        <div style="margin-left:50px">
          <p style="font-size: 40px" class="textTitle">
            My Video
          </p>
        </div>
      </v-col>
      <v-col cols="6">
        <router-link to="/myvideotutor">
          <div style="margin-right: 50px;">
            <p
              style="
                display: flex;
                flex-direction: row-reverse;
                margin-top: 60px;
              "
              class="textDetailAll"
            >
              All
            </p>
          </div>
        </router-link>
      </v-col>
    </v-row>

    <v-row justify="center">
      <div class="tableCard">
        <div class="colCard" v-for="video in videos" :key="video.index">
          <v-hover v-slot="{ hover }">
            <v-card :class="{ 'on-hover': hover }" elevation="8" class="cardCourseSmall">
              <v-img
                class="imgCard"
                :src="videoImage"
              />
              <v-row>
                <v-col class="d-flex pa-5 flex-row justify-space-between">
                  <p class="cardInSmallText">Lesson : {{ video.lesson_name }}</p>
                  <p class="cardInSmallText">{{ video.view }} Views</p>
                </v-col>
              </v-row>
            </v-card>
          </v-hover>
        </div>
      </div>
    </v-row>

    <!-- My Profile short -->
    <v-row align="center" justify="center" style="margin-top:100px;margin-bottom:70px">
      <v-card class="d-flex justify-center pa-10" style="border-radius:20px">
        <v-card
            class="d-flex justify-center"
            style="border-radius:10px; border:2px solid black;"
            max-height="300"
            min-height="200"
            max-width="400"
            min-width="200"
            color="grey darken-1"
          >
            <v-img
              :lazy-src="require('../../assets/image/etc/lazy_loading.png')"
              contain
              :src="user.image"
            />
        </v-card>
        <div class="d-flex flex-column justify-center ml-10">
          <p class="textDetail">{{ this.user.firstName }}  {{ this.user.familyName }}</p>
          <p class="textDetail" style="font-size: 15px;">{{ this.user.bio }}</p>
        </div>
      </v-card>
    </v-row>
  </v-container>
</template>

<script>
import api from "../../service/api"
import { server } from "../../service/constants";
export default {
  name: "Home",
  components: {},
  async mounted () {
    const result = await api.getUser(localStorage.getItem(server.USERNAME))
    if (result.data.status === "1") {
      this.user.image = result.data.result[0].ProfilePic
      this.user.firstName = result.data.result[0].Firstname
      this.user.familyName = result.data.result[0].Familyname
      this.user.bio = result.data.result[0].Bio
    }
  },
  data: () => ({
    videos: [
      {
        index: 1,
        subject_name: "Mathematics",
        lesson_name: "Diff I",
        view: 123
      },
      {
        index: 2,
        subject_name: "Mathematics",
        lesson_name: "Diff II",
        view: 6969
      },
      {
        index: 3,
        subject_name: "Mathematics",
        lesson_name: "Cal III",
        view: 532
      }
    ],
    model: null,
    videoImage: "https://cdn.vuetifyjs.com/images/cards/sunshine.jpg",
    user: {
      image: "",
      firstName: "",
      familyName: "",
      bio: ""
    }
  }),
  methods: {
    onClickCours () {
      this.$router.push("/myvideotutor");
    },
    onClickSubject (n) {
      this.model = n - 1;
      this.$router.push("/myvideotutor");
    },
    onClickUpload () {
      this.$router.push({ name: "UploadVideoTutor" });
    }
  }
};
</script>

<style scoped>
.main {
  background: rgb(239, 239, 239);
  min-height: 80vh;
}

.uploadbtn {
  margin-top: 60px;
  margin-left: 60px;
}

.textTitle {
  font-size: 96px;
  font-family: "Average Sans", sans-serif;
  align-content: center;
  justify-content: center;
  margin-top: 40px;
}

.textDetailAll {
  font-weight: normal;
  color: black;
  font-size: 30px;
  font-family: Delius;
}

.tableCard {
  display: grid;
  grid-template-columns: 30% 30% 30%;
  width: 100%;
  justify-content: space-around;
}

.colCard {
  margin-top: 20px;
  margin-bottom: 30px;
  display: flex;
  align-items: center;
  flex-direction: column;
}

.cardCoruseTitle {
  display: flex;
  justify-items: center;
  width: 303px;
  background-color: rgb(239, 239, 239);
}

.cardCourseSmall {
  border-radius: 10px;
  width: 303px;
  height: 264px;
  background-color: white;
  opacity: 0.6;
  background-color: #70CCFF;
  transition: opacity 0.2s ease-in;
}

.cardCourseSmall:not(.on-hover) {
  opacity: 1;
}

.cardInSmallText {
  font-family: "THSarabunNew";
  font-size: 20px;
  font-weight: bold;
}

.imgCard {
  height: 200px;
  width: 303px;
}

.textDetail{
  font-size: 35px;
  font-family: "Delius"
}

a {
  text-decoration: none;
}

</style>
