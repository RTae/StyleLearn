<template>
  <v-container fluid class="main" id="EditProfileStudent">
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
                    max-width="400"
                    color="grey darken-1"
                  >
                    <v-img
                    :lazy-src="require('../../assets/image/etc/lazy_loading.png')"
                    contain
                    :src="user.image"
                  />
                </v-card>
              <v-btn
                  color="primary"
                  class="text-none"
                  height="50"
                  elevation="3"
                  rounded
                  depressed
                  :loading="isSelectingUploadPic"
                  @click="onButtonClickUploadPic"
                >
                <v-icon left> cloud_upload </v-icon>
                  {{ buttonTextCover }}
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
            <v-col class="d-flex justify-center" cols="10">
              <v-btn type="submit" class="cardButtonContainer" width="200" height="50" color="#70ccff">
                save
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
  name: "EditProfileStudent",
  data () {
    return {
      user: {
        image: "https://picsum.photos/id/11/500/300",
        firstName: "Jungkook",
        familyName: "Jeon",
        birthday: null,
        sex: "Male",
        email: "jk@gmail.com",
        edu: "d"
      },
      isSelectingUploadPic: false,
      defaultButtonTextPic: "Upload Picture",
      selectedFilePic: null,
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
  mounted () {
    this.eduValueMapType(this.user.edu)
    this.user.birthday = new Date("1997-09-01").toISOString().substr(0, 10)
  },
  computed: {
    buttonTextCover () {
      return this.selectedFilePic
        ? "New pic"
        : this.defaultButtonTextPic;
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
      this.$router.push({ name: "ChangePasswordTutor" });
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
      this.selectedFilePic = e.target.files[0];
    },
    save (date) {
      this.$refs.menu.save(date);
    },
    eduTypeMapValue (eduName) {
      Object.values(this.eduType).forEach(value => {
        if (value.name === eduName) {
          this.user.edu = value.id
        }
      });
    },
    eduValueMapType (eduType) {
      Object.values(this.eduType).forEach(value => {
        if (value.id === eduType) {
          this.eduTypeValue = value.name
        }
      });
    },
    submitSave () {
      var state = this.$refs.form.validate();
      if (state) {
        this.eduTypeMapValue(this.eduTypeValue)
        console.log(this.selectedFilePic)
        console.log(this.user)
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
