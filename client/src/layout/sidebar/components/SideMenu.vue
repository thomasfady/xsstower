<template>
  <n-menu
    ref="menu"
    class="side-menu"
    accordion
    :indent="18"
    :collapsed-icon-size="22"
    :collapsed-width="64"
    :options="menuOptions"
    :value="activeKey"
    @update:value="handleMenuSelect"
  />
  
</template>

<script setup>
import { defineComponent, h } from "vue";
import { RouterLink } from "vue-router";
import { NIcon, useMessage } from "naive-ui";
import { useUserStore } from '@/stores/user'


import {
  Terminal as TerminalIcon,
  HomeOutline as HomeIcon,
  FlashOutline as Flash,
  PersonCircleOutline as PersonIcon,
  SettingsOutline as Settings
} from "@vicons/ionicons5";
import {TargetArrow20Filled as FiresIcon } from "@vicons/fluent"

const store = useUserStore()
function renderIcon(icon) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const menuOptions = [
  {
    label: () => h(
      RouterLink,
      {
        to: {
          name: "home",
        }
      },
      { default: () => "Dashboard" }
    ),
    key: "go-back-home",
    icon: renderIcon(HomeIcon)
  },
  {
    key: "divider-1",
    type: "divider",
    props: {
      style: {
        marginLeft: "32px"
      }
    }
  },
  {
    label: () => h(
      RouterLink,
      {
        to: {
          name: "payloads",
        }
      },
      { default: () => "Payloads" }
    ),
    key: "payloads",
    icon: renderIcon(TerminalIcon)
  },
  {
    label: () => h(
      RouterLink,
      {
        to: {
          name: "fires",
        }
      },
      { default: () => "XSS Fires" }
    ),
    key: "fires",
    icon: renderIcon(FiresIcon)
  },
  {
    label: () => h(
      RouterLink,
      {
        to: {
          name: "handlers",
        }
      },
      { default: () => "Handlers" }
    ),
    key: "handlers",
    icon: renderIcon(Flash)
  },
  {
    key: "divider-1",
    type: "divider",
    show: store.isAdmin(),
    props: {
      style: {
        marginLeft: "32px"
      }
    }
  },
  {
    label: () => h(
      RouterLink,
      {
        to: {
          name: "users",
        }
      },
      { default: () => "Users" }
    ),
    key: "users",
    show: store.isAdmin(),
    icon: renderIcon(PersonIcon)
  },
  {
    label: () => h(
      RouterLink,
      {
        to: {
          name: "settings",
        }
      },
      { default: () => "Settings" }
    ),
    key: "settings",
    show: store.isAdmin(),
    icon: renderIcon(Settings)
  }
];

</script>

