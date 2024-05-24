// router.js

import { createRouter, createWebHistory } from 'vue-router';
import homeRoutes from './routes/homeRoutes';
import signRoutes from './routes/signRoutes';
import adminRoutes from './routes/adminRoutes';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        ...homeRoutes,
        ...signRoutes,
        ...adminRoutes,
        {
            path: '/:pathMatch(.*)*',
            name: '404',
            component: () => import('../views/404.vue'),

        },
    ]
});

export default router;
