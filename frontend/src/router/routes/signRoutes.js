// signRoutes.js

export default [
    {
        path: '/sign',
        name: 'sign_login',
        component: () => import('../../views/SignPage/login.vue')
    },
    {
        path: '/sign/register',
        name: 'sign_register',
        component: () => import('../../views/SignPage/register.vue')
    },
    {
        path: '/sign/forgot',
        name: 'sign_forgot',
        component: () => import('../../views/SignPage/forgot.vue')
    }
];
