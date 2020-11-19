<template>
  <v-container fluid class="main" id="LessonPage">
    <v-row align="center" justify="start" style="margin-top:40px">
      <v-col cols="1"></v-col>
      <p class="headtext">CALCULUS I</p>
    </v-row>
    <v-row>
      <v-col cols="1"></v-col>
      <p class="text">Syllabus</p>
    </v-row>
    <v-row align="center" justify="center">
      <div class="d-flex flex-column mb-10" justify-center align-center>
        <div v-for="n in 7" :key="n">
          <v-card class="courseCard">
            <v-row>
              <v-col>
                <div class="detail">
                  <v-card-text class="cardTextTitle">Diff1</v-card-text>
                  <v-card-text class="cardTextDetail"
                    >เรียนเกี่ยวกับ Diff</v-card-text
                  >
                </div>
              </v-col>
              <v-col class="buttonContainer">
                <v-hover v-slot="{ hover }">
                  <button
                    :elevation="hover ? 8 : 0"
                    :class="{ 'on-hover-review': hover }"
                    class="bottonBuy"
                    color="#70ccff"
                  >
                    Buy
                  </button>
                </v-hover>
              </v-col>
            </v-row>
          </v-card>
        </div>
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
    <v-row align="center" justify="center" style="margin-top:70px; margin-bottom: 200px">
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
                  <v-card-text class="cardTextTutorTitle"
                    >ครูพี่นิว</v-card-text
                  >
                </v-row>
                <v-row>
                  <v-col cols="2">
                    <v-card-text class="cardTextTutorDetail"
                      >“ไม่มีอะไรสายเกินไป ถ้าเราคิดที่จะเริ่ม”</v-card-text
                    >
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
                  <v-card-text class="cardTextTutorTitle"
                    >ครูพี่เต้</v-card-text
                  >
                </v-row>
                <v-row>
                  <v-col cols="2">
                    <v-card-text class="cardTextTutorDetail"
                      >“ทุกคนต่างเรียกผมว่า จาน”</v-card-text
                    >
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
                  <v-card-text class="cardTextTutorTitle"
                    >ครูพี่เฟิร์ส</v-card-text
                  >
                </v-row>
                <v-row>
                  <v-col cols="2">
                    <v-card-text class="cardTextTutorDetail"
                      >“ทุกคนต่างเรียกผมว่า เทพ”</v-card-text
                    >
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

<script scope>
// @ is an alias to /src

export default {
  name: "LessonPage",
  components: {}
};
</script>

<style scoped>
.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
  margin-top: 50px;
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
.cardContainer {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  width: 100vw;
}
.cardTextTitle {
  font-weight: bold;
  color: black;
  font-size: 30px;
  font-family: "THSarabunNewRegular";
}
.cardTextDetail {
  font-weight: normal;
  color: black;
  font-size: 16px;
  font-family: "THSarabunNewRegular";
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
  background-color: #70CCFF;
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
  font-family: "Delius", cursive;
  justify-content: end;
}
.cardTextTutorDetail {
  font-weight: normal;
  color: black;
  font-size: 16px;
  font-family: "Delius", cursive;
  justify-content: end;
}
.imgContainer {
  align-items: center;
  justify-content: center;
}
</style>
