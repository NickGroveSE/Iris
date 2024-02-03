import AccessGranted from "../pages/AccessGranted.vue";
import Home from "../pages/Home.vue";
import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const routes: RouteRecordRaw[] = [
  { path: "/", component: Home },
  { path: "/access_granted", component: AccessGranted },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;