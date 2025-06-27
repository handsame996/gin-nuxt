<template>
    <h1>首页</h1>
    1111
    <a-space wrap>
    <a-button type="primary">Primary Button</a-button>
    <a-button>Default Button</a-button>
    <a-button type="dashed">Dashed Button</a-button>
    <a-button type="text">Text Button</a-button>
    <a-button type="link">Link Button</a-button>
  </a-space>
</template>

<script setup lang="ts">
import { useNuxtApp } from '#app'

const { vueApp } = useNuxtApp()
console.log('已注册组件:', Object.keys(vueApp._context.components))

</script>