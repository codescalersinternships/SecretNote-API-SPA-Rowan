<template>
    <div class="note">
        <h2> Secret Note Creation </h2>
        <div>
            <input type="text" v-model="note.Title" placeholder="Your Secret Note Title"><br><br>
        </div>
        <div>
            <textarea class="text-area" v-model="note.Content" placeholder="Content of your secrets"> </textarea>
        </div>
        <button class="note-btn" @click="onSubmit">SUBMIT NOTE</button>
    </div>
</template>


<script setup lang="ts">
import { ref, defineComponent, PropType, reactive, provide, Ref } from "vue";
// import type {Props} from 
interface Note {
  Title: string
  Content: string
}
const note : Ref<Note> = ref({Title:"", Content:""})

function onSubmit(){
    fetch("http://localhost:8080/note", {
        method: "POST",
        body: JSON.stringify({
            Title: note.value.Title,
            Content: note.value.Content
        }),
    }).then(res => res.json()).then( res => console.log(res))

}
</script>

<style >
.note{
    width: 400px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    border: 20px;
    border-radius: 30px;
    margin: 20px, 20px, 20px, 20px;
}
.text-area{
    /* width: 200px; */
    height: 200px;
}
.note-btn{
    background-color: green;
    width: 50%;
    margin-top: 20px;
} 
</style>