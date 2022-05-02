import { createApp } from 'vue'
import App from './App.vue'
import { usePinia } from '@/pinia'
import { setUpRouter } from '@/router'

import "./assets/style/reset.css"


async function Bootstrap () {
   const app = createApp(App)

   await setUpRouter(app)
   
   usePinia(app)
   app.mount('#app')
}

Bootstrap()