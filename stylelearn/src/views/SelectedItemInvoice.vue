<template>
  <v-container fluid class="main" id="CoursesPage">
    <!--title -->
    <v-row align="center" justify="center" style="margin-top: 40px">
      <v-card class="cardContainer">
        <p class="text">Invoice</p>
      </v-card>
    </v-row>
    <!--Order-->
    <v-row align="center" justify="center">
        <v-card class="cardOrder">
          <p class="textTitle">Order</p>
        </v-card>
    </v-row>
    <v-row align="center" justify="center">
        <v-card class="cardOrderContainer">
            <v-row  v-for="lessonBuy in lessonsBuy" :key="lessonBuy.index" align="center" justify="center">
              <v-col>
                <p class="cardTextTitle">{{lessonBuy.name }}</p>
              </v-col>
              <v-col>
                <p class="cardPrice">X {{}} Day</p>
              </v-col>
              <v-col>
                <p class="learnday">50 Bath</p>
              </v-col>
            </v-row>
        </v-card>
    </v-row>
    <v-row align="center" justify="center">
        <div style="width:900px">
            <v-row align="center" justify="end" style="margin-top:50px">
                <v-card class="totalCard">
                    <p class="cardTextTotal">Total : 100 Bath</p>
                </v-card>
            </v-row>
        </div>
    </v-row>
    <v-row align="center" justify="center" style="margin-bottom:100px;">
      <div class="totalCardContainer">
        <v-row align="center" justify="end">
          <v-hover v-slot="{ hover }">
            <router-link to="/coursespage">
              <button
                :elevation="hover ? 8 : 0"
                :class="{ 'on-hover-review': hover }"
                class="bottonCom"
                color="#70ccff"
              >
                Confirm Order
              </button>
            </router-link>
          </v-hover>
        </v-row>
      </div>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "coursesPage",
  components: {},
  data: () => ({
    lessonsBuy: [
      {
        index: 1,
        name: "Differentiation I"
      },
      {
        index: 2,
        name: "Differentiation II"
      },
      {
        index: 3,
        name: "Differentiation III"
      }
    ]
  })
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
  height: 150px;
  width: 900px;
  background-color: #70ccff;
  border-radius: 30px;
  margin-top: 50px;
}
.cardOrder {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-start;
  height: 70px;
  width: 900px;
  background-color: #70ccff;
  border-radius: 30px;
  margin-top: 50px;
}
.cardOrderContainer {
  background-color: white;
  margin-top: 20px;
  border-radius: 30px;
  width: 900px;
}
.text {
  font-weight: normal;
  color: white;
  font-size: 60px;
  font-family: "Average Sans", sans-serif;
}
.textTitle {
  font-weight: normal;
  color: white;
  font-size: 30px;
  font-family: "Delius", cursive;
  margin-left: 20px;
}
.bottonCom {
  background-color: #5cbbf6;
  font-family: "Average Sans", sans-serif;
  border-radius: 100px;
  width: 200px;
  height: 50px;
  font-size: 18px;
  color: white;
  outline: none;
  margin-top: 20px;
}
.bottonCom:hover {
  background: #47a7f5 radial-gradient(circle, transparent 1%, #47a7f5 1%)
    center/15000%;
  color: #000;
  outline: none;
}
.bottonCom:active {
  background-color: #ffce70;
  outline: none;
}
.courseCard {
  border-radius: 30px;
  background-color: white;
  display: flex;
  flex-direction: row;
  justify-content: center;
  text-align: start;
  align-items: center;
  margin-top: 20px;
  width: 800px;
  height: 80px;
  outline-color: none;
}
.totalCardContainer {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
  width: 800px;
  height: 80px;
}
.totalCard {
  border-radius: 30px;
  display: flex;
  justify-content: center;

  width: 400px;
  height: 50px;
  background-color: white;
}
.cardTextTotal {
  color: black;
  font-family: "Delius";
  font-size: 25px;
}
.cardTextTitle {
  width: 400px;
  font-family: "Delius";
  font-size: 20px;
  margin-left: 20px;
  border-radius: 30px;
}
.cardPrice {
  font-family: "Delius";
  font-size: 20px;
}
.buttonContainer {
  display: flex;
  justify-items: center;
  width: 225px;
  height: 123;
}
.learnday {
  font-family: "Delius";
  font-size: 20px;
}
</style>
