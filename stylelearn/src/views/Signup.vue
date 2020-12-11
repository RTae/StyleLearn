<template>
  <v-container fluid class="main" id="signUp">
    <v-row style="margin-top: 10vh" align="center" justify="center">
      <p class="textTitleSignUp">SIGN UP</p>
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
                :rules="[v => !!v || 'You must select role']"
                v-model="user.role"
                row
                required
              >
                <v-radio label="Student" value="a" />
                <v-radio label="Tutor" value="b" />
              </v-radio-group>
            </div>

            <!-- Sex -->
            <div class="radioLineContainer">
              <div class="text_detail">
                <p>Gender</p>
              </div>
              <v-radio-group
                :rules="[v => !!v || 'You must select gender']"
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
              <label>First Name:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field
                ref="FirstName"
                :rules="[v => !!v || 'Firstname is required']"
                v-model="user.firtname"
                solo
                rounded
                outlined
                autocomplete="firstname"
              />
            </div>
          </v-col>
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Family Name:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field
                :rules="[v => !!v || 'Familyname is required!']"
                v-model="user.familyname"
                color="primary"
                solo
                rounded
                outlined
                autocomplete="familyname"
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
                    :rules="[v => !!v || 'Birthday is required']"
                    readonly
                    solo
                    rounded
                    outlined
                    v-bind="attrs"
                    v-on="on"
                    autocomplete="birthday"
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
                v-model="eduValue"
                :items="eduTypeList"
                :rules="[v => !!v || 'Item is required']"
                label="Choose education"
                solo
                rounded
                outlined
                autocomplete="education"
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
                autocomplete="email"
              />
            </div>
          </v-col>
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Comfrim Email:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field
                solo
                rounded
                outlined
                required
                :rules="emailRules"
                v-model="emailCon"
                autocomplete="con-email"
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
                autocomplete="password"
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
                autocomplete="con-password"
              />
            </div>
          </v-col>
        </v-row>

        <!-- Agree -->
        <v-row align="center" justify="center">
          <div class="radioContainer">
            <div class="radioLineContainer">
              <v-checkbox
                :rules="[v => !!v || 'You must agree to continue!']"
                label="Do you agree?"
                required
                name="agree"
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
                Application conditions
              </v-card-title>

              <v-card-text>เงื่อนไขและข้อตกลงในการใช้บริการ บริษัท สไตส์การเรียนรู้ จำกัด (StyleLearn) นี้เป็นข้อผูกพันระหว่าง “ผู้ใช้บริการ” ซึ่งทำการสมัครเรียน การเข้าเรียนทั้งสาขาและระบบการเรียนแบบออนไลน์ ใช้บริการต่างๆ และหรือเพื่อการสื่อสารใดๆ ตามกรอบที่ “StyleLearn” กำหนด ภายใต้กรอบวัตถุประสงค์ในการใช้บริการ รวมทั้งการซื้อสินค้า สมัครเรียน ซึ่งในที่นี้รวมถึง แต่ไม่จำกัดแค่ วิดีโอคอร์สเรียน ตกลงชำระเงิน และหรือล็อกอินเข้าสู่ระบบ แล้ว “ผู้ใช้บริการ” ตกลงและยินยอมที่จะผูกพันตนรวมทั้งตกลงที่จะปฏิบัติตามเงื่อนไขในการใช้งานผลิตภัณฑ์ของ บริษัท สไตส์การเรียนรู้ จำกัด ทุกประการ  “StyleLearn” ขอสงวนสิทธิในการเปลี่ยนแปลง แก้ไข ยกเลิก เงื่อนไขการใช้บริการต่างๆ ได้ทุกเมื่อ โดยที่ “ผู้ใช้บริการ” จะต้องทำการล็อกอินเข้าสู่ระบบอย่างสม่ำเสมอเพื่อให้มั่นใจว่า “ผู้ใช้บริการ” ได้รับรู้ถึงเงื่อนไขล่าสุด ที่มีการประกาศและบังคับใช้ในแต่ละช่วงเวลา และถือเป็นพันธะสัญญาที่ผู้ใช้บริการต้องถือปฏิบัติโดยเคร่งครัด เมื่อ “ผู้ใช้บริการ” ได้ใช้ “บริการ” และหรือกระทำการใด ๆที่เกี่ยวข้องกับ “สินค้า” หลังจากเงื่อนไขใหม่มีผลบังคับใช้ แล้ว ให้ถือว่า “ผู้ใช้บริการ” ได้ยอมรับในเงื่อนไขใหม่นั้น ๆ </v-card-text>
              <v-card-text>ในการใช้ “สินค้า” และหรือ “บริการ” ใดๆ และใช้เว็บไซต์ “ผู้ใช้บริการ” ไม่มีสิทธิ์</v-card-text>
              <v-card-text>1.คัดลอก, เผยแพร่, ดัดแปลง, reverse engineer, ปรับเปลี่ยนหน้าตา, ทำลาย, ขโมยข้อมูล และหรือ ขัดขวางการทำงานของ “สินค้า” และ หรือ “บริการ” และเว็บไซต์</v-card-text>
              <v-card-text>2.แสดงตนว่าเป็นบุคคลอื่น หรือให้ข้อมูลส่วนตัวที่ไม่เป็นความจริง</v-card-text>
              <v-card-text>3.ปล่อย virus, worm, spyware, computer code, file หรือ program ใดๆ ที่จงใจ หรืออาจจะส่งผลเสีย หรือเข้าไปควบคุม การทำงานของ “สินค้า” และหรือ “บริการ” และ เว็บไซต์</v-card-text>

              <v-divider></v-divider>

              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" text @click="dialog = false">
                  ยอมรับ
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-row>

        <!-- Button -->
        <v-row justify="center">
          <button :disabled="!valid" class="signUpBtn" type="submit">
            Sign Up
          </button>
        </v-row>
      </v-form>
    </v-row>

    <!-- Image -->
    <div class="d-flex flex-column justify-bottom align-center">
      <v-img
        alt="bitButton"
        contain
        style="margin-top: 50px"
        src="../assets/image/etc/imgbit.png"
        width="1290"
      />
    </div>
    <!-- Dialog -->
    <PopUpDialog/>
  </v-container>
</template>

<script>
import PopUpDialog from "../components/popupDialog/Dialog"
export default {
  name: "signUp",
  components: {
    PopUpDialog
  },
  data: () => ({
    eduType: [
      {
        id: "a",
        name: "High school"
      },
      {
        id: "b",
        name: "Bachelor"
      },
      {
        id: "c",
        name: "Master"
      },
      {
        id: "d",
        name: "Ph.d"
      }
    ],
    eduTypeList: ["High school", "Bachelor", "Master", "Ph.d"],
    dialog: false,
    valid: true,
    eduValue: "",
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
      v => !!v || "Password is required",
      v => (v && v.length >= 6) || "Password must be more than 6 characters"
    ],
    emailRules: [
      v => !!v || "E-mail is required",
      v => /.+@.+\..+/.test(v) || "E-mail must be valid"
    ],
    date: null,
    menu: false
  }),
  watch: {
    menu (val) {
      val && setTimeout(() => (this.$refs.picker.activePicker = "YEAR"));
    }
  },
  computed: {
  },
  methods: {
    eduTypeMapValue (eduName) {
      Object.values(this.eduType).forEach(value => {
        if (value.name === eduName) {
          this.user.edu = value.id
        }
      });
    },
    save (date) {
      this.$refs.menu.save(date);
    },
    async submitRegister () {
      this.$refs.FirstName.focus();
      var state = this.$refs.form.validate();
      if (this.user.email === this.emailCon) {
        if (this.user.password === this.passwordCon) {
          if (state) {
            this.eduTypeMapValue(this.eduValue)
            this.$store.dispatch({
              type: "doRegister",
              firstname: this.user.firtname,
              familyname: this.user.familyname,
              birthday: this.user.birthday,
              sex: this.user.sex,
              email: this.user.email.trim(),
              password: this.user.password,
              role: this.user.role,
              edu: this.user.edu
            });
          }
        } else {
          this.$store.dispatch({
            type: "dialogPopup",
            value: true,
            msg: "Password must be same"
          });
        }
      } else {
        this.$store.dispatch({
          type: "dialogPopup",
          value: true,
          msg: "Email must be same"
        });
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

p.textTitleSignUp {
  font-weight: bold;
  font-size: 50px;
  color: #5c5c5c;
  font-family: "Average Sans", sans-serif;
  margin-top: 20px;
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
  background: #47a7f5 radial-gradient(circle, transparent 1%, #47a7f5 1%)
    center/15000%;
}

.signUpBtn:active {
  background-color: #6eb9f7;
  background-size: 100%;
  transition: background 0s;
}
</style>
