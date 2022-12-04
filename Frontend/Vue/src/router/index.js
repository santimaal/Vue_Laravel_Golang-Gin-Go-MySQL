import { createRouter, createWebHistory } from "vue-router";

const routes = [
  { path: "", redirect: { name: "home" } },
  { path: "/table", name: "client_table", component: () => import('../views/client/Table/Table') },
  { path: "/atable", name: "admin_table", component: () => import('../views/admin/Table/Table') },
  { path: "/addtable", name: "admin_addtable", component: () => import('../views/admin/Table/AddTable') },
  { path: "/updateTable/:id", name: "updateTable", component: () => import('../views/admin/Table/UpdateTable') },
  { path: "/thematic", name: "client_thematic", component: () => import('../views/client/Thematic/Thematic') },
  { path: "/athematic", name: "admin_thematic", component: () => import('../views/admin/Thematic/Thematic') },
  { path: "/addthematic", name: "admin_addthematic", component: () => import('../views/admin/Thematic/AddThematic') },
  { path: "/updateThematic/:id", name: "updateThematic", component: () => import('../views/admin/Thematic/UpdateThematic') },
  { path: "/reserve", name: "client_reserve", component: () => import('../views/client/Reserve/Reserve.vue') },
  { path: "/home", name: "home", component: () => import('../views/Home') },
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;