<template>
  <div flex items-center>
    
  </div>
  <div style="align-items: center;display: flex;margin-left: auto;">
    <n-button 
      icon-placement="left" 
      @click="openDoc"
      tertiary
      >
      <template #icon>
        <n-icon>
          <code-slash />
        </n-icon>
      </template>
      API Doc
    </n-button>
    <UserAvatar />
  </div>
</template>

<script setup>
import UserAvatar from './components/UserAvatar.vue'
import {CodeSlashOutline as CodeSlash } from "@vicons/ionicons5"

function openDoc(){
  window.open('/api/docs/index.html', '_blank')
}
</script>