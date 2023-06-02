<script setup>
import SideBar from './sidebar/index.vue'
import AppHeader from './header/index.vue'

</script>

<template>
  <n-layout has-sider wh-full>
    <n-layout-sider
      bordered
      collapse-mode="width"
      :collapsed-width="64"
      :width="220"
      :native-scrollbar="false"
      :collapsed="false"
    >
      <SideBar />
    </n-layout-sider>

    <article >
      <header
        class="flex items-center border-b bg-white px-15 bc-eee"
        :style="`height: 60px;`"
      >
        <AppHeader />
      </header>
      <section class="main-section">
        <section class="main-section2 wh-full">
          <router-view></router-view>
        </section>
      </section>
    </article>
  </n-layout>
</template>
