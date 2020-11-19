<template>
  <v-container fluid class="main" id="adminhome">
    <!-- Subject title -->
    <v-row align="center" justify="center" style="margin-top: 40px">
      <v-card elevation="10" class="cardContainer">
        <p class="text">Video</p>
      </v-card>
    </v-row>
    <!-- Body -->
    <v-row style="margin-top: 50px; margin-bottom: 100px">
      <v-col cols="3">
        <v-card class="cardSideContainer">
          <button @click="onClickVideo()" class="cardButtonContainer">
            Video
          </button>
          <button @click="onClickCustomer()" class="cardButtonContainer">
            Customer
          </button>
          <button @click="onClickApporvevideo()" class="cardButtonContainer">
            Approve video
          </button>
          <button @click="onClickApprovepayment()" class="cardButtonContainer">
            Approve payment
          </button>
          <button
            @click="onClickLogout()"
            class="cardButtonContainer"
            style="background-color: red"
          >
            Log out
          </button>
        </v-card>
      </v-col>
      <v-col cols="9">
        <div>
          <v-row v-for="item in linetoGrid()" :key="item.index">
            <v-col
              align="center"
              justify="center"
              v-for="course in item.line"
              :key="course.index"
              cols="3"
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
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "adminhome",
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
      return this.items;
    },

    onClickLesson () {
      console.log("ClickCourse");
    },
    onClickVideo () {
      this.$router.push({ name: "AdminHome" });
    },
    onClickCustomer () {
      this.$router.push({ name: "AdminHome" });
    },
    onClickApporvevideo () {
      this.$router.push({ name: "AdminHome" });
    },
    onClickApprovepayment () {
      this.$router.push({ name: "AdminHome" });
    },
    onClickLogout () {
      this.$router.push({ name: "AdminHome" });
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
  height: 188px;
  width: 890px;
  background-color: #70ccff;
  border-radius: 30px;
}

.text {
  font-weight: normal;
  color: white;
  font-size: 144px;
  font-family: "Average Sans", sans-serif;
}

.cardSideContainer {
  width: 320px;
  height: 800px;
  background-color: #c4c4c4;
  border-radius: 30px;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
}

.cardButtonContainer {
  width: 254px;
  height: 80px;
  background-color: #70ccff;
  border-radius: 30px;
  font-weight: normal;
  color: white;
  font-size: 24px;
  font-family: "Average Sans", sans-serif;
  style: "margin-top: 40px";
}
</style>
