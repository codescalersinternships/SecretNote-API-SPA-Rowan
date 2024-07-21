<template>
    <form class="form">
        <h2> Secret Note Creation </h2>
        
            <input class="title-area" type="text" v-model="note.Title" placeholder="Your Secret Note Title">
        
        
            <textarea class="text-area" v-model="note.Content" placeholder="Content of your secrets"> </textarea>
        
        <button class="note-btn" @click="onSubmit">SUBMIT NOTE</button>
    </form>
</template>


<script setup lang="ts">
import { ref, Ref } from "vue";
// import type {Props} from 
interface Note {
  Title: string
  Content: string
}
const note : Ref<Note> = ref({Title:"", Content:""})

function onSubmit(){
    fetch("http://localhost:8080/note", {
        method: "POST",
        credentials: "include",
        body: JSON.stringify({
            Title: note.value.Title,
            Content: note.value.Content
        }),
    }).then(res => res.json()).then( res => console.log(res))

}
</script>

<style >
.form{
    width: 400px;
    display: flex;
    flex-direction: column;
    
    align-items: center;
    justify-content: center;
    border: 20px;
    border-radius: 30px;
    margin: 20px, 20px, 20px, 20px;
    /* font-size: 18px; */
    /* margin-top: 0; */
}
.text-area{
    /* width: 200px; */
    height: 200px;
    font-size: 18px;
}
.note-btn{
    background-color: green;
    margin-top: 20px;
} 
.title-area{
    margin-bottom: 20px;
    height: 35px;
    font-size: 18px;
}
</style>