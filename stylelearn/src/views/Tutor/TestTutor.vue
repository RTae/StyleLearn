<template>
  <v-container fluid class="main" id="login">
    <v-row style="margin-top:10vh" align="center" justify="center">
      <p class="textTitle">LOGIN</p>
    </v-row>
    <v-row align="center" justify="center">
      <div style="margin-top:10px">
        <p>sex</p>
      </div>
      <v-radio-group v-model="row" row>
        <v-radio label="Option 1" value="radio-1" />
        <v-radio label="Option 2" value="radio-2" />
      </v-radio-group>
    </v-row>
    <v-row align="center" justify="center">
      <button class="button">Sign In</button>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "login"
};
</script>

<style>
.main {
  background: rgb(239, 239, 239);
  min-height: 100vh;
}

.input_button {
  width: 423px;
  height: 60px;
  margin: 20px;
}

p.textTitle {
  font-weight: bold;
  font-size: 46px;
  color: #5c5c5c;
  font-family: "Average Sans", sans-serif;
}

p.text {
  font-weight: bold;
  font-size: 12px;
  font-family: "Average Sans", sans-serif;
}

.button {
  font-family: "Average Sans", sans-serif;
  border-radius: 100px;
  margin-right: 20px;
  width: 140px;
  height: 50px;
  opacity: 1;
  transition: 0.3s;
  font-size: 15px;
  text-transform: uppercase;
  color: rgb(255, 255, 255);
  box-shadow: 0 0 4px rgb(0, 206, 206);
  cursor: pointer;
  outline: none;
}

.button {
  background-position: center;
  font-family: "Average Sans", sans-serif;
  border-radius: 100px;
  margin-right: 20px;
  width: 140px;
  height: 50px;
  opacity: 1;
  transition: 0.3s;
}
.button:hover {
  background: #47a7f5 radial-gradient(circle, transparent 1%, #47a7f5 1%)
    center/15000%;
}
.button:active {
  background-color: #6eb9f7;
  background-size: 100%;
  transition: background 0s;
}
</style>
