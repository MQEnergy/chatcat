import { createRouter, createWebHistory } from 'vue-router';
import routes from "./routes.js";

const router = createRouter({
  history: createWebHistory('/'),
  routes
});

router.beforeEach((to, from, next) => {
  next();
})

router.afterEach((to, from, next) => {
  window.scrollTo(0, 0);
})

export default router;