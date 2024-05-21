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
            path: '/sign',
            name: 'sign_login',
            component: () => import('../views/SignPage/login.vue')
        },
        {
            path: '/sign/register',
            name: 'sign_register',
            component: () => import('../views/SignPage/register.vue')
        },
        {
            path: '/sign/forgot',
            name: 'sign_forgot',
            component: () => import('../views/SignPage/forgot.vue')
        },
        {
            path: '/admin',
            name: 'admin',
            component: () => import('../views/AdminPage/index.vue')
        }
    ]
})

export default router
