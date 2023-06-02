<template>
    <n-modal
      v-model:show="show"
      :mask-closable="false"
      preset="dialog"
      type="warning"
      title="Change password"
      content="Create"
      positive-text="Change password"
      negative-text="Cancel"
      @positive-click="onPositiveClick"
    >
      <template #icon>
        <lock></lock>
      </template>

      <n-form vertical>
        <n-form-item label="Username" >
            <n-input :value="props.user.Username" disabled />
        </n-form-item>
        <n-form-item label="New Password" path="form.password">
            <n-input v-model:value="form.password" placeholder="New Password" />
        </n-form-item>
      </n-form>
    </n-modal>
</template>
<script setup>
import {ref} from 'vue'
import { ShieldLock20Regular as Lock } from "@vicons/fluent";
import { useMessage } from 'naive-ui'

const show = ref(false);

const message = useMessage();

const props = defineProps({
    user: { required: true },
})

const form = ref({
    password: "",
})

function open(){
    show.value = true
}

defineExpose({ open });

async function onPositiveClick(){
    
    const response = await fetch("/api/admin/users/"+props.user.ID,{
        method:"PUT", 
        headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify(form.value)});
    if (!response.ok) {
        message.error(await response.json())
    } else {
        message.success("Password changed")
        show.value = false;
        form.value = {
            password: "",
        }
    }
    return false
}
</script>