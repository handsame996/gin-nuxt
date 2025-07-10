// types/index.d.ts
/// <reference types="@nuxt/types" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}

// 扩展 Nuxt 的 useFetch 返回类型
declare module '#app' {
  interface NuxtApp {
    $apiFetch<T>(params: ApiParams): Promise<resType<T>>;
  }
}