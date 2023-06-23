<template>
    <n-space>
        <n-button type="info" icon-placement="left" @click="show_hit">
            <template #icon>
                <n-icon><eye /></n-icon>
            </template>
        </n-button>
        <n-button type="error" icon-placement="left" @click="delete_hit">
            <template #icon>
                <n-icon><trash /></n-icon>
            </template>
        </n-button>
    </n-space>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { 
    TrashOutline as Trash,
    EyeOutline as Eye
} from "@vicons/ionicons5";
import {useDialog,useMessage } from 'naive-ui'

const router = useRouter()
const dialog = useDialog()
const message = useMessage()

const props = defineProps({
  hit: { required: true },
})

const emit = defineEmits(['ReloadFires'])
function delete_hit() {
    dialog.error({
          title: 'Delete',
          content: 'Delete the hit from '+props.hit.Ip+' ?',
          positiveText: 'Delete',
          negativeText: 'Not Sure',
          onPositiveClick: () => {
            fetch("/api/fire/"+props.hit.CorrelationKey, {method: "DELETE"},)
            .then(response => response.json())
            .then((data) => {
                fire.value = data
            });
            message.success('Hit deleted')
            emit('ReloadFires')
          }
        })
}
function show_hit(){
    router.push('/fire/'+props.hit.CorrelationKey)
}
</script>