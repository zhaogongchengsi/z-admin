import { App } from 'vue'
import { createRouter, createWebHashHistory, RouteLocationNormalized, RouteRecordRaw } from 'vue-router'
import { routerStore } from '@/pinia/index'

export async function setUpRouter (app:App) {

    const r = routerStore()
    await r.getRouterInfo()


    console.log(r.getRouter)

    const homeRouter =  {
        path: '/',
        name: 'home',
        component: () => import('@/views/home/index.vue'),
        children: r.getRouter
    }

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