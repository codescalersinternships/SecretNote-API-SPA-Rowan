import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import './style.css'
import App from './App.vue'
import SignUp from './components/SignUp.vue'
import Login from './components/Login.vue'
import Note from './components/Note.vue'
import Notes from './components/Notes.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: '/signup', name: "SignUp", component: SignUp},
        {path: '/login', name: "Login", component: Login},
        {path: '/note', name: "note", component: Note},
        {path: '/notes', name: "notes", component: Notes},
    ]
})
export default router
createApp(App).use(router).mount('#app')
