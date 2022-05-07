import { createPinia } from "pinia"

import type { App } from "vue"


export function usePinia (app: App) {
  const pinia = createPinia() 
    app.use(pinia)
}

export { userStore } from "./user"
export { mainStore } from "./main"
export { routerStore } from "./router"