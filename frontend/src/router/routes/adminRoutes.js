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
                path: 'categoryMm',
                name: 'categoryMm',
                component: () => import('../../views/AdminPage/ProductMm/CategoryMm/index.vue'),
            },
            {
                // 库存管理
                path: 'inventoryMm',
                name: 'inventoryMm',
                component: () => import('../../views/AdminPage/ProductMm/InventoryMm/index.vue'),
            },

            {
                // 订单管理
                path: 'orderMm',
                name: 'orderMm',
                component: () => import('../../views/AdminPage/OrderMm/OrderMm/index.vue'),
            },
            {
                // 售后管理
                path: 'aftercareMm',
                name: 'aftercareMm',
                component: () => import('../../views/AdminPage/OrderMm/AftercareMm/index.vue'),
            },

            {
                // 广告管理
                path: 'adMm',
                name: 'adMm',
                component: () => import('../../views/AdminPage/AdMm/index.vue'),
            },

            {
                // 会员卡
                path: 'membershipCard',
                name: 'membershipCard',
                component: () => import('../../views/AdminPage/DiscountsMm/MembershipCard/index.vue'),
            },
            {
                // 折扣码
                path: 'discountCode',
                name: 'discountCode',
                component: () => import('../../views/AdminPage/DiscountsMm/DiscountCode/index.vue'),
            },

            {
                // 商户工单
                path: 'merchantTickets',
                name: 'merchantTickets',
                component: () => import('../../views/AdminPage/TicketingMm/MerchantTickets/index.vue'),
            },
            {
                // 用户投诉
                path: 'userComplaints',
                name: 'userComplaints',
                component: () => import('../../views/AdminPage/TicketingMm/UserComplaints/index.vue'),
            },




            {
                // 设置页面
                path: 'allSetting',
                name: 'allSetting',
                component: () => import('../../views/AdminPage/AllSetting/index.vue'),
            }

        ]
    }
];
