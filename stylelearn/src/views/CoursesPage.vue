<template>
  <v-container fluid class="main" id="CoursesPage">
    <v-row align="center" justify="center" style="margin-top:40px">
      <div>
        <v-card width="1300" height="220" rounded="30px" color="#70CCFF">
          <div class="d-flex align-center" align-center justify-center>
            <p>Calulas I</p>
          </div>
        </v-card>
      </div>
    </v-row>
    <v-row>
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:30px"
          border-radius="20px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Calculus I
          </v-btn>
        </v-card>
      </v-col>
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:30px"
          border-radius="20px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Calculus II
          </v-btn>
        </v-card>
      </v-col>
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:30px"
          border-radius="20px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Calculus III
          </v-btn>
        </v-card>
      </v-col>
      <v-col cols="3">
        <v-card
          width="300"
          height="250px"
          style="margin-left:30px"
          border-radius="20px"
          color="#70CCFF"
        >
          <v-img
            class="white--text align-end"
            height="200px"
            src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
          >
          </v-img>
          <v-btn class="botton" color="#70CCFF" width="100%" height="50px">
            Computer architecture
          </v-btn>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
// @ is an alias to /src

export default {
  name: "CoursesPage",
  components: {}
};
</script>

<style scoped>
.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
  margin-top: 50px;
}

.text {
  font-weight: normal;
  color: black;
  margin-top: 20px;
  font-size: 30px;
  font-family: 'Delius', cursive;
  margin-left: 30px;
}
.text2 {
  color: black;
  margin-top: 20px;
  font-size: 20px;
  font-family: 'Delius', cursive;
  margin-left: 90vh;
}
.botton {
  display: block;
  border-style: none;
  position: absolute;
  box-shadow: none;
  background-repeat: no-repeat;
}

</style>
