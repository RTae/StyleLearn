<template>
  <v-container fluid class="main" id="aboutus">
    <v-row justify="center">
      <v-card class="cardContainer">
        <v-row justify="center"><p class="text">ABOUT US</p></v-row>
        <v-row justify="center">
          <p>แพลตฟอร์ม StyleLearn เป็นแพลตฟอร์มที่มีจุดประสงค์เพื่อเป็นสื่อกลาง</p>
        </v-row>
        <v-row justify="center">
          <p>ระหว่างผู้เรียนและผู้สอน โดยที่ผู้เรียนสามารถเลือกเรียนเพียงบทย่อยแต่ละบทที่สนใจได้</p>
        </v-row>
        <v-row justify="center">
          <p>และแต่ละบทย่อยจะมีผู้สอนหลายคนให้ผู้เรียนได้เลือกเรียนตามสไตล์การสอนที่ถูกใจ</p>
        </v-row>
        <v-row justify="center">
          <p>และยังสามารถรับชมคลิปวิดีโอการสอนได้อย่างไม่จำกัดจำนวนครั้ง นอกจากนั้นแล้ว</p>
        </v-row>
        <v-row justify="center">
          <p>แพลตฟอร์ม StyleLearn ยังเป็นอีกหนึ่งช่องทางในการเพิ่มรายได้ให้กับผู้สอนอีกช่องทางหนึ่ง</p>
        </v-row>
      </v-card>
    </v-row>
    <div class="d-flex flex-column justify-bottom align-center">
      <v-img
        alt="bitButton"
        contain
        style="margin-top: 200px"
        src="../assets/image/etc/imgbit.png"
        width="1290"
      />
    </div>
  </v-container>
</template>

<script>
export default {
  name: "aboutus"
};
</script>

<style scoped>
.cardContainer {
  color: black;
  border-radius: 20px;
  width: 900px;
  height: 450px;
  margin-top: 20vh;
 }

.text {
  font-weight: bold;
  font-size: 30px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 50px;
  margin-bottom: 20px;
}

.textdetail {
  font-size: 18px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 40px;
}
</style>
