import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: () => import('@/layout/HomeLayout.vue'),
      meta: { requiresAuth: true },
      children: [
        {
          path: "",
          name: "home",
          component: () => import('@/views/Dashboard.vue'),
        },{
          path: "/payloads",
          name: "payloads",
          component: () => import('@/views/Payloads.vue'),
        },{
          path: "/fires",
          name: "fires",
          component: () => import('@/views/Fires.vue'),
        },
        {
          path: "/fire/:id",
          name: "hit",
          component: () => import('@/views/Fire.vue'),
        },{
          path: "/handlers",
          name: "handlers",
          component: () => import('@/views/Handlers.vue'),
        },
        {
          path: "/handler/:id",
          name: "handler",
          component: () => import('@/views/Handler.vue'),
        }
      ]
    },{
      path: '/admin',
      component: () => import('@/layout/HomeLayout.vue'),
      meta: { requiresAuth: true, requiresAdmin:true },
      children: [
        {
          path: "/admin/users",
          name: "users",
          component: () => import('@/views/admin/Users.vue'),
        },{
          path: "/admin/settings",
          name: "settings",
          component: () => import('@/views/admin/Settings.vue'),
        }
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue')
    },
    {
      path: '/sharing/:id',
      name: 'sharing',
      component: () => import('@/views/PublicFire.vue')
    }
  ]
})

router.beforeEach((to, from) => {
  const userStore = useUserStore()
  if (to.meta.requiresAuth && !userStore.isLoggedIn()) {
    return {
      path: '/login',
      query: { redirect: to.fullPath },
    }
  }
  if (to.meta.requiresAdmin && !userStore.isAdmin()) {
    return {
      path: '/',
    }
  }
  return true
})

export default router
