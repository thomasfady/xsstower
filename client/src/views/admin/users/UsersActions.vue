<template>
    <n-space justify="end" size="small">
        <n-button type="warning" icon-placement="left" @click="setPassword.open()">
            <template #icon>
                <n-icon><lock /></n-icon>
            </template>
            Set Password
        </n-button>
        <n-button type="error" icon-placement="left" @click="delete_user">
            <template #icon>
                <n-icon><trash /></n-icon>
            </template>
            Delete
        </n-button>
        
    </n-space>
    <set-password ref="setPassword" :user="props.row"/>
</template>

<script setup>
import { 
    TrashOutline as Trash,
} from "@vicons/ionicons5";
import { ShieldLock20Regular as Lock } from "@vicons/fluent";
import {useDialog,useMessage } from 'naive-ui'

import SetPassword from './SetPassword.vue'
import { ref } from 'vue'

const dialog = useDialog()
const message = useMessage()

const props = defineProps({
    row: { required: true },
})

const setPassword = ref()

const emit = defineEmits(['UserDeleted'])
function delete_user() {
    dialog.error({
          title: 'Delete',
          content: 'Delete the user '+props.row.Username+' ?',
          positiveText: 'Delete',
          negativeText: 'Not Sure',
          onPositiveClick: () => {
            fetch("/api/admin/users/"+props.row.ID, {method: "DELETE"},)
            .then(response => response.json())
            .then((data) => {
                if(data.length == 0){
                    emit('UserDeleted')
                    message.success('User deleted')
                } else {
                    message.error(data)
                }
            });
            
          }
        })
}
</script>