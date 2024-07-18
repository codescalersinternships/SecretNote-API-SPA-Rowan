
<script setup lang="ts">
// import { defineComponent } from 'vue'
import SignUp from './components/SignUp.vue'
import Login from './components/Login.vue'
// export default defineComponent({
//   name : "app",
//   components: {SignUp},
// });
// const routes = {
//   '/signup': SignUp,
//   '/login': Login,
// }

</script>

<template>
   <!-- <Login/> -->
   <!-- <SignUp/> -->
  <router-link to="/signup"> Sign Up</router-link>
  <router-link to="/login"> Log in</router-link>
   
</template>



<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
