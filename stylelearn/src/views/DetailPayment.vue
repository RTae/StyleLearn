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
        <p class="textTitle">Invoice</p>
      </v-card>
    </v-row>
    <v-row align="center" justify="center">
      <v-card class="cardOrderContainer">
        <v-row align="center" justify="center">
          <v-col>
            <p class="cardTextTitle">Invoice ID : i000000001</p>
          </v-col>
        </v-row>
        <v-row align="center" justify="center">
          <v-col>
            <p class="cardTextTitle">Total : 150 Bath</p>
          </v-col>
        </v-row>
      </v-card>
    </v-row>
    <v-row align="center" justify="center">
      <v-card class="cardOrder">
        <p class="textTitle">Bank Transfer</p>
      </v-card>
    </v-row>
    <v-row>
      <v-img
        height="443"
        width="1010"
        src="../assets/bank.png"
      >
      </v-img>
    </v-row>

    <v-row align="center" justify="center" style="margin-bottom:100px;">
      <div class="buttonCardContainer">
        <button
          class="bottonNext"
          style="background-color:#EB5769;"
          @click="onClickBack()"
        >
          Back
        </button>
        <button class="bottonNext" @click="onClickNext()">
          Next
        </button>
      </div>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "SelectedItemInvoice",
  components: {},
  methods: {
    onClickBack () {
      this.$router.push({ name: "SelectItem" });
    },
    onClickNext () {
      this.$router.push({ name: "SelectItem" });
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
.bottonNext {
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
.bottonNext:active {
  background-color: #ffce70;
  outline: none;
}
.bottonBack {
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
.bottonBack:hover {
  background: #47a7f5 radial-gradient(circle, transparent 1%, #47a7f5 1%)
    center/15000%;
  color: #000;
  outline: none;
}
.bottonBack:active {
  background-color: #ffce70;
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
  margin-top: 10px;
  color: black;
  font-family: "Delius";
  font-size: 25px;
}
</style>
