<template>
  <v-container fluid class="main" id="Home">
    <v-alert transition="scale-transition" :value="$store.getters.getStateLoginDialog" type="success">
      Login Done !!
    </v-alert>
    <!-- Card Course -->
    <v-row align="center" justify="center" style="margin-top: 60px">
      <v-sheet class="cardContainer">
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 16"
            :class="{ 'on-hover': hover }"
            class="courseCard"
            @click="onClickCourse('c000000003')"
            name=Cal1
          >
            <v-img
              class="cardImg"
              src="../assets/image/courses/cal1.png"
            ></v-img>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 16"
            :class="{ 'on-hover': hover }"
            class="courseCard"
            @click="onClickCourse('c000000021')"
            name=Cal2
          >
            <v-img
              class="cardImg"
              src="../assets/image/courses/cal2.png"
            ></v-img>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 16"
            :class="{ 'on-hover': hover }"
            class="courseCard"
            @click="onClickCourse('c000000022')"
            name=Cal3
          >
            <v-img
            class="cardImg"
            src="../assets/image/courses/cal3.png"
          ></v-img>
          </v-card>
        </v-hover>
      </v-sheet>
    </v-row>

    <!-- Course Title -->
    <v-row align="center" justify="center">
      <p class="textTitle">Courses</p>
    </v-row>

    <!-- Newest -->
    <v-row>
      <v-col cols="6" >
        <div style="margin-left:80px">
          <p style="font-size: 40px" class="textDetail">
            Newest
          </p>
        </div>
      </v-col>
      <v-col style="display: flex; justify-content: flex-end; align-items: center;" cols="6">
        <v-btn
          style="margin-right:50px"
          class="textDetailAll"
          text
          @click="onClikcNewest"
        >
          All
        </v-btn>
      </v-col>
    </v-row>
    <v-row justify="center">
      <div class="tableCard">
        <div class="colCard" v-for="item in newestItem" :key="item.CourseID">
          <v-hover v-slot="{ hover }">
            <v-card
              :elevation="hover ? 8 : 16"
              :class="{ 'on-hover': hover }"
              class="cardCourseSmall"
              @click="onClickCourse(item.CourseID)"
            >
              <v-img
                height="200"
                width="303"
                :src="require('../assets/image/subject/cardSmall/' + item.SubjectName + '.png')"
                name=Newest
              />
              <v-row>
                <v-col class="d-flex pa-5 flex-row justify-space-between">
                  <p class="cardInSmallText">{{ item.CourseName }}</p>
                </v-col>
              </v-row>
            </v-card>
          </v-hover>
        </div>
      </div>
    </v-row>

    <!-- Popular -->
    <v-row style="margin-top: 50px">
      <v-col cols="6">
        <p style="margin-left: 80px; font-size: 40px" class="textDetail">
          Popular
        </p>
      </v-col>
      <v-col style="display: flex; justify-content: flex-end; align-items: center;" cols="6">
        <v-btn
          style="margin-right:50px"
          class="textDetailAll"
          text
          @click="onClikcPopular"
        >
          All
        </v-btn>
      </v-col>
    </v-row>

    <v-row justify="center">
      <div class="tableCard">
        <div class="colCard" v-for="item in popularItem" :key="item.CourseID">
          <v-hover v-slot="{ hover }">
            <v-card
              :elevation="hover ? 8 : 16"
              :class="{ 'on-hover': hover }"
              class="cardCourseSmall"
              @click="onClickCourse(item.CourseID)"
            >
              <v-img
                height="200"
                width="303"
                :src="require('../assets/image/subject/cardSmall/' + item.SubjectName + '.png')"
                name=Newest
              />
              <v-row>
                <v-col class="d-flex pa-5 flex-row justify-space-between">
                  <p class="cardInSmallText">{{ item.CourseName }}</p>
                </v-col>
              </v-row>
            </v-card>
          </v-hover>
        </div>
      </div>
    </v-row>

    <!-- Subject -->
    <v-row align="center" justify="center" style="margin-top: 150px">
      <v-sheet class="subjectContainer">
        <v-slide-group
          v-model="model"
          class="pa-4"
          center-active
          show-arrows
        >
          <v-slide-item v-for="subject in subjectItems" :key="subject.SubjectID" >
            <v-hover v-slot="{ hover }">
              <v-card
                :elevation="hover ? 8 : 12"
                class="ma-10 subjectCard"
                :class="{ 'on-hover': hover }"
                @click="onClickSubject(subject.SubjectID);"
                name=Subject
              >
                <v-img
                class="cardImg"
                :src="require('../assets/image/subject/cardBig/'+ subject.Name + '.png')"
              ></v-img>
              </v-card>
            </v-hover>
          </v-slide-item>
        </v-slide-group>
      </v-sheet>
    </v-row>

    <!-- Review -->
    <v-row align="center" justify="center" style="margin-top: 100px">
      <div class="studentReviewContainner">
        <p class="textTitleStudnetReview">ลองมาฟังเสียงนักเรียนของเราหน่อย</p>
        <p class="textDetailStudnetReview">StyleLearn</p>
        <p class="textDetailStudnetReview">
          เพราะเราเชื่อว่าทุกคนมีสไตล์การเรียนรู้เป็นของตัวเอง
        </p>
      </div>
    </v-row>

    <v-row
      align="center"
      justify="center"
      style="margin-top: 70px; margin-bottom: 500px"
    >
      <div class="cardSmallContainer">
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 16"
            :class="{ 'on-hover-review': hover }"
            class="cardReviewContainer"
          >
            <div class="cardInReviewContainer">
              <v-card-text class="cardInSmallText">ดีมากๆ เลยครับ</v-card-text>
            </div>
            <div class="cardInReviewNameContainer">
              <v-list-item>
                <v-list-item-avatar color="grey darken-3">
                  <v-img
                    alt="pravit"
                    src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
                  ></v-img>
                </v-list-item-avatar>

                <v-list-item-content>
                  <v-list-item-title>ประวิต ยืมเพื่อน</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </div>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 12"
            :class="{ 'on-hover-review': hover }"
            class="cardReviewContainer"
          >
            <div class="cardInReviewContainer">
              <v-card-text class="cardInSmallText"
                >เป็น Platfrom ที่ผมรักมากครับ</v-card-text
              >
            </div>
            <div class="cardInReviewNameContainer">
              <v-list-item>
                <v-list-item-avatar color="grey darken-3">
                  <v-img
                    alt="pravit"
                    src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
                  ></v-img>
                </v-list-item-avatar>

                <v-list-item-content>
                  <v-list-item-title>ประยุทธ์ สีบุญเรือง</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </div>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 12"
            :class="{ 'on-hover-review': hover }"
            class="cardReviewContainer"
          >
            <div class="cardInReviewContainer">
              <v-card-text class="cardInSmallText"
                >สร้างประสบการณ์ใหม่ ในการเรียนรู้มากเลยครับ</v-card-text
              >
            </div>
            <div class="cardInReviewNameContainer">
              <v-list-item>
                <v-list-item-avatar color="grey darken-3">
                  <v-img
                    alt="pravit"
                    src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
                  ></v-img>
                </v-list-item-avatar>

                <v-list-item-content>
                  <v-list-item-title>ธนาธร สวนพึ่ง</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </div>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 12"
            :class="{ 'on-hover-review': hover }"
            class="cardReviewContainer"
          >
            <div class="cardInReviewContainer">
              <v-card-text class="cardInSmallText"
                >เรียนเข้าใจมากขึ้นเลยครับ</v-card-text
              >
            </div>
            <div class="cardInReviewNameContainer">
              <v-list-item>
                <v-list-item-avatar color="grey darken-3">
                  <v-img
                    alt="pravit"
                    src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
                  ></v-img>
                </v-list-item-avatar>
                <v-list-item-content>
                  <v-list-item-title>ไพรบูร วันวานยังหวานชืน</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </div>
          </v-card>
        </v-hover>
      </div>
    </v-row>
    <v-overlay :value="loading">
      <v-progress-circular
        indeterminate
        size="64"
      ></v-progress-circular>
    </v-overlay>
  </v-container>
</template>

<script>
import api from "../service/api"
export default {
  name: "Home",
  components: {},
  data: () => ({
    newestItem: [],
    popularItem: [],
    subjectItems: [],
    model: null,
    loading: true
  }),
  async mounted () {
    const resultSubject = await api.getAllSubject()
    const resultNew = await api.getNewCourseLimit()
    const resultPop = await api.getPopCourseLimit()
    if (resultSubject.data.status === "1") {
      this.subjectItems = resultSubject.data.result
      if (resultNew.data.status === "1") {
        var temp = resultNew.data.result
        this.newestItem = temp.sort(this.compareACED)
        if (resultPop.data.status === "1") {
          var temp2 = resultPop.data.result
          this.popularItem = temp2.sort(this.compareACED)
          this.loading = false
        }
      }
    }
  },
  methods: {
    compareACED (a, b) {
      if (a.CourseName < b.CourseName) {
        return -1;
      }
      if (a.CourseName > b.CourseName) {
        return 1;
      }
      return 0;
    },
    onClickCourse (id) {
      this.$router.push({ name: "LessonPage", query: { id: id } })
    },
    onClickSubject (id) {
      this.$router.push({ name: "CoursesPage", query: { id: id } })
    },
    onClikcNewest () {
      this.$router.push({ name: "CoursePandN", query: { titleName: "Newest" } })
    },
    onClikcPopular () {
      this.$router.push({ name: "CoursePandN", query: { titleName: "Popular" } })
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
  justify-content: space-around;
  align-items: center;
  height: 280px;
  width: 100vw;
  background-color: #70CCFF;
  border-radius: 0px;
}
.courseCard {
  border-radius: 32px;
  width: 464px;
  height: 233px;
  background-color: white;
  display: flex;
  flex-direction: row;
  text-align: center;
  align-items: center;
  opacity: 0.6;
  transition: opacity 0.2s ease-in;
}

.courseCard:not(.on-hover) {
  opacity: 1;
}

.cardTextTitle {
  font-size: 40px;
  font-family: "Average Sans", sans-serif;
}

.cardImg{
  height: 100%;
  width: 100%;
}

.tableCard {
  display: grid;
  grid-template-columns: 25% 25% 25% 25%;
  width: 100vw;
  justify-content: space-around
}

.colCard {
  margin-top: 20px;
  margin-bottom: 30px;
  display: flex;
  justify-content: center;
}

a {
  text-decoration: none;
}

.textTitle {
  font-size: 70px;
  font-family: "Average Sans", sans-serif;
  align-content: center;
  justify-content: center;
  margin-top: 40px;
}

.textDetail {
  font-weight: normal;
  color: black;
  font-size: 30px;
  font-family: Delius;
}

.cardSmallContainer {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  width: 100vw;
}

.cardCourseSmall {
  margin-top: 30px;
  margin-bottom: 10px;
  border-radius: 10px;
  width: 303px;
  height: 264px;
  background-color: white;
  opacity: 0.6;
  background-color: #70CCFF;
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
  font-family: "THSarabunNew";
  font-size: 23px;
  font-weight: bold;
}

.subjectContainer {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  height: 474px;
  width: 100vw;
  background-color: #70CCFF;
}

.subjectCard {
  height: 300px;
  width: 642px;
  border-radius: 50px;
  display: flex;
  flex-direction: row;
  align-items: center;
  text-align: center;
  opacity: 0.6;
  transition: opacity 0.2s ease-in;
}

.subjectCard:not(.on-hover) {
  opacity: 1;
}

.studentReviewContainner {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.textTitleStudnetReview {
  font-size: 40px;
  font-family: "Delius"
}

.textDetailStudnetReview {
  font-size: 18px;
  font-family: "Delius"
}

.cardReviewContainer {
  height: 343px;
  width: 373px;
  border-radius: 5px;
  opacity: 1;
  transition: opacity 0.2s ease-in;
}

.cardReviewContainer:not(.on-hover-review) {
  height: 343px;
  width: 373px;
  border-radius: 5px;
  opacity: 0.5;
}

.cardInReviewContainer {
  height: 243px;
  width: 373px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.cardInReviewNameContainer {
  background-color: #70CCFF;
  height: 100px;
  width: 373px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.textDetailAll {
  font-weight: normal;
  color: black;
  font-size: 30px;
  font-family: Delius;
}
</style>
