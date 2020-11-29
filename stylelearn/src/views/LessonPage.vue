<template>
  <v-container fluid class="main" id="LessonPage">
    <v-snackbar
      top
      v-model="popUpBuy"
      color="#ffce70"
      elevation="7"
      :timeout="2000"
    >
      <p style="color:black; text-align:center">เพิ่มเข้าตะกร้าเรียบร้อย</p>
      <template v-slot:action="{ attrs }">
        <v-btn
          color="blue"
          text
          v-bind="attrs"
          @click="snackbar = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
    <!-- Title -->
    <v-row align="center" justify="start" style="margin-top:40px">
      <v-col cols="1"></v-col>
      <p class="headText">{{ this.title }}</p>
    </v-row>
    <v-row>
      <v-col cols="1"></v-col>
      <p class="text">Syllabus</p>
    </v-row>
    <!-- Card -->
    <v-row align="center" justify="center">
      <div class="d-flex flex-column mb-10" justify-center align-center>
        <div v-for="lesson in lessons" :key="lesson.LessonID">
          <v-card class="courseCard">
            <v-row>
              <v-col>
                <div class="detail">
                  <v-card-text class="cardTextTitle">{{ lesson.Name }}</v-card-text>
                  <v-card-text class="cardTextDetail">{{ lesson.Description }}</v-card-text>
                </div>
              </v-col>
              <v-col class="buttonContainer">
                <v-hover v-slot="{ hover }">
                  <button
                    :elevation="hover ? 8 : 0"
                    :class="{ 'on-hover-review': hover }"
                    class="bottonBuy"
                    color="#70ccff"
                    @click="onClickBuy(lesson.LessonID, lesson.Name)"
                  >
                    Buy
                  </button>
                </v-hover>
              </v-col>
            </v-row>
          </v-card>
        </div>
        <!-- Buy whole -->
        <v-card class="courseCard">
          <v-row>
            <v-col>
              <div class="detail">
                <v-card-text class="cardTextTitle"> Whole lesson</v-card-text>
              </div>
            </v-col>
            <v-col class="buttonContainer">
              <v-hover v-slot="{ hover }">
                <button
                  :elevation="hover ? 8 : 0"
                  :class="{ 'on-hover-review': hover }"
                  class="bottonBuy"
                  color="#70ccff"
                  @click="onClickBuyWhole()"
                >
                  Buy
                </button>
              </v-hover>
            </v-col>
          </v-row>
        </v-card>
      </div>
    </v-row>
    <v-row>
      <v-col cols="1"></v-col>
      <v-hover v-slot="{ hover }">
        <router-link to="/coursespage">
          <button
            :elevation="hover ? 8 : 0"
            :class="{ 'on-hover-review': hover }"
            class="bottonBuy"
            color="#70ccff"
          >
            Back
          </button>
        </router-link>
      </v-hover>
    </v-row>

    <!-- teacher -->
    <v-row align="center" justify="center" style="margin-top:150px;margin-bottom: 200px;">
      <div class="cardContainer">
        <v-card class="tutorCard">
          <v-row>
            <v-col cols="2" class="imgContainer">
              <v-img
                width="128px"
                height="128px"
                alt="pravit"
                src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
              ></v-img>
            </v-col>
            <v-col cols="2">
              <div class="detail">
                <v-row>
                  <v-card-text class="cardTextTutorTitle">ครูพี่เฟิร์ส</v-card-text>
                </v-row>
                <v-row>
                  <v-col cols="2">
                    <v-card-text class="cardTextTutorDetail">“ทุกคนต่างเรียกผมว่า เทพ”</v-card-text>
                  </v-col>
                </v-row>
              </div>
            </v-col>
          </v-row>
        </v-card>
        <v-card class="tutorCard">
          <v-row>
            <v-col cols="2" class="imgContainer">
              <v-img
                width="128px"
                height="128px"
                alt="pravit"
                src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
              ></v-img>
            </v-col>
            <v-col cols="2">
              <div class="detail">
                <v-row>
                  <v-card-text class="cardTextTutorTitle">ครูพี่นิว</v-card-text>
                </v-row>
                <v-row>
                  <v-col cols="2">
                    <v-card-text class="cardTextTutorDetail">“เลขง่ายๆ กับครูพี่นิว”</v-card-text>
                  </v-col>
                </v-row>
              </div>
            </v-col>
          </v-row>
        </v-card>
        <v-card class="tutorCard">
          <v-row>
            <v-col cols="2" class="imgContainer">
              <v-img
                width="128px"
                height="128px"
                alt="pravit"
                src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
              ></v-img>
            </v-col>
            <v-col cols="2">
              <div class="detail">
                <v-row>
                  <v-card-text class="cardTextTutorTitle">ครูพี่อี้หมิง</v-card-text>
                </v-row>
                <v-row>
                  <v-col cols="2">
                    <v-card-text class="cardTextTutorDetail">“มาเรียนภาษาจีนด้วยกันเถอะ”</v-card-text>
                  </v-col>
                </v-row>
              </div>
            </v-col>
          </v-row>
        </v-card>
      </div>
    </v-row>
  </v-container>
</template>

<script>
import api from "../service/api"
export default {
  name: "LessonPage",
  components: {},
  async mounted () {
    this.title = this.$route.query.titleName
    const result = await api.getLessonByCourse(this.$route.query.id)
    if (result.data.status === "1") {
      this.lessons = result.data.result
    } else {
      this.lessons = []
    }
  },
  data () {
    return {
      title: "",
      popUpBuy: false,
      lessons: []
    }
  },
  methods: {
    onClickBuy (id, name) {
      if (this.$store.getters.getLoginHeaderStudent) {
        this.popUpBuy = true
        this.$store.dispatch({
          type: "addItemToBukect",
          id: id,
          name: name
        })
      } else {
        this.$router.push({ name: "Login" })
      }
    },
    onClickBuyWhole () {
      if (this.$store.getters.getLoginHeaderStudent) {
        this.popUpBuy = true
        for (var i = 0; i < this.lessons.length; i++) {
          this.$store.dispatch({
            type: "addItemToBukect",
            id: this.lessons[i].LessonID,
            name: this.lessons[i].Name
          })
        }
      } else {
        this.$router.push({ name: "Login" })
      }
    }
  }
};
</script>

<style scoped>
.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
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

.headText {
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

.cardContainer {
  display:flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  width: 90vw;
  background-color: rgb(239, 239, 239);
}

.cardTextTitle {
  font-weight: bold;
  color: black;
  font-size: 40px;
  font-family: "THSarabunNew";
}

.cardTextDetail {
  font-weight: normal;
  color: black;
  font-size: 23px;
  font-family: "THSarabunNew";
}

.detail {
  width: 1000px;
  height: 123;
  margin-left: 10px;
}

.buttonContainer {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 200px;
  height: 123;
}

.bottonBuy {
  background-color: #5cbbf6;
  font-family: "Average Sans", sans-serif;
  border-radius: 100px;
  width: 120px;
  height: 50px;
  font-size: 18px;
  color: white;
  outline: none;
}
.bottonBuy:hover {
  background: #47a7f5 radial-gradient(circle, transparent 1%, #47a7f5 1%)
    center/15000%;
  color: #000;
  outline: none;
}
.bottonBuy:active {
  background-color: #ffce70;
  outline: none;
}
.tutorCard {
  border-radius: 10px;
  background-image: linear-gradient(to right, #3888FF , #70CCFF);
  display: flex;
  flex-direction: row;
  text-align: start;
  align-items: center;
  margin-top: 20px;
  width: 383px;
  height: 238px;
}
.cardTextTutorTitle {
  font-weight: bold;
  color: black;
  font-size: 30px;
  font-family: "Delius";
  justify-content: end;
}
.cardTextTutorDetail {
  font-weight: normal;
  color: black;
  font-size: 16px;
  font-family: "Delius";
  justify-content: end;
}
.imgContainer {
  align-items: center;
  justify-content: center;
}
</style>
