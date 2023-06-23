<template>
    <header style="justify-content: space-between;">
        <h2>Profile</h2>
    </header>
    <n-space vertical>
        <change-password />
        <notification />
  </n-space>
</template>

<script setup>
import ChangePassword from "./profile/ChangePassword.vue"
import Notification from "./profile/Notification.vue"
</script>