<template>
    <header style="justify-content: space-between;">
        <h2>Handler</h2>
        <div>
            <n-space>
                <n-button type="success" icon-placement="left" @click="saveHandle">
                    <template #icon>
                        <n-icon><save /></n-icon>
                    </template>
                    Save
                </n-button>
            </n-space>
        </div>
    </header>
    <n-tabs>
      <n-tab-pane name="Information" tab="Information">
        <n-card>
            <n-form v-if="handler">
                <n-form-item path="domain" label="Domain">
                    <n-input v-model:value="handler.Domain" disabled />
                </n-form-item>
                <n-form-item path="path" label="Path">
                    <n-input v-model:value="handler.Path" disabled />
                </n-form-item>
                <n-form-item path="screenshot" label="Screenshot">
                    <n-switch v-model:value="handler.Screenshot"/>
                </n-form-item>
                <n-form-item path="dom" label="Dom">
                    <n-switch v-model:value="handler.Dom"/>
                </n-form-item>
                <n-form-item path="CollectedPages" label="Collected Page">
                    <n-dynamic-input
                        v-model:value="handler.CollectedPages"
                        placeholder="Type URL (relative or absolute) or filename (file:///etc/passwd)"
                    />
                </n-form-item>
            </n-form>
        </n-card>
      </n-tab-pane>
      <n-tab-pane name="Members" tab="Members">
        <n-card title="Add Member">
            <n-form inline>
            <n-form-item label="User" path="user_mail">
                <n-select
                    v-model:value="user_mail"
                    filterable
                    placeholder="Search Users"
                    :options="users"
                    :loading="loading"
                    clearable
                    remote
                    @search="handleSearch"
                />
            </n-form-item>
            <n-form-item label="Role" path="user_role">
                <n-select :options="options" v-model:value="user_role"/>
            </n-form-item>
            <n-form-item label="" path="user_role">
                <n-button type="success" secondary @click="createHandle">Create</n-button>
            </n-form-item>
        </n-form>
        </n-card>
        <n-card title="Member List">
            <n-data-table
                :columns="columns"
                :data="handler.Members"
                :bordered="true"
            >
                
            </n-data-table>
        </n-card>
      </n-tab-pane>
    </n-tabs>
    
</template>

<style scoped>
header {
    margin-bottom: 15px;
}
.n-card {
    margin-bottom: 15px;
}
.n-select {
    min-width: 100px;
}
</style>

<script setup>
import { useRoute } from 'vue-router'
import { ref, h } from 'vue'
import {SaveOutline as Save} from "@vicons/ionicons5"
import { NSelect, NButton, useMessage, useDialog } from "naive-ui"

const message = useMessage();
const dialog = useDialog();
const route = useRoute();
const handler = ref(null)
const loading = ref(false)

function createHandle(){
    fetch("/api/handler/"+handler.value.ID+"/members", {
        method:"POST",
        headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify({"email":user_mail.value,"role":user_role.value})
    })
    .then(() => {
        fetchHandle()
    });
}

function fetchHandle() {
    fetch("/api/handler/"+route.params.id)
    .then(response => response.json())
    .then((data) => {
        handler.value = data
    });
}
function saveHandle(){
    fetch("/api/handler/"+handler.value.ID, {
        method:"PUT",
        headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify(handler.value)
    })
    .then(response => response.json())
    .then((data) => {
        handler.value = data
    });
}

function handleSearch(search){
    if(search.length > 2){
        fetch("/api/users?search="+search, {
            method:"GET",
            headers: {
            "Content-Type": "application/json",
            },
        })
        .then(response => response.json())
        .then((data) => {
            users.value = []
            data.forEach(function(element){
                users.value.push({"value":element.Email, "label":element.Email})
            })
        });
    }
    
}

const columns= [
    {
      title: "Email",
      key: "User.Email"
    },
    {
      title: "Role",
      key: "Permission",
      render (row) {
        return h(
            NSelect,
          {
            options: options,
            value: row.Permission,
            onChange: (value) => change_role(row, value)
          }
        )
      }
    },
    {
      title: "Action",
      render (row) {
        return h(
            NButton,
          {
            type: "error",
            onClick: () => remove_member(row)
          },
          { default: () => "Delete" }
        )
      }
    }
]


fetchHandle()

const options = [
    {"value":"READ", "label":"Reader"},
    {"value":"WRITE", "label":"Writer"},
    {"value":"OWNER", "label":"Owner"}
]
const users = ref([])
const user_role = ref("READ")
const user_mail = ref("")

async function change_role(row, value){
    let res = await fetch("/api/handler/"+handler.value.ID+"/members", {
        method:"PUT",
        headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify({"user_id": row.UserID, "role": value})
    })
    if(res.status == 200){
        row.Permission = value
    } else {
        const msg = await res.json()
        message.error(msg)
    }
    
}

function remove_member(row){
    dialog.error({
        title: 'Delete',
        content: 'Delete the member '+row.User.Email+' ?',
        positiveText: 'Delete',
        negativeText: 'Not Sure',
        onPositiveClick: () => {
        fetch("/api/handler/"+handler.value.ID+"/members", 
        {
            method:"DELETE",
            headers: {
            "Content-Type": "application/json",
            },
            body: JSON.stringify({"user_id": row.UserID})
        })
        .then(response => response.json())
        .then((data) => {
            if(data.length == 0){
                fetchHandle()
                message.success('Member deleted')
                
            } else {
                message.error(data)
            }
        });
        
        }
    })
}
</script>