<template>
  <v-container fluid class="main" id="CoursesPage">
    <!-- Subject title -->
    <v-row align="center" justify="center" style="margin-top: 40px">
      <v-card elevation="10" class="cardContainer">
        <p class="textTitle">Upload My Video</p>
      </v-card>
    </v-row>
    <v-row align="center" justify="center">
      <v-form
        ref="form"
        v-model="valid"
        @submit.prevent="submitRegister"
        lazy-validation
      >
        <!-- Upload -->
        <v-row align="center" justify="center">
          <v-col>
            <v-row justify="start">
              <label>Vidoe</label>
            </v-row>
            <div>
              <v-btn
                color="primary"
                class="text-none"
                rounded
                depressed
                :loading="isSelectingUploadVideo"
                @click="onButtonClick"
              >
                <v-icon left> cloud_upload </v-icon>
                {{ buttonText }}
              </v-btn>
              <input
                ref="uploaderVideo"
                class="d-none"
                type="file"
                accept="video/*"
                @change="onFileChanged"
              />
            </div>
          </v-col>
          <v-col>
            <v-row justify="start">
              <label>Cover picture</label>
            </v-row>
            <div>
              <v-btn
                color="primary"
                class="text-none"
                rounded
                depressed
                :loading="isSelectingUploadCover"
                @click="onButtonClick"
              >
                <v-icon left> cloud_upload </v-icon>
                Upload Cover
              </v-btn>
              <input
                ref="uploader"
                class="d-none"
                type="file"
                accept="image/*"
                @change="onFileChanged"
              />
            </div>
          </v-col>
        </v-row>

        <!-- Email -->
        <v-row align="center" justify="start">
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Email:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field
                solo
                rounded
                outlined
                required
                :rules="emailRules"
                v-model="user.email"
              />
            </div>
          </v-col>
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Comfrim email:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field
                solo
                rounded
                outlined
                required
                :rules="emailRules"
                v-model="emailCon"
              />
            </div>
          </v-col>
        </v-row>
        <!-- Button -->
        <v-row justify="center">
          <button :disabled="!valid" class="signUpBtn" type="submit">
            Sign Up
          </button>
        </v-row>
      </v-form>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "coursesPage",
  components: {},
  mounted () {
    this.title = this.$route.params.titleName;
  },
  computed: {
    buttonText () {
      return this.selectedFileVideo
        ? this.selectedFileVideo.name
        : this.defaultButtonText;
    }
  },
  data: () => ({
    defaultButtonText: "Upload vidoe",
    selectedFileVideo: null,
    selectedFileCover: null,
    isSelectingUploadVideo: false,
    isSelectingUploadCover: false,
    valid: true,
    emailCon: "",
    user: {
      role: "",
      sex: "",
      firtname: "",
      familyname: "",
      birthday: null,
      edu: "",
      email: "",
      password: ""
    },
    passwordRules: [
      (v) => !!v || "Password is required",
      (v) => (v && v.length >= 6) || "Password must be more than 6 characters"
    ],
    emailRules: [
      (v) => !!v || "E-mail is required",
      (v) => /.+@.+\..+/.test(v) || "E-mail must be valid"
    ]
  }),
  methods: {
    linetoGrid () {
      return this.items;
    },

    onClickLesson () {
      console.log("ClickCourse");
    },
    onButtonClick () {
      this.isSelectingUploadVideo = true;
      window.addEventListener(
        "focus",
        () => {
          this.isSelectingUploadVideo = false;
        },
        { once: true }
      );

      this.$refs.uploaderVideo.click();
    },
    onFileChanged (e) {
      this.selectedFileVideo = e.target.files[0];
    }
  }
};
</script>

<style scoped>
.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
  margin-top: 50px;
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
.textTitle {
  font-weight: normal;
  color: white;
  font-size: 80px;
  font-family: "Average Sans", sans-serif;
}
</style>
