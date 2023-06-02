<script setup>
import { ClipboardCode16Regular as Clipboard } from "@vicons/fluent";

import { useMessage } from 'naive-ui'
import { useRouter } from 'vue-router'
import { ref } from 'vue'

const router = useRouter()
const message = useMessage()


function copyToCliboard(value){
    navigator.clipboard.writeText(value);
    message.success("Payload copied to clipboard")
}

function fetchPayloads() {
        fetch("/api/payloads")
        .then(response => response.json())
        .then((data) => {
            if(data != null){
                payloads.value = data
                Object.keys(data).forEach(element => {
                    options.value.push(
                        {
                            label: element,
                            value: element
                        }
                    )
                });
                selected.value = options.value[0]['label']

            }
            
        });
    }

const payloads = ref([])
const selected = ref(undefined)

const options = ref([])
fetchPayloads()

function ChangeHandler(a){
    selected.value = a
}
</script>
<template>
    <header>
        <n-space>
            <h2>Payloads</h2>
            <n-select
                v-model:value="selected"
                filterable
                placeholder="Select an handler"
                :options="options"
                @change="ChangeHandler"
            />
        </n-space>
        
    </header>
    <n-space vertical>
        <n-card v-for="payload in payloads[selected]" v-bind:key="name" :title="payload.name">
        <template #header>
            <n-space>
                <n-button secondary icon-placement="left" @click="copyToCliboard(payload.Value)">
                    <template #icon>
                        <n-icon><clipboard /></n-icon>
                    </template>
                    Copy to Clipboard
                </n-button>
                <p>{{ payload.Name }}</p>
            </n-space>
        </template>
        
        <n-input
      :value="payload.Value"
      type="text"
      readonly
    />
    </n-card>
    </n-space>
    
    <n-result v-if="selected == undefined" status="error" title="No Handler Registered" description="You need an handler to get payloads">
    <template #footer>
      <n-button @click="router.push('/handlers')">Create one</n-button>
    </template>
  </n-result>
</template>
<style scoped>
header {
    margin-bottom: 15px;
}
</style>