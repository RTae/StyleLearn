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
            <p class="cardTextTitle">Invoice ID : {{ invoiceID }}</p>
          </v-col>
        </v-row>
        <v-row align="center" justify="center">
          <v-col>
            <p class="cardTextTitle">Total : {{ total }} Bath</p>
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
import api from "../../service/api"
import { server } from "../../service/constants";
export default {
  name: "DetailPayment",
  components: {},
  async mounted () {
    this.$store.commit("SET_DIALOG_LOADING", true)
    const id = localStorage.getItem(server.USERNAME)
    var result = await api.getUnPaidInvoice(id)
    if (result[1]) {
      this.invoiceID = result[0][0].Invoice_id
      this.total = result[0][0].Total
      this.$store.commit("SET_DIALOG_LOADING", false)
    } else {
      this.$router.push({ name: "Home" });
    }
  },
  data () {
    return {
      invoiceID: null,
      total: null
    }
  },
  methods: {
    async onClickCancle () {
      this.$store.commit("SET_DIALOG_LOADING", true)
      var result = await api.cancelInvoice(this.invoiceID)
      if (result.status === "1") {
        this.$store.commit("SET_DIALOG_LOADING", false)
        this.$router.push({ name: "Home" });
        this.$store.commit("SET_UNPAID_STATE", false)
      } else {
        this.$store.commit("SET_DIALOG_LOADING", false)
        this.$router.push({ name: "Home" });
        this.$store.dispatch({ type: "dialogPopup", value: true, msg: result.msg })
      }
    },
    onClickCom () {
      this.$router.push({ name: "ConfirmPayment", query: { invoiceID: this.invoiceID, total: this.total } });
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
