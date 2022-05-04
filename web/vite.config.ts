import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import { resolve } from 'path'
import { config } from 'dotenv'
import { env } from 'process'

function resolvePath(path: string) :string {
  return resolve(__dirname, path)
}

export default defineConfig(() => {
  // 先读取环境变量
  config({
    path: resolvePath('.env.development')
  })

  const port = parseInt(env["VITE_SERVER_PORT"])
  const host = env["VITE_SERVER_HOST"]

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
      port,
      host,
      proxy
    }
  }
})
