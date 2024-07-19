
<script setup lang="ts">
// import { defineComponent } from 'vue'
import SignUp from './components/SignUp.vue'
import Login from './components/Login.vue'
import Note from './components/Note.vue'
import Notes from './components/Notes.vue';
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
    <h1> Choose Your Fighter</h1>
    <div class="home">
      <ul>
        <li>
          <!-- <router-link to="/signup"> Sign Up</router-link> -->
          <a href="/signup"> Sign UP</a>
        </li>
        <li>
          <router-link to="/login"> Log in</router-link>
        </li>
        <li>
          <router-link to="/note"> Create Your Secret Note</router-link>
        </li>
        <li>
          <router-link to="/notes"> Check out Notes</router-link>
        </li>
      </ul>
    </div>
    
  
  
    <!-- <router-view></router-view> -->
   <!-- <Note/> -->
  <!-- <Notes/> -->
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
ul{
  list-style-type: none;
}
.home{
  font-size: 20px;
}
</style>
