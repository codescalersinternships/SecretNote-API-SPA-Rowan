<template>
    <form>
        <div class="signup">
            <label for="username" class="labeling"> <b>Username</b></label>
            <input placeholder="Username" v-model="user.Username" required class="input">
            <label class="labeling"> <b>Password</b></label>
            <input placeholder="Password" required v-model="user.Password" class="input">
            <button class="signup_btn" @click="onSubmit">Sign Up</button>
        </div>
    </form>
</template>


<script lang="ts">
import { defineComponent, PropType, reactive } from "vue";
interface User {
  Username: string
  Password: string
}
export default defineComponent({
    name: "SignUp",
    data(){
        return {
            user: {
                Username : '',
                Password : '',
            } as User
        };
    },
    methods:{
        onSubmit(){
            console.log(this.user)
        }
    }
    // return {form, ons}
});

</script>

<style scoped>
/* input[type=text] */
.signup{
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
.signup_btn{
    background-color: green;
    width: 50%;
    margin-top: 20px;
}
</style>