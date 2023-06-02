import { defineStore } from 'pinia'

import { ref } from 'vue'

export const useAdminConfigStore = defineStore('admin_config', () => {

    const config = ref(undefined);
    function fetchConfig() {
        fetch("/api/admin/config")
        .then(response => response.json())
        .then((data) => {
          config.value = data
        });
    }

    function GetConfig(keys) {
        let v = config.value
        keys.split('.').forEach(function (e){
            v = v[e]
        })
        return v
    }
    function SaveConfig(keys) {
        fetch("/api/admin/config",{
            method:"PUT", 
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(keys)})
            .then(response => response.json())
            .then((data) => {
                config.value = data
            });
    }
    return { config, fetchConfig, GetConfig, SaveConfig }
})