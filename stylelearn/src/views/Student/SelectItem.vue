<template>
  <v-container fluid class="main" id="SelectItem">
    <!-- Subject title -->
    <v-row align="center" justify="center" style="margin-top: 10px">
      <v-card class="cardContainer">
        <p class="text">Selected Item</p>
      </v-card>
    </v-row>
    <!-- Selected Item -->
    <v-row align="center" justify="center">
      <div class="d-flex flex-column mb-10" justify-center align-center>
        <div v-for="n in 2" :key="n">
          <v-card class="courseCard">
            <v-row align="center" justify="center">
              <v-col>
                <div class="detail">
                  <v-card-text class="cardTextTitle"
                    >Differentiation I</v-card-text
                  >
                </div>
              </v-col>
              <v-col>
                <div class="detail">
                  <v-card-text class="cardPrice">50 Bath</v-card-text>
                </div>
              </v-col>
              <v-col class="buttonDayContainer">
                <v-select
                  label="Day"
                  solo
                  rounded
                  background-color="#A8DFFE"
                  :items="items"
                ></v-select>
              </v-col>
              <v-col class="buttonContainer">
                <v-btn icon color="red">
                  <v-icon class="material-icons">cancel</v-icon>
                </v-btn>
              </v-col>
            </v-row>
          </v-card>
        </div>
      </div>
    </v-row>
    <v-row align="center" justify="center" style="margin-top:50px">
        <div style="width:800px">
            <v-row align="center" justify="end">
                <v-card class="totalCard">
                    <p class="cardTextTotal">Total : 100 Bath</p>
                </v-card>
            </v-row>
        </div>
    </v-row>
    <v-row align="center" justify="center" style="margin-bottom:100px;">
      <div class="buttonCardContainer" align="center" justify="end">
          <button
            class="bottonCom"
            @click="onClickComfirmOrder()"
          >
          Comfirm Order
          </button>
      </div>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "SelectItem",
  components: {},
  data: () => ({
    items: [1, 2, 3, 4]
  }),
  methods: {
    onClickComfirmOrder () {
      this.$router.push({ name: "SelectedItemInvoice" })
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
  width: 800px;
  background-color: #70ccff;
  border-radius: 30px;
  margin-bottom: 50px;
  margin-top: 50px;
}
.text {
  font-weight: normal;
  color: white;
  font-size: 60px;
  font-family: "Average Sans", sans-serif;
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
  justify-content: center;
  align-self: center;
  border-radius: 30px;
  width: 400px;
  height: 50px;
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
  color: black;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  font-family: "Delius", cursive;
  font-size: 25px;
  margin-top:10px;
}
.cardTextTitle {
  display: flex;
  width: 300px;
  height: 80px;
  font-family: "Delius", cursive;
  font-size: 20px;
  margin-left: 10px;
  background-color: white;
  align-items: center;
  border-radius: 30px;
}
.cardPrice {
  display: flex;
  width: 175px;
  height: 123;
  font-family: "Delius", cursive;
  font-size: 20px;
}
.buttonContainer {
  display: flex;
  justify-items: center;
  width: 225px;
  height: 123;
}
.buttonDayContainer {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 50px;
  height: 100px;
  margin-top: 30px;
}
.buttonCardContainer {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  margin-top: 50px;
  width: 800px;
  height: 80px;
}
</style>
