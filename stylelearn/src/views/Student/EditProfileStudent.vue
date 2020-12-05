<template>
  <v-container fluid class="main" id="adminhome">
    <!-- Title -->
    <v-row align="center" justify="center" style="margin-top: 40px">
      <v-card elevation="4" class="cardContainer">
        <p class="text">Edit Profile</p>
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
        >
          <v-row>
            <div class="d-flex justify-center pa-2" style="width:55vw">
              <!-- Picture -->
              <v-col class="d-flex flex-column  justify-start" cols="5">
                <v-card
                    class="d-flex flex-column  justify-center"
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
                    :src="imageURL"
                  />
                </v-card>
              <v-btn
                  color="primary"
                  class="text-none"
                  height="50"
                  elevation="3"
                  rounded
                  depressed
                  max-width="300"
                  :loading="isSelectingUploadPic"
                  @click="onButtonClickUploadPic"
                >
                <v-icon left> cloud_upload </v-icon>
                  Upload File
                </v-btn>
                <input
                  ref="uploaderPic"
                  class="d-none"
                  type="file"
                  accept="image/*"
                  @change="onFileChangedPic"
                />
              </v-col>
              <v-col cols="7">
                <!-- First Name -->
                <v-row class="ml-2 mb-3" justify="start">
                  <label>First Name</label>
                </v-row>
                <v-text-field
                  ref="FirstName"
                  :rules="[v => !!v || 'Firstname is required']"
                  v-model="user.firstName"
                  solo
                  rounded
                  outlined
                />
                <!-- Family Name -->
                <v-row class="ml-2 mb-3" justify="start">
                  <label>Family Name</label>
                </v-row>
                <v-text-field
                  ref="FirstName"
                  :rules="[v => !!v || 'Firstname is required']"
                  v-model="user.familyName"
                  solo
                  rounded
                  outlined
                />
                <!-- Birthday -->
                <v-row class="ml-2 mb-3" justify="start">
                  <label>Birthday</label>
                </v-row>
                <v-menu
                  ref="menu"
                  v-model="menu"
                  :close-on-content-click="false"
                  transition="scale-transition"
                  offset-y
                  min-width="290px"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="user.birthday"
                      :rules="[v => !!v || 'Birthday is required']"
                      readonly
                      solo
                      rounded
                      outlined
                      v-bind="attrs"
                      v-on="on"
                    />
                  </template>
                  <v-date-picker
                    ref="picker"
                    v-model="user.birthday"
                    :max="new Date().toISOString().substr(0, 10)"
                    min="1900-01-01"
                    @change="save"
                  />
                </v-menu>
                <!-- Sex -->
                <v-row class="ml-2 mb-3" justify="start">
                  <label>Sex</label>
                </v-row>
                <v-select
                  :rules="[v => !!v || 'Item is required']"
                  v-model="user.sex"
                  :items="sexItemList"
                  solo
                  rounded
                  outlined
                />
                <!-- Education -->
                <v-row class="ml-2 mb-3" justify="start">
                  <label>Education</label>
                </v-row>
                <v-select
                  :rules="[v => !!v || 'Item is required']"
                  v-model="eduTypeValue"
                  :items="eduTypeList"
                  solo
                  rounded
                  outlined
                />
              </v-col>
            </div>
          </v-row>
          <v-row>
            <v-col class="d-flex justify-center" cols="7">
              <v-btn type="submit" class="sideButtonContainer" width="200" height="50" color="#70ccff">
                save
              </v-btn>
            </v-col>
          </v-row>
        </v-form>
      </v-col>
    </v-row>
    <PopUpDialog/>
    <v-overlay :value="$store.getters.getDialogLoading">
      <v-progress-circular
        indeterminate
        size="64"
      />
    </v-overlay>
  </v-container>
</template>

<script>
import api from "../../service/api"
import { server } from "../../service/constants";
import PopUpDialog from "../../components/popupDialog/Dialog"
export default {
  name: "ProfileTutor",
  components: {
    PopUpDialog
  },
  data () {
    return {
      user: {
        image: null,
        firstName: "",
        familyName: "",
        birthday: "",
        sex: "",
        email: "",
        edcation: ""
      },
      isSelectingUploadPic: false,
      defaultButtonTextPic: "Upload Picture",
      imageURL: null,
      valid: true,
      date: null,
      menu: false,
      eduTypeValue: "",
      sexItemList: ["Male", "Female"],
      eduType: [
        {
          id: "a",
          name: "High school"
        },
        {
          id: "b",
          name: "Bachelor"
        },
        {
          id: "c",
          name: "Master"
        },
        {
          id: "d",
          name: "Ph.d"
        }
      ],
      eduTypeList: ["High school", "Bachelor", "Master", "Ph.d"]
    }
  },
  async mounted () {
    const id = localStorage.getItem(server.USERNAME)
    const result = await api.getUser(id)
    if (result.data.status === "1") {
      this.user.image = null
      this.user.firstName = result.data.result[0].Firstname
      this.user.familyName = result.data.result[0].Familyname
      this.user.birthday = new Date(result.data.result[0].Birthday.slice(0, 10)).toISOString().substr(0, 10)
      this.user.sex = result.data.result[0].Sex
      this.user.email = result.data.result[0].Email
      this.eduTypeValue = result.data.result[0].EduName
      this.imageURL = "https://stylelearn.s3-ap-southeast-1.amazonaws.com/image/" + result.data.result[0].ProfilePic
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
    onButtonClickUploadPic () {
      this.isSelectingUploadPic = true;
      window.addEventListener(
        "focus",
        () => {
          this.isSelectingUploadPic = false;
        },
        { once: true }
      );

      this.$refs.uploaderPic.click();
    },
    onFileChangedPic (e) {
      const reader = new FileReader();
      reader.onload = event => {
        // for preview
        this.imageURL = event.target.result;
      };
      reader.readAsDataURL(event.target.files[0]);
      // This for upload
      this.user.image = e.target.files[0];
    },
    save (date) {
      this.$refs.menu.save(date);
    },
    eduTypeMapValue (eduName) {
      Object.values(this.eduType).forEach(value => {
        if (value.name === eduName) {
          this.user.edcation = value.id
        }
      });
    },
    async submitSave () {
      var state = this.$refs.form.validate();
      this.$store.commit("SET_DIALOG_LOADING", true)
      if (state) {
        this.eduTypeMapValue(this.eduTypeValue)
        this.user.id = this.$store.getters.getUserName
        this.$store.dispatch({
          type: "editProfile",
          id: this.user.id,
          firstName: this.user.firstName,
          familyname: this.user.familyName,
          birthday: this.user.birthday,
          sex: this.user.sex,
          edu: this.user.edcation,
          pic: this.user.image
        });
      }
    }
  },
  watch: {
    menu (val) {
      val && setTimeout(() => (this.$refs.picker.activePicker = "YEAR"));
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
