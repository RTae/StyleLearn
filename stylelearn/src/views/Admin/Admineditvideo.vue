<template>
  <v-container fluid class="main" id="admineditvideo">
    <v-row align="center" justify="center" style="margin-top: 40px">
      <v-card elevation="10" class="cardContainer">
        <p class="text">Edit Video</p>
      </v-card>
    </v-row>
    <!-- Title -->
    <v-row>
      <div class=" mt-4 mr-4 pd-6 pl-12">
        <p class="headtext">Lesson : Differentiation I</p>
        <p class="headtext">Teacher : Cherprang</p>
      </div>
    </v-row>
    <!-- VideoPlayer -->
    <v-row align="center" justify="center" style="margin-top:40px;">
        <div v-cloak>
            <v-card class="video">
                <iframe
                    id="ytplayer"
                    type="text/html"
                    width="938"
                    height="528"
                    :src="'https://www.youtube.com/watch?v=XdE1u9lHie0'"
                    frameborder="0"
            ></iframe>
            <hr />
            </v-card>
        </div>
    </v-row>
    <!-- dsecribe -->
    <v-col align="center" justify="center">
      <v-row>
        <label for="x" class="textStory">Describe:</label>
        <v-card class="inputbox">
          <input type="text" name="x" v-model="dataedit.describe_input" class="inputbox" />
        </v-card>
      </v-row>
    </v-col>
    <!-- Subject -->
    <v-col class="chooseButton">
      <v-row >
        <label class="textStory">Subject:</label>
            <v-select v-model="dataedit.subjectType" :items="subjectType" label="Choose Subject" solo rounded outlined />
      </v-row>
    </v-col>
    <!-- Course -->
    <v-col class="chooseButton">
      <v-row >
        <label class="textStory">Course:</label>
        <v-select v-model="dataedit.courseType" label="Choose Course" solo rounded outlined/>
      </v-row>
    </v-col>
    <!-- Lesson -->
    <v-col class="chooseButton">
      <v-row >
        <label class="textStory">Lesson:</label>
        <v-select v-model="dataedit.lessonType" label="Choose Lesson" solo rounded outlined/>
      </v-row>
    </v-col>
    <!-- SaveButton -->
    <v-col>
      <v-row>
        <button class="signUpBtn" type="submit">Save</button>

      </v-row>
    </v-col>
  </v-container>
</template>
<script>
export default {
  name: "admineditvideo",
  data () {
    return {
      dataedit: {
        describe_input: "",
        subjectType: ["Econ", "Math", "Ele"],
        courseType: ["first", "Tae", "new"],
        lessonType: ["555", "666", "777"]
      }
    };
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
  height: 188px;
  width: 890px;
  background-color: #70ccff;
  border-radius: 30px;
}
.headtext {
  font-weight: normal;
  color: black;
  font-size: 54px;
  font-family: "Average Sans", sans-serif;
}

.text {
  font-weight: normal;
  color: white;
  font-size: 144px;
  font-family: "Average Sans", sans-serif;
}
.textStory {
  font-weight: normal;
  color:black;
  font-size: 36px;
  font-family: "Average Sans", sans-serif;
  margin-right: 20px;
}

.signUpBtn {
  background-color: #5cbbf6;
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
  margin-top: 50px;
}
.inputbox {
  background-color:white;
  width: 983px;
  height: 122px;

}
.chooseButton {
  width: 424px;
  height: 50px;
  margin-top: 20px;
}
</style>
