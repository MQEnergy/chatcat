import LayoutIndex from '@components/layouts/index.vue'

const pre = 'settings-';

export default {
  path: '/settings',
  name: 'settings',
  component: LayoutIndex,
  redirect: {
    name: `${pre}index`
  },
  children: [
    {
      path: 'index',
      name: `${pre}index`,
      meta: {
        title: '设置'
      },
      component: () => import('@views/settings/index.vue')
    }
  ]
}