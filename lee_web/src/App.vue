<template>
  <div id="app">

    <Home/>
    <!--    <img src="./assets/logo.png">-->
    <!--    <div>-->
    <!--      <p>-->
    <!--        If Element is successfully added to this project, you'll see an-->
    <!--        <code v-text="'<el-button>'"></code>-->
    <!--        below-->
    <!--      </p>-->
    <!--      <el-button>el-button</el-button>-->
    <!--    </div>-->
    <!--    <HelloWorld msg="Welcome to Your Vue.js App"/>-->
  </div>
</template>

<script>
import Home from './components/Home'

export default {
  name: 'app',
  components: {
    Home
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  /*color: #2c3e50;*/
  /*margin-top: 60px;*/
  height: 100vh;
  width: 100%;
}

body, html {
  margin: 0;
  height: 100%;
}

.gao100 {
  height: 100%;
  background: #CCC;
}
</style>
