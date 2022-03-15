<template>
  <div id="app">
    <Home :style="conTop" />
  </div>
</template>

<script>
import Home from './components/Home'

export default {
  name: 'app',
  data() {
    return {
      conTop: {
        backgroundImage:'url(' + require('./assets/HomeBack.jpg') + ')',
        backgroundRepeat:'no-repeat',
        backgroundSize: `cover`
      }

    }
  },
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
  height: 100%;
  width: 100%;
}

body, html {
  margin: 0;
  height: 100%;
  width: 100%;
}

</style>
