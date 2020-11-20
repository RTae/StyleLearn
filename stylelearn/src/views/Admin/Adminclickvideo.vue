<template>
    <v-container fluid class="main" id="admineditvideo">
        <v-row align="center" justify="center" style="margin-top: 40px">
            <v-card elevation="10" class="cardContainer">
                <p class="text">Video</p>
            </v-card>
        </v-row>
        <v-row style="margin-top:50px">
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:200px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://sv1.picz.in.th/images/2020/11/16/bHaFJ2.png"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Differential I
          </v-btn>
        </v-card>
      </v-col>
      <!--course picture 2-->
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:200px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://sv1.picz.in.th/images/2020/11/16/bHaFJ2.png"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Differential II
          </v-btn>
        </v-card>
      </v-col>
      <!--course picture 3-->
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:200px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://sv1.picz.in.th/images/2020/11/16/bHaFJ2.png"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Differential III
          </v-btn>
        </v-card>
      </v-col>
    </v-row>
    <!--course picture 4-->
    <v-row>
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:200px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://sv1.picz.in.th/images/2020/11/16/bHaFJ2.png"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Data model
          </v-btn>
        </v-card>
      </v-col>
      <!--course picture 5-->
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:200px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://sv1.picz.in.th/images/2020/11/16/bHaFJ2.png"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Integrate I
          </v-btn>
        </v-card>
      </v-col>
      <!--course picture 6-->
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:200px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://sv1.picz.in.th/images/2020/11/16/bHaFJ2.png"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Integrate II
          </v-btn>
        </v-card>
      </v-col>
    </v-row>
    </v-container>
</template>
<script>
export default {
  name: "admineditvideo",
  components: {},
  mounted () {
    this.title = this.$route.params.titleName;
  },
  data: () => ({
  }),
  methods: {
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
  align-items: center;
  justify-content: center;
  height: 188px;
  width: 890px;
  background-color: #70ccff;
  border-radius: 30px;
}

.text {
  font-weight: normal;
  color: white;
  font-size: 144px;
  font-family: "Average Sans", sans-serif;
}

.cardSideContainer {
  width: 320px;
  height: 800px;
  background-color: #c4c4c4;
  border-radius: 30px;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
}

.cardButtonContainer {
  width: 254px;
  height: 80px;
  background-color: #70ccff;
  border-radius: 30px;
  font-weight: normal;
  color: white;
  font-size: 24px;
  font-family: "Average Sans", sans-serif;
  style: "margin-top: 40px";
}
</style>
