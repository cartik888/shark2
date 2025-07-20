<template>
  <DashboardLayout>
    <div class="admin-dashboard-page container">
      <h1>Admin Dashboard</h1>
      <p class="subtitle">Overview of your SaaS platform.</p>

      <div class="dashboard-grid">
        <Card class="dashboard-card">
          <template #title>Total Users</template>
          <template #content>
            <div class="metric-value">{{ adminStore.pagination.total }}</div>
            <Button label="View All Users" class="p-button-text" @click="router.push('/admin/users')" />
          </template>
        </Card>

        <Card class="dashboard-card">
          <template #title>Active Subscriptions</template>
          <template #content>
            <div class="metric-value">{{ activeSubscriptionsCount }}</div>
            <Button label="View Subscriptions" class="p-button-text" @click="router.push('/admin/subscriptions')" />
          </template>
        </Card>

        <Card class="dashboard-card">
          <template #title>Total Revenue (Mock)</template>
          <template #content>
            <div class="metric-value">${{ totalRevenue.toFixed(2) }}</div>
            <Button label="View Payments" class="p-button-text" @click="router.push('/admin/payments')" />
          </template>
        </Card>
      </div>

      <div class="recent-activity-section">
        <h2>Recent User Activity</h2>
        <DataTable :value="adminStore.users.slice(0, 5)" :loading="adminStore.loading" tableStyle="min-width: 50rem">
          <Column field="id" header="ID"></Column>
          <Column field="email" header="Email"></Column>
          <Column field="role" header="Role"></Column>
          <Column field="created_at" header="Joined Date">
            <template #body="slotProps">
              {{ formatDate(slotProps.data.created_at) }}
            </template>
          </Column>
        </DataTable>
        <div class="flex justify-content-end mt-3">
          <Button label="View All Users" class="p-button-text" @click="router.push('/admin/users')" />
        </div>
      </div>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAdminStore } from '@/store/admin'
import DashboardLayout from '@/layouts/DashboardLayout.vue'

const adminStore = useAdminStore()
const router = useRouter()

const activeSubscriptionsCount = computed(() => {
  return adminStore.subscriptions.filter(sub => sub.status === 'active').length
})

const totalRevenue = computed(() => {
  return adminStore.payments.reduce((sum, payment) => sum + payment.amount, 0)
})

const formatDate = (dateString) => {
  if (!dateString) return 'N/A'
  const options = { year: 'numeric', month: 'long', day: 'numeric' }
  return new Date(dateString).toLocaleDateString(undefined, options)
}

onMounted(() => {
  adminStore.fetchUsers()
  adminStore.fetchSubscriptions()
  adminStore.fetchPayments()
})
</script>

<style scoped>
.admin-dashboard-page {
  padding: 2rem;
}

.admin-dashboard-page h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
  color: var(--text-color);
}

.admin-dashboard-page .subtitle {
  font-size: 1.1rem;
  margin-bottom: 2rem;
  color: var(--text-color-secondary);
}

.dashboard-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 2rem;
  margin-bottom: 3rem;
}

.dashboard-card {
  text-align: left;
}

.dashboard-card .p-card-title {
  color: var(--primary-color);
}

.metric-value {
  font-size: 3rem;
  font-weight: bold;
  margin-bottom: 1rem;
  color: var(--text-color);
}

.recent-activity-section {
  margin-top: 3rem;
}

.recent-activity-section h2 {
  font-size: 2rem;
  margin-bottom: 1.5rem;
  color: var(--text-color);
}
</style>
