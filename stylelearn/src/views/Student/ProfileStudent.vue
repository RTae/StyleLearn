<template>
  <v-container fluid class="main" id="adminhome">
    <!-- Title -->
    <v-row align="center" justify="center" style="margin-top: 40px">
      <v-card elevation="4" class="cardContainer">
        <p class="text">Profile</p>
      </v-card>
    </v-row>
    <!-- Silde bar -->
    <v-row style="margin-top: 50px; margin-bottom: 100px">
      <v-col cols="3" offset="1">
        <v-card class="cardSideContainer">
          <v-btn @click="onClickProfile()" class="sideButtonContainer" width="200" height="50" color="#70ccff">
            Profile
          </v-btn>
          <v-btn @click="onClickEditProfile()" class="sideButtonContainer" width="200" height="50" color="#70ccff">
            Edit Profile
          </v-btn>
          <v-btn @click="onClickAccount()" class="sideButtonContainer" width="200" height="50" color="#70ccff">
            Account
          </v-btn>
        </v-card>
      </v-col>
      <!-- Body -->
      <v-col cols="8">
        <!-- Picture -->
        <div class="d-flex flex-column align-center pa-2" style="width:55vw">
          <v-card
              class="d-flex flex-column justify-center"
              style="border-radius:10px; border:2px solid black; margin-bottom:30px"                   max-height="300"
              min-height="200"
              max-width="300"
              min-width="200"
              color="grey darken-1"
            >
              <v-img
              :lazy-src="require('../../assets/image/etc/lazy_loading.png')"
              contain
              max-height="300"
              max-width="300"
              :src="user.image"
            />
          </v-card>
          <!-- Detail -->
          <div>
            <v-card class="cardDetailContainer">
              <p class="textDetail">First name: {{ this.user.firstName }}</p>
              <p class="textDetail">Family name: {{ this.user.familyName  }}</p>
              <p class="textDetail">Birthday: {{ this.user.birthday }}</p>
              <p class="textDetail">Sex: {{ this.user.sex }}</p>
              <p class="textDetail">Email: {{ this.user.email }}</p>
              <p class="textDetail">Education: {{ this.user.edcation }}</p>
            </v-card>
          </div>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import api from "../../service/api"
import { server } from "../../service/constants";
export default {
  name: "ProfileTutor",
  async mounted () {
    const result = await api.getUser(localStorage.getItem(server.USERNAME))
    if (result.data.status === "1") {
      this.user.image = "https://stylelearn.s3-ap-southeast-1.amazonaws.com/image/" + result.data.result[0].ProfilePic.slice(0, 10) + ".png"
      this.user.firstName = result.data.result[0].Firstname
      this.user.familyName = result.data.result[0].Familyname
      this.user.birthday = result.data.result[0].Birthday.slice(0, 10)
      this.user.sex = result.data.result[0].Sex
      this.user.email = result.data.result[0].Email
      this.user.edcation = result.data.result[0].EduName
    }
  },
  data () {
    return {
      user: {
        image: "",
        firstName: "",
        familyName: "",
        birthday: "",
        sex: "",
        email: "",
        edcation: ""
      }
    }
  },
  methods: {
    onClickProfile () {
      this.$router.push({ name: "ProfileStudent" });
    },
    onClickEditProfile () {
      this.$router.push({ name: "EditProfileStudent" });
    },
    onClickAccount () {
      this.$router.push({ name: "ChangePasswordStudent" });
    }
  }
};
</script>
<style scoped>

.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
}

.cardProfile {
  height: 315px;
  width: 377px;
  border-radius: 50px;
  margin-top: 50px;
}

.cardContainer {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  height: 127px;
  width: 80vw;
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
  width: 250px;
  height: 400px;
  background-color: #DDDDDD;
  border-radius: 30px;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
}

.sideButtonContainer {
  border-radius: 30px;
  font-weight: normal;
  color: white;
  font-size: 24px;
  font-family: "Average Sans", sans-serif;
  margin-top: 40px;
}

.cardDetailContainer{
  width: 50vw;
  background-color: white;
  border:  2px solid black;
  border-radius: 10px;
  display: grid;
  grid-template-columns: auto auto;
  padding: 20px;
  margin-bottom: 20px;
}

.textDetail{
  font-size: 18px;
  font-family: "Delius"
}
</style>
