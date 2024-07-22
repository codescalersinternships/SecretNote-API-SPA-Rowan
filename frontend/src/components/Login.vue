<template>
    <form>
        <div class="login">
            <label for="username" class="labeling"> <b>Username</b></label>
            <input placeholder="Username" v-model="user.Username" required class="input">
            <label class="labeling"> <b>Password</b></label>
            <input placeholder="Password" required v-model="user.Password" class="input">
            <button class="login-btn" @click="onSubmit">Log In</button>
        </div>
    </form>
</template>


<script setup lang="ts">
import { ref, Ref } from "vue";
import router from "../main";
import axios from "axios"
// import type {Props} from 
interface User {
  Username: string
  Password: string
}
const user : Ref<User> = ref({Username:"", Password:""})

async function onSubmit(){
    // try{
    console.log(user.value)
    console.log(user)
        const response = await axios.post("http://localhost:8080/login", user.value)
        const data = await response.statusText
        console.log(data);
        console.log(response);
        console.log("hellooooooo")
    // } catch(error){
    //     console.log(error);
    // }
}
// async function onSubmit(){
//     console.log("hello")
//     try{
//         const res = await fetch("http://localhost:8080/login", {
//             method: "POST",
//             headers: {
//                 'Content-Type': 'application/json',
//                 // 'Cache-Control': 'no-cache'
//             },
//             credentials: "include",
//             body: JSON.stringify({
//                 Username: user.value.Username,
//                 Password: user.value.Password,
//             }),
//         })
//         if (!res.ok) {
//             console.log(res.status);
//         }
//         console.log("hello before text")
//         console.log(res.text)
//         try {
//             const json = await res.json();
//             console.log(json);
//             router.replace('/note');
//         } catch(error) {
//             console.log(error);
//         }
        
//         console.log("mmmmmmm");
        
//     } catch (error) {
//         console.error(error);
//         console.log(error);
//     }
//     console.log("aaaaaaaaaaaaaaa");
//     router.replace('/note');
// }
</script>

<style scoped>
.login{
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    /* padding: 100px; */
    width: 150%;
    height: 150%;
    /* border-color: cornsilk; */
    border: 20px;
    border-radius: 30px;
    margin: 20px, 20px, 20px, 20px;
}
.input{
    padding: 12px 20px;
    margin: 20px, 20px;
    width: 100%;
    /* margin: 8px 0; */
    box-sizing: border-box;
}
.labeling{
    font-size: 30;
    margin-top: 20px;
    margin-bottom: 20px;
    padding: 10;
}
.login-btn{
    background-color: green;
    width: 50%;
    margin-top: 20px;
}
</style>