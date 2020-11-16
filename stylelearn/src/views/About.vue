<template>
  <v-container fluid class="main" id="about">
    <h1>This is an about page</h1>
    <v-row justify="center">
      <v-dialog v-model="dialog" width="600px">
        <template v-slot:activator="{ on, attrs }">
          <p color="primary" dark v-bind="attrs" v-on="on">
            Application conditions
          </p>
        </template>
        <v-card>
          <v-card-title>
            <span class="headline">Application conditions</span>
          </v-card-title>
          <v-card-text
            >Lorem ipsum dolor sit amet, semper quis, sapien id natoque elit.
            Nostra urna at, magna at neque sed sed ante imperdiet, dolor mauris
            cursus velit, velit non, sem nec. Volutpat sem ridiculus placerat
            leo, augue in, duis erat proin condimentum in a eget, sed fermentum
            sed vestibulum varius ac, vestibulum volutpat orci ut elit eget
            tortor. Ultrices nascetur nulla gravida ante arcu. Pharetra rhoncus
            morbi ipsum, nunc tempor debitis, ipsum pellentesque, vitae id quam
            ut mauris dui tempor, aptent non. Quisque turpis. Phasellus quis
            lectus luctus orci eget rhoncus. Amet donec vestibulum mattis
            commodo, nulla aliquet, nibh praesent, elementum nulla. Sit lacus
            pharetra tempus magna neque pellentesque, nulla vel erat.
            +</v-card-text
          >
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="green darken-1" text @click="dialog = false"
              >OK</v-btn
            >
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-row>
  </v-container>
</template>
