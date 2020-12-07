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
      <v-form
          ref="form"
          v-model="valid"
          @submit.prevent="submitUpload"
          lazy-validation
      >
        <!-- Transfer To -->
        <v-row align="center" justify="start" >
          <v-col>
            <v-card-text class="textLabel">Transfer To</v-card-text>
          </v-col>
          <v-col name=bank>
            <v-select
              :rules="[v => !!v || 'Item is required']"
              class="selectField"
              label="Please select Bank"
              :items="items"
              v-model="receipt.bankTransferTo"
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
          <v-col name=yourbank>
            <v-select
              :rules="[v => !!v || 'Item is required']"
              class="selectField"
              label="Please select Your Bank "
              :items="items"
              v-model="receipt.bankTransferFrom"
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
              v-model="menuDate"
              :close-on-content-click="false"
              transition="scale-transition"
              offset-y
              min-width="290px"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-text-field
                  v-model="dateOfTransfer"
                  label="Date of Transfer"
                  :rules="[v => !!v || 'Date Transfer is required']"
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
              v-model="menuTime"
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
                  :rules="[v => !!v || 'Time Transfer is required']"
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
                v-if="menuTime"
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
              :rules="numberRule"
              v-model="receipt.amountTranfer"
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
                Upload Receipt
              </v-btn>
              <input
                ref="uploaderReceipt"
                class="d-none"
                type="file"
                accept="image/*"
                @change="onFileChangedPic"
              />
          </v-col>
        </v-row>
        <v-row style="justify-content:center">
          <v-img
            max-height="500"
            max-width="300"
            :src=imagePreview />
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
    <PopUpDialog/>
  </v-container>
</template>

<script>
import PopUpDialog from "../../components/popupDialog/Dialog"
export default {
  name: "ConfirmPayment",
  components: {
    PopUpDialog
  },
  mounted () {
    this.receipt.invoiceID = this.$route.query.invoiceID
    this.total = this.$route.query.total
  },
  computed: {},
  data: () => ({
    items: ["Kasikorn Bank (K-BANK)", "Bangkok Bank (BBL)", "Government Savings Bank (GSB)", "Krung Thai Bank (KTB)", "Siam Commercial Bank (SCB)", "Krungsri Bank (BAY)"],
    valid: true, // Form Valid status
    menuDate: false, // Popup Date status
    menuTime: false, // Popup time status
    dateOfTransfer: null, // Date
    time: null, // Time
    imagePreview: null, // Image preview
    isSelectingUploadReceipt: false, // State check image is upload
    total: null, // Total
    receipt: {
      invoiceID: null,
      bankTransferTo: null,
      bankTransferFrom: null,
      amountTranfer: "",
      image: null
    },
    numberRule: [
      v => {
        if (!v.trim()) return true;
        if (!isNaN(parseFloat(v)) && v >= 0 && v <= 999) return true;
        return "Number has to be between 0 and 999";
      },
      v => !!v || "Amount Transfer is required"
    ],
    bankType: [
      {
        id: "kbank",
        name: "Kasikorn Bank (K-BANK)"
      },
      {
        id: "bbl",
        name: "Bangkok Bank (BBL)"
      },
      {
        id: "gsb",
        name: "Government Savings Bank (GSB)"
      },
      {
        id: "ktb",
        name: "Krung Thai Bank (KTB)"
      },
      {
        id: "scb",
        name: "Siam Commercial Bank (SCB)"
      },
      {
        id: "bay",
        name: "Krungsri Bank (BAY)"
      }
    ]
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
    bankTypeMapValue (bankName) {
      for (var idx = 0; idx < this.bankType.length; idx++) {
        if (this.bankType[idx].name === bankName) {
          return this.bankType[idx].id
        }
      }
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
    },
    onFileChangedPic (e) {
      const reader = new FileReader();
      reader.onload = event => {
        // for preview
        this.imagePreview = event.target.result;
      };
      reader.readAsDataURL(event.target.files[0]);
      // This for upload
      this.receipt.image = e.target.files[0];
    },
    submitUpload () {
      var state = this.$refs.form.validate()
      if (state) {
        if (this.receipt.image !== null) {
          if (parseFloat(this.receipt.amountTranfer) === parseFloat(this.total)) {
            var datetime = new Date(this.dateOfTransfer + " " + this.time).toISOString()
            this.receipt.tranferDateTime = datetime
            this.receipt.userid = this.$store.getters.getUserName
            this.receipt.total = parseFloat(this.total)
            this.receipt.bankTransferTo = this.bankTypeMapValue(this.receipt.bankTransferTo)
            this.receipt.bankTransferFrom = this.bankTypeMapValue(this.receipt.bankTransferFrom)
            console.log(this.receipt)
          } else {
            this.$store.dispatch({
              type: "dialogPopup",
              value: true,
              msg: "The amount tranfer is incorrect"
            });
          }
        } else {
          this.$store.dispatch({
            type: "dialogPopup",
            value: true,
            msg: "Please upload recipt"
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
