<template>
  <div id="lee-menu">
    <el-menu
        mode="horizontal"
        background-color="#"
        text-color="#fff"
        active-text-color="#fff"
    >
      <el-menu-item index="1">Home</el-menu-item>
      <el-menu-item index="2">Personal Center</el-menu-item>
<!--      <el-sub-menu index="2">-->
<!--        <template #title>Workspace</template>-->
<!--        <el-menu-item index="2-1">item one</el-menu-item>-->
<!--        <el-menu-item index="2-2">item two</el-menu-item>-->
<!--        <el-menu-item index="2-3">item three</el-menu-item>-->
<!--        <el-sub-menu index="2-4">-->
<!--          <template #title>item four</template>-->
<!--          <el-menu-item index="2-4-1">item one</el-menu-item>-->
<!--          <el-menu-item index="2-4-2">item two</el-menu-item>-->
<!--          <el-menu-item index="2-4-3">item three</el-menu-item>-->
<!--        </el-sub-menu>-->
<!--      </el-sub-menu>-->
<!--&lt;!&ndash;      <el-menu-item index="3" disabled>Info</el-menu-item>&ndash;&gt;-->
<!--      <el-menu-item index="4">Orders</el-menu-item>-->
    </el-menu>

  </div>

</template>

<script>
export default {
name: "Navigation",

}

</script>

<style scoped>
#lee-menu{
  width: 100%;
}


</style>