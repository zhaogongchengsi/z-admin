import { App } from 'vue'
import { createRouter, createWebHashHistory, RouteLocationNormalized, RouteRecordRaw } from 'vue-router'


const staticRouter = [
    {
        path: '/',
        name: 'home',
        component: () => import('@/views/home/index.vue'),
    },
    {
        path: '/login',
        name: 'login',
        component: () => import("@/views/login/index.vue"),
    }
]


export async function setUpRouter (app:App) {

    const routes: ConcatArray<{ path: string; name: string; component: () => Promise<typeof import("*.vue")> }> = []

    const router = createRouter({
        history: createWebHashHistory(),
        routes: staticRouter.concat(routes), // `routes: routes` 的缩写
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