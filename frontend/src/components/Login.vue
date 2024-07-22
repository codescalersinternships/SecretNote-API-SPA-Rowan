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
    try{
        const res = await fetch("http://localhost:8080/login", {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
                'Cache-Control': 'no-cache'
            },
            credentials: "include",
            body: JSON.stringify({
                Username: user.value.Username,
                Password: user.value.Password
            }),
        })
        if (!res.ok) {
            console.log(res.status);
        }
        console.log("hello before text")
        console.log(res.text)
        try {
            const json = await res.json()
            console.log(json);
            router.replace('/note')
        } catch(error) {
            console.log(error)
        }
        
        console.log("mmmmmmm")
        
    } catch (error) {
        console.error(error)
        console.log(error)
    }

    // const res = fetch("http://localhost:8080/login", {
    //     method: "POST",
    //     headers: {
    //         'Content-Type': 'application/json',
    //     },
    //     credentials: "include",
    //     body: JSON.stringify({
    //         Username: user.value.Username,
    //         Password: user.value.Password
    //     }),
    // }).then((response) => {
    //     response.json()
    // }).then((data) =>{
    //     console.log(data)
    // }).then(()=> {
    //     console.log(res)
    // })
    // .then((response) => {
    //       response.json().then((data) => {
    //         console.log(data)
    //       });
    //     })
    //     .catch((err) => {
    //       console.error(err);
    //     });
    // const res = await fetch("http://localhost:8080/login", {
    //     method: "POST",
    //     body: JSON.stringify({
    //         Username: user.value.Username,
    //         Password: user.value.Password
    //     }),
    // })
    // const data = await res
    // if ((await response).ok){
        
    // }
    // console.log(res)
    // console.log(data)
    console.log("aaaaaaaaaaaaaaa")
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