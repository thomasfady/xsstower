<script setup>
import { RefreshOutline as Refresh, AddOutline as Add, AddCircleOutline as AddCircle } from "@vicons/ionicons5";

import { NSwitch } from 'naive-ui'
import HandlersActions from './handlers/HandlersActions.vue'
import CreateHandler from './handlers/CreateHandler.vue'
import { ref,h } from 'vue'
import {
  FlashOutline as Flash
} from "@vicons/ionicons5";


const pagination= false
const create = ref()

function openModal(){
  create.value.open()
}

const columns= [
    {
      title: "Domain",
      key: "Domain"
    },
    {
      title: "Path",
      key: "Path"
    },
    {
      title: "Screenshot",
      key: "Screenshot",
      render(row) {
        return h(
          NSwitch,
          {
            value: row.Screenshot,
            "on-update:value": function (){ updateSwitch(row, "Screenshot")}
          },
        );
      }
    },
    {
      title: "Dom",
      key: "Dom",
      render(row) {
        return h(
          NSwitch,
          {
            value: row.Dom,
            "on-update:value": function (){ updateSwitch(row, "Dom")}
          },
        );
      }
    },
    {
      title: "Action",
      key: "actions",
      render(row) {
        return h(
          HandlersActions,
          {
            handler: row,
            onHandleDeleted: fetchHandlers
          },
        );
      }
    }
  ];

function fetchHandlers() {
    fetch("/api/handlers")
    .then(response => response.json())
    .then((data) => {
      handlers.value = data
    });
}

function updateSwitch(row, type){
  
  let data = {}
  data['Screenshot'] = row['Screenshot']
  data['CollectedPages'] = row['CollectedPages']
  data['Dom'] = row['Dom']
  data[type] = !row[type]
  fetch("/api/handler/"+row.ID,{
    method:"PUT", 
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data)})
    .then(response => response.json())
    .then((d) => {
      row['Screenshot'] = d['Screenshot']
      row['Dom'] = d['Dom']
    });
}

const handlers = ref([])
fetchHandlers()

</script>
<template>
    <header style="justify-content: space-between;">
        <h2>Handlers</h2>
        <div>
          <n-space>
            <n-button type="success" icon-placement="left" @click="openModal()">
                    <template #icon>
                        <n-icon><add /></n-icon>
                    </template>
                    Add
                </n-button>
            <n-button type="info" icon-placement="left" @click="fetchHandlers">
                    <template #icon>
                        <n-icon><refresh /></n-icon>
                    </template>
                    Refresh
            </n-button>
          </n-space>
        </div>
    </header>
    <n-card >
        <n-data-table
            :columns="columns"
            :data="handlers"
            :pagination="pagination"
            :bordered="true"
        >
          <template #empty>
            <n-empty description="No handler. Create one to receive XSS Fires !">
              <template #icon>
                <n-icon>
                  <flash />
                </n-icon>
              </template>
            </n-empty>
          </template>
        </n-data-table>
    </n-card>
    <create-handler @handlerCreated="fetchHandlers" ref="create"/>
</template>
<style scoped>
header {
    margin-bottom: 15px;
}
</style>