<template>
    <n-data-table
    :columns="columns"
    :data="data"
    >
    </n-data-table>
</template>
<script setup>
const props = defineProps({
    cookies: { required: true },
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
    if(props.cookies == ""){
        return d
    }
    if(props.cookies.split('; ').length != 0){
        for (const value of props.cookies.split('; ')) {
            d.push({
                "key": value.split('=')[0],
                "value": value.split('=')[1]
            })
        }
    }
    
    return d
}

const data = props_to_data();
</script>