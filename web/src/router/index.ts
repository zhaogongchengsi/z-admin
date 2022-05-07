import { App } from 'vue'
import { createRouter, createWebHashHistory, RouteLocationNormalized, RouteRecordRaw } from 'vue-router'
import { routerStore } from '@/pinia/index'

const staticRouter = [
    
]


export async function setUpRouter (app:App) {

    const r = routerStore()
    await r.getRouterInfo()

    const homeRouter =  {
        path: '/',
        name: 'home',
        component: () => import('@/views/home/index.vue'),
        children: [
            {
                path: '/menu',
                name: 'menu',
                component: () => import('@/views/menu/index.vue'),
            }
        ].concat(r.getRouter)
    }

    console.log(homeRouter)

    const router = createRouter({
        history: createWebHashHistory(),
        routes: [
            {
                path: '/login',
                name: 'login',
                component: () => import("@/views/login/index.vue"),
            },
            homeRouter,
        ],
    })

    router.beforeEach(beforeEach)

    app.use(router)
}


function beforeEach (to: RouteLocationNormalized, from: RouteLocationNormalized) {

    const token = localStorage.getItem('z_token');

    if (to.name != "login" && !token) {
        return "/login"
    } else {
        return true
    }

}