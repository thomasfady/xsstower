<script setup>
import {SaveOutline as Save} from "@vicons/ionicons5"
import { ref } from "vue"
import { useAdminConfigStore } from '@/stores/admin_config'

const store = useAdminConfigStore()
const keys = ref({
    "handlers.wildcard.domain": store.GetConfig("handlers.wildcard.domain"),
    "handlers.wildcard.path": store.GetConfig("handlers.wildcard.path"),
})
</script>
    
<template>
    <n-card title="Handlers">
        <n-form>
            <n-form-item path="enabled" label="Allow domain wildcard">
                <n-switch v-model:value="keys['handlers.wildcard.domain']"/>
            </n-form-item>
            <n-form-item path="enabled" label="Allow path wildcard">
                <n-switch v-model:value="keys['handlers.wildcard.path']"/>
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