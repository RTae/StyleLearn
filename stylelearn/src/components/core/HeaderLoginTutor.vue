<template>
  <div id="header">
    <v-app-bar app color="secondary" height="70">
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="../../assets/image/main/logo.png"
          transition="scale-transition"
          width="200"
        />
      </div>
      <v-spacer></v-spacer>
      <v-row class="hidden-sm-and-down" align="center" justify="end">
        <v-toolbar-items>
          <v-btn text>
            <router-link to="/">
              <p class="text">Home</p>
            </router-link>
          </v-btn>
          <v-btn text>
            <router-link to="/basicuse">
              <p class="text">Basic Use</p>
            </router-link>
          </v-btn>
          <v-btn text>
            <router-link to="/about"> <p class="text">About Us</p></router-link>
          </v-btn>
          <v-menu botton left>
            <template v-slot:activator="{ on, attrs }">
              <button dark icon v-bind="attrs" v-on="on">
                <v-img
                  class="iconBarPro"
                  contain
                  outline:none
                  src="../../assets/image/etc/profile.png"
                />
              </button>
            </template>

            <v-list>
              <v-list-item v-for="(item, i) in items" :key="i">
                <v-list-item-title
                  ><v-hover v-slot="{ hover }"
                    ><button
                      :elevation="hover ? 8 : 0"
                      :class="{ 'on-hover-review': hover }"
                      class="btnPro"
                      color="#70ccff"
                      @click="onClickHover(item.title)"
                    >
                      {{ item.title }}
                    </button></v-hover
                  ></v-list-item-title
                >
              </v-list-item>
            </v-list>
          </v-menu>
        </v-toolbar-items>
      </v-row>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  name: "Header",
  data: () => ({
    items: [
      { title: "Account" },
      { title: "Log Out" }
    ]
  }),
  methods: {
    onClickHover (functionName) {
      if (functionName === "Account") {
      } else if (functionName === "Log Out") {
        this.$store.dispatch({
          type: "doLogout"
        });
      }
    }
  }
};
</script>

<style scope>
.iconBar {
  width: 34px;
  height: 34px;
  margin-left: 20px;
  margin-top: 8px;
}
.iconBarPro {
  width: 40px;
  height: 40px;
  margin-left: 20px;
  margin-right: 20px;
  margin-top: 8px;
}
.iconBarLine {
  width: 10px;
  height: 34px;
  margin-left: 20px;
  margin-top: 8px;
}
.btnPro {
  background-color: white;
  width: 150px;
  height: 50px;
  outline: none;
}
.btnPro:hover {
  color: #47a7f5;
  outline: none;
}
.text {
  color: black;
  margin-top: 20px;
  font-family: "Average Sans", sans-serif;
  font-size: 20px;
}
</style>
