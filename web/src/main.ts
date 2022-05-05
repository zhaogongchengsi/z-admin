import { createApp } from 'vue'
import App from './App.vue'
import { usePinia } from '@/pinia'
import { setUpRouter } from '@/router'
import useIcon from '@/plugin/icon'

import 'element-plus/dist/index.css'
import "./assets/style/reset.css"


async function Bootstrap () {

   const app = createApp(App)

   usePinia(app)

   useIcon(app)

   await setUpRouter(app)

   app.mount('#app')
}

Bootstrap()