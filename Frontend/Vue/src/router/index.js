import { createRouter, createWebHistory } from "vue-router";

const routes = [
  { path: "", redirect: { name: "home" } },
  { path: "/table/:filter", name: "client_table", component: () => import('../views/client/ReserveTables') },
  { path: "/atable", name: "admin_table", component: () => import('../views/admin/Table/Table') },
  { path: "/addtable", name: "admin_addtable", component: () => import('../views/admin/Table/AddTable') },
  { path: "/updateTable/:id", name: "updateTable", component: () => import('../views/admin/Table/UpdateTable') },
  { path: "/athematic", name: "admin_thematic", component: () => import('../views/admin/Thematic/Thematic') },
  { path: "/addthematic", name: "admin_addthematic", component: () => import('../views/admin/Thematic/AddThematic') },
  { path: "/register", name: "register", component: () => import('../views/login/Register') },
  { path: "/login", name: "login", component: () => import('../views/login/Login') },
  { path: "/profile", name: "profile", component: () => import('../views/profile/SetProfile') },
  { path: "/updateThematic/:id", name: "updateThematic", component: () => import('../views/admin/Thematic/UpdateThematic') },
  { path: "/home", name: "home", component: () => import('../views/Home') },
  { path: "/:catchAll(.*)", component: () => import('../views/Home') },

];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;