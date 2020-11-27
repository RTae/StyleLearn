<template>
  <v-container fluid class="main" id="basicuse">
    <v-row justify="center">
      <v-card class="cardContainer">
        <v-row justify="center"><p class="text">BASIC USE</p></v-row>
        <v-row justify="center" class="textdetail">
          <p>เมื่อทำการเข้าสู่ระบบในส่วนของ Student จะขึ้นแถบดังรูปด้านล่าง</p>
        </v-row>
        <v-row  justify="center">
            <v-card class="imgContain">
                <v-img
                class="cardImg"
                src="../assets/image/barstudent.png"
                ></v-img>
            </v-card>
        </v-row>
        <v-row class="textdetail1">
          <p>หมายเลข 1 แสดงหน้าแรกของ StyleLearn</p>
          </v-row>
        <v-row class="textdetail2">
        <p>หมายเลข 2 แสดงหน้า Basic use ของ StyleLearn</p>
        </v-row>
        <v-row class="textdetail2">
        <p>หมายเลข 3 แสดงหน้า About us ของ StyleLearn</p>
        </v-row>
        <v-row  class="textdetail2">
        <p>หมายเลข 4 แสดงหน้ารายการของคอร์สเรียนที่ได้เลือก</p>
        </v-row>
        <v-row class="textdetail2">
        <p>หมายเลข 5 หน้าต่าง pop up ที่แสดงดังรูปถัดไป</p>
        </v-row>
        <v-row justify="center">
            <v-card class="imgContain1">
                <v-img
                class="cardImg"
                src="../assets/image/barstudent2.png"
                ></v-img>
            </v-card>
        </v-row>
        <v-row class="textdetail1">
          <p>หมายเลข 1 แสดงมายังหน้าประวัติของผู้เรียน</p>
          </v-row>
        <v-row class="textdetail2">
        <p>หมายเลข 2 แสดงมายังหน้า course ทั้งหมดของผู้เรียน</p>
        </v-row>
        <v-row class="textdetail2">
        <p>หมายเลข 3 ออกจากระบบ</p>
        </v-row>
        <v-row justify="center" class="textdetail">
          <p>เมื่อทำการเข้าสู่ระบบในส่วนของ Tutor จะขึ้นแถบดังรูปด้านล่าง</p>
        </v-row>
        <v-row  justify="center">
            <v-card class="imgContain2">
                <v-img
                class="cardImg"
                src="../assets/image/bartutor.png"
                ></v-img>
            </v-card>
        </v-row>
        <v-row class="textdetail1">
          <p>หมายเลข 1 แสดงหน้าแรกของ StyleLearn</p>
          </v-row>
        <v-row class="textdetail2">
        <p>หมายเลข 2 แสดงหน้า Basic use ของ StyleLearn</p>
        </v-row>
        <v-row class="textdetail2">
        <p>หมายเลข 3 แสดงหน้า About us ของ StyleLearn</p>
        </v-row>
        <v-row  class="textdetail2">
        <p>หมายเลข 4 หน้าต่าง pop up ที่แสดงดังรูปถัดไป</p>
        </v-row>
        <v-row  justify="center">
            <v-card class="imgContain3">
                <v-img
                class="cardImg"
                src="../assets/image/bartutor2.png"
                ></v-img>
            </v-card>
        </v-row>
        <v-row class="textdetail1">
          <p>หมายเลข 1 แสดงมายังหน้าประวัติของผู้สอน</p>
          </v-row>
        <v-row class="textdetail2">
        <p>หมายเลข 2 ออกจากระบบ</p>
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
  name: "basicuse"
};
</script>

<style scoped>
.cardContainer {
  color: black;
  border-radius: 20px;
  width: 900px;
  height: 2000px;
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
  font-weight: bold;
  font-size: 18px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 40px;
}
.textdetail1 {
  font-size: 14px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-top: 40px;
  margin-left: 200px;
}
.textdetail2 {
  font-size: 14px;
  color: #5c5c5c;
  font-family: "ABeeZee";
  margin-left: 200px;
}
.imgContain{
    height: 110px;
    width: 700px;
}
.imgContain1{
    height: 300px;
    width: 450px;
    margin-top: 40px;
}
.imgContain2{
    height: 130px;
    width: 700px;
}
.imgContain3{
    height: 270px;
    width: 520px;
    margin-top: 40px;
}
.cardImg {
  height: 100%;
  width: 100%;
  border-radius: 30px;
}
</style>
