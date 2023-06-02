<script setup>
import Registration from "./settings/Registration.vue"
import Handlers from "./settings/Handlers.vue"

import { useAdminConfigStore } from '@/stores/admin_config'

const store = useAdminConfigStore()
store.fetchConfig()
</script>
    
<template>
  <n-space vertical v-if="store.config">
    <Registration />
    <Handlers />
  </n-space>
  
</template>