<template>
  <div class="min-h-screen bg-gray-100">
    <div class="container mx-auto px-4 py-8">
      <h1 class="text-3xl font-bold mb-8">{{ t('menu.dashboard') }}</h1>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div class="bg-white rounded-lg shadow p-6">
          <h2 class="text-xl font-semibold mb-2">{{ t('server.online') }}</h2>
          <p class="text-3xl font-bold text-green-500">0</p>
        </div>
        <div class="bg-white rounded-lg shadow p-6">
          <h2 class="text-xl font-semibold mb-2">{{ t('server.offline') }}</h2>
          <p class="text-3xl font-bold text-red-500">0</p>
        </div>
        <div class="bg-white rounded-lg shadow p-6">
          <h2 class="text-xl font-semibold mb-2">{{ t('monitoring.alerts') }}</h2>
          <p class="text-3xl font-bold text-yellow-500">0</p>
        </div>
        <div class="bg-white rounded-lg shadow p-6">
          <h2 class="text-xl font-semibold mb-2">{{ t('menu.tasks') }}</h2>
          <p class="text-3xl font-bold text-blue-500">0</p>
        </div>
      </div>
      <p class="mt-8 text-gray-600">{{ t('common.info') }}: Dashboard implementation coming soon</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
</script>
