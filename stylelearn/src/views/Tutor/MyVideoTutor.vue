<template>
  <v-container fluid class="main" id="MyCourse">
    <v-row align="center" justify="center" style="margin-top: 40px">
      <v-card elevation=10 class="cardContainer">
        <p class="cardTextTitle">My Video</p>
      </v-card>
    </v-row>
    <!-- Card -->
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
  </v-container>
</template>

<script scope>
// @ is an alias to /src

export default {
  name: "MyCourse",
  components: {},
  mounted () {
    this.title = this.$route.params.titleName;
  },
  data: () => ({
    videoImage: "https://cdn.vuetifyjs.com/images/cards/cooking.png",
    title: null,
    videos: [
      {
        index: 1,
        subject_name: "Mathematics",
        lesson_name: "Diff I",
        view: 123,
      },
      {
        index: 2,
        subject_name: "Mathematics",
        lesson_name: "Diff II",
        view: 12322,
      },
      {
        index: 3,
        subject_name: "Mathematics",
        lesson_name: "Diff III",
        view: 122,
      },
      {
        index: 4,
        lesson_name: "Geometric I",
        subject_name: "Mathematics",
        view: 221,
      },
      {
        index: 5,
        lesson_name: "Geometric II",
        subject_name: "Mathematics",
        view: 202,
      }
    ]
  }),
};
</script>

<style scoped>
.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
}

.cardContainer {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  height: 122px;
  width: 520px;
  background-color: #70ccff;
  border-radius: 30px;
}

.cardTextTitle {
  font-weight: normal;
  color: white;
  font-size: 70px;
  font-family: "Average Sans", sans-serif;
}
.cardCourseContainer {
  margin-bottom: 200px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
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
  justify-content: center;
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

.cardInSmallText {
  font-family: "THSarabunNew";
  font-size: 20px;
  font-weight: bold;
}

.imgCard {
  height: 200px;
  width: 303px;
}
</style>
