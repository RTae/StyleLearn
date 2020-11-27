<template>
    <v-hover v-slot="{ hover }" id="cardshow">
        <v-card
        :elevation="hover ? 8 : 16"
        :class="{ 'on-hover': hover }"
        class="cardCourseSmallContainer"
        >
            <v-img
            height="200"
            width="303"
            :src="require('../../assets/image/subject/cardSmall/' +  subject + '.png')"
            />
            <v-sheet class="cardFooterContainer" style="background-color: #70CCFF;">
                <p class="cardInSmallText">{{ title }}</p>
            </v-sheet>
        </v-card>
    </v-hover>
</template>

<script>
export default {
  name: "CardShow",
  props: ["subject", "title"]
};
</script>

<style>
.cardCourseSmallContainer {
  border-radius: 10px;
  width: 303px;
  height: 279px;
  background-color: white;
  opacity: 0.6;
  transition: opacity 0.2s ease-in;
}

.cardCourseSmallContainer:not(.on-hover) {
  opacity: 1;
}

.cardFooterContainer {
  background-color: #70CCFF;
  display: flex;
  align-items: center;
  width: 303px;
  height: 79px;
}

.cardInSmallText {
  margin-left: 40px;
  font-family: "THSarabunNew";
  font-size: 23px;
  font-weight: bold;
}
</style>
