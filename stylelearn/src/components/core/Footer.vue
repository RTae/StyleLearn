<template>
  <div id="footer">
      <v-footer
    dark
    padless
  >
    <v-card
      class="flex"
      flat
      tile
    >
      <v-card-title class="black">
        <div>
          <v-img
            alt="Vuetify Logo"
            class="shrink mr-2"
            contain
            src="../../assets/logo_footer.png"
            transition="scale-transition"
            width="200"
          />
          <v-btn
            v-for="icon in icons"
            :key="icon"
            icon
          >
            <v-icon size="24px">
              {{ icon }}
            </v-icon>
          </v-btn>
        </div>
        <v-spacer></v-spacer>
        <div class="d-flex flex-column justify-center">
          <pre>
            Contact Us
          </pre>
          <pre class="flex-start">
            09 226 38737
            Info@stylelearn.com
            Bangkok Nuchan
          </pre>
        </div>
      </v-card-title>

      <v-card-text class="py-2 white--text text-center">
        {{ new Date().getFullYear() }} — <strong>StyleLearn</strong>
      </v-card-text>
    </v-card>
  </v-footer>
  </div>
</template>

<script>
export default {
  name: 'Footer',
  data: () => ({
    icons: [
      'mdi-facebook',
      'mdi-twitter',
      'mdi-linkedin',
      'mdi-instagram'
    ]
  })
}
</script>

<style>

pre{
  font-size: 10px;
  text-align: left;
}
</style>
