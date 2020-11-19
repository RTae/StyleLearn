<template>
  <v-container fluid class="main" id="login">
    <v-row style="margin-top: 15vh" align="center" justify="center">
      <p class="textTitle">My Video</p>
    </v-row>
    <!--sort by btn-->
    <v-row justify="center">
      <v-col cols="1">
        <v-btn class="botton2" color="#FFFFFF" width="8%" height="50px">
          Sort by
        </v-btn>
      </v-col>
    </v-row>
    <!--course picture 1-->
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
export default {};
</script>

<style scoped>
.text2 {
  color: black;
  margin-top: 20px;
  font-size: 20px;
  font-family: "Delius", cursive;
  margin-left: 90vh;
}

.botton {
  display: block;
  border-style: none;
  position: absolute;
  box-shadow: none;
  background-repeat: no-repeat;
}

.botton2 {
  display: block;
  border-style: none;
  position: absolute;
  box-shadow: none;
  background-repeat: no-repeat;
  margin-left: 60vh;
  border-radius: 20px;
}
</style>
