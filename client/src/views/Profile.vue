<template>
    <header style="justify-content: space-between;">
        <h2>Profile</h2>
    </header>
    <n-space vertical>
        <n-grid x-gap="12" :cols="2">
            <n-gi>
                <change-password />
            </n-gi>
            <n-gi>
                <second-factor />
            </n-gi>
        </n-grid>
        <notification />
  </n-space>
</template>

<script setup>
import ChangePassword from "./profile/ChangePassword.vue"
import SecondFactor from "./profile/SecondFactor.vue"
import Notification from "./profile/Notification.vue"
</script>