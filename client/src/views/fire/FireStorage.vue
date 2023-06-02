<template>
    <n-data-table
    :columns="columns"
    :data="data"
    >
    </n-data-table>
</template>
<script setup>
const props = defineProps({
    storage: { required: true },
})
const columns = [
    {
        title: "Key",
        key: "key"
    },
    {
        title: "Value",
        key: "value"
    }
];


function props_to_data(){
    let d = []
    if(props.storage == ""){
        return d
    }
    for (const [key, value] of Object.entries(JSON.parse(props.storage))) {
        d.push({
            "key": key,
            "value": value
        })
        console.log(key, value);
    }
    return d
}

const data = props_to_data();
</script>