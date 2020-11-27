<template>
  <v-container fluid class="main" id="ProfileStudent">
    <!-- Subject title -->
    <v-row align="center" justify="center" style="margin-top: 40px">
      <v-card elevation="10" class="cardContainer">
        <p class="text">Profile</p>
      </v-card>
    </v-row>
    <!-- Body -->
    <v-row style="margin-top: 50px; margin-bottom: 100px">
      <v-col cols="0.5"></v-col>
      <v-col cols="2.5">
        <v-card class="cardSideContainer">
          <button @click="onClickProfile()" class="cardButtonContainer">
            Profile
          </button>
          <button @click="onClickEditProfile()" class="cardButtonContainer">
            Edit Profile
          </button>
          <button @click="onClickAccount()" class="cardButtonContainer">
            Account
          </button>
        </v-card>
      </v-col>
      <!-- PICTURE -->
      <v-col cols="8" align="center">
            <v-row align="center" justify="center">
              <v-img
                  src="https://upload.wikimedia.org/wikipedia/commons/9/9e/Jungkook_for_Dispatch_%22Boy_With_Luv%22_MV_behind_the_scene_shooting%2C_15_March_2019_07_(cropped).jpg"
                  height="300px"
                  max-height="315"
                  max-width="367"
                  style="border-radius: 50px;"
                  contain
              />
            </v-row>
        <v-card class="cardProfileDetail" align="start" justify-center>
          <v-row
            ><v-col cols="1"></v-col>
            <v-col>
              <p class="textDetail">First name : Jungkook</p>
            </v-col>
            <v-col>
              <p class="textDetail">Family name : Jeon</p>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="1"></v-col>
            <v-col>
              <p class="textDetail">Birthday : 1997/09/01</p>
            </v-col>
            <v-col class="textDetail">
              <p>Gender : Male</p>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="1"></v-col>
            <v-col class="textDetail">
              <p>Email : jeonjk@gmail.com</p>
            </v-col>
            <v-col class="textDetail">
              <p>Education : Bachelor's degree</p>
            </v-col>
          </v-row>
        </v-card>
      </v-col>
      <v-col cols="0.5"></v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  methods: {
    linetoGrid () {
      return this.items;
    },
    onClickProfile () {
      this.$router.push({ name: "ProfileStudent" });
    },
    onClickEditProfile () {
      this.$router.push({ name: "EditProfileStudent" });
    },
    onClickAccount () {
      this.$router.push({ name: "ProfileTutor" });
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
  height: 400px;
  width: 1000px;
  border-radius: 50px;
}

.cardImgProfile {
  height: 400px;
  width: 500px;
}

.cardProfileDetail {
  height: 230px;
  width: 1000px;
  border-radius: 50px;
  margin-top: 30px;
}

.cardContainer {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  height: 130px;
  width: 1350px;
  background-color: #70ccff;
  border-radius: 30px;
}

.text {
  font-weight: normal;
  color: white;
  font-size: 70px;
  font-family: "Average Sans", sans-serif;
}

.cardSideContainer {
  width: 250px;
  height: 400px;
  background-color: #c4c4c4;
  border-radius: 30px;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
}

.cardButtonContainer {
  width: 200px;
  height: 70px;
  background-color: #70ccff;
  border-radius: 30px;
  font-weight: normal;
  color: white;
  font-size: 22px;
  font-family: "Average Sans", sans-serif;
  margin-top: 30px;
  margin-bottom: 30px;
}
.textDetail {
  font-size: 25px;
  font-family: Delius;
}
</style>
