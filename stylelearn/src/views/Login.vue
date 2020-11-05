<template>
  <v-container>
    <v-row class="title-row">
      <v-col cols="12">
        <div class="title">
          <h1>Login</h1>
        </div>
      </v-col>
    </v-row>

    <v-row class="username-input-row">
      <v-col cols="12">
        <div class="username_input">
          <v-text-field 
            type="text" 
            placeholder="Enter your email">
          </v-text-field>
        </div>
      </v-col>
    </v-row>

    <v-row class="password-input-row">
      <v-col cols="12">
        <div class="password_input">
          <v-text-field 
            type="password" 
            id="myPsw" 
            value="" 
            placeholder="Enter your Password">
          </v-text-field>
        </div>
      </v-col>
    </v-row>

    <v-row class="help-text-row">
      <v-col cols="3">
        <div class="help-text">
          <h3>Don't have an account?</h3>
        </div>
      </v-col>
      
      <v-col cols="1">
        <div class="signup_input">
          <router-link to="/signup">Sign Up</router-link>
        </div>
      </v-col>

      <v-col cols="8">
        <div class="password_input">
          <router-link to="/forgetPassword">Forget Password?</router-link>
        </div>
      </v-col>

    </v-row>


  </v-container>
</template>

<script>
export default  {
  name: "login"
  
};
</script>


<style>
.username-input-row{
    border: 0px , black;
    font-size: 15px;
    border-radius: 30px;
    width: 30%;
    padding: 12px 20px;
    box-sizing: border-box;
    margin: 10px 0;  
}
.password-input-row{
    border: 0px , black;
    font-size: 15px;
    border-radius: 30px;
    width: 30%;
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
div{
  font-family: 'CSChatThaiUI';
  text-align: center;
}
 h1{
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

            
