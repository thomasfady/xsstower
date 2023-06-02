<template>
    <n-modal
      v-model:show="show"
      :mask-closable="false"
      preset="dialog"
      title="Create new user"
      content="Create"
      positive-text="Confirm"
      negative-text="Cancel"
      @positive-click="onPositiveClick"
    >
      <template #icon>
        <add-circle></add-circle>
      </template>

      <n-form vertical>
        <n-form-item label="Username" path="form.Username">
            <n-input v-model:value="form.Username" placeholder="Username" />
        </n-form-item>
        <n-form-item label="Password" path="form.Password">
            <n-input v-model:value="form.Password" placeholder="Password" />
        </n-form-item>
        <n-form-item label="Email" path="form.Email">
            <n-input v-model:value="form.Email" placeholder="Email" />
        </n-form-item>
        <n-form-item label="IsAdmin" path="form.IsAdmin">
            <n-switch v-model:value="form.IsAdmin" />
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
    Username: "",
    Email: "",
    Password: "",
    IsAdmin: false,
})

function open(){
    show.value = true
}

defineExpose({ open });

const emit = defineEmits(['userCreated'])
async function onPositiveClick(){
    
    const response = await fetch("/api/admin/users",{
        method:"POST", 
        headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify(form.value)});
    if (!response.ok) {
        message.error(await response.json())
    } else {
        message.success("User created")
        show.value = false;
        form.value = {
            Username: "",
            Email: "",
            Password: "",
            IsAdmin: false,
        }
        emit('userCreated')
    }
    return false
}
</script>