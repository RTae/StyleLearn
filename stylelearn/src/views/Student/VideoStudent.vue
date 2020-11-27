<template>
  <v-container fluid class="main" id="LearnCourseTutorPage">
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
        <p class="textHead">Lesson : {{ video.lesson_name }}</p>
        <p class="textHead">Tutor : {{ video.user_firstname }}</p>
      </div>
    </v-row>
    <!-- Video -->
    <v-row justify="center" align="center">
      <div class="videoContainer">
        <vue-core-video-player name="video" src="https://media.vued.vanthink.cn/sparkle_your_name_am720p.mp4"></vue-core-video-player>
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
        <div class="colCard" v-for="course in courses" :key="course.index">
          <v-sheet class="cardCoruseTitle">
            <p class="cardInSmallText">Teacher : {{ course.tutor }}</p>
          </v-sheet>
          <v-hover v-slot="{ hover }">
            <v-card elevation="8" class="cardCourseSmall">
              <v-img
                :class="{ 'on-hover': hover }"
                class="imgCard"
                :src="math"
                @click="onClickTutor()"
              />
              <v-card-actions style="background-color: #70ccff; height:64px">
                <v-list-item class="grow">
                  <v-row>
                    <v-col  align="center" justify="end" cols="8" offset="4">
                        <p class="cardInSmallText">{{ course.view }} Views</p>
                    </v-col>
                  </v-row>
                </v-list-item>
              </v-card-actions>
            </v-card>
          </v-hover>
        </div>
      </div>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "LearnCourseTutorPage",
  mounted () {
    this.title = this.$route.params.titleName;
  },
  data: () => ({
    math: "https://cdn.vuetifyjs.com/images/cards/cooking.png",
    title: null,
    video: {
      user_firstname: "New",
      lesson_name: "Differntail I",
      description: "Differntail วิชาแห่งการลืนทไล"
    },
    courses: [
      {
        index: 1,
        tutor: "Cherprang",
        view: 15657
      },
      {
        index: 2,
        tutor: "June",
        view: 14757
      },
      {
        index: 3,
        tutor: "Pun",
        view: 13474
      },
      {
        index: 4,
        tutor: "Jennis",
        view: 12657
      },
      {
        index: 5,
        tutor: "Wee",
        view: 9157
      },
      {
        index: 6,
        tutor: "Mobile",
        view: 8157
      }
    ]
  }),
  methods: {
    onClickBack () {
      this.$router.push({ name: "TutorPage" })
    },
    onClickLike (n) {
      this.courses[n].likeState = !this.courses[n].likeState
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
  background-color: white;
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
