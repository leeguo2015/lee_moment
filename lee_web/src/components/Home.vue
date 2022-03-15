<template>
  <div id="Home" class="common-layout">
    <el-container>
      <el-header id="header" ><Navigation/></el-header>
      <el-main id="desk"></el-main>
    </el-container>
  </div>
</template>

<script>
import Navigation from "@/components/home/Navigation";

export default {
  name: "Home",
  components: {
    Navigation,
  }
}
</script>

<style scoped>
#Home{
  height: 100%;
}
#header{
  background-color: rgba(54, 53, 53, 0.5);
  width: 100%;
  padding: 0;
  margin: 0;

}
#desk{
  /*background-color: darkseagreen;*/
  height: 100%;
}
</style>