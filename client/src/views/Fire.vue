<script setup>
import { useRoute } from 'vue-router'
import { ref } from 'vue'

import FireInfos from './fire/FireInfos.vue'
import FirePages from './fire/FirePages.vue'
import FireShare from './fire/Share.vue'

function fetchFire() {
    fetch("/api/fire/"+route.params.id)
    .then(response => response.json())
    .then((data) => {
        fire.value = data
        
    });
}
const route = useRoute();
const fire = ref(undefined)
fetchFire()

</script>
    
<template>
    <header v-if="fire"  style="justify-content: space-between;">
      <h2>Xss Fire on <small>{{ fire.Url }}</small></h2>
      <fire-share @reloadFire="fetchFire" v-if="fire" :fire="fire"/>
    </header>
    <n-tabs>
      <n-tab-pane name="Information" tab="Information">
        <fire-infos :fire="fire"/>
      </n-tab-pane>
      <n-tab-pane name="pages" tab="Collected Pages">
        <fire-pages :hit="fire"/>
      </n-tab-pane>
    </n-tabs>
</template>
<style scoped>
header {
    margin-bottom: 15px;
}
</style>