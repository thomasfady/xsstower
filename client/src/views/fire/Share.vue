<template>
  <div>
          <n-space>
            <n-button v-if="!props.fire.PublicId" type="success" icon-placement="left" @click="enableShare()">
                    <template #icon>
                        <n-icon><add /></n-icon>
                    </template>
                    Share
            </n-button>
            <n-button v-if="fire.PublicId" type="info" icon-placement="left" @click="showModal = true">
                    <template #icon>
                        <n-icon><share-social /></n-icon>
                    </template>
                    Share
            </n-button>
          </n-space>
        </div>
    <n-modal 
      style="max-width: calc(100vw - 32px); width: 446px;"
      v-model:show="showModal" 
      preset="card"
      v-if="fire"
      >
      <template #header>
        Copy Public Link
      </template>
        <n-space vertical>
          <n-button secondary icon-placement="left" @click="copyToCliboard(publicLink())">
                <template #icon>
                    <n-icon><clipboard /></n-icon>
                </template>
                Copy to Clipboard
            </n-button>
          <n-input
            :value="publicLink()"
            type="text"
            readonly
          />
        </n-space>
      <template #footer>
        <n-button v-if="props.fire.PublicId" type="error" icon-placement="left" @click="removeShare()">
          <template #icon>
              <n-icon><trash /></n-icon>
          </template>
          Remove Sharing
        </n-button>
      </template>
  </n-modal>
</template>
<script setup>
import {
  ShareSocialOutline as ShareSocial,
  AddCircleOutline as Add,
  TrashOutline as Trash,
} from "@vicons/ionicons5"
import { ClipboardCode16Regular as Clipboard } from "@vicons/fluent";
import { ref } from "vue"
import { useMessage } from "naive-ui";

const emit = defineEmits(['reloadFire'])
const message = useMessage()
const showModal = ref(false)
const props = defineProps({
    fire: { required: true },
})

function enableShare() {
  fetch("/api/fire/"+props.fire.CorrelationKey+"/share",{method:"POST"})
  .then(() => {
    emit('reloadFire')
  });
}
function copyToCliboard(value){
    navigator.clipboard.writeText(value);
    message.success("Public link copied to clipboard")
    showModal.value = false
}

function publicLink(){
  let loc = window.location
  return loc.protocol + "//" + loc.host + "/app/sharing/" + props.fire.PublicId
}
function removeShare(){
  fetch("/api/fire/"+props.fire.CorrelationKey+"/share",{method:"DELETE"})
  .then(() => {
    emit('reloadFire')
    showModal.value = false
  });
}
</script>
<style scoped>
.n-dialog.n-modal {
    width: 446px;
    max-width: calc(100vw - 32px);
}
</style>