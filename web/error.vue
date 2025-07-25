<template>
  <div class="min-h-screen flex flex-col items-center justify-center p-6 text-center">
    <h1 class="text-6xl font-bold mb-4 text-gray-800 dark:text-gray-200">404</h1>
    <p class="text-xl mb-8 text-gray-600 dark:text-gray-400">页面未找到</p>
    <p class="mb-8 max-w-md">
      抱歉，您访问的页面不存在或已被移除。请检查您输入的网址是否正确。
    </p>
    <NuxtLink 
      to="/admin" 
      class="px-6 py-3 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"
    >
      返回首页
    </NuxtLink>
  </div>
</template>

<script setup>
// 可以在这里添加额外的逻辑，比如统计404错误
useHead({
  title: '404 - 页面未找到',
  meta: [
    { name: 'description', content: '您访问的页面不存在' }
  ]
})
</script>

<style scoped>
/* 可以添加自定义样式 */
</style>
