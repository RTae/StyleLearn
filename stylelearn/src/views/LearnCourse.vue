<template>
  <v-container fluid class="main" id="MyCourse">
    <v-row align="center" justify="start" style="margin-top:40px">
      <v-col align="end" justify="center" col="1">
        <v-hover v-slot="{ hover }">
          <v-btn
              depressed
              height="60px"
              width="60px"
              :class="{ 'on-hover-review': hover }"
              class="btnBack material-icons"
              @click="onClickBack()"
          >
            <v-icon>
              keyboard_arrow_left
            </v-icon>
          </v-btn>
        </v-hover>
      </v-col>
      <v-col align="start" justify="center" cols="11">
        <p class="headtext">CALCULUS I</p>
      </v-col>
    </v-row>
    <!-- Card -->
    <v-row justify="center">
      <div class="tableCard">
        <div class="colCard" v-for="course in courses" :key="course.index">
          <v-hover v-slot="{ hover }">
            <v-card
              :elevation="hover ? 8 : 12"
              :class="{ 'on-hover': hover }"
              class="cardCourseSmall"
              @click="onClickLesson()"
            >
              <v-img height="200" width="303" :src="math" />
              <v-sheet class="cardInSmallContainer">
                <p class="cardInSmallText">{{ course.name }}</p>
              </v-sheet>
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
    math: "https://cdn.vuetifyjs.com/images/cards/cooking.png",
    title: null,
    courses: [
      {
        index: 1,
        name: "Calculus I"
      },
      {
        index: 2,
        name: "Calculus II"
      },
      {
        index: 3,
        name: "Calculus III"
      },
      {
        index: 4,
        name: "Geometric I"
      },
      {
        index: 5,
        name: "Geometric II"
      },
      {
        index: 6,
        name: "Geometric III"
      },
      {
        index: 7,
        name: "Algebra I"
      },
      {
        index: 8,
        name: "Algebra II"
      },
      {
        index: 9,
        name: "Statistic I"
      },
      {
        index: 10,
        name: "Statistic II"
      }
    ]
  }),
  methods: {
    linetoGrid () {
      return this.items;
    },
    onClickLesson () {
      console.log("ClickCourse");
    },
    onClickBack () {
      this.$router.push({ name: "MyCourse" })
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
.btnBack {
  background-color: white;
  border-radius: 100px;
  outline: none;
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
  font-size: 80px;
  font-family: "Average Sans", sans-serif;
}
.text {
  font-weight: normal;
  color: black;
  font-size: 46px;
  font-family: "Delius", cursive;
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
  margin-top: 50px;
  margin-bottom: 50px;
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
  background-color: #70CCFF;
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
</style>
