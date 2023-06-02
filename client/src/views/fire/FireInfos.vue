<template>
    <n-space vertical v-if="fire">
        <n-grid x-gap="12" :cols="3">
            <n-gi>
                <n-card title="IP">
                {{ props.fire.Ip }}
                </n-card>
            </n-gi>
            <n-gi>
                <n-card title="Referer"  style="height: 100%;">
                {{ props.fire.Referer }}
                </n-card>
            </n-gi>
            <n-gi>
                <n-card title="Origin"  style="height: 100%;">
                {{ props.fire.Origin }}
                </n-card>
            </n-gi>
        </n-grid>
        <n-card title="UserAgent">
            <n-grid x-gap="12" :cols="3">
                <n-gi>
                    <n-statistic label="Browser" :value="parseUA().browser.name"/>
                </n-gi>
                <n-gi>
                    <n-statistic label="OS" :value="parseUA().os.name"/>
                </n-gi>
                <n-gi>
                    <n-statistic label="Arch" :value="parseUA().os.version"/>
                </n-gi>
            </n-grid>
        </n-card>
        
        <n-card title="Cookies">
            <fire-cookie :cookies="props.fire.Cookies"/>
        </n-card>
        <n-grid x-gap="12" :cols="2">
            <n-gi>
            <n-card title="LocalStorage" style="height: 100%;">
                <fire-storage :storage="props.fire.LocalStorage" type="local" />
            </n-card></n-gi><n-gi>
            <n-card title="SessionStorage" style="height: 100%;">
                <fire-storage :storage="props.fire.SessionStorage" type="local" />
            </n-card></n-gi>
        </n-grid>
        
        <n-card title="Screenshot" >
            <div style="text-align: center;"><img :src="props.fire.Screenshot" style="max-width: 50%;"/></div>
        </n-card>
        <n-card title="Dom">
            <code>{{ unb64(props.fire.Dom) }}</code>
        </n-card>
    </n-space>
</template>

<script setup>
import FireStorage from './FireStorage.vue'
import FireCookie from './FireCookie.vue'
import { UAParser } from 'ua-parser-js';
const props = defineProps({
    fire: { required: true },
})

function unb64(code){
    return atob(code)
}
function parseUA(){
    return new UAParser().setUA(props.fire.UserAgent).getResult()
}
</script>