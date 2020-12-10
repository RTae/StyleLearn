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
                :color="video.file === null ? 'primary' : 'green'"
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
                :items="subjectItem"
                v-model="subject"
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
              :items="courseItem"
              v-model="course"
              solo
              rounded
              outlined
              :disabled="this.subject === null ? true : false"
              :background-color="this.subject === null ? 'grey lighten-2' : 'white'"
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
                :items="lessonItem"
                v-model="video.lessonID"
                solo
                rounded
                outlined
                :disabled="this.course === null ? true : false"
                :background-color="this.course === null ? 'grey lighten-2' : 'white'"
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
              v-model="video.description"
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
    <PopUpDialog />
    <!-- Popup overlay -->
    <v-overlay :value="loading">
      <v-progress-circular
        indeterminate
        size="64"
      />
    </v-overlay>

  </v-container>
</template>

<script>
import api from "../../service/api"
import PopUpDialog from "../../components/popupDialog/Dialog"
export default {
  name: "uploadVideo",
  components: {
    PopUpDialog
  },
  async mounted () {
    const resultSubject = await api.getAllSubject()
    const resultCourse = await api.getAllCouse()
    const resultLesson = await api.getAllLeson()
    this.subjectList = resultSubject.data.result
    this.courseList = resultCourse.data.result
    this.lessonList = resultLesson.data.result
    this.loading = false
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
    },
    subjectItem () {
      var temp = []
      for (var idx = 0; idx < this.subjectList.length; idx++) {
        temp.push(this.subjectList[idx].Name)
      }
      return temp
    },
    courseItem () {
      const getSubject = this.subjectList.find(s => s.Name === this.subject);
      if (getSubject !== undefined) {
        var course = this.courseList.filter(function (sl) {
          return sl.SubjectID === getSubject.SubjectID
        });
        var temp = []
        for (var idx = 0; idx < course.length; idx++) {
          temp.push(course[idx].Name)
        }
        return temp
      } else {
        return []
      }
    },
    lessonItem () {
      const getCourse = this.courseList.find(c => c.Name === this.course);
      if (getCourse !== undefined) {
        var lesson = this.lessonList.filter(function (ll) {
          return ll.CourseID === getCourse.CourseID
        });
        var temp = []
        for (var idx = 0; idx < lesson.length; idx++) {
          temp.push(lesson[idx].Name)
        }
        return temp
      } else {
        return []
      }
    }
  },
  data: () => ({
    loading: true,
    defaultButtonTextVideo: "Upload Video",
    selectedFileVideo: null,
    isSelectingUploadVideo: false,
    valid: true,
    subject: null,
    course: null,
    lesson: null,
    subjectList: [],
    courseList: [],
    lessonList: [],
    videoPreview: null,
    video: {
      userID: "",
      lessonID: "",
      file: null,
      description: ""
    }
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
    onFileChangedVideo (e) {
      const reader = new FileReader();
      reader.readAsDataURL(event.target.files[0]);
      // This for upload
      this.video.file = e.target.files[0];
      if (this.video.file.type.slice(0, 5) !== "video") {
        this.$store.dispatch({
          type: "dialogPopup",
          value: true,
          msg: "Please upload video"
        });
        this.video.file = null
      }
    },
    submitUpload () {
      var state = this.$refs.form.validate()
      if (state) {
        if (this.video.file !== null) {
          const getLesson = this.lessonList.find(l => l.Name === this.video.lessonID);
          this.video.lessonID = getLesson.LessonID
          this.video.userID = this.$store.getters.getUserName
          this.$store.dispatch({
            type: "uploadVideo",
            userID: this.video.userID,
            lessonID: this.video.lessonID,
            description: this.video.description,
            videoFile: this.video.file,
            status: "1"
          });
        } else {
          this.$store.dispatch({
            type: "dialogPopup",
            value: true,
            msg: "You must upload video"
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
