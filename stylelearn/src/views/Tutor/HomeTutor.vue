<template>
  <v-container fluid class="main" id="Home">
    <!-- home btn -->
    <v-container>
      <v-row>
        <router-link to="/uploadvideotutor">
          <v-btn>
            + Upload Video
          </v-btn>
        </router-link>
      </v-row>
    </v-container>
    <!-- My video -->
    <v-row>
      <v-col cols="6">
        <div style="margin-left:50px">
          <p style="font-size: 50px" class="textDetail">
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
                margin-top: 25px;
              "
              class="textDetail"
            >
              All
            </p>
          </div>
        </router-link>
      </v-col>
    </v-row>

    <v-row align="center" justify="center">
      <div class="cardSmallContainer">
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 16"
            :class="{ 'on-hover': hover }"
            class="cardCourseSmall"
            @click="onClickCourse()"
          >
            <v-img height="200" width="303" :src="math" />
            <v-sheet class="cardInSmallContainer">
              <p class="cardInSmallText">Calculus I</p>
            </v-sheet>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 12"
            :class="{ 'on-hover': hover }"
            class="cardCourseSmall"
            @click="onClickCourse()"
          >
            <v-img height="200" width="303" :src="math" />
            <v-sheet class="cardInSmallContainer">
              <p class="cardInSmallText">Calculus II</p>
            </v-sheet>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 12"
            :class="{ 'on-hover': hover }"
            class="cardCourseSmall"
            @click="onClickCourse()"
          >
            <v-img height="200" width="303" :src="math" />
            <v-sheet class="cardInSmallContainer">
              <p class="cardInSmallText">Calculus III</p>
            </v-sheet>
          </v-card>
        </v-hover>
      </div>
    </v-row>

    <v-row align="center" justify="center" style="margin-top:100px">
      <v-card class="cardProfile">
          <v-img
            class="white--text align-end"
              height="300"
              width="700"
            src="https://www.img.in.th/images/464684b598cb4e84e69d186f0f5c1478.png"
          >
          </v-img>
      </v-card>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Home",
  components: {},
  data: () => ({
    items: ["Mathematics", "Economics", "Chemistry", "Computer", "Electric"],
    model: null,
    math: "https://sv1.picz.in.th/images/2020/11/19/bvfryV.png",
    computer: "https://sv1.picz.in.th/images/2020/11/19/bvfryV.png",
    langauge:
      "https://cdn.pixabay.com/photo/2020/07/12/07/47/bee-5396362_1280.jpg",
    econ: "https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
  }),
  methods: {
    onClickCours () {
      this.$router.push("/myvideotutor");
    },
    onClickSubject (n) {
      this.model = n - 1;
      this.$router.push("/myvideotutor");
    }
  }
};
</script>

<style scoped>
.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
  margin-top: 50px;
}

.cardContainer {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  height: 280px;
  width: 100vw;
  background-color: #70ccff;
}
.courseCard {
  border-radius: 32px;
  width: 464px;
  height: 233px;
  background-color: white;
  display: flex;
  flex-direction: row;
  text-align: center;
  align-items: center;
  opacity: 0.6;
  transition: opacity 0.2s ease-in;
}

.courseCard:not(.on-hover) {
  opacity: 1;
}

.cardTextTitle {
  font-size: 72px;
  font-family: "Average Sans", sans-serif;
}

.cardProfile {
  height: 300px;
  width: 700px;
}

a {
  text-decoration: none;
}

.textTitle {
  font-size: 96px;
  font-family: "Average Sans", sans-serif;
  align-content: center;
  justify-content: center;
  margin-top: 40px;
}

.textDetail {
  font-weight: normal;
  color: black;
  font-size: 30px;
  font-family: Delius;
}

.cardSmallContainer {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  width: 100vw;
}

.cardCourseSmall {
  border-radius: 10px;
  width: 303px;
  height: 279px;
  background-color: white;
  opacity: 0.6;
  transition: opacity 0.2s ease-in;
}

.cardCourseSmall:not(.on-hover) {
  opacity: 1;
}

.cardInSmallContainer {
  background-color: #70ccff;
  display: flex;
  align-items: center;
  width: 303px;
  height: 79px;
}

.cardInSmallText {
  margin-left: 40px;
  font-family: "THSarabunNewRegular";
  font-size: 15px;
  font-weight: bold;
}

.subjectContainer {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  height: 474px;
  width: 100vw;
  background-color: #70ccff;
}

.subjectCard {
  height: 300px;
  width: 642px;
  border-radius: 50px;
  display: flex;
  flex-direction: row;
  align-items: center;
  text-align: center;
  opacity: 0.6;
  transition: opacity 0.2s ease-in;
}

.subjectCard:not(.on-hover) {
  opacity: 1;
}

.studentReviewContainner {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.textTitleStudnetReview {
  font-size: 40px;
  font-family: "Delius", cursive;
}

.textDetailStudnetReview {
  font-size: 18px;
  font-family: "Delius", cursive;
}

.cardReviewContainer {
  height: 343px;
  width: 373px;
  border-radius: 5px;
  opacity: 1;
  transition: opacity 0.2s ease-in;
}

.cardReviewContainer:not(.on-hover-review) {
  height: 343px;
  width: 373px;
  border-radius: 5px;
  opacity: 0.5;
}

.cardInReviewContainer {
  height: 243px;
  width: 373px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.cardInReviewNameContainer {
  background-color: #70ccff;
  height: 100px;
  width: 373px;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
