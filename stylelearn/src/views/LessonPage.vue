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
      <div class="d-flex flex-column mb-10">
        <v-card
          v-for="n in 7"
          :key="n"
          class="pa-2"
          outlined
          tile
          style="width: 100px; hight: ;"
        >
          <v-hover v-slot="{ hover }">
            <v-card
              :elevation="hover ? 8 : 16"
              :class="{ 'on-hover': hover }"
              class="courseCard"
              @click="onClickCourse()"
            >
              <v-card-text class="cardTextTitle"> CALCULUS III </v-card-text>
            </v-card>
          </v-hover>
        </v-card>
      </div>
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
