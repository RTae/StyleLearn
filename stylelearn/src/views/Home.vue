<template>
  <v-container fluid class="main" id="Home">
    <!-- Card Course -->
    <v-row align="center" justify="center" style="margin-top: 60px">
      <v-sheet class="cardContainer">
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 16"
            :class="{ 'on-hover': hover }"
            class="courseCard"
            @click="onClickCourse()"
          >
            <v-card-text class="cardTextTitle"> CALCULUS I </v-card-text>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 16"
            :class="{ 'on-hover': hover }"
            class="courseCard"
            @click="onClickCourse()"
          >
            <v-card-text class="cardTextTitle"> CALCULUS II </v-card-text>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 16"
            :class="{ 'on-hover': hover }"
            class="courseCard"
            @click="onClickCourse()"
          >
            <v-card-text class="cardTextTitle"> CALCULUS III </v-card-text>
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
      <v-col cols="6">
        <div style="margin-left:50px">
          <p style="font-size: 50px" class="textDetail">
            Newest
          </p>
        </div>
      </v-col>
      <v-col cols="6">
        <router-link to="/coursespage">
          <div style="margin-right: 50px;">
            <p
              style="
                display: flex;
                flex-direction: row-reverse;
                margin-top: 25px;
              "
              class="textDetail"
            >
              All
            </p>
          </div>
        </router-link>
      </v-col>
    </v-row>

    <v-row align="center" justify="center">
      <div class="cardSmallContainer">
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 16"
            :class="{ 'on-hover': hover }"
            class="cardCourseSmall"
            @click="onClickCourse()"
          >
            <v-img
              height="200"
              width="303"
              :src= "math"
            />
            <v-sheet class="cardInSmallContainer">
              <p class="cardInSmallText">Calculus I</p>
            </v-sheet>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 12"
            :class="{ 'on-hover': hover }"
            class="cardCourseSmall"
            @click="onClickCourse()"
          >
            <v-img
              height="200"
              width="303"
              :src= "math"
            />
            <v-sheet class="cardInSmallContainer">
              <p class="cardInSmallText">Calculus II</p>
            </v-sheet>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 12"
            :class="{ 'on-hover': hover }"
            class="cardCourseSmall"
            @click="onClickCourse()"
          >
            <v-img
              height="200"
              width="303"
              :src= "math"
            />
            <v-sheet class="cardInSmallContainer">
              <p class="cardInSmallText">Calculus III</p>
            </v-sheet>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 12"
            :class="{ 'on-hover': hover }"
            class="cardCourseSmall"
            @click="onClickCourse()"
          >
            <v-img
              height="200"
              width="303"
              :src="computer"
            />
            <v-sheet class="cardInSmallContainer">
              <p class="cardInSmallText">Computer Architecture</p>
            </v-sheet>
          </v-card>
        </v-hover>
      </div>
    </v-row>

    <!-- Popular -->
    <v-row style="margin-top: 50px">
      <v-col cols="6">
        <div style="margin-left:50px">
          <p style="font-size: 50px" class="textDetail">
            Popular
          </p>
        </div>
      </v-col>
      <v-col cols="6">
        <router-link to="/coursespage">
          <div style="margin-right: 50px;">
            <p
              style="
                display: flex;
                flex-direction: row-reverse;
                margin-top: 25px;
              "
              class="textDetail"
            >
              All
            </p>
          </div>
        </router-link>
      </v-col>
    </v-row>

    <v-row align="center" justify="center">
      <div class="cardSmallContainer">
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 12"
            :class="{ 'on-hover': hover }"
            class="cardCourseSmall"
            @click="onClickCourse()"
          >
            <v-img
              height="200"
              width="303"
              :src="langauge"
            />
            <v-sheet class="cardInSmallContainer">
              <p class="cardInSmallText">English for oral communication</p>
            </v-sheet>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 12"
            :class="{ 'on-hover': hover }"
            class="cardCourseSmall"
            @click="onClickCourse()"
          >
            <v-img
              height="200"
              width="303"
              :src="computer"
            />
            <v-sheet class="cardInSmallContainer">
              <p class="cardInSmallText">Design Thinking</p>
            </v-sheet>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 12"
            :class="{ 'on-hover': hover }"
            class="cardCourseSmall"
            @click="onClickCourse()"
          >
            <v-img
              height="200"
              width="303"
              :src="econ"
            />
            <v-sheet class="cardInSmallContainer">
              <p class="cardInSmallText">Economic I</p>
            </v-sheet>
          </v-card>
        </v-hover>
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 8 : 12"
            :class="{ 'on-hover': hover }"
            class="cardCourseSmall"
            @click="onClickCourse()"
          >
            <v-img
              height="200"
              width="303"
              :src="computer"
            />
            <v-sheet class="cardInSmallContainer">
              <p class="cardInSmallText">Image Processing</p>
            </v-sheet>
          </v-card>
        </v-hover>
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
          <v-slide-item
            v-for="n in 5"
            :key="n"
          >
            <v-hover v-slot="{ hover }">
              <v-card
                :elevation="hover ? 8 : 12"
                class="ma-10 subjectCard"
                :class="{ 'on-hover': hover }"
                @click="onClickSubject(n);"
              >
                <v-card-text class="cardTextTitle">
                  {{ items[n - 1] }}
                </v-card-text>
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
  </v-container>
</template>

<script>
export default {
  name: "Home",
  components: {},
  data: () => ({
    items: ["Mathematics", "Economics", "Chemistry", "Computer", "Electric"],
    model: null,
    math: "https://cdn.vuetifyjs.com/images/cards/cooking.png",
    computer: "https://cdn.pixabay.com/photo/2020/06/24/19/12/cabbage-5337431_1280.jpg",
    langauge: "https://cdn.pixabay.com/photo/2020/07/12/07/47/bee-5396362_1280.jpg",
    econ: "https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
  }),
  methods: {
    onClickCourse () {
      this.$router.push("/coursespage")
    },
    onClickSubject (n) {
      this.model = n - 1;
      this.$router.push("/coursespage")
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
  justify-content: space-around;
  align-items: center;
  height: 280px;
  width: 100vw;
  background-color: #70CCFF;
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
  font-size: 72px;
  font-family: "Average Sans", sans-serif;
}

a {
  text-decoration: none;
}

.textTitle {
  font-size: 96px;
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
  width: 303px;
  height: 79px;
}

.cardInSmallText {
  margin-left: 40px;
  font-family: "THSarabunNewRegular";
  font-size: 15px;
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
  font-family: "Delius", cursive;
}

.textDetailStudnetReview {
  font-size: 18px;
  font-family: "Delius", cursive;
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
</style>
