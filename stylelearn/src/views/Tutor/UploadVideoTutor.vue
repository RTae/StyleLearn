<template>
  <v-container fluid class="main" id="signUp">
    <!-- Title -->
    <v-row style="margin-top: 30px" align="center" justify="center">
      <v-card elevation="10" class="cardContainer">
        <p class="textTitle">Upload My Video</p>
      </v-card>
    </v-row>
    <!-- Form -->
    <v-row justify="center" style="margin-top: 30px">
      <v-form
        ref="form"
        v-model="valid"
        @submit.prevent="submitUpload"
        lazy-validation
      >
        <!-- Upload -->
        <v-row align="center" justify="center" style="margin-top:10px">
          <v-col>
            <v-row class="ml-2">
              <label class="textLabel">Video</label>
            </v-row>
            <div>
            <v-btn
                color="primary"
                class="text-none"
                height="50"
                elevation="5"
                rounded
                depressed
                :loading="isSelectingUploadVideo"
                @click="onButtonClickUploadVideo"
              >
              <v-icon left> cloud_upload </v-icon>
              {{ buttonTextVideo }}
              </v-btn>
              <input
                ref="uploaderVideo"
                class="d-none"
                type="file"
                accept="video/*"
                @change="onFileChangedVideo"
                />
            </div>
          </v-col>
          <v-col>
            <v-row class="ml-2" justify="start">
              <label class="textLabel">Cover picture</label>
            </v-row>
            <v-btn
                color="primary"
                class="text-none"
                height="50"
                elevation="5"
                rounded
                depressed
                :loading="isSelectingUploadCover"
                @click="onButtonClickUploadCover"
              >
              <v-icon left> cloud_upload </v-icon>
                {{ buttonTextCover }}
              </v-btn>
              <input
                ref="uploaderCover"
                class="d-none"
                type="file"
                accept="image/*"
                @change="onFileChangedCover"
              />
          </v-col>
        </v-row>
        <!-- Subject Course -->
        <v-row align="center" justify="start">
          <v-col>
            <v-row class="ml-2" justify="start">
              <label class="textLabel">Subject</label>
            </v-row>
             <v-select
                :rules="[v => !!v || 'Item is required']"
                class="selectField"
                :items="items"
                solo
                rounded
                outlined
              />
          </v-col>
          <v-col>
            <v-row class="ml-2" justify="start">
              <label class="textLabel">Course</label>
            </v-row>
            <v-select
              :rules="[v => !!v || 'Item is required']"
              class="selectField"
              :items="items"
              solo
              rounded
              outlined
              background-color="grey lighten-2"
              :disabled="true"
            />
          </v-col>
        </v-row>

        <!-- Lesson -->
        <v-row align="center" justify="start">
          <v-col cols="12">
            <v-row class="ml-2" justify="start">
              <label class="textLabel">Lesson</label>
            </v-row>
             <v-select
                :rules="[v => !!v || 'Item is required']"
                class="selectField"
                :items="items"
                solo
                rounded
                outlined
                background-color="grey lighten-2"
                :disabled="true"
              />
          </v-col>
        </v-row>

        <!-- Desciption -->
        <v-row align="center" justify="center">
          <v-col cols="12">
            <v-row class="ml-2" justify="start">
              <label class="textLabel">Description</label>
            </v-row>
            <v-textarea
              v-model="title"
              counter
              rounded
              maxlength="100"
              full-width
              single-line
              solo
            />
          </v-col>
        </v-row>

        <!-- Button -->
        <v-row align="center" justify="center">
          <v-col cols="2">
            <button :disabled="!valid" class="submitBtn" type="submit">
              Submit
            </button>
          </v-col>
        </v-row>
      </v-form>
    </v-row>

    <!-- Dialog -->
    <v-dialog v-model="$store.getters.getDialogState" width="500">
      <v-card>
        <v-card-title class="primary mb-6"> Alert </v-card-title>
        <v-card-text class="popUpText">
          {{ $store.getters.getDialogMsg }}
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="primary"
            text
            @click="$store.dispatch({ type: 'dialogPopup', value: false, msg: '' })"
          >
            OK
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Popup overlay -->
    <v-overlay :value="$store.getters.getDialogLoading">
      <v-progress-circular
        indeterminate
        size="64"
      />
    </v-overlay>

  </v-container>
</template>

<script>
export default {
  name: "uploadVideo",
  components: {},
  mounted () {
    this.title = this.$route.params.titleName;
  },
  computed: {
    buttonTextVideo () {
      return this.selectedFileVideo
        ? this.selectedFileVideo.name
        : this.defaultButtonTextVideo;
    },
    buttonTextCover () {
      return this.selectedFileCover
        ? this.selectedFileCover.name
        : this.defaultButtonTextCover;
    }
  },
  data: () => ({
    defaultButtonTextVideo: "Upload Vidoe",
    defaultButtonTextCover: "Upload Cover",
    selectedFileVideo: null,
    selectedFileCover: null,
    isSelectingUploadVideo: false,
    isSelectingUploadCover: false,
    valid: true,
    title: "",
    emailCon: "",
    items: ["a", "b", "c", "d"],
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
    onButtonClickUploadVideo () {
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
    onButtonClickUploadCover () {
      this.isSelectingUploadCover = true;
      window.addEventListener(
        "focus",
        () => {
          this.isSelectingUploadCover = false;
        },
        { once: true }
      );

      this.$refs.uploaderCover.click();
    },
    onFileChangedVideo (e) {
      this.selectedFileVideo = e.target.files[0];
    },
    onFileChangedCover (e) {
      this.selectedFileCover = e.target.files[0];
    },
    submitUpload () {
      var state = this.$refs.form.validate()
      if (state) {
        if (this.selectedFileVideo !== null) {
        } else {
          this.$store.dispatch({
            type: "dialogPopup",
            value: true,
            msg: "Yout must upload video"
          });
        }
      }
    }
  }
};
</script>

<style scoped>
.main {
  background-color: rgb(239, 239, 239);
  min-height: 100vh;
}
.titleContainer{
  display: flex;
  justify-content: center;
  margin-top: 100px;
}
.formContainer{
  width: 70vw;
}
.cardContainer {
  display: flex;
  flex-direction: column;
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

.selectField{
  width: 445px;
}

.textField{
  border-radius: 10px;
}

.submitBtn {
  background-color: #6eb9f7;
  background-position: center;
  font-family: "Average Sans", sans-serif;
  border-radius: 100px;
  margin-right: 20px;
  width: 130px;
  height: 45px;
  opacity: 1;
  transition: 0.3s;
  font-size: 13px;
  text-transform: uppercase;
  color: white;
  box-shadow: 0 0 4px #999;
  cursor: pointer;
  outline: none;
}

.submitBtn:hover {
  background: #47a7f5 radial-gradient(circle, transparent 1%, #47a7f5 1%)
    center/15000%;
}

.submitBtn:active {
  background-color: #6eb9f7;
  background-size: 100%;
  transition: background 0s;
}

.textLabel{
  font-weight: normal;
  color: black;
  font-size: 20px;
  font-family: "Delius";
  margin-bottom: 10px;
}
</style>
