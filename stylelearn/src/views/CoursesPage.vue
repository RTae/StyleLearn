<template>
  <v-container fluid class="main" id="CoursesPage">
    <v-row align="center" justify="center" style="margin-top:40px">
      <v-card class="cardcontainer">
        <p class="text">MATHEMATIC</p>
      </v-card>
    </v-row>
    <v-row>
         <div class="inputFiled">
              <v-select
                :items="items"
                label="Sort by"
                solo
                rounded
                outlined
              />
            </div>
    </v-row>
    <v-row>
      <v-col v-for="n in 4" :key="n" cols="6" md="3">
        <v-card
          widt
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

.cardcontainer {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  height: 257px;
  width: 1300px;
  background-color: #70ccff;
  border-radius: 30px;
}

.text {
  font-weight: normal;
  color: white;
  font-size: 100px;
  font-family: "Average Sans", sans-serif;
}
.botton {
  display: block;
  border-style: none;
  position: absolute;
  box-shadow: none;
  background-repeat: no-repeat;
}
</style>
