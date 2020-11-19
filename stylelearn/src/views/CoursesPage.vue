<template>
  <v-container fluid class="main" id="CoursesPage">
    <!-- Subject title -->
    <v-row align="center" justify="center" style="margin-top: 40px">
      <v-card elevation=10 class="cardContainer">
        <p class="text">{{ title }}</p>
      </v-card>
    </v-row>
    <!-- Sort -->
    <v-row align="center" justify="end" style="margin-top: 50px">
      <div class="inputFiled">
        <v-select label="Sort by" solo rounded outlined />
      </div>
    </v-row>
    <!-- Card  -->
    <div style="margin-top: 20px; margin-bottom: 200px">
      <v-row v-for="item in linetoGrid()" :key="item.index">
        <v-col align="center" justify="center" v-for="course in item.line" :key="course.index" cols=3>
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

<script>
export default {
  name: "coursesPage",
  components: {},
  mounted () {
    this.title = this.$route.params.titleName
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
          },
          {
            index: 4,
            course: "Geometric I"
          }
        ]
      },
      {
        index: 2,
        line: [
          {
            index: 1,
            course: "Geometric II"
          },
          {
            index: 2,
            course: "Geometric III"
          },
          {
            index: 3,
            course: "Algebra I"
          },
          {
            index: 4,
            course: "Algebra II"
          }
        ]
      },
      {
        index: 3,
        line: [
          {
            index: 2,
            course: "Statistic I"
          },
          {
            index: 3,
            course: "Statistic II"
          }
        ]
      }
    ],
    course: ["Geometric I", "Geometric II", "Geometric III", "Geometric IV"]
  }),
  methods: {
    linetoGrid () {
      return this.items
    },

    onClickLesson () {
      console.log("ClickCourse")
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
  align-items: center;
  justify-content: center;
  height: 317px;
  width: 1300px;
  background-color: #70ccff;
  border-radius: 30px;
}

.cardSmallContainer {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
  width: 100vw;
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

.text {
  font-weight: normal;
  color: white;
  font-size: 100px;
  font-family: "Average Sans", sans-serif;
}
.botton {
  display: block;
  border-style: none;
  position: absolute;
  box-shadow: none;
  background-repeat: no-repeat;
}
</style>
