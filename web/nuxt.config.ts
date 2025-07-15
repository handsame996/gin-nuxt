// https://nuxt.com/docs/api/configuration/nuxt-config
import path from 'path'
import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
  ssr: false,
  alias:{
    '@': process.cwd(),
    '~': process.cwd(),
  },
  typescript: {
    typeCheck: true,
    tsConfig: {
      include: ['types/**/*.d.ts', 'types/**/*.ts']
    }
  },
  imports:{
    dirs:[
      'composables/api'
    ]
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
    '@ant-design-vue/nuxt','@pinia/nuxt','@nuxtjs/i18n'
  ],
  antd:{},
  plugins:[],
  css: [
    '@/assets/less/main.less',
    "@/assets/css/main.css"
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
        // 关键：使用 hack 语法确保自定义变量优先加载
        modifyVars: {},
        javascriptEnabled: true
      }
      }
    },
    plugins: [
      tailwindcss(),
      ],
    server:{
      proxy:{
        '/api':{
        target: 'http://127.0.0.1:8888',
        changeOrigin: true,
        rewrite: path => path.replace(/^\/api/, '')
        }
      }
    }
  },
  i18n:{
    // 配置默认语言
    defaultLocale: 'zh',
    
    // 支持的语言列表
    locales: [
      {
        code: 'en',
        name: 'English',
        file: 'en.ts' // 对应语言文件
      },
      {
        code: 'zh',
        name: '中文',
        file: 'zh.ts'
      }
    ],
    
    // 语言文件存放路径
    langDir: './',
    
    // 启用路由国际化
    strategy: 'prefix_except_default',
    
    // 允许在URL中自动检测语言
    detectBrowserLanguage: {
      useCookie: true,
      cookieKey: 'nuxt-i18n',
      redirectOn: 'root'
    }
  },
  build: {
    transpile: ['ant-design-vue'],
  }
})