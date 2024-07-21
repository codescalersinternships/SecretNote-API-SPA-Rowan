<template>
    <div class="notes">
        <h2 @click="onSubmit">Your Not So Secret Secret-Notes</h2>
        <ul>
            <li>
                <h3> note title</h3>
                <p>
                    Note content
                </p>
            </li>
            <li>
                <h3> note title</h3>
                <p>
                    Note content
                </p>
            </li>
            <li>
                <h3> note title</h3>
                <p>
                    Note content
                </p>
            </li>
        </ul>
    </div>
</template>


<script setup lang="ts">
import { ref, Ref } from "vue";
// import type {Props} from 
interface Note {
  Title: string
  Content: string
}
// const note : Ref<Note[]> = ref()

function onSubmit(){
    fetch("http://localhost:8080/notes", {
        method: "GET",
        credentials: "include",
        // body: JSON.stringify({
        //     Title: note.value.Title,
        //     Content: note.value.Content
        // }),
    }).then(res => res.json()).then( res => console.log(res))
    // const data = res
}
</script>

<style >
.notes{
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