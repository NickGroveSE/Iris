import User from "../pages/User.vue";
import Home from "../pages/Home.vue";
import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const routes: RouteRecordRaw[] = [
  { path: "/", component: Home },
  { path: "/:username", component: User },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;