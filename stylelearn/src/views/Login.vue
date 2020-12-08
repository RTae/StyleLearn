<template>
  <v-container fluid class="main" id="login">
    <v-row style="margin-top: 10vh" align="center" justify="center">
      <p class="textTitle">LOGIN</p>
    </v-row>
    <v-row align="center" justify="center">
      <v-form
        ref="form"
        v-model="valid"
        @submit.prevent="submit"
        lazy-validation
      >
        <!-- Username field -->
        <div class="input_button">
          <v-text-field
            color="primary"
            solo
            rounded
            outlined
            v-model="account.email"
            required
            autocomplete="username"
            :rules="emailRules"
            placeholder="Email"
            name="email"
          />
        </div>
        <!-- Password field -->
        <div class="input_button">
          <v-text-field
            solo
            rounded
            outlined
            placeholder="Password"
            required
            min="9"
            autocomplete="password"
            :rules="passwordRules"
            v-model="account.password"
            :type="showPassword ? 'text' : 'password'"
            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            @click:append="showPassword = !showPassword"
            name="password"
          />
        </div>

        <!-- ETC -->
        <v-row justify="center">
          <p style="margin-right: 10px" class="text">Don't have an account?</p>
          <router-link to="/signup">
            <p class="text">Sign up</p>
          </router-link>
        </v-row>

        <!-- Button -->
        <v-row justify="center">
          <button :disabled="!valid" class="signInBtn" type="submit">Sign In</button>
        </v-row>

      </v-form>
    </v-row>

    <!-- Image -->
    <div class="d-flex flex-column justify-bottom align-center">
      <v-img
        alt="bitButton"
        contain
        style="margin-top: 100px"
        src="../assets/image/etc/imgbit.png"
        width="1290"
      />
    </div>
    <PopUpDialog />
  </v-container>
</template>

<script>
import PopUpDialog from "../components/popupDialog/Dialog"
export default {
  name: "login",
  components: {
    PopUpDialog
  },
  data () {
    return {
      showPassword: false,
      account: {
        email: "",
        password: ""
      },
      valid: true,
      passwordRules: [
        (v) => !!v || "Password is required",
        (v) => (v && v.length >= 6) || "Password must be more than 6 characters"
      ],
      emailRules: [
        (v) => !!v || "E-mail is required",
        (v) => /.+@.+\..+/.test(v) || "E-mail must be valid"
      ]
    };
  },
  methods: {
    submit () {
      var state = this.$refs.form.validate();
      if (state) {
        this.$store.dispatch({
          type: "doLogin",
          email: this.account.email,
          password: this.account.password
        })
      }
    }
  }
};
</script>

<style scope>
.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
}

.input_button {
  width: 423px;
  height: 60px;
  margin: 20px;
}

.textTitle {
  font-weight: bold;
  font-size: 70px;
  color: #5c5c5c;
  font-family: "Average Sans", sans-serif;
  margin-top: 20px;
}

.text {
  font-size: 12px;
  font-family: "Average Sans", sans-serif;
}

.signInBtn {
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

.signInBtn:hover {
  background: #47a7f5
    radial-gradient(circle, transparent 1%, #47a7f5 1%) center/15000%;
}
.signInBtn:active {
  background-color: #6eb9f7;
  background-size: 100%;
  transition: background 0s;
}

.popUpText {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 20px;
}

</style>
