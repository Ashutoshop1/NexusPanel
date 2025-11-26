<template>
  <div class="min-h-screen bg-gray-100">
    <div class="container mx-auto px-4 py-8">
      <h1 class="text-3xl font-bold mb-8">{{ t('monitoring.title') }}</h1>
      <div class="bg-white rounded-lg shadow p-6">
        <p class="text-gray-600">{{ t('common.info') }}: Monitoring implementation coming soon</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
</script>
