<template>
  <v-container fluid class="main" id="LearnCourseTutorPage">
    <v-row align="center" justify="start" style="margin-top:40px">
        <v-hover v-slot="{ hover }">
          <v-btn
            depressed
            height="60px"
            width="60px"
            :elevation="hover ? 0 : 0"
            :class="{ 'on-hover-review': hover }"
            class="btnBack"
            @click="onClickBack()"
            >
              <span class="material-icons md-48">
                keyboard_arrow_left
              </span>
          </v-btn>
        </v-hover>
      <div class=" mt-4 mr-4 pd-6 pl-12">
        <p class="headtext">{{ title }}</p>
      </div>
    </v-row>
    <!-- Card group of tutor -->
    <v-row justify="center">
      <div class="tableCard">
        <div class="colCard" v-for="video in videos" :key="video.index">
          <v-sheet class="cardCoruseTitle">
            <p class="cardInSmallText">Teacher : {{ video.tutor }}</p>
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
                        <p class="cardInSmallText">{{ video.view }} Views</p>
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

<script scope>
// @ is an alias to /src

export default {
  name: "LearnCourseTutorPage",
  components: {},
  mounted () {
    this.$store.commit("SET_DIALOG_LOADING", true)
    this.title = this.$route.query.titleName;
    this.$store.commit("SET_DIALOG_LOADING", false)
  },
  data: () => ({
    math: "https://cdn.vuetifyjs.com/images/cards/cooking.png",
    title: null,
    videos: []
  }),
  methods: {
    onClickTutor () {
      this.$router.push({ name: "Video" })
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
