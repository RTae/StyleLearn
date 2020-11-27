<template>
  <v-container fluid class="main" id="VideoStudent">
    <v-row align="center" justify="start" style="margin-top:40px">
      <router-link to="/learncoursetutorpage">
        <v-hover v-slot="{ hover }">
          <v-btn
            depressed
            height="60px"
            width="60px"
            :elevation="hover ? 0 : 0"
            :class="{ 'on-hover-review': hover }"
            class="btnBack"
            ><span class="material-icons md-48">
              keyboard_arrow_left
            </span></v-btn
          ></v-hover
        ></router-link
      >
      <div class=" mt-4 mr-4 pd-6 pl-12">
        <p class="headtext">Lesson : Differentiation I</p>
        <p class="headtext">Teacher : Cherprang</p>
      </div>
    </v-row>
    <!-- VIDEO -->
    <v-row align="center" justify="center" style="margin-top:40px;">
      <div v-cloak>
        <v-card class="video">
          <iframe
            width="640"
            height="360"
            :src="'https://www.youtube.com/watch?v=XdE1u9lHie0'"
            frameborder="0"
          ></iframe>
          <hr />
        </v-card>
      </div>
    </v-row>
    <!-- Describe : สอน Diff นั้นง่ายจะตายไป -->
    <v-row justify="center" style="margin-top:40px;">
      <div class="describeCard" >
        <p class="Describetext">Describe : สอน Diff นั้นง่ายจะตายไป</p>
      </div>
    </v-row>
    <!-- card video of tutor -->
    <v-row justify="center" style="margin-top:40px;">
      <div class="tableCard">
        <div class="colCard" v-for="course in courses" :key="course.index">
          <v-sheet class="cardCoruseTitle">
            <p class="cardInSmallText">Teacher : {{ course.tutor }}</p>
          </v-sheet>
          <v-hover v-slot="{ hover }">
            <v-card elevation="8" class="cardCourseSmall">
              <v-img
                :class="{ 'on-hover': hover }"
                class="imgCard"
                :src="math"
                @click="onClickLesson()"
              />
              <v-card-actions style="background-color: #70ccff;">
                <v-list-item class="grow">
                  <v-row>
                    <v-col cols="2">
                      <v-btn
                        icon
                        v-bind:color="course.likeState ? 'red' : 'grey'"
                        @click="onClickLike(course.index - 1)"
                      >
                        <v-icon>mdi-heart</v-icon>
                      </v-btn>
                    </v-col>
                    <v-col align="center" justify="end" cols="5" offset="4">
                      <p class="cardInSmallText">{{ course.like }} Views</p>
                    </v-col>
                  </v-row>
                </v-list-item>
              </v-card-actions>
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
  name: "VideoStudent",
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
        tutor: "Cherprang",
        like: 15657,
        likeState: false
      },
      {
        index: 2,
        tutor: "June",
        like: 14757,
        likeState: false
      },
      {
        index: 3,
        tutor: "Pun",
        like: 13474,
        likeState: false
      },
      {
        index: 4,
        tutor: "Jennis",
        like: 12657,
        likeState: false
      },
      {
        index: 5,
        tutor: "Wee",
        like: 9157,
        likeState: false
      },
      {
        index: 6,
        tutor: "Mobile",
        like: 8157,
        likeState: false
      }
    ]
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
  background-color: rgb(239, 239, 239);
  min-height: 100vh;
}
.video {
  align-items: center;
  text-align: center;
  width: 1200px;
  height: 400px;
}
.describeCard {
  width: 1200px;
  justify-content: start;
  align-items: center;
}
.Describetext{
    font-family: "Delius", cursive;
    font-size: 35px;
    text-align: start;
}
.textTutor {
  font-family: "Delius", cursive;
  font-size: 20px;
  background-color: green;
}
.btnBack {
  background-color: white;
  border-radius: 100px;
  margin-left: 50px;
}
.btnBack:hover {
  color: #47a7f5;
  background-color: white;
  outline: none;
  opacity: 1;
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
  font-size: 54px;
  font-family: "Average Sans", sans-serif;
}
.text {
  font-weight: normal;
  color: black;
  font-size: 46px;
  font-family: "Delius", cursive;
}

.tableCard {
  display: grid;
  grid-template-columns: 40% 40% 40%;
  width: 1000px;
  justify-content: space-around;
}

.colCard {
  margin-top: 20px;
  margin-bottom: 30px;
  display: flex;
  flex-direction: column;
}

.cardCoruseTitle {
  display: flex;
  justify-items: center;
  width: 303px;
  background-color: rgb(239, 239, 239);
}

.cardInSmallText {
  font-family: 'ABeeZee', sans-serif;
  font-size: 15px;
  font-weight: bold;
  margin-top: 10px;
}

.cardCourseSmall {
  margin-bottom: 50px;
  border-radius: 10px;
  width: 303px;
  height: 264px;
  background-color: white;
  transition: opacity 0.2s ease-in;
}

.imgCard {
  height: 200px;
  width: 303px;
  opacity: 0.6;
  transition: opacity 0.2s ease-in;
}

.imgCard:not(.on-hover) {
  opacity: 1;
}
</style>
