<template>
  <div id="header">
    <v-app-bar
      absolute
      color="secondary"
      hide-on-scroll
    >
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="../../assets/logo.png"
          transition="scale-transition"
          width="200"
        />
      </div>
      <v-spacer></v-spacer>
      <div>
        <router-link to="/signup"><button class="btnHeader">Sign Up</button></router-link>
        <router-link to="/login"><button class="btnHeader">Login In</button></router-link>
      </div>

    </v-app-bar>
  </div>
</template>

<script>
export default {
  name: 'Header'
}
</script>

<style>
.btnHeader {
  font-family: 'Average Sans',sans-serif;
  border-radius: 100px;
  margin-right:20px;
  width: 140px;
  height: 50px;
  opacity: 1;
  transition: 0.3s;
  font-size: 15px;
  text-transform: uppercase;
  color: white;
  box-shadow: 0 0 4px #999;
  cursor: pointer;
  outline: none;
}

.btnHeader {
  background-position: center;
  font-family: 'Average Sans',sans-serif;
  border-radius: 100px;
  margin-right:20px;
  width: 140px;
  height: 50px;
  opacity: 1;
  transition: 0.3s;
}
.btnHeader:hover {
  background: #47a7f5 radial-gradient(circle, transparent 1%, #47a7f5 1%) center/15000%;
}
.btnHeader:active {
  background-color: #6eb9f7;
  background-size: 100%;
  transition: background 0s;
}
</style>
