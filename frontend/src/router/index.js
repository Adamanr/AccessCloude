import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      redirect: '/sign_in'
    },
    {
      path: '/sign_in',
      name: 'signIn',
      component: () => import('../views/SignInPage.vue')
    },
    {
      path: '/sign_up',
      name: 'signUp',
      component: () => import('../views/SignUpPage.vue')
    },
    {
      path: '/forget',
      name: 'forgetPassword',
      component: () => import('../views/ForgetPasswordPage.vue')
    },
    {
      path: '/update',
      name: 'update',
      component: () => import('../views/UpdateInfoPage.vue')
    },
  ]
})

export default router
