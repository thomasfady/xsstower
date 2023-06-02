<script setup>
import { useRoute } from 'vue-router'
import { ref } from 'vue'

import FireInfos from './fire/FireInfos.vue'
import FirePages from './fire/FirePages.vue'
import FireShare from './fire/Share.vue'

async function fetchFire() {
    let res = await fetch("/api/sharing/"+route.params.id)
    if(res.status == 200){
        fire.value = await res.json()
    }
}
const route = useRoute();
const fire = ref(undefined)
fetchFire()

</script>
    
<template>
    <n-layout v-if="fire" style="background-color: #f5f6fb; padding:20px;" wh-full>
        <article >
            <header  style="text-align: center;">
                <h2 style="width: 100%;">Xss Fire on <small>{{ fire.Url }}</small></h2>
            </header>
            <n-space justify="center" style="height: 100%;"> 
                <n-tabs >
                <n-tab-pane name="Information" tab="Information" style="min-width: 1000px;">
                    <fire-infos :fire="fire"/>
                </n-tab-pane>
                <n-tab-pane name="pages" tab="Collected Pages" style="min-width: 1000px;">
                    <n-space vertical>
                        <n-card title="Collected Pages" >
                            <fire-pages :hit="fire"/>
                        </n-card>
                    </n-space>
                </n-tab-pane>
                </n-tabs>
            </n-space>
        </article>
    </n-layout>
    <n-space v-else justify="center" style="height: 100%;background-color: #f5f6fb;">
            <n-space vertical justify="center" style="height: 100%;">
                <n-result 
                    status="404"
                    title="404 Not Found"
                    description="Your XSS Hit public link is not available."
                >
            </n-result>
        </n-space>
    </n-space>
</template>
<style scoped>
header {
    margin-bottom: 15px;
}
body {
    padding: 20px
}
</style>