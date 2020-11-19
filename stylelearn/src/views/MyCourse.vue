<template>
  <v-container fluid class="main" id="MyCourse">
    <v-row align="center" justify="start" style="margin-top:40px">
      <v-col cols="1"></v-col>
      <p class="headtext">My Course</p>
    </v-row>
    <div style="margin-top: 20px; margin-bottom: 200px">
      <v-row v-for="item in linetoGrid()" :key="item.index">
        <v-col
          align="center"
          justify="center"
          v-for="course in item.line"
          :key="course.index"
          cols="4"
        >
          <v-hover v-slot="{ hover }">
            <v-card
              :elevation="hover ? 8 : 12"
              :class="{ 'on-hover': hover }"
              class="cardCourseSmall"
              @click="onClickLesson()"
            >
              <v-img height="200" width="303" :src="math" />
              <v-sheet class="cardInSmallContainer">
                <p class="cardInSmallText">{{ course.course }}</p>
              </v-sheet>
            </v-card>
          </v-hover>
        </v-col>
      </v-row>
    </div>
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
    items: [
      {
        index: 1,
        line: [
          {
            index: 1,
            course: "Calculus I"
          },
          {
            index: 2,
            course: "Calculus II"
          },
          {
            index: 3,
            course: "Calculus III"
          }
        ]
      },
      {
        index: 2,
        line: [
          {
            index: 1,
            course: "Geometric I"
          },
          {
            index: 2,
            course: "Geometric II"
          },
          {
            index: 3,
            course: "Geometric III"
          }
        ]
      },
      {
        index: 3,
        line: [
          {
            index: 1,
            course: "Algebra I"
          },
          {
            index: 2,
            course: "Algebra II"
          },
          {
            index: 3,
            course: "Statistic I"
          }
        ]
      },
      {
        index: 4,
        line: [
          {
            index: 1,
            course: "Statistic II"
          }
        ]
      }
    ],
    course: ["Geometric I", "Geometric II", "Geometric III", "Geometric IV"]
  }),
  methods: {
    linetoGrid () {
      return this.items;
    },
    onClickLesson () {
      console.log("ClickCourse");
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
.courseCard {
  border-radius: 10px;
  background-color: white;
  display: flex;
  flex-direction: row;
  text-align: start;
  align-items: center;
  margin-top: 20px;
  width: 1200px;
  height: 100px;
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
.cardSmallContainer {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
  width: 80vw;
}
.cardCourseSmall {
  margin-top: 50px;
  margin-bottom: 50px;
  border-radius: 10px;
  width: 303px;
  height: 260px;
  background-color: #70ccff;
  opacity: 0.6;
  transition: opacity 0.2s ease-in;
}
.cardCourseSmall:not(.on-hover) {
  opacity: 1;
}
.cardInSmallText {
  font-family: "THSarabunNewRegular";
  font-size: 15px;
  font-weight: bold;
  background-color: #70ccff;
  margin-top: 20px;
}
</style>
