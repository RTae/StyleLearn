<template>
  <v-container fluid class="main" id="login">
    <v-row style="margin-top: 10vh" align="center" justify="center">
      <p class="textTitle">Upload My Video</p>
    </v-row>
    <v-row justify="center">
      <v-form ref="form" lazy-validation>
        <!--name-->
        <v-row align="center" justify="center">
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Frist name:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Family name:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
        </v-row>
        <!--birthday and education-->
        <v-row align="center" justify="start">
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Birthday:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Education:</label>
            </v-row>
            <div class="inputFiled">
              <v-select
                :items="items"
                label="Choose education"
                solo
                rounded
                outlined
              />
            </div>
          </v-col>
        </v-row>
        <v-row align="center" justify="start">
          <v-col>
            <v-row class="ml-8" justify="start">
              <label>Email:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
          <v-col>
            <!--email-->
            <v-row class="ml-8" justify="start">
              <label>Comfrim email:</label>
            </v-row>
            <div class="inputFiled">
              <v-text-field color="primary" solo rounded outlined />
            </div>
          </v-col>
        </v-row>
        <v-row align="center" justify="center">
          <div class="radioContainer">
            <div class="radioLineContainer">
              <v-radio-group v-model="row" row>
                <v-radio label="I agree" value="radio-1" />
              </v-radio-group>
            </div>
          </div>
              <v-dialog v-model="dialog" width="600px">
                <template v-slot:activator="{ on, attrs }">
                  <div class="tectcondit">
                    <p class="text-decoration-underline" v-bind="attrs" v-on="on" style="margin-top:15px">
                       Application conditions
                     </p>
                  </div>

                </template>
                <v-card>
                  <v-card-title>
                    <span class="headline">Application conditions</span>
                  </v-card-title>
                  <v-card-text
                    >Lorem ipsum dolor sit amet, semper quis, sapien id natoque
                    elit. Nostra urna at, magna at neque sed sed ante imperdiet,
                    dolor mauris cursus velit, velit non, sem nec. Volutpat sem
                    ridiculus placerat leo, augue in, duis erat proin
                    condimentum in a eget, sed fermentum sed vestibulum varius
                    ac, vestibulum volutpat orci ut elit eget tortor. Ultrices
                    nascetur nulla gravida ante arcu. Pharetra rhoncus morbi
                    ipsum, nunc tempor debitis, ipsum pellentesque, vitae id
                    quam ut mauris dui tempor, aptent non. Quisque turpis.
                    Phasellus quis lectus luctus orci eget rhoncus. Amet donec
                    vestibulum mattis commodo, nulla aliquet, nibh praesent,
                    elementum nulla. Sit lacus pharetra tempus magna neque
                    pellentesque, nulla vel erat.
                  </v-card-text>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="#5cbbf6" text @click="dialog = false"
                      >OK</v-btn
                    >
                  </v-card-actions>
                </v-card>
              </v-dialog>
        </v-row>
        <v-row justify="center">
          <router-link to="/signupsuccessfully">
            <button class="signUpBtn">SIGN UP</button>
          </router-link>
        </v-row>
      </v-form>
    </v-row>
  </v-container>
</template>

<script>
export default {};
</script>

<style>
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

.inputFiled {
  width: 375px;
  height: 32px;
  margin-top: 20px;
  margin-left: 20px;
  margin-right: 20px;
  margin-bottom: 35px;
}
</style>
