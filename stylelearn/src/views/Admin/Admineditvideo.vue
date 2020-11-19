<template>
    <v-container fluid class="main" id="admineditvideo">
        <v-row align="center" justify="center" style="margin-top: 40px">
            <v-card elevation="10" class="cardContainer">
                <p class="text">Edit Video</p>
            </v-card>
        </v-row>
    </v-container>
</template>
<script>
export default {
  name: "admineditvideo",
  components: {},
  mounted () {
    this.title = this.$route.params.titleName;
  },
  data: () => ({
  }),
  methods: {
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

.text {
  font-weight: normal;
  color: white;
  font-size: 144px;
  font-family: "Average Sans", sans-serif;
}

.cardSideContainer {
  width: 320px;
  height: 800px;
  background-color: #c4c4c4;
  border-radius: 30px;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
}

.cardButtonContainer {
  width: 254px;
  height: 80px;
  background-color: #70ccff;
  border-radius: 30px;
  font-weight: normal;
  color: white;
  font-size: 24px;
  font-family: "Average Sans", sans-serif;
  style: "margin-top: 40px";
}
</style>
