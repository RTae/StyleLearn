<template>
  <v-container fluid class="main" id="signUp">
    <v-row style="margin-top: 10vh" align="center" justify="center">
      <p class="textTitle">SIGN UP</p>
    </v-row>
    <v-row align="center" justify="center">
      <v-form
        ref="form"
        v-model="valid"
        @submit.prevent="submitRegister"
        lazy-validation
      >
        <!-- Role -->
        <v-row align="center" justify="center">
          <div class="radioContainer">
            <div class="radioLineContainer">
              <div class="text_detail">
                <p>I'm</p>
              </div>
              <v-radio-group
                :rules="[(v) => !!v || 'You must select role']"
                v-model="user.role"
                row
                required
              >
                <v-radio label="Student" value="1" />
                <v-radio label="Teacher" value="0" />
              </v-radio-group>
            </div>

            <!-- Sex -->
            <div class="radioLineContainer">
              <div class="text_detail">
                <p>Gender</p>
              </div>
              <v-radio-group
                :rules="[(v) => !!v || 'You must select gender']"
                v-model="user.sex"
                row
                required
              >
                <v-radio label="Male" value="Male" />
                <v-radio label="Female" value="Female" />
              </v-radio-group>
            </div>
          </div>
        </v-row>

        <!-- Name -->
        <v-row align="center" justify="center">
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Firstname:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field
                :rules="[(v) => !!v || 'Firstname is required']"
                v-model="user.firtname"
                solo
                rounded
                outlined
              />
            </div>
          </v-col>
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Familyname:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field
                :rules="[(v) => !!v || 'Familyname is required!']"
                v-model="user.familyname"
                color="primary"
                solo
                rounded
                outlined
              />
            </div>
          </v-col>
        </v-row>

        <v-row align="center" justify="start">
          <!-- Birthday -->
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Birthday:</label>
            </v-row>
            <v-menu
              ref="menu"
              v-model="menu"
              :close-on-content-click="false"
              transition="scale-transition"
              offset-y
              min-width="290px"
            >
              <template v-slot:activator="{ on, attrs }">
                <div class="inputFiled">
                  <v-text-field
                    v-model="user.birthday"
                    :rules="[(v) => !!v || 'Birthday is required']"
                    readonly
                    solo
                    rounded
                    outlined
                    v-bind="attrs"
                    v-on="on"
                  />
                </div>
              </template>
              <v-date-picker
                ref="picker"
                v-model="user.birthday"
                :max="new Date().toISOString().substr(0, 10)"
                min="1950-01-01"
                @change="save"
              ></v-date-picker>
            </v-menu>
          </v-col>
          <!-- Education -->
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Education:</label>
            </v-row>
            <div class="inputFiled">
              <v-select
                v-model="user.edu"
                :items="eduTypes"
                :rules="[(v) => !!v || 'Item is required']"
                label="Choose education"
                solo
                rounded
                outlined
              />
            </div>
          </v-col>
        </v-row>

        <!-- Email -->
        <v-row align="center" justify="start">
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Email:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field
                solo
                rounded
                outlined
                required
                :rules="emailRules"
                v-model="user.email"
              />
            </div>
          </v-col>
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Comfrim email:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field
                solo
                rounded
                outlined
                required
                :rules="emailRules"
                v-model="emailCon"
              />
            </div>
          </v-col>
        </v-row>

        <!-- Password -->
        <v-row align="center" justify="start">
          <v-col col="6">
            <v-row class="ml-8" justify="start">
              <label class="">Password:</label>
            </v-row>
            <div></div>
            <div class="inputFiled">
              <v-text-field
                solo
                rounded
                outlined
                v-model="user.password"
                :rules="passwordRules"
                type="password"
              />
            </div>
          </v-col>
          <v-col col="6">
            <v-row class="ml-8" justify="start">
              <label>Comfrim Password:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field
                solo
                rounded
                outlined
                :rules="passwordRules"
                v-model="passwordCon"
                type="password"
              />
            </div>
          </v-col>
        </v-row>

        <!-- Agree -->
        <v-row align="center" justify="center">
          <div class="radioContainer">
            <div class="radioLineContainer">
              <v-checkbox
                :rules="[(v) => !!v || 'You must agree to continue!']"
                label="Do you agree?"
                required
              />
            </div>
          </div>
          <!-- Dialog -->
          <v-dialog v-model="dialog" width="500">
            <template v-slot:activator="{ on, attrs }">
              <p
                class="text_detail2 text-decoration-underline"
                v-bind="attrs"
                v-on="on"
                style="margin-top: 15px; margin-left:20px"
              >
                Application conditions
              </p>
            </template>

            <v-card>
              <v-card-title class="headline grey lighten-2">
                Privacy Policy
              </v-card-title>

              <v-card-text> คิดว่านิวไม่สวยใช่มัย </v-card-text>

              <v-divider></v-divider>

              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" text @click="dialog = false">
                  ใช่
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-row>

        <!-- Button -->
        <v-row justify="center">
          <button :disabled="!valid" class="signUpBtn" type="submit">Sign Up</button>
        </v-row>
      </v-form>
    </v-row>

    <!-- Image -->
    <div class="d-flex flex-column justify-bottom align-center">
      <v-img
        alt="bitButton"
        contain
        style="margin-top: 50px"
        src="../assets/imgbit.png"
        width="1290"
      />
    </div>
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

<script scope>
import api from "../service/api";
export default {
  name: "signUp",
  components: {},
  data: () => ({
    eduTypes: ["a", "b", "c", "d", "e", "f"],
    dialog: false,
    valid: true,
    emailCon: "",
    passwordCon: "",
    user: {
      role: "",
      sex: "",
      firtname: "",
      familyname: "",
      birthday: null,
      edu: "",
      email: "",
      password: ""
    },
    passwordRules: [
      (v) => !!v || "Password is required",
      (v) => (v && v.length >= 6) || "Password must be more than 6 characters"
    ],
    emailRules: [
      (v) => !!v || "E-mail is required",
      (v) => /.+@.+\..+/.test(v) || "E-mail must be valid"
    ],
    menu: false
  }),
  mounted () {
    this.$store.dispatch({ type: "enterSignUp" });
  },
  watch: {
    menu (val) {
      val && setTimeout(() => (this.$refs.picker.activePicker = "YEAR"))
    }
  },
  methods: {
    save (date) {
      this.$refs.menu.save(date);
    },
    async submitRegister () {
      var state = this.$refs.form.validate();
      if (this.user.email == this.emailCon) {
        if (this.user.password == this.passwordCon) {
          if (state) {
            await api.register(this.user)
          }
        }else{
          this.$store.dispatch({ type: 'dialogPopup', value: true, msg: 'Password must be same' })
        }
      }else{
          this.$store.dispatch({ type: 'dialogPopup', value: true, msg: 'Email must be same' })
      }
    }
  }
};
</script>

<style scope>
.tectcondit {
  color: red;
}
.radioContainer {
  display: flex;
  flex-wrap: wrap;
  align-content: center;
  flex-direction: column;
}

.radioLineContainer {
  display: flex;
  flex-wrap: wrap;
  align-content: flex-start;
  flex-direction: row;
}

p.textTitle {
  font-weight: bold;
  font-size: 20px;
  color: #5c5c5c;
  font-family: "Average Sans", sans-serif;
}

label {
  font-weight: bold;
  font-size: 15px;
  color: #5c5c5c;
  font-family: "Average Sans", sans-serif;
}

.text_detail {
  margin-top: 20px;
  margin-right: 20px;
}

.text_detail2 {
  margin-top: 20px;
  margin-right: 20px;
  color: red;
}

.inputFiled {
  width: 375px;
  height: 32px;
  margin-top: 20px;
  margin-left: 20px;
  margin-right: 20px;
  margin-bottom: 35px;
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

.signUpBtn:hover {
  background: #47a7f5
    radial-gradient(circle, transparent 1%, #47a7f5 1%) center/15000%;
}

.signUpBtn:active {
  background-color: #6eb9f7;
  background-size: 100%;
  transition: background 0s;
}
</style>
