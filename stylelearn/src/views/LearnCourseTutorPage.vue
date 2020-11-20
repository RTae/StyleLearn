<template>
  <v-container fluid class="main" id="MyCourse">
    <v-row align="center" justify="start" style="margin-top:40px">
      <router-link to="/mycourse">
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
        <p class="headtext">CALCULUS I : Differentiation I</p>
      </div>
    </v-row>
    <div style="margin-top: 20px; margin-bottom: 200px">
      <v-row v-for="item in linetoGrid()" :key="item.index">
        <v-col cols="1.5"></v-col>
        <v-col
          align="center"
          justify="center"
          v-for="course in item.line"
          :key="course.index"
          cols="3"
        >
        <div style="background-color:rgb(239, 239, 239)">
            <p class="textTutor">{{ course.tutor }}</p>
        </div>
        <v-card class="cardBig">
            <v-hover v-slot="{ hover }">
              <v-card
                :elevation="hover ? 8 : 0"
                :class="{ 'on-hover': hover }"
                class="cardCourseSmall"
                @click="onClickLesson()"
              >
                <v-img height="200" width="303" :src="math" />
              </v-card>
            </v-hover>
            <v-sheet>
              <v-row class="cardInSmallContainer">
                <v-col cols="6">
                  <v-btn fab depressed small
                    ><icon class="material-icons right">
                      favorite_border
                    </icon>
                  </v-btn>
                </v-col>
                <v-col cols="6">
                  <p class="cardInSmallText">{{ course.like }}</p>
                </v-col>
              </v-row>
            </v-sheet>
          </v-card>
        </v-col>
        <v-col cols="1.5"></v-col>
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
            tutor: "Teacher : Cherprang",
            like: "15,657 Views"
          },
          {
            index: 2,
            tutor: "Teacher : June",
            like: "14,757 Views"
          },
          {
            index: 3,
            tutor: "Teacher : Pun",
            like: "13,474 Views"
          }
        ]
      },
      {
        index: 2,
        line: [
          {
            index: 1,
            tutor: "Teacher : Jennis",
            like: "12,657 Views"
          },
          {
            index: 2,
            tutor: "Teacher : Wee",
            like: "9,157 Views"
          },
          {
            index: 3,
            tutor: "Teacher : Mobile",
            like: "8,157 Views"
          }
        ]
      }
    ],
    like: ["15,657 Views", "14,757 Views", "13,474 Views", "12,657 Views", "9,157 Views", "8,157 Views"]
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
.textTutor {
  font-family: "Delius", cursive;
  font-size: 24px;
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
  font-size: 80px;
  font-family: "Average Sans", sans-serif;
}
.text {
  font-weight: normal;
  color: black;
  font-size: 46px;
  font-family: "Delius", cursive;
}
.cardBig {
  width: 303px;
  height: 300px;
  background-color: #70ccff;
  border-radius: 10px;
  margin-bottom: 20px;
}
.cardInSmallContainer {
  border-radius: 10px;
  background-color: #70ccff;
  width: 303px;
  height: 60px;
  align-items: center;
  justify-content: start;
}
.cardCourseSmall {
  width: 303px;
  height: 200px;
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
  margin-top: 10px;
}
</style>
