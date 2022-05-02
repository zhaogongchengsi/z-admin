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

  // 这个代理配置不生效 真奇怪
  const proxy = {
    '/api': {
      target: 'http://localhost:9090',
      changeOrigin: true,
      rewrite: (path: string) => path.replace(/^\/api/, '')
    },
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
    },
    server: {
      port: 3050,
      host: '0.0.0.0',
      proxy
    },
    proxy
  }
})
