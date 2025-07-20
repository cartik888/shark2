<template>
  <DashboardLayout>
    <div class="admin-logs-page container">
      <h1>System Logs</h1>
      <p class="subtitle">View recent system activities and errors.</p>

      <div v-if="loading" class="loading-state">
        <ProgressBar mode="indeterminate" style="height: 6px"></ProgressBar>
        <p>Loading logs...</p>
      </div>

      <div v-else class="logs-table">
        <DataTable :value="logs" paginator :rows="10" :rowsPerPageOptions="[5, 10, 20]" tableStyle="min-width: 50rem">
          <Column field="timestamp" header="Timestamp">
            <template #body="slotProps">
              {{ formatDate(slotProps.data.timestamp) }}
            </template>
          </Column>
          <Column field="level" header="Level">
            <template #body="slotProps">
              <Tag :value="slotProps.data.level" :severity="getSeverity(slotProps.data.level)" />
            </template>
          </Column>
          <Column field="message" header="Message"></Column>
          <Column field="source" header="Source"></Column>
        </DataTable>
      </div>

      <div v-else class="no-logs-state">
        <Card>
          <template #content>
            <p>No system logs available.</p>
          </template>
        </Card>
      </div>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import DashboardLayout from '@/layouts/DashboardLayout.vue'

const logs = ref([])
const loading = ref(true)

onMounted(() => {
  // Simulate fetching logs
  setTimeout(() => {
    logs.value = [
      { id: 1, timestamp: '2024-07-13T10:00:00Z', level: 'INFO', message: 'User test@example.com logged in', source: 'AuthService' },
      { id: 2, timestamp: '2024-07-13T09:55:00Z', level: 'WARN', message: 'Failed payment attempt for user 123', source: 'PaymentGateway' },
      { id: 3, timestamp: '2024-07-13T09:50:00Z', level: 'ERROR', message: 'Database connection lost', source: 'DBService' },
      { id: 4, timestamp: '2024-07-13T09:45:00Z', level: 'INFO', message: 'New subscription created for plan Basic', source: 'SubscriptionService' },
      { id: 5, timestamp: '2024-07-13T09:40:00Z', level: 'DEBUG', message: 'API request: GET /api/me', source: 'GinRouter' },
    ]
    loading.value = false
  }, 1500)
})

const formatDate = (dateString) => {
  if (!dateString) return 'N/A'
  const options = { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit', second: '2-digit' }
  return new Date(dateString).toLocaleDateString(undefined, options)
}

const getSeverity = (level) => {
  switch (level) {
    case 'INFO': return 'info'
    case 'WARN': return 'warning'
    case 'ERROR': return 'danger'
    case 'DEBUG': return 'secondary'
    default: return 'info'
  }
}
</script>

<style scoped>
.admin-logs-page {
  padding: 2rem;
}

.admin-logs-page h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
  color: var(--text-color);
}

.admin-logs-page .subtitle {
  font-size: 1.1rem;
  margin-bottom: 2rem;
  color: var(--text-color-secondary);
}

.loading-state {
  text-align: center;
  padding: 2rem;
}

.logs-table {
  margin-top: 2rem;
}

.no-logs-state {
  text-align: center;
  padding: 2rem;
}
</style>
