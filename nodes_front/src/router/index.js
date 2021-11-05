import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from "@/views/Home.vue";
import ArticleList from "@/views/ArticleList.vue";
import Login from "@/views/Login.vue";
import Dashboard from "@/views/Dashboard.vue";
import Error404 from "@/views/Error404.vue";

Vue.use(VueRouter)


const routes = [
  {
    path: "/",
    name: "home",
    component: Home
  },
  {
    path: "/articles",
    name: "articles",
    component: ArticleList
  },
  {
    path: "/login",
    name: "login",
    component: Login
  },
  {
    path: "/dashboard",
    name: "dashboard",
    component: Dashboard
  },
  {
    path: "/:catchAll(.*)",
    name: "Error",
    component: Error404
  }
];

const router = new VueRouter({
    mode: "history",
    routes
});

router.beforeEach((to, from, next) => {
  // redirect to login page if not logged in and trying to access a restricted page
  const publicPages = ['/', '/login'];
  const authRequired = !publicPages.includes(to.path);
  const loggedIn = localStorage.getItem('user');

  if (authRequired && !loggedIn) {
    return next({ 
      path: '/login', 
      query: { returnUrl: to.path } 
    });
  }

  if (loggedIn && to.path=='/login') {
    return next({ 
      path: '/dashboard', 
    });
  }

  next();
})


export default router;