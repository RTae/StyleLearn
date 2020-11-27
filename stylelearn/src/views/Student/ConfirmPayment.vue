<template>
  <v-container fluid class="main" id="ConfirmPayment">
    <!-- Title -->
    <v-row style="margin-top: 10vh" align="center" justify="center">
      <v-card elevation="4" class="cardContainer">
        <p class="textTitle">Confirm Payment</p>
      </v-card>
    </v-row>
    <!-- Form -->
    <v-row justify="center" style="margin-top: 30px">
      <v-form>
        <!-- Transfer To -->
        <v-row align="center" justify="start" >
          <v-col>
            <v-card-text class="textLabel">Transfer To</v-card-text>
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
            <v-card-text class="textLabel">Transfer From</v-card-text>
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
            <v-card-text class="textLabel">Date Of Transfer</v-card-text>
          </v-col>
          <v-col>
            <v-menu
              ref="menu"
              v-model="menu"
              :close-on-content-click="false"
              transition="scale-transition"
              offset-y
              min-width="290px"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-text-field
                  v-model="dateOfTransfer"
                  label="Date of Transfer"
                  prepend-inner-icon="mdi-calendar"
                  readonly
                  v-bind="attrs"
                  v-on="on"
                  class="selectField"
                  solo
                  rounded
                  outlined
                ></v-text-field>
              </template>
              <v-date-picker
                ref="picker"
                v-model="dateOfTransfer"
                :max="new Date().toISOString().substr(0, 10)"
                min="1950-01-01"
                @change="save"
              ></v-date-picker>
            </v-menu>
          </v-col>
        </v-row>

        <!-- Time of Transfer  -->
        <v-row align="center" justify="start">
          <v-col>
            <v-card-text class="textLabel">Time of Transfer</v-card-text>
          </v-col>
          <v-col>
            <v-menu
              ref="menu"
              v-model="menu2"
              :close-on-content-click="false"
              :nudge-right="40"
              :return-value.sync="time"
              transition="scale-transition"
              offset-y
              max-width="290px"
              min-width="290px"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-text-field
                  v-model="time"
                  label="select Time"
                  prepend-inner-icon="mdi-clock-time-four-outline"
                  readonly
                  v-bind="attrs"
                  v-on="on"
                  class="selectField"
                  solo
                  rounded
                  outlined
                ></v-text-field>
              </template>
              <v-time-picker
                v-if="menu2"
                v-model="time"
                full-width
                @click:minute="$refs.menu.save(time)"
              ></v-time-picker>
            </v-menu>
          </v-col>
        </v-row>

        <!-- Amount Transfer  -->
        <v-row align="center" justify="start">
          <v-col>
            <v-card-text class="textLabel">Amount Transfer</v-card-text>
          </v-col>
          <v-col>
            <v-text-field
              class="selectField"
              solo
              rounded
              outlined
            />
          </v-col>
        </v-row>

        <!-- Upload receipt  -->
        <v-row align="center" justify="start">
          <v-col>
            <v-card-text class="textLabel">Upload receipt</v-card-text>
          </v-col>
          <v-col>
            <v-btn
                color="primary"
                class="text-none"
                height="60"
                width="445px"
                elevation="0"
                rounded
                depressed
                :loading="isSelectingUploadReceipt"
                @click="onButtonClickUploadReceipt"
              >
              <v-icon left> cloud_upload </v-icon>
              {{ buttonTextReciept }}
              </v-btn>
              <input
                ref="uploaderReceipt"
                class="d-none"
                type="file"
                accept="image/*"
                />
          </v-col>
        </v-row>
        <!-- Button -->
        <v-row justify="center">
          <v-dialog
            v-model="dialog"
            persistent
            max-width="300"
          >
            <template v-slot:activator="{ on, attrs }">
              <v-btn
                color="blue lighten-2"
                class="submitBtn"
                height="50px"
                dark
                v-bind="attrs"
                v-on="on"
              >
                Submit
              </v-btn>
            </template>
            <v-card>
              <v-card-title class="headline">
                Confirm Payment
              </v-card-title>
              <v-card-text>Are you sure to comfirm this payment?</v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn
                  color="#EB5757"
                  text
                  @click="dialog = false"
                >
                  No
                </v-btn>
                <v-btn
                  color="#70ccff"
                  text
                  @click="dialog = false"
                >
                  Yes
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-row>
      </v-form>
    </v-row>
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
    items: ["KASIKORN BANK (K-BANK)", "BANGKOK BANK (BBL)", "Government Savings Bank (GSB)", "Krung Thai Bank (KTB)", "Siam Commercial Bank (SCB)", "Krungsri Bank (BAY)"],
    dateOfTransfer: null,
    menu: false,
    time: null,
    menu2: false,
    modal2: false,
    files: [],
    dialog: false
  }),
  watch: {
    menu (val) {
      val && setTimeout(() => (this.$refs.picker.activePicker = "YEAR"))
    }
  },
  methods: {
    save (date) {
      this.$refs.menu.save(date)
    },
    onButtonClickUploadReceipt () {
      this.isSelectingUploadReceipt = true;
      window.addEventListener(
        "focus",
        () => {
          this.isSelectingUploadReceipt = false;
        },
        { once: true }
      );

      this.$refs.uploaderReceipt.click();
    }
  }
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
  margin-bottom: 30px;
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
  margin-top:50px;
  margin-bottom: 100px;
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
