<template>
  <v-container fluid class="main" id="ChangePasswordStudent">
    <!-- Subject title -->
    <v-row align="center" justify="center" style="margin-top: 40px">
      <v-card elevation="4" class="cardContainer">
        <p class="text">Account</p>
      </v-card>
    </v-row>
    <!-- Body -->
    <v-row style="margin-top: 50px; margin-bottom: 100px">
      <v-col cols="3" offset="1">
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
      <!-- BODY -->
      <v-col cols="8">
        <!-- Detail -->
        <v-form
          ref="form"
          v-model="valid"
          @submit.prevent="submitSave"
          lazy-validation
        >
          <v-row>
            <div class="d-flex justify-center pa-2" style="width:55vw">
              <v-col cols="7">
                  <!-- Current Password -->
                  <v-row class="ml-2 mb-3" justify="start">
                    <label>Current Password</label>
                  </v-row>
                  <v-text-field
                    ref="currentPassword"
                    :rules="this.passwordRules"
                    v-model="user.currentPassword"
                    solo
                    rounded
                    outlined
                    type="password"
                  />
                  <!-- New Password -->
                  <v-row class="ml-2 mb-3" justify="start">
                    <label>New Password</label>
                  </v-row>
                  <v-text-field
                    ref="newPassword"
                    :rules="this.passwordRules"
                    v-model="user.newPassword"
                    solo
                    rounded
                    outlined
                    type="password"
                  />
                  <!-- Con Password -->
                  <v-row class="ml-2 mb-3" justify="start">
                    <label>Confirm New Password</label>
                  </v-row>
                  <v-text-field
                    ref="FirstName"
                    :rules="this.passwordRules"
                    v-model="conPassword"
                    solo
                    rounded
                    outlined
                    type="password"
                  />
              </v-col>
            </div>
          </v-row>
          <v-row>
            <v-col class="d-flex justify-center" cols="10">
              <v-btn type="submit" class="sideButtonContainer" width="200" height="50" color="#70ccff">
                Save
              </v-btn>
            </v-col>
          </v-row>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "ProfileTutor",
  data () {
    return {
      valid: true,
      conPassword: "",
      passwordRules: [
        v => !!v || "Password is required",
        v => (v && v.length >= 6) || "Password must be more than 6 characters"
      ],
      user: {
        userName: "",
        currentPassword: "",
        newPassword: ""
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
    },
    submitSave () {
      var state = this.$refs.form.validate();
      if (this.user.newPassword === this.conPassword) {
        var checkPassword = "123qwe"
        if (this.user.currentPassword === checkPassword) {
          if (state) {
            this.user.userName = this.$store.getters.getUserName
            console.log(this.user)
          }
        } else {
          console.log("Current password is not true")
        }
      } else {
        console.log("Password must be same")
      }
    }
  }
};
</script>
<style scoped>
.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
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

.sideButtonContainer {
  border-radius: 30px;
  font-weight: normal;
  color: white;
  font-size: 24px;
  font-family: "Average Sans", sans-serif;
  margin-top: 40px;
}

.textDetail {
  font-size: 25px;
  font-family: Delius;
}
</style>
