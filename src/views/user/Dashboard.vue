<template>
  <DashboardLayout>
    <div class="dashboard-page container">
      <h1>Welcome, {{ userStore.profile?.first_name || authStore.user?.email }}!</h1>
      <p class="subtitle">Here's a quick overview of your account.</p>

      <div class="dashboard-grid">
        <Card class="dashboard-card">
          <template #title>Current Plan</template>
          <template #content>
            <div v-if="userStore.subscription">
              <h3>{{ userStore.subscription.Plan?.name || 'N/A' }}</h3>
              <p>Status: <Tag :value="userStore.subscription.status" :severity="getSeverity(userStore.subscription.status)" /></p>
              <p v-if="userStore.subscription.next_billing_date">Next Billing: {{ formatDate(userStore.subscription.next_billing_date) }}</p>
              <Button label="Manage Subscription" class="p-button-text" @click="router.push('/subscription')" />
            </div>
            <div v-else>
              <p>You don't have an active subscription.</p>
              <Button label="View Plans" class="p-button-primary" @click="router.push('/pricing')" />
            </div>
          </template>
        </Card>

        <Card class="dashboard-card">
          <template #title>Recent Activity</template>
          <template #content>
            <p>No recent activity to display.</p>
            <Button label="View Payments" class="p-button-text" @click="router.push('/payments')" />
          </template>
        </Card>

        <Card class="dashboard-card">
          <template #title>Profile Status</template>
          <template #content>
            <p>Complete your profile for a better experience.</p>
            <Button label="Update Profile" class="p-button-text" @click="router.push('/profile')" />
          </template>
        </Card>
      </div>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import { useUserStore } from '@/store/user'
import DashboardLayout from '@/layouts/DashboardLayout.vue'

const authStore = useAuthStore()
const userStore = useUserStore()
const router = useRouter()

onMounted(() => {
  userStore.fetchProfile()
  userStore.fetchSubscription()
})

const formatDate = (dateString) => {
  if (!dateString) return 'N/A'
  const options = { year: 'numeric', month: 'long', day: 'numeric' }
  return new Date(dateString).toLocaleDateString(undefined, options)
}

const getSeverity = (status) => {
  switch (status) {
    case 'active': return 'success'
    case 'cancelled': return 'danger'
    case 'expired': return 'warning'
    default: return 'info'
  }
}
</script>

<style scoped>
.dashboard-page {
  padding: 2rem;
}

.dashboard-page h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
  color: var(--text-color);
}

.dashboard-page .subtitle {
  font-size: 1.1rem;
  margin-bottom: 2rem;
  color: var(--text-color-secondary);
}

.dashboard-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.dashboard-card {
  text-align: left;
}

.dashboard-card .p-card-title {
  color: var(--primary-color);
}

.dashboard-card h3 {
  margin-bottom: 0.5rem;
  color: var(--text-color);
}

.dashboard-card p {
  margin-bottom: 0.75rem;
  color: var(--text-color-secondary);
}
</style>
