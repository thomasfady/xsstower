<template>
    <n-card title="Change Password">
        <n-form>
            <n-form-item label="Current Password" path="form.current_password">
                <n-input type="password" v-model:value="form.current_password" placeholder="Current Password" />
              </n-form-item>
              <n-form-item label="New Password" path="form.new_password">
                <n-input type="password" v-model:value="form.new_password" placeholder="New Password" />
              </n-form-item>
              <n-form-item label="Confirm" path="form.confirm_password">
                <n-input type="password" v-model:value="form.confirm_password" placeholder="Confirm" />
              </n-form-item>
            <n-form-item>
                <n-button type="success" icon-placement="left" @click="ChangePassword()">
                    <template #icon>
                        <n-icon><save /></n-icon>
                    </template>
                    Save
                </n-button>
            </n-form-item>
        </n-form>
    </n-card>   
</template>

<script setup>
import { useMessage } from 'naive-ui'
import {SaveOutline as Save} from "@vicons/ionicons5"
import { ref } from "vue"

const message = useMessage()

const form = ref({
    current_password: "",
    new_password: "",
    confirm_password: ""
    })
async function ChangePassword() {
    let res = await fetch("/api/profile/set_password", 
    {
        method:"POST",
        headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify(form.value)
    })
    if (!res.ok){
        message.error(await res.json())
    } else {
        message.success("Password Updated")
    }
}
</script>