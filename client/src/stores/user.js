import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import {
    createDiscreteApi,
    darkTheme,
    lightTheme
  } from "naive-ui";

const themeRef = ref("light");
const configProviderPropsRef = computed(() => ({
  theme: themeRef.value === "light" ? lightTheme : darkTheme
}));

const { message } = createDiscreteApi(
  ["message"],
  {
    configProviderProps: configProviderPropsRef
  }
);

export const useUserStore = defineStore('user', () => {

    const userInfos = ref(initInfos());
    const router = useRouter()
    
    function initInfos(){
        let user =  localStorage.getItem('user');
        if(user) return JSON.parse(user)
        else return {}
    }
    async function login(email, password, loading){
        loading = true;
        const response = await fetch("/api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email: email, password: password }),
        })
            loading = false;
        if (!response.ok) {
            message.error(await response.json())
        } else {
            userInfos.value = await response.json()
            localStorage.setItem('user', JSON.stringify(userInfos.value))
            router.push('/')
        }
    }
    function isLoggedIn() {
        return userInfos.value.Username != undefined
    }
    function isAdmin() {
        return userInfos.value.IsAdmin
    }

    async function logout(){
        const response = await fetch("/api/logout", {
            method: "POST",
        })
        if (!response.ok) {
            message.error(await response.json())
        } else {
            userInfos.value = {}
            localStorage.removeItem('user')
            router.push('/login')
        }
    }

    return { userInfos, login, logout, isLoggedIn, isAdmin }
})