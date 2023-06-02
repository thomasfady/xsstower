<template>
    <n-space>
        <n-button type="info" icon-placement="left" @click="edit_handle">
            <template #icon>
                <n-icon><edit /></n-icon>
            </template>
        </n-button>
        <n-button type="error" icon-placement="left" @click="delete_handle">
            <template #icon>
                <n-icon><trash /></n-icon>
            </template>
        </n-button>
    </n-space>
</template>

<script setup>
import { 
    TrashOutline as Trash,
} from "@vicons/ionicons5";
import {
    Edit32Regular as Edit
} from "@vicons/fluent"
import {useDialog,useMessage } from 'naive-ui'

import { useRouter } from 'vue-router'

const router = useRouter()

const dialog = useDialog()
const message = useMessage()

const props = defineProps({
    handler: { required: true },
})
function edit_handle(){
    router.push('/handler/'+props.handler.ID)
}
const emit = defineEmits(['HandleDeleted'])
function delete_handle() {
    dialog.error({
          title: 'Delete',
          content: 'Delete the handler '+props.handler.Domain+'/'+props.handler.Path+' ?',
          positiveText: 'Delete',
          negativeText: 'Not Sure',
          onPositiveClick: () => {
            fetch("/api/handler/"+props.handler.ID, {method: "DELETE"},)
            .then(response => response.json())
            .then((data) => {
                emit('HandleDeleted')
                message.success('Handler deleted')
            });
            
          }
        })
}
</script>