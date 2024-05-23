// adminRoutes.js

export default [
    {
        path: '/admin',
        name: 'admin',
        component: () => import('../../views/AdminPage/index.vue'),
        children: [
            {
                // 仪表盘
                path: '',
                name: 'dashboard',
                component: () => import('../../views/AdminPage/Dashboard/index.vue')
            },
            {
                // 所有用户
                path: 'userMm',
                name: 'userMm',
                component: () => import('../../views/AdminPage/UserMm/UserMm/index.vue'),
            },
            {
                // 提现管理
                path: 'Withdraw',
                name: 'Withdraw',
                component: () => import('../../views/AdminPage/UserMm/WithdrawMm/index.vue'),
            }
        ]
    }
];
