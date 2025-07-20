<template>
  <div class="dashboard-layout">
    <Sidebar 
      v-model:visible="sidebarVisible" 
      :modal="isMobile"
      :dismissable="!isMobile"
      :showCloseIcon="isMobile"
      :pt="{
        root: 'dashboard-sidebar',
        content: 'flex-1 overflow-y-auto p-0'
      }"
    >
      <div class="sidebar-header">
        <router-link to="/" class="brand-link">
          <img src="/images/orinix-logo.svg" alt="Orinix Logo" />
          <h3>{{ appName }}</h3>
        </router-link>
      </div>
      
      <Menu :model="menuItems" class="sidebar-menu" />
    </Sidebar>
    
    <div class="dashboard-main" :class="{ 'sidebar-collapsed': !sidebarVisible && !isMobile }">
      <header class="dashboard-header">
        <div class="header-left">
          <Button 
            icon="pi pi-bars" 
            @click="toggleSidebar"
            text
            rounded
            class="sidebar-toggle-button"
          />
          <h2>{{ pageTitle }}</h2>
        </div>
        
        <div class="header-right">
          <Dropdown 
            v-model="selectedUser" 
            :options="userMenuItems" 
            option-label="label"
            @change="handleUserMenuAction"
            :pt="{
              root: 'user-dropdown-root',
              trigger: 'user-dropdown-trigger',
              label: 'user-dropdown-label',
              panel: 'user-dropdown-panel'
            }"
          >
            <template #value>
              <div class="user-info">
                <Avatar 
                  :label="userInitials" 
                  shape="circle"
                  class="user-avatar"
                />
                <span class="user-email">{{ authStore.user?.email }}</span>
                <i class="pi pi-chevron-down user-dropdown-icon"></i>
              </div>
            </template>
            <template #option="slotProps">
              <div class="user-menu-option">
                <i :class="slotProps.option.icon"></i>
                <span>{{ slotProps.option.label }}</span>
              </div>
            </template>
          </Dropdown>
        </div>
      </header>
      
      <main class="dashboard-content">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const authStore = useAuthStore()
const route = useRoute()
const router = useRouter()

const appName = import.meta.env.VITE_APP_NAME || 'Orinix'
const sidebarVisible = ref(true)
const selectedUser = ref(null)
const isMobile = ref(false)

const pageTitle = computed(() => {
  return route.meta.title || route.name || 'Dashboard'
})

const userInitials = computed(() => {
  const email = authStore.user?.email || ''
  return email.charAt(0).toUpperCase()
})

const menuItems = computed(() => {
  const baseItems = [
    {
      label: 'Dashboard',
      icon: 'pi pi-home',
      command: () => navigateTo('/dashboard')
    },
    {
      label: 'Subscription',
      icon: 'pi pi-credit-card',
      command: () => navigateTo('/subscription')
    },
    {
      label: 'Payments',
      icon: 'pi pi-money-bill',
      command: () => navigateTo('/payments')
    },
    {
      label: 'Profile',
      icon: 'pi pi-user',
      command: () => navigateTo('/profile')
    }
  ]

  if (authStore.isAdmin) {
    return [
      {
        label: 'Admin Dashboard',
        icon: 'pi pi-chart-line',
        command: () => navigateTo('/admin')
      },
      {
        label: 'Users',
        icon: 'pi pi-users',
        command: () => navigateTo('/admin/users')
      },
      {
        label: 'Subscriptions',
        icon: 'pi pi-list',
        command: () => navigateTo('/admin/subscriptions')
      },
      {
        label: 'Payments',
        icon: 'pi pi-dollar',
        command: () => navigateTo('/admin/payments')
      },
      {
        label: 'Plans',
        icon: 'pi pi-cog',
        command: () => navigateTo('/admin/plans')
      },
      {
        label: 'Logs',
        icon: 'pi pi-file',
        command: () => navigateTo('/admin/logs')
      }
    ]
  }

  return baseItems
})

const userMenuItems = [
  { label: 'Profile', value: 'profile', icon: 'pi pi-user' },
  { label: 'Settings', value: 'settings', icon: 'pi pi-cog' },
  { label: 'Logout', value: 'logout', icon: 'pi pi-sign-out' }
]

const toggleSidebar = () => {
  sidebarVisible.value = !sidebarVisible.value
}

const navigateTo = (path) => {
  router.push(path)
  if (isMobile.value) {
    sidebarVisible.value = false // Close sidebar on mobile after navigation
  }
}

const handleUserMenuAction = (event) => {
  const action = event.value?.value
  if (action === 'logout') {
    authStore.logout()
  } else if (action === 'profile') {
    navigateTo('/profile')
  } else if (action === 'settings') {
    // Handle settings navigation or dialog
    console.log('Navigate to settings or open settings dialog')
  }
  selectedUser.value = null
}

const checkMobile = () => {
  isMobile.value = window.innerWidth < 768
  if (isMobile.value) {
    sidebarVisible.value = false // Hide sidebar by default on mobile
  } else {
    sidebarVisible.value = true // Show sidebar by default on desktop
  }
}

onMounted(() => {
  checkMobile()
  window.addEventListener('resize', checkMobile)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', checkMobile)
})
</script>

<style scoped>
.dashboard-layout {
  display: flex;
  min-height: 100vh;
  background: var(--orinix-dark-blue); /* Main background for the entire dashboard area */
}

/* Sidebar Styling */
.dashboard-sidebar {
  width: 280px;
  background: var(--orinix-card-bg); /* Darker background for sidebar */
  border-right: 1px solid var(--orinix-border-color);
  box-shadow: 2px 0 10px rgba(0, 0, 0, 0.3);
  transition: transform 0.3s ease-in-out;
  flex-shrink: 0; /* Prevent sidebar from shrinking */
}

.sidebar-header {
  padding: 1.5rem;
  border-bottom: 1px solid var(--orinix-border-color);
  display: flex;
  align-items: center;
  justify-content: center;
}

.sidebar-header .brand-link {
  font-size: 1.8rem;
  font-weight: bold;
  text-decoration: none;
  color: var(--text-color);
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.sidebar-header .brand-link img {
  height: 30px;
}

.sidebar-header h3 {
  margin: 0;
  color: var(--text-color);
}

.sidebar-menu {
  background: transparent; /* Remove default menu background */
  border: none;
  width: 100%;
  padding: 1rem 0;
}

/* PrimeVue Menu Overrides for Sidebar */
.sidebar-menu :deep(.p-menuitem-link) {
  padding: 0.75rem 1.5rem;
  color: var(--text-color-secondary);
  transition: background-color 0.2s, color 0.2s;
  border-radius: 0; /* Remove default border-radius */
}

.sidebar-menu :deep(.p-menuitem-link:hover) {
  background: rgba(0, 123, 255, 0.1); /* Light blue hover */
  color: var(--text-color);
}

.sidebar-menu :deep(.p-menuitem-link.router-link-active) {
  background: var(--orinix-accent-blue); /* Active item background */
  color: white;
  font-weight: 600;
}

.sidebar-menu :deep(.p-menuitem-link.router-link-active .p-menuitem-icon),
.sidebar-menu :deep(.p-menuitem-link.router-link-active .p-menuitem-text) {
  color: white; /* Ensure icon and text are white for active */
}

.sidebar-menu :deep(.p-menuitem-icon) {
  color: var(--text-color-secondary);
  margin-right: 0.75rem;
}

/* Main Content Area */
.dashboard-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  transition: margin-left 0.3s ease-in-out;
  margin-left: 280px; /* Default margin for open sidebar */
}

.dashboard-main.sidebar-collapsed {
  margin-left: 0; /* No margin when sidebar is collapsed on desktop */
}

/* Header Styling */
.dashboard-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem 2rem;
  background: var(--orinix-card-bg); /* Header background */
  border-bottom: 1px solid var(--orinix-border-color);
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
  flex-shrink: 0; /* Prevent header from shrinking */
}

.header-left {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.sidebar-toggle-button {
  color: var(--text-color-secondary);
  background: transparent;
  border: none;
}

.sidebar-toggle-button:hover {
  background: rgba(255, 255, 255, 0.1);
}

.header-left h2 {
  margin: 0;
  font-size: 1.5rem;
  color: var(--text-color);
}

/* User Dropdown Styling */
.user-dropdown-root {
  background: var(--orinix-dark-blue);
  border: 1px solid var(--orinix-border-color);
  border-radius: 50px; /* Pill shape */
  padding: 0.5rem 1rem;
  cursor: pointer;
  transition: background-color 0.2s;
}

.user-dropdown-root:hover {
  background: rgba(0, 123, 255, 0.1);
}

.user-dropdown-label {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  color: var(--text-color);
}

.user-avatar {
  background: var(--orinix-accent-blue);
  color: white;
  font-weight: bold;
  width: 36px;
  height: 36px;
  font-size: 1rem;
}

.user-email {
  color: var(--text-color);
  font-weight: 500;
}

.user-dropdown-icon {
  color: var(--text-color-secondary);
  font-size: 0.8rem;
}

.user-dropdown-panel {
  background: var(--orinix-card-bg);
  border: 1px solid var(--orinix-border-color);
  border-radius: var(--border-radius);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.4);
  padding: 0.5rem 0;
}

.user-menu-option {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem 1.25rem;
  color: var(--text-color-secondary);
  cursor: pointer;
  transition: background-color 0.2s, color 0.2s;
}

.user-menu-option:hover {
  background: rgba(0, 123, 255, 0.1);
  color: var(--text-color);
}

.user-menu-option i {
  color: var(--text-color-secondary);
}

/* Content Area */
.dashboard-content {
  flex: 1;
  padding: 2rem;
  background: var(--orinix-dark-blue); /* Content area background */
  overflow-y: auto; /* Enable scrolling for content */
}

/* Responsive Adjustments */
@media (max-width: 768px) {
  .dashboard-sidebar {
    position: fixed;
    height: 100vh;
    z-index: 1000;
    transform: translateX(-100%); /* Hidden by default on mobile */
  }

  .p-sidebar-active .dashboard-sidebar {
    transform: translateX(0); /* Show when active */
  }

  .dashboard-main {
    margin-left: 0; /* No margin on mobile */
  }

  .sidebar-toggle-button {
    display: block; /* Always show toggle button on mobile */
  }

  .dashboard-header {
    padding: 1rem;
  }

  .header-left h2 {
    font-size: 1.2rem;
  }

  .dashboard-content {
    padding: 1rem;
  }
}
</style>
