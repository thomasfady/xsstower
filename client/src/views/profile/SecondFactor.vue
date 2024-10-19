<template>
    <n-card title="2FA">
        <n-form v-if="!second_factor.enabled">
            <n-grid x-gap="12" :cols="2">
            <n-gi>
                <n-qr-code size="200" type="svg" :value="second_factor.totp_uri" />
            </n-gi>
            <n-gi>
                <n-form-item label="Secret" >
                    <n-input type="password" show-password-on="mousedown" v-model:value="second_factor.secret" placeholder="New Password" />
                </n-form-item>
                <n-form-item label="TOTP" >
                    <n-input type="password" show-password-on="mousedown" v-model:value="second_factor.otp_code" placeholder="Type valid OTP code to enable 2FA" />
                </n-form-item>
                <n-form-item>
                    <n-button type="success" icon-placement="left" @click="EnableOTP()">
                        <template #icon>
                            <n-icon><save /></n-icon>
                        </template>
                        Enable 2FA
                    </n-button>
                </n-form-item>
            </n-gi>
        </n-grid>
            
            
        </n-form>
        <n-form v-else>
            <n-form-item>
                <n-popconfirm positive-text="Disable" negative-text="Cancel" @positive-click="RemoveOTP">
                    <template #trigger>
                        <n-button type="error" icon-placement="left">
                        <template #icon>
                            <n-icon><trash /></n-icon>
                        </template>
                        Disable 2FA
                    </n-button>
                    </template>
                    Are you sure you want to remove two-factor authentication?
                </n-popconfirm>
                
            </n-form-item>
        </n-form>
        
    </n-card>   
</template>

<script setup>
import { useMessage } from 'naive-ui'
import {SaveOutline as Save} from "@vicons/ionicons5"
import {TrashBinOutline as Trash} from "@vicons/ionicons5"
import { ref } from "vue"


const message = useMessage()

const second_factor = ref({
    enabled: false,
    totp_uri	: "",
    secret	: "",
    otp_code : "",
    })


async function EnableOTP() {
    let res = await fetch("/api/profile/2fa", 
    {
        method:"POST",
        headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify({"otp_code":second_factor.value.otp_code, "secret":second_factor.value.secret, "enable":true})
    })
    if (!res.ok){
        message.error(await res.json())
    } else {
        message.success("OTP Enable")
        second_factor.value.enabled = true;
    }
}

async function RemoveOTP() {
    let res = await fetch("/api/profile/2fa", 
    {
        method:"DELETE",
        headers: {
        "Content-Type": "application/json",
        }
    })
    if (!res.ok){
        message.error(await res.json())
    } else {
        message.success("OTP Disabled")
        fetchSecondFactor();
    }
}


function fetchSecondFactor() {
    fetch("/api/profile/2fa")
    .then(response => response.json())
    .then((data) => {
        second_factor.value = data
    });
}
fetchSecondFactor()
</script>