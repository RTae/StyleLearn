<template>
  <v-container fluid class="main" id="Home">
    <v-row align="center" justify="center" style="margin-top:40px">
      <div>
        <v-card width="1500" height="220" color="#70CCFF">
          <div class="d-flex align-center" align-center justify-center>
            <v-col cols="4">
              <router-link to="/coursespage"
                ><button class="course">CALCULUS I</button></router-link
              >
              <!--<v-card class="course">
                    <v-card-text>
                      <p class="text_coure" align="center" justify="center">CALCULUS I</p>
                    </v-card-text>
              </v-card>!-->
            </v-col>
            <v-col cols="4">
              <router-link to="/coursespage"
                ><button class="course">CALCULUS II</button></router-link
              >
            </v-col>
            <v-col cols="4">
              <router-link to="/coursespage"
                ><button class="course">CALCULUS III</button></router-link
              >
            </v-col>
          </div>
        </v-card>
      </div>
    </v-row>
    <v-row align="center" justify="center">
      <p class="Head_text">Courses</p>
    </v-row>
    <v-row align="center" justify="start">
      <v-col cols="6"><p class="text">Newest</p></v-col>
      <router-link to="/coursespage"><p class="text2">All</p></router-link>
    </v-row>
    <v-row>
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:30px"
          border-radius="20px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Calculus I
          </v-btn>
        </v-card>
      </v-col>
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:30px"
          border-radius="20px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Calculus II
          </v-btn>
        </v-card>
      </v-col>
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:30px"
          border-radius="20px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Calculus III
          </v-btn>
        </v-card>
      </v-col>
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:30px"
          border-radius="20px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Computer architecture
          </v-btn>
        </v-card>
      </v-col>
    </v-row>
    <v-row align="center" justify="start">
      <v-col cols="6"><p class="text">Popular</p></v-col>
      <router-link to="/about"><p class="text2">All</p></router-link>
    </v-row>
    <v-row>
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:30px"
          border-radius="20px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            English for oral communication
          </v-btn>
        </v-card>
      </v-col>
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:30px"
          border-radius="20px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Design Thinking
          </v-btn>
        </v-card>
      </v-col>
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:30px"
          border-radius="20px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Digital and Logic Design
          </v-btn>
        </v-card>
      </v-col>
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:30px"
          border-radius="20px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Signal Processing
          </v-btn>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-card width="1500" height="300" color="#70CCFF" style="margin-top:50px">
        <div align-center justify-center>
          <v-carousel hide-delimiters>
              <v-carousel-item
                v-for="(item,i) in items"
                :key="i"
                :src="item.src"
              ></v-carousel-item>
              <v-carousel-item
                v-for="(item,i) in items"
                :key="i"
                :src="item.src"
              ></v-carousel-item>
            </v-carousel>
        </div>
      </v-card>
    </v-row>
    <div align="center" justify="center" style="margin-top:100px">
      <p class="headreveiw">ลองมาฟังเสียงนักเรียนของเราหน่อย</p>
      <p class="reveiw">Style Learn</p>
      <p class="reveiw">เพราะเราเชื่อว่าทุกคนมีสไตล์การเรียนรู้เป็นของตัวเอง</p>
    </div>

  </v-container>
</template>

<script>
// @ is an alias to /src

export default {
  name: "Home",
  components: {}
};
</script>

<style scoped>
.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
  margin-top: 50px;
}
.course {
  font-family: "Average Sans", sans-serif;
  border-radius: 20px;
  width: 480px;
  height: 200px;
  opacity: 1;
  transition: 0.3s;
  font-size: 12px;
  text-transform: uppercase;
  cursor: pointer;
  outline: none;
  background-color: white;
  color: #000;
  font-size: 55px;
  font-family: "Average Sans", sans-serif;
}
.Head_text {
  font-size: 50px;
  font-family: "Average Sans", sans-serif;
  align-content: center;
  justify-content: center;
  margin-top: 40px;
  font-weight: bold;
}
.text {
  font-weight: normal;
  color: black;
  margin-top: 20px;
  font-size: 30px;
  font-family: 'Delius', cursive;
  margin-left: 30px;
}
.text2 {
  color: black;
  margin-top: 20px;
  font-size: 20px;
  font-family: 'Delius', cursive;
  margin-left: 90vh;
}
.botton {
  display: block;
  border-style: none;
  position: absolute;
  box-shadow: none;
  background-repeat: no-repeat;
}
.headreveiw {
  font-size: 35px;
  font-family: 'Delius', cursive;
}
.reveiw {
  font-size: 20px;
  font-family: 'Delius', cursive;
}
</style>
