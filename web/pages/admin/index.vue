<!-- pages/index.vue -->
<template>
  <div>
    <a-space wrap>
    <a-button type="primary">Primary Button</a-button>
    <a-button>Default Button</a-button>
    <a-button type="dashed">Dashed Button</a-button>
    <a-button type="text">Text Button</a-button>
    <a-button type="link">Link Button</a-button>
    </a-space>
    <h1>主题切换示例</h1>
    <a-button class="btn-primary" @click="toggleTheme">
      {{ isDark ? '切换到浅色' : '切换到深色' }}
    </a-button>
  </div>
</template>

<script setup>
const { isDark, toggleTheme } = useTheme()
</script>