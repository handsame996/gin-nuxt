// https://nuxt.com/docs/api/configuration/nuxt-config
import path from 'path'

export default defineNuxtConfig({
  ssr: false,
  alias:{
    '@': path.resolve(__dirname, './'),
  },
  runtimeConfig: {
    public: {
      apiBase: process.env.NODE_ENV === 'production'
          ? process.env.PROD_API_BASE_URL
          : process.env.DEV_API_BASE_URL
    }
  },
  compatibilityDate: '2025-05-15',
  devtools: { enabled: true },
   modules: [
    '@ant-design-vue/nuxt','@pinia/nuxt'
  ],
  antd:{},
  plugins:[],
  css: [
    '@/assets/less/main.less'
  ],
  router: {
   options:{
     hashMode: true
   }
  },
  vite:{
    css:{
       preprocessorOptions: {
        less: {
          // 全局注入 Less 变量（无需每个文件手动导入）
          additionalData: `@import "@/assets/less/variables.less"; @import "@/assets/less/antd-variables.less";`,
          // 允许 Less 中使用 JavaScript（如需动态计算）
          javascriptEnabled: true,
          modifyVars: {
            hack: `true; @import "@/assets/less/antd-variables.less";`,
            javascriptEnabled: true
          }
        }
      }
    },
    server:{
      proxy:{
        '/api':{
        target: 'https://api.example.com',
        changeOrigin: true,
        rewrite: path => path.replace(/^\/api/, '')
        }
      }
    }
  },
  build: {
    transpile: ['ant-design-vue'],
   
  }
})