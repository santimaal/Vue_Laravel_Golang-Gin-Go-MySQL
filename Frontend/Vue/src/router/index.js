import { createRouter, createWebHistory } from "vue-router";
import UpdateTodo from '../views/UpdateTodo';

const routes = [
  { path: "", redirect: { name: "home" } },
  { path: "/table", name: "client_table", component: () => import('../views/client/Table/Table') },
  { path: "/atable", name: "admin_table", component: () => import('../views/admin/Table/Table') },
  { path: "/addtable", name: "admin_addtable", component: () => import('../views/admin/Table/AddTable') },
  { path: "/home", name: "home", component: () => import('../views/Home') },
  { path: "/todos/update/:id", name: "updateTodo", component: UpdateTodo },
  { path: "/:catchAll(.*)", component: () => import('../views/NotFound') },
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;