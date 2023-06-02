<script setup>
import { RefreshOutline as Refresh } from "@vicons/ionicons5";

import { NButton } from 'naive-ui'
import FiresActions from './fires/FiresActions.vue'
import { ref,h,watch } from 'vue'

import { useRouter } from 'vue-router'
import {
  Terminal as TerminalIcon,
} from "@vicons/ionicons5";

const router = useRouter()


const pagination= false
const page = ref(1)

watch(page, (page, prevPage) => {
    fetchFires()
})

const page_count = ref(1)
const total_items = ref(0)

function crop_text(value, size){
    if (value.length > size) {
        return value.substring(0, size-3) + "...";
    } else {
        return value
    }
    
}

const columns= [
    {
      title: "IP",
      key: "Ip"
    },
    {
      title: "Handler",
      key: "Handler"
    },
    {
      title: "Url",
      key: "Url",
      render (row) {
        return h(
          'p',
          {},
          { default: () => crop_text(row.Url, 50) }
        )
      }
    },
    {
      title: "UserAgent",
      key: "UserAgent",
      render (row) {
        return h(
          'p',
          {},
          { default: () => crop_text(row.UserAgent, 50) }
        )
      }
    },
    {
      title: "Referer",
      key: "Referer"
    },
    {
      title: "Action",
      key: "actions",
      render(row) {
        return h(
            FiresActions,
          {
            hit: row,
          },
        );
      }
    }
  ];

function fetchFires() {
    fetch("/api/fires?" + new URLSearchParams({
        page: page.value,
    }))
    .then(response => response.json())
    .then((data) => {
        fires.value = data['fires']
        page.value = data['page']
        total_items.value = data['total_rows']
        page_count.value = data['total_pages']
    });
}

const fires = ref([])
fetchFires()


</script>
<template>
    <header style="justify-content: space-between;">
        <h2>XSS Fires</h2>
        <div>
            <n-button type="info" icon-placement="left" @click="fetchFires">
                    <template #icon>
                        <n-icon><refresh /></n-icon>
                    </template>
                    Refresh
                </n-button>
        </div>
    </header>
    <n-card >
        <n-space vertical>      
            <n-data-table
                :columns="columns"
                :data="fires"
                :pagination="pagination"
                :bordered="true"
            >
            <template #empty>
                <n-empty description="No fires received. Use payload on vulnerable endpoint and enjoy !">
                <template #icon>
                    <n-icon>
                    <terminal-icon />
                    </n-icon>
                </template>
                <template #extra>
                    <n-button size="small" @click="router.push('/payloads')">
                        View Payloads
                    </n-button>
                </template>
                </n-empty>
            </template>
            </n-data-table>
            <n-pagination 
                v-model:page="page" 
                :page-size="10"
                :item-count="total_items"
                :page="page"
            >
                <template #suffix="{ itemCount }">
                    Total items {{ itemCount }}
                </template>
            </n-pagination>
        </n-space>
    </n-card>
</template>
<style scoped>
header {
    margin-bottom: 15px;
}
</style>