import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import { resolve } from 'path'

function resolvePath(path: string) {
  return resolve(__dirname, path)
}

// https://vitejs.dev/config/
export default defineConfig(() => {

  const alias = {
    '@': resolvePath('src'),
  }

  return {
    plugins: [
      vue(),
      AutoImport({
        resolvers: [ElementPlusResolver()],
      }),
      Components({
        resolvers: [ElementPlusResolver()]
      }),
    ],
    resolve: {
      alias,
    }
  }
})
