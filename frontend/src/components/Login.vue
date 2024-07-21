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
// import type {Props} from 
interface User {
  Username: string
  Password: string
}
const user : Ref<User> = ref({Username:"", Password:""})

async function onSubmit(){
    console.log("hello")
    const response = fetch("http://localhost:8080/login", {
        method: "POST",
        // headers: {
        //     'Content-Type': 'application/json',
        // },
        credentials: "include",
        body: JSON.stringify({
            Username: user.value.Username,
            Password: user.value.Password
        }),
    });
    // const res = await fetch("http://localhost:8080/login", {
    //     method: "POST",
    //     body: JSON.stringify({
    //         Username: user.value.Username,
    //         Password: user.value.Password
    //     }),
    // })
    const data = (await response).json
    if ((await response).ok){
        
    }
    console.log(response)
    console.log(data)
    router.replace('/note')
}
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