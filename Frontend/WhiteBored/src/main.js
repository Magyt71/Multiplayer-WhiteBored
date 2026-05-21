import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { createRouter, createWebHistory } from 'vue-router';

import 'virtual:uno.css' // لو ما اشتغلت احذف هذا السطر
import App from './App.vue'
import Login from './Login.vue'
import Bored from './bored.vue' // تم تصحيح الاسم هنا

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: Login },
    { path: '/board/:id', component: Bored }
  ]
})

createApp(App).use(createPinia()).use(router).mount('#app')