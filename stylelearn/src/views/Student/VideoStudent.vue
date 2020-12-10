<template>
  <v-container fluid class="main" id="VideoPage">
    <!-- Title -->
    <v-row align="center" justify="start" style="margin-top:40px">
        <v-hover v-slot="{ hover }">
          <v-btn
            depressed
            height="60px"
            width="60px"
            :elevation="hover ? 0 : 0"
            :class="{ 'on-hover-review': hover }"
            @click="onClickBack()"
            class="btnBack"
            >
            <span class="material-icons md-48">
              keyboard_arrow_left
            </span>
          </v-btn>
        </v-hover>
      <div class="d-flex flex-column justify-start ml-10">
        <p class="textHead">Lesson : {{ video.lessonName | Capitalize }}</p>
        <p class="textHead">Tutor : {{ video.name }}</p>
      </div>
    </v-row>
    <!-- Video -->
    <v-row justify="center" align="center">
      <div class="videoContainer">
        <vue-core-video-player v-if="video.video" name="video" :src=video.video />
      </div>
    </v-row>
    <!-- Description -->
    <v-row>
      <v-col cols="9" offset="1">
        <p class="textDescription">Description : {{ video.description }}</p>
      </v-col>
    </v-row>
    <!-- Card group of tutor -->
    <v-row justify="center">
      <div class="tableCard">
        <div class="colCard" v-for="video in vidoes" :key="video.VideoID">
          <v-hover v-slot="{ hover }">
            <v-card
              :elevation="hover ? 8 : 16"
              :class="{ 'on-hover': hover }"
              class="cardCourseSmall"
              @click="onClickTutor(video.VideoID)"
            >
              <v-img
                height="200"
                width="303"
                src="../../assets/image/video/video.png"
              />
              <v-row>
                <v-col class="d-flex pa-5 flex-row justify-space-between">
                  <p class="cardInSmallText">Tutor:  {{ video.Name }}</p>
                </v-col>
              </v-row>
            </v-card>
          </v-hover>
        </div>
      </div>
    </v-row>
  </v-container>
</template>

<script>
import api from "../../service/api"
export default {
  name: "VideoPage",
  async mounted () {
    this.$store.commit("SET_DIALOG_LOADING", true)
    var videoID = this.$route.params.id
    if (videoID === undefined) {
      var myCourseVideoID = localStorage.getItem("myCourseVideoID")
      if (myCourseVideoID !== null) {
        var result = await api.getShowVideoByVideoID(myCourseVideoID)
        if (result.data.status === "1") {
          this.video.name = result.data.result[0].Name
          this.video.lessonName = result.data.result[0].LessonName
          this.video.description = result.data.result[0].Description
          this.video.video = result.data.result[0].Video
          var resultVideo = await api.getAllVideoByLessonID(result.data.result[0].LessonID)
          if (resultVideo.data.status === "1") {
            this.vidoes = resultVideo.data.result
            this.$store.commit("SET_DIALOG_LOADING", false)
          } else {
            this.video = []
            this.$store.commit("SET_DIALOG_LOADING", false)
          }
        }
      } else {
        this.$store.commit("SET_DIALOG_LOADING", false)
        this.$router.push({ name: "MyCourse" })
      }
    } else {
      localStorage.setItem("myCourseVideoID", videoID);
      const result = await api.getShowVideoByVideoID(videoID)
      if (result.data.status === "1") {
        this.video.name = result.data.result[0].Name
        this.video.lessonName = result.data.result[0].LessonName
        this.video.description = result.data.result[0].Description
        this.video.video = result.data.result[0].Video
        resultVideo = await api.getAllVideoByLessonID(result.data.result[0].LessonID)
        if (resultVideo.data.status === "1") {
          this.vidoes = resultVideo.data.result
          this.$store.commit("SET_DIALOG_LOADING", false)
        } else {
          this.video = []
          this.$store.commit("SET_DIALOG_LOADING", false)
        }
      } else {
        this.$router.push({ name: "MyCourse" })
        this.$store.commit("SET_DIALOG_LOADING", false)
      }
    }
  },
  data: () => ({
    video: {
      name: "",
      lessonName: "",
      description: "",
      video: false
    },
    vidoes: []
  }),
  filters: {
    Capitalize (value) {
      if (typeof value !== "string") return ""
      return value.charAt(0).toUpperCase() + value.slice(1)
    }
  },
  methods: {
    onClickBack () {
      this.$router.go(-1)
    },
    onClickTutor () {
      console.log("Tutor")
    }
  }
};
</script>

<style scoped>
.main {
  background-color: rgb(239, 239, 239);
  min-height: 100vh;
}
.btnBack {
  background-color: white;
  border-radius: 100px;
  margin-left: 50px;
}
.btnBack:hover {
  color: #47a7f5;
  background-color: white;
  outline: none;
  opacity: 1;
}

.textHead {
  font-weight: normal;
  color: black;
  font-size: 40px;
  font-family: "Average Sans", sans-serif;
}

.videoContainer{
  width: 80vw;
}

.textDescription {
  font-size: 35px;
  color: black;
  font-family: "THSarabunNew";
}

.tableCard {
  display: grid;
  grid-template-columns: 30% 30% 30%;
  width: 80%;
  justify-content: space-around;
}

.colCard {
  margin-top: 20px;
  margin-bottom: 30px;
  display: flex;
  flex-direction: column;
}

.cardCoruseTitle {
  display: flex;
  justify-items: center;
  width: 303px;
  background-color: rgb(239, 239, 239);
}

.cardInSmallText {
  font-family: "THSarabunNew";
  font-size: 20px;
  font-weight: bold;
}

.cardCourseSmall {
  margin-bottom: 50px;
  border-radius: 10px;
  width: 303px;
  height: 264px;
  background-color: #6eb9f7;
  transition: opacity 0.2s ease-in;
}

.imgCard {
  height: 200px;
  width: 303px;
  opacity: 0.6;
  transition: opacity 0.2s ease-in;
}

.imgCard:not(.on-hover) {
  opacity: 1;
}
</style>
