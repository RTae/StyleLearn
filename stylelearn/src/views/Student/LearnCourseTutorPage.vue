<template>
  <v-container fluid class="main" id="LearnCourseTutorPage">
    <v-row align="center" justify="start" style="margin-top:40px">
      <v-col align="start" justify="center" class="titleHead" col="12">
        <v-btn
          depressed
          height="60px"
          width="60px"
          class="btnBack"
          @click="onClickBack()"
        >
            <span class="material-icons md-48">
              keyboard_arrow_left
            </span>
        </v-btn>
        <p class="headtext">{{ title }}</p>
      </v-col>
    </v-row>
    <!-- Card group of tutor -->
    <v-row justify="center">
      <div class="tableCard">
        <div class="colCard" v-for="video in videos" :key="video.VideoID">
          <v-hover v-slot="{ hover }">
            <v-card
              :elevation="hover ? 8 : 16"
              :class="{ 'on-hover': hover }"
              class="cardCourseSmall"
              @click="onClickTutor(video.VideoID)"
              :id= video.VideoID
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
    <v-dialog v-model="dialog" width="500">
      <v-card>
        <v-card-title class="primary mb-6"> Alert </v-card-title>
        <v-card-text class="textDetail">
          Sorry we don't have video right now
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="primary"
            text
            @click="onClickDialog"
          >
            OK
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import api from "../../service/api";
export default {
  name: "LearnCourseTutorPage",
  components: {},
  async mounted () {
    this.$store.commit("SET_DIALOG_LOADING", true)
    var lessonID = this.$route.params.id;
    if (lessonID === undefined) {
      var myCourseLessonID = localStorage.getItem("myCourseLessonID")
      if (myCourseLessonID !== null) {
        var result = await api.getAllVideoByLessonID(myCourseLessonID)
        if (result.data.status === "1") {
          this.videos = result.data.result
          this.$store.commit("SET_DIALOG_LOADING", false)
        }
      } else {
        this.$router.push({ name: "MyCourse" })
      }
    } else {
      localStorage.setItem("myCourseLessonID", lessonID);
      result = await api.getAllVideoByLessonID(lessonID)
      if (result.data.status === "1") {
        this.videos = result.data.result
        this.$store.commit("SET_DIALOG_LOADING", false)
      }
    }
    if (result.data.status === "1") {
      this.videos = result.data.result
      this.$store.commit("SET_DIALOG_LOADING", false)
    } else {
      this.$store.commit("SET_DIALOG_LOADING", false)
    }
    if (this.videos === null) {
      this.dialog = true
    } else {
      this.title = this.videos[0].CourseName + " : " + this.videos[0].LessonName
    }
  },
  data: () => ({
    title: null,
    dialog: false,
    videos: []
  }),
  methods: {
    onClickTutor (videoID) {
      this.$router.push({
        name: "Video",
        params: { id: videoID }
      });
    },
    onClickDialog () {
      this.dialog = false
      this.$router.push({ name: "MyCourse" })
    },
    onClickBack () {
      this.$router.go(-1)
    }
  }
};
</script>

<style scoped>
.main {
  background-color: rgb(239, 239, 239);
  min-height: 100vh;
}
.titleHead {
  display: flex;
  justify-content: flex-start;
  align-content: flex-start;
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

.headtext {
  font-weight: normal;
  color: black;
  font-size: 50px;
  font-family: "Average Sans", sans-serif;
  margin-left: 30px;
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

.cardInSmallText {
  font-family: "THSarabunNew";
  font-size: 20px;
  font-weight: bold;
}

.cardCourseSmall {
  margin-top: 30px;
  margin-bottom: 10px;
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
