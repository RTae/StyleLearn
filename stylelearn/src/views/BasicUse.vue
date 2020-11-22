<template>
  <v-container fluid class="main" id="basicuse">
    <v-row justify="center">
      <v-card class="cardContainer">
        <v-row justify="center"><p class="text">BASIC USE</p></v-row>
        <v-row justify="center" class="textdetail">
          <p>แพลตฟอร์ม StyleLearn เป็นแพลตฟอร์มที่มีจุดประสงค์เพื่อเป็นสื่อกลาง</p>
        </v-row>
        <v-row  justify="center">
            <v-card class="imgContain">
                <v-img
                class="cardImg"
                src="../assets/courses/cal1.png"
                ></v-img>
            </v-card>
        </v-row>
        <v-row justify="center" class="textdetail">
          <p>แพลตฟอร์ม StyleLearn เป็นแพลตฟอร์มที่มีจุดประสงค์เพื่อเป็นสื่อกลาง</p>
        </v-row>
        <v-row  justify="center">
            <v-card class="imgContain">
                <v-img
                class="cardImg"
                src="../assets/courses/cal2.png"
                ></v-img>
            </v-card>
        </v-row>
        <v-row justify="center" class="textdetail">
          <p>แพลตฟอร์ม StyleLearn เป็นแพลตฟอร์มที่มีจุดประสงค์เพื่อเป็นสื่อกลาง</p>
        </v-row>
        <v-row  justify="center">
            <v-card class="imgContain">
                <v-img
                class="cardImg"
                src="../assets/courses/cal3.png"
                ></v-img>
            </v-card>
        </v-row>
      </v-card>
    </v-row>
    <div class="d-flex flex-column justify-bottom align-center">
      <v-img
        alt="bitButton"
        contain
        style="margin-top: 200px"
        src="../assets/imgbit.png"
        width="1290"
      />
    </div>
  </v-container>
</template>

<script>
export default {
  name: "basicuse"
};
</script>

<style scoped>
.cardContainer {
  color: black;
  border-radius: 20px;
  width: 900px;
  height: 1500px;
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
.imgContain{
    height: 300px;
    width: 700px;
}
.cardImg {
  height: 100%;
  width: 100%;
  border-radius: 30px;
}
</style>
