<template>
  <v-container id="AdminapprovevideoPop1">
    <v-row style="margin-top: 100px" align="center" justify="center">
      <v-card class="bgcard">
        <v-row><label class="text">Send to Tutor :</label></v-row>
        <v-row><label for="x" class="textStory">Describe:</label></v-row>
        <v-row align="center" justify="center">
            <input type="text" name="x" v-model="dataedit1.describe_input2" class="inputbox"/>
        </v-row>
        <v-row><router-link to="/Adminapprovevideo"><button class="mybtn" style="background-color: red">Cancle</button>
        </router-link><router-link to="/Adminapprovevideo"><button class="mybtn">Submit</button></router-link>
    </v-row>
      </v-card>
    </v-row>

  </v-container>
</template>

<script>
export default {
  name: "AdminapprovevideoPop1",
  data () {
    return {
      dataedit1: {
        describe_input2: ""
      }
    };
  }
};
</script>
<style scoped>
.bgcard {
  color: white;
  border-radius: 20px;
  width: 722px;
  height: 549px;
  align: center;
  justify: center;
}
.text {
  font-weight: bold;
  font-size: 40px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 30px;
  margin-left: 190px;
  margin-bottom: 20px;
  align: center;
  justify: center;
}
.textStory {
  font-weight: normal;
  color:black;
  font-size: 24px;
  font-family: "Average Sans", sans-serif;
  margin-left: 50px;
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
  margin-top: 30px;
  background-color: #6eb9f7;
}
.inputbox {
  background-color:white;
  width: 583px;
  height: 214px;
}

</style>
