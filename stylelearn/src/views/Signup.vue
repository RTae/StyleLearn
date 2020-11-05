<template>
  <v-container fluid class='main'>
  <h1>This is an Signup page Yooo</h1>
  <div id='SignUp'>
    <h1>SIGN UP</h1>
      First name:<input v-model='message' type='text'><br/>
      Family name:<input v-model='message' type='text'><br/>
      Birthday:<input v-model='message' type='text'><br/>
      Education:<input v-model='message' type='text'><br/>
      Email:<input v-model='message' type='text'><br/>
      Comfrim Email:<input v-model='message' type='text'><br/>
      Password:<input type='password' id='myPsw' value=''><br/>
      Confrim Password:<input type='password' id='myPsw' value=''><br/>
      <ul><li><router-link to='/home' tag='button1'>Sign up</router-link></li></ul>
  </div>
  </v-container>
</template>

<script>
export default {
  name: 'signUp'
}
</script>

<style>
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
