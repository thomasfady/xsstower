<template>
    <n-card title="Notifications">
        <n-tabs justify-content="space-evenly" type="line">
            <n-tab-pane v-for="notifier in notifiers" v-bind:key="notifier.Name" :name="notifier.Name.key" :tab="notifier.Name">
                <n-form>
                    <n-form-item :label="param.Text" path="form.current_password" v-for="param in notifier.Config" v-bind:key="param.text">
                        <n-input v-if="param.Sensitive" type="password" v-model:value="param.Value" :placeholder="param.Text" />
                        <n-input v-else type="text" v-model:value="param.Value" :placeholder="param.Text" />
                    </n-form-item>
                    <n-form-item>
                        <n-button type="success" icon-placement="left" @click="SetConfig()">
                            <template #icon>
                                <n-icon><save /></n-icon>
                            </template>
                            Save {{ notifier.Name }} config
                        </n-button>
                    </n-form-item>
                </n-form>
            </n-tab-pane>
        </n-tabs>
        <n-form>
            
        </n-form>
    </n-card>   
</template>

<script setup>
import { useMessage } from 'naive-ui'
import {SaveOutline as Save} from "@vicons/ionicons5"
import { ref } from "vue"

const message = useMessage()

const notifiers = ref([])

const form = ref({
    current_password: "",
    new_password: "",
    confirm_password: ""
    })
async function SetConfig() {
    let res = await fetch("/api/profile/notifications", 
    {
        method:"POST",
        headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify(notifiers.value)
    })
    if (!res.ok){
        message.error(await res.json())
    } else {
        message.success("Config Updated")
    }
}
function fetchConfig() {
    fetch("/api/profile/notifications")
    .then(response => response.json())
    .then((data) => {
        notifiers.value = data
    });
}
fetchConfig()
</script>