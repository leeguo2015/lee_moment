<template>
<div>
    <h1>{{ msg }}</h1>
  <button @click="count++">count is: {{ data.counter }}</button>
  <p>Edit <code>components/HelloWorld.vue</code> to test hot module replacement.</p>
  <p> {{data.counter}}</p>
  <p> {{count}}</p>
</div>
</template>
 
<script>
import { reactive} from 'vue'


export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  setup() {
    const data  = reactive({
      counter:1,

    })
    return {data}
  },
  // data() {
  //   return {
  //     count: 0
  //   }
  // }
}
</script>

50864987
