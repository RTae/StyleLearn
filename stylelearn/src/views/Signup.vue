<template>
  <v-container fluid class='main' id='signUp'>
    <v-row style="margin-top:10vh" align="center" justify="center">
        <p class="textTitle">SIGN UP</p>
    </v-row>
    <v-row align="center" justify="center">
    <v-form
        ref="form"
        lazy-validation
        >
      <v-row
        class="d-flex flex-row align-center"
      >
        <v-col>
          <div class = "input_button">
            First name:<v-text-field
            color="primary"
            solo
            rounded
            outlined
            />
          </div>
        </v-col>
            <div class = "input_button">
                Family name:<v-text-field
                color="primary"
                solo
                rounded
                outlined
                />
            </div>
        </v-row>
        <v-row
        class="d-flex flex-row align-center"
        >
          <v-col>
            <div class = "input_button">
              Birthday:<v-text-field
              color="primary"
              solo
              rounded
              outlined
              />
            </div>
          </v-col>
          <v-col>
            <div class = "input_button">
              Education:<v-text-field
              color="primary"
              solo
              rounded
              outlined
              />
            </div>
          </v-col>
        </v-row>
        <v-row
        class="d-flex flex-row align-center"
        >
          <v-col>
            <div class = "input_button">
              Email:<v-text-field
              color="primary"
              solo
              rounded
              outlined
              />
            </div>
          </v-col>
          <v-col>
            <div class = "input_button">
              Comfrim Email:<v-text-field
              color="primary"
              solo
              rounded
              outlined
              />
            </div>
          </v-col>
        </v-row>
        <v-row
          class="d-flex flex-row align-center"
        >
          <v-col>
            <div class = "input_button">
                Password:<v-text-field
                color="primary"
                solo
                rounded
                outlined
                />
            </div>
          </v-col>
          <v-col>
            <div class = "input_button">
                Confrim Password:<v-text-field
                color="primary"
                solo
                rounded
                outlined
                />
            </div>
          </v-col>
        </v-row>
            <button class="button">Sign Up</button>
          </v-form>
      </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'signUp'
}
</script>

<style>
  p.textTitle{
      font-weight: bold;
      font-size: 46px;
      color: #5C5C5C;
      font-family: 'Average Sans',sans-serif;
  }

  input[type=text] {
      border: 0px;
      font-size: 15px;
      border-radius: 30px;
      width: 20%;
      padding: 12px 20px;
      box-sizing: border-box;
      margin: 10px 0;
  }
  input[type=password] {
      border: 0px;
      font-size: 15px;
      border-radius: 30px;
      width: 20%;
      padding: 12px 20px;
      box-sizing: border-box;
      margin: 1px 0;
  }
  button{
      background-color: #5cbbf6;
      border: none;
      color: white;
      padding: 15px 32px;
      text-align: center;
      text-decoration: none;
      display: inline-block;
      font-size: 15px;
      margin: 4px 2px;
      cursor: pointer;
      border-radius: 20px;
  }
  div {
    font-family: 'CSChatThaiUI';
    text-align: center;
  }
  h1{
    text-align: center;
    font-weight:bold;
    font-size:50px;
  }
 h2{
   font-weight:normal;
   font-size:15px;
 }
 li{
   margin:0;
   padding: 0;
   color:#5cbbf6;
   display: inline;
   margin-right:10px;
  }
 body {
    background-color: #ecede8;
}
</style>
