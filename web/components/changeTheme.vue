<!-- pages/index.vue -->
<template>
  <div>
    <h1>主题切换示例</h1>
    <a-button class="btn-primary" @click="toggleTheme">
      {{ isDark ? '切换到浅色' : '切换到深色' }}
    </a-button>
  </div>
</template>

<script setup>
const { isDark, toggleTheme } = useTheme()
</script>