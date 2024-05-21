import {createRouter, createWebHistory} from 'vue-router'

import HomeView from '../views/HomePage.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView
        },
        {
            path:'/sign',
            name:'sign',
            component:()=>import('../views/SignPage/login.vue')

        }
    ]
})

export default router
