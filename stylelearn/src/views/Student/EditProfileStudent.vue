<template>
  <v-container fluid class="main" id="EditProfileStudent">
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
      <!-- PICTURE -->
      <v-col cols="3.5" align="center">
        <v-row align="center" justify="center">
          <v-card
            width="300px"
            max-height="300"
            align="center"
            style="border-radius: 50px;"
          >
            <v-img
              src="https://upload.wikimedia.org/wikipedia/commons/9/9e/Jungkook_for_Dispatch_%22Boy_With_Luv%22_MV_behind_the_scene_shooting%2C_15_March_2019_07_(cropped).jpg"
              height="300px"
              contain
            ></v-img>
          </v-card>
        </v-row>
        <v-row align="center" justify="start">
          <v-col>
            <v-card-text class="textLabel">Upload receipt</v-card-text>
          </v-col>
          <v-col>
            <v-btn
              color="primary"
              class="text-none"
              height="60"
              width="200px"
              elevation="5"
              rounded
              depressed
              :loading="isSelectingUploadReceipt"
              @click="onButtonClickUploadReceipt"
            >
              <v-icon left> cloud_upload </v-icon>
              {{ buttonTextReciept }}
            </v-btn>
            <input
              ref="uploaderReceipt"
              class="d-none"
              type="file"
              accept="image/*"
              @change="onFileChangedReceipt"
            />
          </v-col>
        </v-row>
      </v-col>
      <v-col cols="5" align="center">
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
              <p>Sex : Male</p>
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
  height: 400px;
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
  width: 300px;
  height: 700px;
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
  margin-bottom: 30px;
}
.textDetail {
  font-size: 25px;
  font-family: Delius;
}
</style>
