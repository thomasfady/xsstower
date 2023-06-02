<template>
    <n-data-table
    :columns="columns"
    :data="props.hit.CollectedPages"
    >
    </n-data-table>
</template>
<script setup>
import { NButton } from 'naive-ui'
import { h } from 'vue'
const props = defineProps({
    hit: { required: true },
})

function download(row){
    window.open(window.location.pathname.replace("app","api")+"/file/"+row.ID, '_blank')
}

const columns = [
    {
        title: "Path",
        key: "Path"
    },
    {
        title: "Download",
        key: "download",
        render(row) {
        return h(
            NButton,
          {
            strong: true,
            tertiary: true,
            onClick: () => download(row)
          },{ default: () => 'Download' }
        );
      }
    }
];
</script>