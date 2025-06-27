// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  runtimeConfig: {
    public: {
      apiBase: process.env.NODE_ENV === 'production'
          ? process.env.PROD_API_BASE_URL
          : process.env.DEV_API_BASE_URL
    }
  },
  compatibilityDate: '2025-05-15',
  devtools: { enabled: true },
  modules: ['@nuxt/test-utils','@ant-design-vue/nuxt'],
  plugins:[],
  css: [],
  build: {
    transpile: ['ant-design-vue']
  }
})