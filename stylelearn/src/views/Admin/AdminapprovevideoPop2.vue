<template>
  <v-container fluid class="main" id="AdminapprovevideoPop2">
    <v-row align="center" justify="center" style="margin-top:20vh">
      <v-card class="bgcard">
        <v-row align="center" justify="center"><label class="textdetail">Are you sure to approve video.</label></v-row>
        <v-row align="center" justify="center" ><router-link to="/Adminapprovevideo"><button class="mybtn" style="background-color: red">No</button></router-link>
        <router-link to="/Adminapprovevideo"><button class="mybtn">Yes</button></router-link></v-row>
      </v-card>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "AdminapprovevideoPop2"
};
</script>
<style scoped>
.bgcard {
  color: white;
  border-radius: 20px;
  width: 891px;
  height: 269px;
  align: center;
  justify: center;
}
.text {
  font-weight: bold;
  font-size: 48px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 30px;
  margin-left: 190px;
  margin-bottom: 20px;
}
.textdetail {
  font-size: 24px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 50px;
}

.mybtn {
  font-family: "Average Sans", sans-serif;
  border-radius: 100px;
  margin-right: 300px;
  margin-top: 100px;
  width: 130px;
  height: 45px;
  opacity: 1;
  transition: 0.3s;
  font-size: 12px;
  text-transform: uppercase;
  color: white;
  box-shadow: 0 0 4px #999;
  cursor: pointer;
  outline: none;
  background-color: #6eb9f7;
  display: flex;
  align-items: center;
  justify-content: center;
}
body {
  background-color: #ecede8;
}
</style>
