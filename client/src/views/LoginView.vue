<script setup>
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'

const store = useUserStore()

async function handleLogin(){
  store.login(email.value, password.value, loading.value)
}


const email = ref("")
const password = ref("")

const loading = ref(false)
</script>

<template>
  <n-space justify="center" style="height: 100%;background-color: #f5f6fb;">
    <n-space vertical justify="center" style="height: 100%;">
      <div style="background-color: white;">
        <n-grid x-gap="0" :cols="2">
          <n-gi>
            <img height="350" width="350" style="" src="../assets/logo.svg">
          </n-gi>
          <n-gi>
            <n-form>
              <h2 style="text-align: center;">
                Xss Tower
              </h2>
              <n-form-item label="Email" path="email">
                <n-input v-model:value="email" placeholder="Email" />
              </n-form-item>
              <n-form-item label="Password" path="password">
                <n-input 
                  v-model:value="password" 
                  placeholder="Password" 
                  type="password"
                  show-password-on="mousedown"
                  :maxlength="20"
                  @keydown.enter="handleLogin"
                  />
              </n-form-item>
              <n-form-item >
                <n-button type="primary"
                :loading="loading"
                style="width: 100%;"
                @click="handleLogin">
                  Login
                </n-button>
              </n-form-item>
            </n-form>
            <n-divider />
            <p>You don't have an account ? <router-link h-60 f-c-c to="/">Create one</router-link></p>
          </n-gi>
        </n-grid>
        
        
        
      </div>
      
    </n-space>
  </n-space>
</template>
<style scoped>
p {
  text-align: center;
  color: gray;
  font-size: 12px
}
.n-grid > div {
  padding: 25px;
}
.n-grid {
  height: 400px;width: 800px;
  box-shadow: 0 1px 2px -2px #00000029, 0 3px 6px #0000001f, 0 5px 12px 4px #00000017;
  border-radius: 10px;
}

</style>