import { App } from "vue";
import * as ElementPlusIconsVue from '@element-plus/icons-vue'


export default function useIcon(app: App) {

    for (const [key, value] of Object.entries(ElementPlusIconsVue)) {
        app.component(key, value)
    }
    
}