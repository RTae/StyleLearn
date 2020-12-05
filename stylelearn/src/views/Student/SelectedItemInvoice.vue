<template>
  <v-container fluid class="main" id="SelectedItemInvoice">
    <!--title -->
    <v-row align="center" justify="center" style="margin-top: 10px">
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
            <v-row  v-for="lesson in lessonBukect" :key="lesson.id" align="center" justify="center">
              <v-col>
                <p class="cardTextTitle">{{lesson.name }}</p>
              </v-col>
              <v-col>
                <p class="cardPrice">{{lesson.day}} Day</p>
              </v-col>
              <v-col>
                <p class="learnday">{{parseInt(lesson.day) * 50}} Bath</p>
              </v-col>
            </v-row>
        </v-card>
    </v-row>
    <v-row align="center" justify="center" style="margin-top:50px">
        <div style="width:900px">
            <v-row align="center" justify="end">
                <v-card class="totalCard">
                    <p class="cardTextTotal">Total: {{ calculateTotal }} Baht</p>
                </v-card>
            </v-row>
        </div>
    </v-row>

    <v-row align="center" justify="center" style="margin-bottom:100px;">
      <div class="buttonCardContainer">
          <button
            class="bottonNext"
            style="background-color:#EB5769;"
            @click="onClickBack()"
            name=btnCancle
          >
          Back
          </button>
          <button
            class="bottonNext"
            @click="onClickNext()"
            name=btnNext
          >
          Next
          </button>
      </div>
    </v-row>
  </v-container>
</template>

<script>
import { server } from "../../service/constants"
export default {
  name: "SelectedItemInvoice",
  components: {},
  async mounted () {
    this.lessonBukect = await JSON.parse(localStorage.getItem(server.BUKECT))
    for (var i = 0; i < this.lessonBukect.length; i++) {
      this.calculateTotal += parseInt(this.lessonBukect[i].day)
    }
    this.calculateTotal = this.calculateTotal * 50
  },
  data: () => ({
    lessonBukect: [],
    calculateTotal: 0
  }),
  methods: {
    onClickBack () {
      this.$router.go(-1)
    },
    onClickNext () {
      this.$router.push({ name: "DetailPayment" })
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
  font-family: "Delius";
  margin-left: 20px;
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
.buttonCardContainer {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  margin-top: 50px;
  width: 800px;
  height: 80px;
}
.bottonNext{
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
.bottonNext:hover {
  background: #47a7f5 radial-gradient(circle, transparent 1%, #47a7f5 1%)
    center/15000%;
  color: #000;
  outline: none;
}
.totalCard {
  display: flex;
  justify-content: center;
  align-self: center;
  border-radius: 30px;
  width: 400px;
  height: 50px;
  background-color: white;
}
.cardTextTotal {
  margin-top:10px;
  color: black;
  font-family: "Delius";
  font-size: 25px;
}
</style>
