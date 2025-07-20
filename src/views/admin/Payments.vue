<template>
  <DashboardLayout>
    <div class="admin-payments-page container">
      <h1>All Payments</h1>
      <p class="subtitle">View all payment transactions on the platform.</p>

      <div v-if="adminStore.loading" class="loading-state">
        <ProgressBar mode="indeterminate" style="height: 6px"></ProgressBar>
        <p>Loading payments...</p>
      </div>

      <div v-else class="payments-table">
        <DataTable :value="adminStore.payments" paginator :rows="adminStore.pagination.limit" :totalRecords="adminStore.pagination.total"
                   @page="onPageChange" tableStyle="min-width: 50rem">
          <Column field="id" header="ID"></Column>
          <Column field="User.email" header="User Email"></Column>
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
          <Column field="transaction_id" header="Transaction ID"></Column>
          <Column field="created_at" header="Date">
            <template #body="slotProps">
              {{ formatDate(slotProps.data.created_at) }}
            </template>
          </Column>
          <Column header="Actions">
            <template #body="slotProps">
              <Button icon="pi pi-info-circle" class="p-button-text p-button-sm" @click="viewPaymentDetails(slotProps.data)" />
            </template>
          </Column>
        </DataTable>
      </div>

      <!-- Payment Details Dialog -->
      <Dialog v-model:visible="showPaymentDetailsDialog" header="Payment Details" :modal="true" class="p-fluid">
        <div v-if="currentPayment">
          <p><strong>Payment ID:</strong> {{ currentPayment.id }}</p>
          <p><strong>User Email:</strong> {{ currentPayment.User?.email }}</p>
          <p><strong>Amount:</strong> ${{ currentPayment.amount?.toFixed(2) }} {{ currentPayment.currency }}</p>
          <p><strong>Status:</strong> <Tag :value="currentPayment.status" :severity="getSeverity(currentPayment.status)" /></p>
          <p><strong>Payment Method:</strong> {{ currentPayment.payment_method }}</p>
          <p><strong>Transaction ID:</strong> {{ currentPayment.transaction_id }}</p>
          <p><strong>Date:</strong> {{ formatDate(currentPayment.created_at) }}</p>
        </div>
        <template #footer>
          <Button label="Close" icon="pi pi-times" class="p-button-text" @click="showPaymentDetailsDialog = false" />
        </template>
      </Dialog>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useAdminStore } from '@/store/admin'
import DashboardLayout from '@/layouts/DashboardLayout.vue'

const adminStore = useAdminStore()

const showPaymentDetailsDialog = ref(false)
const currentPayment = ref(null)

// Ensure all hooks are called at the top level
onMounted(() => {
  adminStore.fetchPayments()
})

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

const onPageChange = (event) => {
  adminStore.fetchPayments(event.page + 1, event.rows)
}

const viewPaymentDetails = (payment) => {
  currentPayment.value = payment
  showPaymentDetailsDialog.value = true
}
</script>

<style scoped>
.admin-payments-page {
  padding: 2rem;
}

.admin-payments-page h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
  color: var(--text-color);
}

.admin-payments-page .subtitle {
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

.p-field {
  margin-bottom: 1.5rem;
}

.p-field label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: bold;
  color: var(--text-color);
}
</style>
