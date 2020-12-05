<template>
  <v-container fluid class="main" id="DetailPayment">
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
    <v-row align="center" justify="center" >
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
    <!-- Detail Bank -->
    <v-row align="center" justify="center">
      <v-card class="cardOrder">
        <p class="textTitle">Bank Transfer</p>
      </v-card>
    </v-row>
    <v-row align="center" justify="center">
      <v-card class="cardBank" >
        <v-img
          src="../../assets/image/etc/bank.png"
        >
      </v-img>
        </v-card>
    </v-row>
    <v-row align="center" justify="center" style="margin-bottom:100px;">
      <div class="buttonCardContainer">
        <button
          class="bottonNext"
          style="background-color:#EB5769;"
          @click="onClickCancle()"
          name=btnCancel
        >
          Cancle
        </button>
        <button class="bottonNext" @click="onClickCom()" name=btnNext>
          Confirm Payment
        </button>
      </div>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "DetailPayment",
  components: {},
  methods: {
    onClickCancle () {
      this.$router.push({ name: "SelectItem" });
    },
    onClickCom () {
      this.$router.push({ name: "ConfirmPayment" });
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
.cardBank{
    width: 900px;
    height: 500;
    border-radius: 30px;
    margin-top: 20px;
}
.cardTextTitle {
  width: 400px;
  font-family: "Delius";
  font-size: 25px;
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
