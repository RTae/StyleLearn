<template>
  <v-container fluid class="main" id="MyCourse">
    <v-row align="center" justify="start" style="margin-top:40px">
      <v-col align="start" justify="center" class="titleHead" col="12">
        <v-btn
            depressed
            height="60px"
            width="60px"
            class="btnBack material-icons"
            style="margin-left:50px"
            @click="onClickBack()"
        >
          <v-icon>
            keyboard_arrow_left
          </v-icon>
        </v-btn>
        <p class="headtext">{{ title }}</p>
      </v-col>
    </v-row>
    <!-- Card -->
    <v-row justify="center">
      <div class="tableCard">
        <div class="colCard" v-for="lesson in lessons" :key="lesson.LessonID">
          <v-hover v-slot="{ hover }">
            <v-card
              :elevation="hover ? 8 : 12"
              :class="{ 'on-hover': hover }"
              class="cardCourseSmall"
              @click="onClickLesson(lesson.LessonID, lesson.LessonName)"
            >
              <v-img height="200" width="303" src="../../assets/image/lesson/Lesson.png" />
              <v-sheet class="cardInSmallContainer">
                <p class="cardInSmallText">{{ lesson.LessonName | Capitalize }}</p>
              </v-sheet>
            </v-card>
          </v-hover>
        </div>
      </div>
    </v-row>
  </v-container>
</template>

<script scope>
import api from "../../service/api"
import { server } from "../../service/constants";
export default {
  name: "MyCourse",
  components: {},
  async mounted () {
    this.$store.commit("SET_DIALOG_LOADING", true)
    this.title = this.$route.query.titleName;
    const id = localStorage.getItem(server.USERNAME)
    const result = await api.getProgressLesson(id, this.$route.query.id)
    if (result.data.status === "1") {
      this.lessons = result.data.result
      this.$store.commit("SET_DIALOG_LOADING", false)
    }
  },
  data: () => ({
    title: null,
    lessons: []
  }),
  methods: {
    linetoGrid () {
      return this.items;
    },
    onClickLesson (id, name) {
      this.$router.push({ name: "TutorPage", query: { titleName: this.title + " : " + name, id: id } })
    },
    onClickBack () {
      this.$router.go(-1)
    }
  },
  filters: {
    Capitalize (value) {
      if (typeof value !== "string") return ""
      return value.charAt(0).toUpperCase() + value.slice(1)
    }
  }
};
</script>

<style scoped>
.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
}
.titleHead {
  display: flex;
  justify-content: flex-start;
  align-content: flex-start;
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
  font-size: 50px;
  font-family: "Average Sans", sans-serif;
  margin-left: 30px;
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
  font-size: 18px;
  font-weight: bold;
}
</style>
