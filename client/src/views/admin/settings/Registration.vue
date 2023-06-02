<script setup>
import {SaveOutline as Save} from "@vicons/ionicons5"
import { ref } from "vue"
import { useAdminConfigStore } from '@/stores/admin_config'

const store = useAdminConfigStore()
const keys = ref({
    "registration.enabled": store.GetConfig("registration.enabled")
})

</script>
    
<template>
    <n-card title="Registration">
        <n-form>
            <n-form-item path="enabled" label="Enabled">
                <n-switch v-model:value="keys['registration.enabled']"/>
            </n-form-item>
            <n-form-item>
                <n-button type="success" icon-placement="left" @click="store.SaveConfig(keys)">
                    <template #icon>
                        <n-icon><save /></n-icon>
                    </template>
                    Save
                </n-button>
            </n-form-item>
        </n-form>
    </n-card>
</template>