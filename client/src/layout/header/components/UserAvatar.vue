<template>
  <n-dropdown :options="options" @select="handleSelect">
    <div flex cursor-pointer items-center>
      <n-button quaternary>
        <template #icon>
          <n-icon><person-icon /></n-icon>
        </template>
        {{ store.userInfos.Username }}
      </n-button>
    </div>
  </n-dropdown>
</template>

<script setup>
import {
  LogOutOutline as LogoutIcon,
  PersonCircleOutline as PersonIcon,
} from '@vicons/ionicons5'
import { h } from "vue";
import { NIcon } from "naive-ui";
import {useDialog } from 'naive-ui'

import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'


const router = useRouter()
const store = useUserStore()

const dialog = useDialog()

const renderIcon = (icon) => {
  return () => {
    return h(NIcon, null, {
      default: () => h(icon)
    });
  };
};
const options = [
  {
    label: 'Profile',
    key: 'profile',
    icon: renderIcon(PersonIcon),
  },
  {
    label: 'Logout',
    key: 'logout',
    icon: renderIcon(LogoutIcon),
  }
  
]

function handleSelect(key) {
  if (key === 'logout') {
    dialog.warning({
      title: 'Logout',
      type: 'info',
      content: 'Logout from your account ?',
      negativeText: 'Cancel',
      positiveText: 'Logout',
      onPositiveClick: () => {
        store.logout()
      },
    })
  } else if (key === 'profile') {
    router.push("/profile")
  }
}
</script>
