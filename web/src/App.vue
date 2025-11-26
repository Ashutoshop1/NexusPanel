<template>
  <div id="app" class="min-h-screen bg-gray-50 dark:bg-gray-900">
    <router-view />
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useI18n } from 'vue-i18n'

const { locale } = useI18n()

onMounted(() => {
  // Load user language preference
  const savedLocale = localStorage.getItem('locale')
  if (savedLocale) {
    locale.value = savedLocale
  }
})
</script>

<style>
@import './assets/main.css';
</style>
