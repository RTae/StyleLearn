<template>
    <v-dialog persistent v-model="$store.getters.getDialogState" width="500">
      <v-card>
        <v-card-title class="primary mb-6"> Alert </v-card-title>
        <v-card-text class="textDetail">
          {{ $store.getters.getDialogMsg }}
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="primary"
            text
            @click="$store.dispatch({ type: 'dialogPopup', value: false, msg: '' })"
          >
            OK
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
</template>

<script>
export default {
  name: "PopUpDialog"
};
</script>

<style>
.textDetail{
  font-size: 30px;
  font-family: "THSarabunNew"
}
</style>
