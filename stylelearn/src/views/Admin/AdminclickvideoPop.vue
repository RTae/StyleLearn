<template>
  <v-container fluid class="main" id="adminclickvideoPop">
    <v-row align="center" justify="center" style="margin-top:20vh">
      <v-card class="bgcard" width="500px" height="285px">
        <v-row><label class="text">Delete</label></v-row>
        <v-row><label class="textdetail">Are you sure to delete.</label></v-row>
          <router-link to="/AdminclickvideoPop"><button class="mybtn">Yes</button></router-link>
          <router-link to="/AdminclickvideoPop"><button class="mybtn" style="background-color: red">No</button></router-link>
      </v-card>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "adminclickvideoPop"
};
</script>
<style scoped>
.bgcard {
  color: white;
  border-radius:20px;
}
.text {
  font-weight: bold;
  font-size: 20px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 30px;
  margin-left: 190px;
  margin-bottom: 20px;
}
.textdetail {
  font-size: 12px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 10px;
  margin-left: 140px;
}

.mybtn {
  font-family: "Average Sans", sans-serif;
  border-radius: 100px;
  margin-right: 20px;
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
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
body {
  background-color: #ecede8;
}
</style>
