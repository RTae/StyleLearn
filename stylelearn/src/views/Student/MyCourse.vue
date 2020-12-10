<template>
  <v-container fluid class="main" id="MyCourse">
    <v-row align="center" justify="start" style="margin-top:40px">
      <v-col align="start" justify="center" cols="11" offset="1">
        <p class="textHead">My Course</p>
      </v-col>
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
              @click="onClickCourse(course.CourseID, course.CourseName)"
            >
              <v-img height="200" width="303" :src="require('../../assets/image/subject/cardSmall/' + course.SubjectName + '.png')" />
              <v-sheet class="cardInSmallContainer">
                <p class="cardInSmallText">{{ course.CourseName }}</p>
              </v-sheet>
            </v-card>
          </v-hover>
        </div>
      </div>
    </v-row>
  </v-container>
</template>

<script>
import api from "../../service/api"
import { server } from "../../service/constants";
export default {
  name: "MyCourse",
  components: {},
  async mounted () {
    this.$store.commit("SET_DIALOG_LOADING", true)
    const id = localStorage.getItem(server.USERNAME)
    const result = await api.getCourseFromProgressLesson(id, name)
    if (result.data.status === "1") {
      this.courses = result.data.result
      this.$store.commit("SET_DIALOG_LOADING", false)
    }
  },
  data: () => ({
    math: "https://cdn.vuetifyjs.com/images/cards/cooking.png",
    title: null,
    courses: []
  }),
  methods: {
    onClickCourse (id, name) {
      this.$router.push({ name: "LearnCourse", params: { titleName: name, id: id } })
    }
  }
};
</script>

<style scoped>
.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
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
.textHead {
  font-weight: normal;
  color: black;
  font-size: 55px;
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
  margin-bottom: 50px;
  border-radius: 10px;
  width: 303px;
  height: 279px;
  background-color: #70CCFF;
  opacity: 0.6;
  transition: opacity 0.2s ease-in;
}

.cardCourseSmall:not(.on-hover) {
  opacity: 1;
}

.cardInSmallContainer {
  background-color: #70CCFF;
  margin-left: 15px;
  height: 79px;
  display: flex;
  align-items: center;
  justify-content: flex-start;
}

.cardInSmallText {
  font-family: "THSarabunNew";
  font-size: 20px;
  font-weight: bold;
}
</style>
