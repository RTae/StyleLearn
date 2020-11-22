<template>
  <v-container fluid class="main" id="adminhome">
    <!-- Subject title -->
    <v-row align="center" justify="center" style="margin-top: 40px">
      <v-card elevation="10" class="cardContainer">
        <p class="text">Profile</p>
      </v-card>
    </v-row>
    <!-- Body -->
    <v-row style="margin-top: 50px; margin-bottom: 100px">
      <v-col cols="3">
        <v-card class="cardSideContainer">
          <button @click="onClickVideo()" class="cardButtonContainer">
            Profile
          </button>
          <button @click="onClickCustomer()" class="cardButtonContainer">
            Edit Profile
          </button>
          <button @click="onClickApporvevideo()" class="cardButtonContainer">
            Account
          </button>
          <button @click="onClickApprovepayment()" class="cardButtonContainer">
            Manage Bank Account
          </button>
        </v-card>
      </v-col>
      <v-col cols="9" align-start style="background-color: gray;">
        <v-card class="cardProfile">
            <v-row>
                <v-col cols="1"></v-col>
                <v-col>
                    <p class="textDetail">First name : Sarunya</p>
                </v-col>
                <v-col>
                    <p class="textDetail">Family name : Prasert</p>
                </v-col>
            </v-row>
            <v-row>
                <v-col cols="1"></v-col>
                <v-col>
                    <p class="textDetail">Birthday : 1999/05/04</p>
                </v-col>
                <v-col class="textDetail">
                    <p>Sex : Female</p>
                </v-col>
            </v-row>
            <v-row>
                <v-col cols="1"></v-col>
                <v-col class="textDetail">
                    <p>Email : sarunya@hotmail.com</p>
                </v-col>
                <v-col class="textDetail">
                    <p>Education : Bachelor's degree</p>
                </v-col>
            </v-row>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  methods: {
    linetoGrid () {
      return this.items;
    },

    onClickLesson () {
      console.log("ClickCourse");
    },
    onClickVideo () {
      this.$router.push({ name: "ProfileTutor" });
    },
    onClickCustomer () {
      this.$router.push({ name: "ProfileTutor" });
    },
    onClickApporvevideo () {
      this.$router.push({ name: "ProfileTutor" });
    },
    onClickApprovepayment () {
      this.$router.push({ name: "ProfileTutor" });
    }
  }
};
</script>
<style scoped>

.cardProfile {
  height: 500px;
  width: 1200px;
  border-radius: 50px;
  margin-top: 50px;
}

.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
}

.cardContainer {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  height: 127px;
  width: 890px;
  background-color: #70ccff;
  border-radius: 30px;
}

.text {
  font-weight: normal;
  color: white;
  font-size: 80px;
  font-family: "Average Sans", sans-serif;
}

.cardSideContainer {
  width: 320px;
  height: 500px;
  background-color: #c4c4c4;
  border-radius: 30px;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
}

.cardButtonContainer {
  width: 250px;
  height: 70px;
  background-color: #70ccff;
  border-radius: 30px;
  font-weight: normal;
  color: white;
  font-size: 22px;
  font-family: "Average Sans", sans-serif;
  margin-top: 30px;
}
.textDetail{
    font-size: 25px;
    font-family: Delius;
}
</style>
