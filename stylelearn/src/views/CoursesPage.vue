<template>
  <v-container fluid class="main" id="CoursesPage">
    <!-- Subject title -->
    <v-row align="center" justify="center" style="margin-top: 40px">
      <v-card elevation=10 class="cardContainer">
        <p class="text">{{ title }}</p>
      </v-card>
    </v-row>
    <!-- Sort -->
    <v-row align="center" justify="end" style="margin-right:50px;margin-top: 50px">
      <v-menu
        transition="slide-y-transition"
        bottom
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            class="btnSort"
            v-bind="attrs"
            v-on="on"
          >
          <p style="font-family: 'Delius';font-size: 15px; margin-top:15px">Sort by</p>
          <v-icon right> mdi-arrow-down-drop-circle-outline </v-icon>
          </v-btn>
        </template>
        <v-list>
          <v-list-item
            v-for="(item, i) in sortItem"
            :key="i"
          >
            <v-list-item-title>
              <button
                class="btnPro"
                @click="sort=item"
              >
                {{ item }}
              </button>
            </v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-row>
    <!-- Card -->
    <v-row justify="center">
      <div class="tableCard">
        <div class="colCard" v-for="course in courseList" :key="course.CourseID">
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
    this.$store.commit("SET_DIALOG_LOADING", true)
    this.title = this.$route.query.titleName
    const result = await api.getCourseBySubject(this.$route.query.id)
    if (result.data.status === "1") {
      this.courses = result.data.result
      this.$store.commit("SET_DIALOG_LOADING", false)
    } else {
      this.$store.commit("SET_DIALOG_LOADING", false)
    }
  },
  data: () => ({
    title: null,
    courses: [],
    sortItem: ["a - z", "z - a"],
    sort: "a - z"
  }),
  computed: {
    courseList () {
      var temp = this.courses
      if (this.sort === "a - z") {
        return temp.sort(this.compareACED)
      } else {
        return temp.sort(this.compareDEC)
      }
    }
  },
  methods: {
    compareACED (a, b) {
      if (a.Name < b.Name) {
        return -1;
      }
      if (a.Name > b.Name) {
        return 1;
      }
      return 0;
    },
    compareDEC (a, b) {
      if (a.Name > b.Name) {
        return -1;
      }
      if (a.Name < b.Name) {
        return 1;
      }
      return 0;
    },
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
  height: 250px;
  width: 1100px;
  background-color: #70ccff;
  border-radius: 30px;
}

.btnSort {
  width: 214px;
  height: 70px;
  border-radius: 15px;
  background-color: white;
}

.btnPro {
  background-color: white;
  width: 150px;
  height: 50px;
  outline: none;
}
.btnPro:hover {
  color: #47a7f5;
  outline: none;
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
