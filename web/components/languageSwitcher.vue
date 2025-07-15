<!-- components/LanguageSwitcher.vue -->
<template>
  <div class="language-switcher">
     <a-select
      ref="select"
      v-model:value="currentLocale"
      @change="switchLocale"
    >
      <a-select-option v-for="locale in availableLocales" :key="locale.code" :value="locale.code">{{ locale.name }}</a-select-option>
    </a-select>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import { useNuxtApp } from '#app'
import type {Locale} from "~/types/language"
import {supportedLocales} from "~/types/language"
import type { DefaultOptionType,SelectValue } from 'ant-design-vue/es/select'

const { locale, locales } = useI18n()
const nuxtApp = useNuxtApp()

// 当前语言
const currentLocale = computed({
  get() {
    return locale.value
  },
  set(value) {
    locale.value = value
  }
})

// 可用语言列表
const availableLocales = computed(() => {
  return locales.value.filter(locale =>  supportedLocales.map(item => item == locale.code) )
})

// 切换语言
const switchLocale = async (value: SelectValue, option: DefaultOptionType) => {
  const newLocale = value
  
  try {
    // 更新路由前缀（如果启用了路由前缀策略）
    await nuxtApp.$i18n.setLocale(newLocale as Locale)
    
    // 保存语言偏好到 cookie
    useCookie('i18n_redirected').value = newLocale as string
    
    console.log(`已切换到 ${newLocale} 语言`)
  } catch (error) {
    console.error('语言切换失败:', error)
  }
}
</script>