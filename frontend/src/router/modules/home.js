import LayoutIndex from '@components/layouts/index.vue'

const pre = 'home-';

export default {
  path: '/',
  component: LayoutIndex,
  redirect: {
    name: `${pre}index`
  },
  children: [
    {
      path: 'index',
      name: `${pre}index`,
      meta: {
        title: '首页'
      },
      component: () => import('@views/home/index.vue')
    }
  ]
}