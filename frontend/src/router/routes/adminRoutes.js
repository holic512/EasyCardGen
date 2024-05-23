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
            },
            {
                // 商店管理
                path: 'storeMm',
                name: 'storeMm',
                component: () => import('../../views/AdminPage/StoreMm/StoreMm/index.vue'),
            },
            {
                // 权限管理
                path: 'accessMm',
                name: 'accessMm',
                component: () => import('../../views/AdminPage/StoreMm/AccessMm/index.vue'),
            },
            {
                // 商品管理
                path: 'productMm',
                name: 'productMm',
                component: () => import('../../views/AdminPage/ProductMm/ProductMm/index.vue'),
            },
            {
                // 分类管理
                path:'categoryMm',
                name: 'categoryMm',
                component: () => import('../../views/AdminPage/ProductMm/CategoryMm/index.vue'),
            },
            {
                // 库存管理
                path:'inventoryMm',
                name: 'inventoryMm',
                component: () => import('../../views/AdminPage/ProductMm/InventoryMm/index.vue'),
            },
        ]
    }
];
