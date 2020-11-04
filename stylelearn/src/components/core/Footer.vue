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
          color='font_b'
        >
          <v-card-text class="py-2 white--text text-center">
            {{ new Date().getFullYear() }} — <strong>Vuetify</strong>
          </v-card-text>
        </v-card>
    </v-footer>
  </div>
</template>

<script>
export default {
  name: 'Footer'
}
</script>
