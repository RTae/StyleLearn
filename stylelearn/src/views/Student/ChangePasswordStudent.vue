<template>
  <v-container fluid class="main" id="adminhome">
    <!-- Title -->
    <v-row align="center" justify="center" style="margin-top: 40px">
      <v-card elevation="4" class="cardContainer">
        <p class="text">Account</p>
      </v-card>
    </v-row>
    <!-- Silde bar -->
    <v-row style="margin-top: 50px; margin-bottom: 100px">
      <v-col cols="3" offset="1">
        <v-card class="cardSideContainer">
          <v-btn @click="onClickProfile()" class="sideButtonContainer" width="200" height="50" color="#70ccff" name=profile>
            Profile
          </v-btn>
          <v-btn @click="onClickEditProfile()" class="sideButtonContainer" width="200" height="50" color="#70ccff" name=editProfile>
            Edit Profile
          </v-btn>
          <v-btn @click="onClickAccount()" class="sideButtonContainer" width="200" height="50" color="#70ccff" name=password>
            Account
          </v-btn>
        </v-card>
      </v-col>
      <!-- Body -->
      <v-col cols="8">
        <!-- Detail -->
        <v-form
          ref="form"
          v-model="valid"
          @submit.prevent="submitSave"
          lazy-validation
          ng-hide="true"
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
                  autocomplete="current-password"
                  type="password"
                  ng-hide="true"
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
                  autocomplete="new-password"
                  type="password"
                  ng-hide="true"
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
                  autocomplete="con-password"
                  type="password"
                  ng-hide="true"
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
    <PopUpDialog/>
  </v-container>
</template>

<script>
import PopUpDialog from "../../components/popupDialog/Dialog"
export default {
  name: "ProfileTutor",
  components: {
    PopUpDialog
  },
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
        if (state) {
          this.user.id = this.$store.getters.getUserName
          this.$store.dispatch({
            type: "changePassword",
            id: this.user.id,
            oldPassword: this.user.currentPassword,
            newPassword: this.user.newPassword
          });
        }
      } else {
        this.$store.dispatch({
          type: "dialogPopup",
          value: true,
          msg: "New Password must be same"
        });
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
