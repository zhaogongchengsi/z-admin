import { createApp } from 'vue'
import App from './App.vue'
import { usePinia } from '@/pinia'
import { setUpRouter } from '@/router'

import 'element-plus/dist/index.css'
import "./assets/style/reset.css"


import { ElCollapseTransition } from 'element-plus'

async function Bootstrap () {
   const app = createApp(App)

   await setUpRouter(app)

   usePinia(app)
   app.mount('#app')
}

Bootstrap()