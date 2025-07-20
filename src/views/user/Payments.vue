<template>
  <DashboardLayout>
    <div class="payments-page container">
      <h1>Payment History</h1>
      <p class="subtitle">View your past transactions and invoices.</p>

      <div v-if="userStore.loading" class="loading-state">
        <ProgressBar mode="indeterminate" style="height: 6px"></ProgressBar>
        <p>Loading payment history...</p>
      </div>

      <div v-else-if="userStore.payments.length > 0" class="payments-table">
        <DataTable :value="userStore.payments" paginator :rows="10" :rowsPerPageOptions="[5, 10, 20]" tableStyle="min-width: 50rem">
          <Column field="transaction_id" header="Transaction ID"></Column>
          <Column field="amount" header="Amount">
            <template #body="slotProps">
              ${{ slotProps.data.amount.toFixed(2) }} {{ slotProps.data.currency }}
            </template>
          </Column>
          <Column field="status" header="Status">
            <template #body="slotProps">
              <Tag :value="slotProps.data.status" :severity="getSeverity(slotProps.data.status)" />
            </template>
          </Column>
          <Column field="payment_method" header="Method"></Column>
          <Column field="created_at" header="Date">
            <template #body="slotProps">
              {{ formatDate(slotProps.data.created_at) }}
            </template>
          </Column>
          <Column header="Invoice">
            <template #body="slotProps">
              <Button icon="pi pi-file-pdf" class="p-button-text" @click="viewInvoice(slotProps.data.id)" />
            </template>
          </Column>
        </DataTable>
      </div>

      <div v-else class="no-payments-state">
        <Card>
          <template #content>
            <p>No payment history found.</p>
            <Button label="View Plans" class="p-button-primary" @click="router.push('/pricing')" />
          </template>
        </Card>
      </div>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/store/user'
import DashboardLayout from '@/layouts/DashboardLayout.vue'

const userStore = useUserStore()
const router = useRouter()

const formatDate = (dateString) => {
  if (!dateString) return 'N/A'
  const options = { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit' }
  return new Date(dateString).toLocaleDateString(undefined, options)
}

const getSeverity = (status) => {
  switch (status) {
    case 'completed': return 'success'
    case 'pending': return 'warning'
    case 'failed': return 'danger'
    default: return 'info'
  }
}

const viewInvoice = (paymentId) => {
  console.log('View invoice for payment ID:', paymentId)
  // In a real app, you would navigate to an invoice view or download a PDF
}

onMounted(() => {
  userStore.fetchPayments()
})
</script>

<style scoped>
.payments-page {
  padding: 2rem;
}

.payments-page h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
  color: var(--text-color);
}

.payments-page .subtitle {
  font-size: 1.1rem;
  margin-bottom: 2rem;
  color: var(--text-color-secondary);
}

.loading-state {
  text-align: center;
  padding: 2rem;
}

.payments-table {
  margin-top: 2rem;
}

.no-payments-state {
  text-align: center;
  padding: 2rem;
}
</style>
