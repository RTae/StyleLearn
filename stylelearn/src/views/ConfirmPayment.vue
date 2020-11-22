<template>
  <v-container fluid class="main" id="ConfirmPayment">
    <!-- Title -->
    <v-row style="margin-top: 10vh" align="center" justify="center">
      <v-card elevation="10" class="cardContainer">
        <p class="textTitle">ConfirmPayment</p>
      </v-card>
    </v-row>
    <!-- Form -->
    <v-row justify="center" style="margin-top: 30px">
      <v-form>
        <!-- Transfer To -->
        <v-row align="center" justify="start">
          <v-col>
              <v-card-text class="textLabel" >Transfer To</v-card-text>
          </v-col>
          <v-col>
            <v-select
              :rules="[v => !!v || 'Item is required']"
              class="selectField"
              label="select Bank "
              :items="items"
              solo
              rounded
              outlined
            />
          </v-col>
        </v-row>

        <!-- Transfer From -->
        <v-row align="center" justify="start">
          <v-col>
              <v-card-text class="textLabel" >Transfer From</v-card-text>
          </v-col>
          <v-col>
            <v-select
              :rules="[v => !!v || 'Item is required']"
              class="selectField"
              label="select Your Bank "
              :items="items"
              solo
              rounded
              outlined
            />
          </v-col>
        </v-row>

        <!-- Date Of Transfer  -->
        <v-row align="center" justify="start">
          <v-col>
              <v-card-text class="textLabel" >Date Of Transfer</v-card-text>
          </v-col>
          <v-col>
            <v-select
              :rules="[v => !!v || 'Item is required']"
              class="selectField"
              label="select Your Bank "
              :items="items"
              solo
              rounded
              outlined
            />
          </v-col>
        </v-row>

        <!-- Time of Transfer  -->
        <v-row align="center" justify="start">
          <v-col>
              <v-card-text class="textLabel" >Time of Transfer</v-card-text>
          </v-col>
          <v-col>
            <v-select
              :rules="[v => !!v || 'Item is required']"
              class="selectField"
              label="select Your Bank "
              :items="items"
              solo
              rounded
              outlined
            />
          </v-col>
        </v-row>

        <!-- Amount Transfer  -->
        <v-row align="center" justify="start">
          <v-col>
              <v-card-text class="textLabel" >Amount Transfer</v-card-text>
          </v-col>
          <v-col>
            <v-select
              :rules="[v => !!v || 'Item is required']"
              class="selectField"
              label="select Your Bank "
              :items="items"
              solo
              rounded
              outlined
            />
          </v-col>
        </v-row>

        <!-- Upload receipt  -->
        <v-row align="center" justify="start">
          <v-col>
              <v-card-text class="textLabel" >Upload receipt</v-card-text>
          </v-col>
          <v-col>
            <v-select
              :rules="[v => !!v || 'Item is required']"
              class="selectField"
              label="select Your Bank "
              :items="items"
              solo
              rounded
              outlined
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
            @click="
              $store.dispatch({ type: 'dialogPopup', value: false, msg: '' })
            "
          >
            OK
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
export default {
  name: "ConfirmPayment",
  components: {},
  mounted () {
    this.title = this.$route.params.titleName;
  },
  computed: {},
  data: () => ({
    items: ["a", "b", "c", "d"]
  })
};
</script>

<style scoped>
.main {
  background-color: rgb(239, 239, 239);
  min-height: 100vh;
  min-width: 100vw;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
}
.titleContainer {
  display: flex;
  justify-content: center;
  margin-top: 100px;
}
.formContainer {
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

.selectField {
  width: 445px;
}

.textField {
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

.textLabel {
    height: 80px;
  font-weight: bold;
  color: black;
  font-size: 20px;
  font-family: "Delius";
  margin-bottom: 10px;
}
</style>
