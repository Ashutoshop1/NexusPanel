<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="max-w-md w-full bg-white rounded-lg shadow-lg p-8">
      <h1 class="text-3xl font-bold text-center mb-8">{{ t('auth.login') }}</h1>
      <p class="text-center text-gray-600 mb-8">NexusPanel</p>
      
      <form @submit.prevent="handleLogin">
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2">
            {{ t('auth.username') }}
          </label>
          <input 
            v-model="username"
            type="text" 
            class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
            required
          >
        </div>
        
        <div class="mb-6">
          <label class="block text-gray-700 text-sm font-bold mb-2">
            {{ t('auth.password') }}
          </label>
          <input 
            v-model="password"
            type="password" 
            class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
            required
          >
        </div>
        
        <button 
          type="submit"
          class="w-full bg-blue-500 text-white py-2 px-4 rounded-lg hover:bg-blue-600 transition"
        >
          {{ t('auth.login') }}
        </button>
      </form>
      
      <p class="mt-4 text-center text-sm text-gray-600">
        {{ t('common.info') }}: This is a placeholder login view
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'

const router = useRouter()
const { t } = useI18n()

const username = ref('')
const password = ref('')

const handleLogin = () => {
  // TODO: Implement actual login logic
  // SECURITY NOTE: localStorage is used here for development convenience.
  // For production, use httpOnly cookies or secure session storage to prevent XSS attacks.
  localStorage.setItem('token', 'dummy-token')
  router.push('/dashboard')
}
</script>
