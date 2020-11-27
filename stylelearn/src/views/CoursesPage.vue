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
    <!-- Card -->
    <v-row justify="center">
      <div class="tableCard">
        <div class="colCard" v-for="course in courses" :key="course.CourseID">
          <v-hover v-slot="{ hover }">
            <v-card
              :elevation="hover ? 8 : 12"
              :class="{ 'on-hover': hover }"
              class="cardCourseSmall"
              @click="onClickCourse(course.Name, course.CourseID)"
            >
              <v-img height="200" width="303" :src="require('../assets/image/subject/cardSmall/' + title + '.png')" />
              <v-sheet class="cardInSmallContainer">
                <p class="cardInSmallText">{{ course.Name }}</p>
              </v-sheet>
            </v-card>
          </v-hover>
        </div>
      </div>
    </v-row>
  </v-container>
</template>

<script>
import api from "../service/api";
export default {
  name: "coursesPage",
  components: {},
  async mounted () {
    this.title = this.$route.query.titleName
    const result = await api.getCourseBySubject(this.$route.query.id)
    if (result.data.status === "1") {
      this.courses = result.data.result
    } else {
      this.courses = []
    }
  },
  data: () => ({
    math: "https://cdn.vuetifyjs.com/images/cards/cooking.png",
    title: null,
    courses: []
  }),
  methods: {
    onClickCourse (name, id) {
      this.$router.push({ name: "LessonPage", query: { titleName: name, id: id } })
    }
  }
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
  height: 317px;
  width: 1300px;
  background-color: #70ccff;
  border-radius: 30px;
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
  grid-template-columns: 20% 20% 20% 20%;
  width: 90%;
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
  width: 100%;
  height: 79px;
}

.cardInSmallText {
  margin-left: 40px;
  font-family: "THSarabunNew";
  font-size: 23px;
  font-weight: bold;
}

.text {
  font-weight: normal;
  color: white;
  font-size: 100px;
  font-family: "Average Sans", sans-serif;
}
</style>
