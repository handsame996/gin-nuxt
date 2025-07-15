<!-- pages/index.vue -->
<template>
  <div>
    <a-button type="link" @click="toggleTheme">
      {{ isDark ? '切换到浅色' : '切换到深色' }}
    </a-button>
  </div>
</template>

<script setup>
import { watch } from 'vue';
const { isDark,themeVars, toggleTheme } = useAntTheme()

defineExpose({
  themeVars
})
</script>