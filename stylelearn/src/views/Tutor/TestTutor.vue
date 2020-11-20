<template>
  <v-container fluid class="main" id="login">
    <v-row style="margin-top: 15vh" align="center" justify="center">
      <p class="textTitle">My Video</p>
    </v-row>
    <!--sort by btn-->
    <v-row justify="center">
      <v-col cols="1">
        <v-btn class="botton2" color="#FFFFFF" width="8%" height="50px">
          Sort by
        </v-btn>
      </v-col>
    </v-row>
    <v-row style="margin-top: 10vh">
      <p></p>
    </v-row>
    <!--card-->
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

<script>
export default {
  name: "coursesPage",
  components: {},
  mounted () {
    this.title = this.$route.params.titleName
  },
  data: () => ({
    math: "https://sv1.picz.in.th/images/2020/11/16/bHaFJ2.png",
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
            course: "Geometric II"
          },
          {
            index: 2,
            course: "Geometric III"
          },
          {
            index: 3,
            course: "Algebra I"
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
.text2 {
  color: black;
  margin-top: 20px;
  font-size: 20px;
  font-family: "Delius", cursive;
  margin-left: 90vh;
}

.botton {
  display: block;
  border-style: none;
  position: absolute;
  box-shadow: none;
  background-repeat: no-repeat;
}

.botton2 {
  display: block;
  border-style: none;
  position: absolute;
  box-shadow: none;
  background-repeat: no-repeat;
  margin-left: 60vh;
  border-radius: 20px;
}
.cardCourseSmall {
  margin-top: 45px;
  margin-bottom: 45px;
  border-radius: 15px;
  width: 303px;
  height: 250px;
  background-color: white;
  opacity: 0.8;
  transition: opacity 0.2s ease-in;
}
</style>
