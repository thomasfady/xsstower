<template>
    <n-modal
      v-model:show="show"
      :mask-closable="false"
      preset="dialog"
      title="Create new handler"
      content="Create"
      positive-text="Confirm"
      negative-text="Cancel"
      @positive-click="onPositiveClick"
    >
      <template #icon>
        <add-circle></add-circle>
      </template>

      <n-form vertical>
        <n-form-item label="Domain" path="form.domain">
            <n-input v-model:value="form.domain" placeholder="Domain" />
        </n-form-item>
        <n-form-item label="Path" path="form.path">
            <n-input v-model:value="form.path" placeholder="Path" />
        </n-form-item>
        <n-form-item label="Take Screenshot" path="form.screenshot">
            <n-switch v-model:value="form.screenshot" />
        </n-form-item>
        <n-form-item label="Get Dom" path="form.dom">
            <n-switch v-model:value="form.dom" />
        </n-form-item>
      </n-form>
    </n-modal>
</template>
<script setup>
import {ref} from 'vue'
import { AddCircleOutline as AddCircle } from "@vicons/ionicons5";
import { useMessage } from 'naive-ui'

const show = ref(false);

const message = useMessage();

const form = ref({
    domain: "*",
    path: "admin",
    screenshot: true,
    dom: true,
})

function open(){
    show.value = true
}

defineExpose({ open });

const emit = defineEmits(['handleCreated'])
async function onPositiveClick(){
    
    const response = await fetch("/api/handlers",{
        method:"POST", 
        headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify(form.value)});
    if (!response.ok) {
        message.error(await response.json())
    } else {
        message.success("Handler created")
        show.value = false;
        form.value = {
            domain: "*",
            path: "admin",
            screenshot: true,
            dom: true,
        }
        emit('handleCreated')
    }
    return false
}
</script>