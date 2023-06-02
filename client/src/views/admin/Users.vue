<script setup>
import { RefreshOutline as Refresh, AddOutline as Add, AddCircleOutline as AddCircle } from "@vicons/ionicons5";

import { NSwitch, useMessage } from 'naive-ui'
import UsersActions from './users/UsersActions.vue'
import CreateUser from './users/CreateUser.vue'
import { ref,h } from 'vue'
const message = useMessage()

const pagination= false
const create = ref()
const setPassword = ref()

function openModal(){
  create.value.open()
}

const columns= [
    {
      title: "Username",
      key: "Username"
    },
    {
      title: "Email",
      key: "Email"
    },
    {
      title: "IsAdmin",
      key: "IsAdmin",
      render(row) {
        return h(
          NSwitch,
          {
            value: row.IsAdmin,
            "on-update:value": function (){ updateSwitch(row)}
          },
        );
      }
    },
    {
      title: "Action",
      key: "actions",
      align: "right",
      width: 300,
      render(row) {
        return h(
          UsersActions,
          {
            row: row,
            onUserDeleted: fetchUsers,
          },
        );
      }
    }
  ];

function fetchUsers() {
    fetch("/api/admin/users")
    .then(response => response.json())
    .then((data) => {
      users.value = data
    });
}

function updateSwitch(row){
  
  let data = {}
  data['IsAdmin'] = !row['IsAdmin']
  fetch("/api/admin/users/"+row.ID,{
    method:"PUT", 
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data)})
    .then(response => response.json())
    .then((d) => {
      console.log(d.length)
      if(d.length == 0){
        row['IsAdmin'] = d['IsAdmin']
      } else {
          message.error(d)
      }
      
    });
}

const users = ref([])
fetchUsers()

</script>
<template>
    <header style="justify-content: space-between;">
        <h2>Users</h2>
        <div>
          <n-space>
            <n-button type="success" icon-placement="left" @click="openModal()">
                    <template #icon>
                        <n-icon><add /></n-icon>
                    </template>
                    Add
                </n-button>
            <n-button type="info" icon-placement="left" @click="fetchUsers">
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
            :data="users"
            :pagination="pagination"
            :bordered="true"
        >
        </n-data-table>
    </n-card>
    <create-user @userCreated="fetchUsers" ref="create"/>
</template>
<style scoped>
header {
    margin-bottom: 15px;
}
</style>