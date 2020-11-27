<template>
  <v-container fluid class="main" id="Home">
    <!-- home btn -->
      <v-row class="uploadbtn">
        <v-btn
          color="white"
          class="text-none"
          style="height: 85px; border-radius: 30px; font-size: 24px;"
          elevation="5"
          rounded
          depressed
          @click="onClickUpload"
        >
        <v-icon left> add </v-icon>
          Upload video
        </v-btn>
      </v-row>
    <!-- My video -->
    <v-row>
      <v-col cols="6">
        <div style="margin-left:50px">
          <p style="font-size: 50px" class="textTitle">
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
        <v-hover v-for="n in 3" :key="n" v-slot="{ hover }">
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
      </div>
    </v-row>

    <v-row align="center" justify="center" style="margin-top:100px;margin-bottom:70px">
      <v-card class="cardProfile">
          <v-img
              class="white--text align-end"
              height="443"
              width="1010"
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
    model: null,
    math: "https://sv1.picz.in.th/images/2020/11/19/bvfryV.png",
    computer: "https://sv1.picz.in.th/images/2020/11/19/bvfryV.png",
    langauge: "https://cdn.pixabay.com/photo/2020/07/12/07/47/bee-5396362_1280.jpg",
    econ: "https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
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
.uploadbtn {
  margin-top: 60px;
  margin-left: 60px;
}

.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
}

.cardProfile {
  height: 443px;
  width: 1010px;
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
  font-family: "Average Sans", sans-serif;
}

.cardSmallContainer {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  width: 80vw;
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

</style>
