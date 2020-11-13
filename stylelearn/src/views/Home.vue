<template>
  <v-container fluid class="main">
    <v-row class="course" align="center" justify="center">
      <div >
        <v-card width="1500" height="220" color="#70CCFF">
          <div class="d-flex align-center" align-center justify-center>
            <v-col cols="4" >
              <v-card width="500px" height="200" color="#FFFFFF">CALCULUS I</v-card>
            </v-col>
            <v-col cols="4">
              <v-card width="500px" height="200" color="#FFFFFF">CALCULUS II </v-card>
            </v-col>
            <v-col cols="4">
              <v-card width="500px" height="200" color="#FFFFFF">CALCULUS III </v-card>
            </v-col>
          </div>
        </v-card>
      </div>
    </v-row>
    <p>Courses</p>
    <p>Newest</p>
    <v-row>
      <v-col cols="3">
        <v-card width="500px" height="200" color="#FFFFFF"></v-card>
      </v-col>
      <v-col cols="3">

      </v-col>
      <v-col cols="3">

      </v-col>
      <v-col cols="3">

      </v-col>
    </v-row>
    <p>Popular</p>
  </v-container>
</template>

<script>
// @ is an alias to /src

export default {
  name: "Home",
  components: {}
};
</script>

<style scoped>
.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
  margin-top: 50px;
}
.course {
  margin-top: 50px;
}

</style>
