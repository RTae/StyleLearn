<template>
  <v-container fluid class="main" id="signUpSec">
    <v-row justify="center">
      <v-card class="cardContainer">
        <v-row><label class="text">Successfully Register</label></v-row>
        <v-row
          ><label class="textdetail"
            >Hope you will find your style</label
          ></v-row
        >
      </v-card>
    </v-row>
    <div class="d-flex flex-column justify-bottom align-center">
      <v-img
        alt="bitButton"
        contain
        style="margin-top: 200px"
        src="../assets/imgbit.png"
        width="1290"
      />
    </div>
  </v-container>
</template>

<script>
export default {
  name: "signUpSec"
};
</script>

<style scoped>
.cardContainer {
  color: black;
  border-radius: 20px;
  width: 461px;
  height: 184px;
  margin-top: 20vh;
 }

.text {
  font-weight: bold;
  font-size: 30px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 30px;
  margin-left: 110px;
  margin-bottom: 20px;
}

.textdetail {
  font-size: 18px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 10px;
  margin-left: 140px;
}

.summitBtn {
  background-color: #5cbbf6;
  background-position: center;
  font-family: "Average Sans", sans-serif;
  border-radius: 100px;
  margin-right: 20px;
  width: 130.5px;
  height: 45px;
  opacity: 1;
  transition: 0.3s;
  font-size: 11px;
  text-transform: uppercase;
  color: white;
  box-shadow: 0 0 4px #999;
  cursor: pointer;
  outline: none;
  margin-top: 40px;
}
body {
  background-color: #ecede8;
}
</style>
