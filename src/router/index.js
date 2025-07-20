import { createRouter, createWebHistory } from "vue-router"
import { useAuthStore } from "@/store/auth"

// Public Layout
import PublicLayout from "@/layouts/PublicLayout.vue"
import Home from "@/views/public/Home.vue"
import About from "@/views/public/About.vue"
import Contact from "@/views/public/Contact.vue"
import Pricing from "@/views/public/Pricing.vue"

// Auth Views
import Login from "@/views/auth/Login.vue"
import Register from "@/views/auth/Register.vue"
import ForgotPassword from "@/views/auth/ForgotPassword.vue"

// Dashboard Layout
import DashboardLayout from "@/layouts/DashboardLayout.vue"

// User Views
import Dashboard from "@/views/user/Dashboard.vue"
import Profile from "@/views/user/Profile.vue"
import Subscription from "@/views/user/Subscription.vue"
import Payments from "@/views/user/Payments.vue"
import Checkout from "@/views/user/Checkout.vue"

// Admin Views
import AdminDashboard from "@/views/admin/Dashboard.vue"
import AdminUsers from "@/views/admin/Users.vue"
import AdminPlans from "@/views/admin/Plans.vue"
import AdminSubscriptions from "@/views/admin/Subscriptions.vue"
import AdminPayments from "@/views/admin/Payments.vue"
import AdminLogs from "@/views/admin/Logs.vue"
import AdminCredKeys from "@/views/admin/CredKeys.vue"

const routes = [
  // Public routes
  {
    path: "/",
    component: PublicLayout,
    children: [
      { path: "", name: "Home", component: Home },
      { path: "about", name: "About", component: About },
      { path: "contact", name: "Contact", component: Contact },
      { path: "pricing", name: "Pricing", component: Pricing },
    ],
  },

  // Auth routes
  { path: "/login", name: "Login", component: Login },
  { path: "/register", name: "Register", component: Register },
  { path: "/forgot-password", name: "ForgotPassword", component: ForgotPassword },

  // Checkout route (requires auth)
  {
    path: "/checkout",
    name: "Checkout",
    component: Checkout,
    meta: { requiresAuth: true },
  },

  // Dashboard routes
  {
    path: "/dashboard",
    component: DashboardLayout,
    meta: { requiresAuth: true },
    children: [
      { path: "", name: "Dashboard", component: Dashboard },
      { path: "profile", name: "Profile", component: Profile },
      { path: "subscription", name: "Subscription", component: Subscription },
      { path: "payments", name: "Payments", component: Payments },
    ],
  },

  // Admin routes
  {
    path: "/admin",
    component: DashboardLayout,
    meta: { requiresAuth: true, requiresAdmin: true },
    children: [
      { path: "", name: "AdminDashboard", component: AdminDashboard },
      { path: "users", name: "AdminUsers", component: AdminUsers },
      { path: "plans", name: "AdminPlans", component: AdminPlans },
      { path: "subscriptions", name: "AdminSubscriptions", component: AdminSubscriptions },
      { path: "payments", name: "AdminPayments", component: AdminPayments },
      { path: "logs", name: "AdminLogs", component: AdminLogs },
      { path: "cred-keys", name: "AdminCredKeys", component: AdminCredKeys },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

// Navigation guard
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next("/login")
  } else if (to.meta.requiresAdmin && !authStore.user?.is_admin) {
    next("/dashboard")
  } else {
    next()
  }
})

export default router
