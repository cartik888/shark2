<template>
  <DashboardLayout>
    <div class="admin-subscriptions-page container">
      <h1>All Subscriptions</h1>
      <p class="subtitle">View and manage all platform subscriptions.</p>

      <div v-if="adminStore.loading" class="loading-state">
        <ProgressBar mode="indeterminate" style="height: 6px"></ProgressBar>
        <p>Loading subscriptions...</p>
      </div>

      <div v-else class="subscriptions-table">
        <DataTable :value="adminStore.subscriptions" paginator :rows="adminStore.pagination.limit" :totalRecords="adminStore.pagination.total"
                   @page="onPageChange" tableStyle="min-width: 50rem">
          <Column field="id" header="ID"></Column>
          <Column field="User.email" header="User Email"></Column>
          <Column field="Plan.name" header="Plan"></Column>
          <Column field="status" header="Status">
            <template #body="slotProps">
              <Tag :value="slotProps.data.status" :severity="getSeverity(slotProps.data.status)" />
            </template>
          </Column>
          <Column field="start_date" header="Start Date">
            <template #body="slotProps">
              {{ formatDate(slotProps.data.start_date) }}
            </template>
          </Column>
          <Column field="next_billing_date" header="Next Billing">
            <template #body="slotProps">
              {{ formatDate(slotProps.data.next_billing_date) }}
            </template>
          </Column>
          <Column header="Actions">
            <template #body="slotProps">
              <Button icon="pi pi-pencil" class="p-button-text p-button-sm" @click="editSubscription(slotProps.data)" />
              <Button icon="pi pi-times" class="p-button-text p-button-sm p-button-danger" @click="confirmCancelSubscription($event, slotProps.data)" />
            </template>
          </Column>
        </DataTable>
      </div>

      <!-- Edit Subscription Dialog -->
      <Dialog v-model:visible="showEditSubscriptionDialog" header="Edit Subscription" :modal="true" class="p-fluid">
        <div class="p-field">
          <label for="editSubUserEmail">User Email</label>
          <InputText id="editSubUserEmail" v-model="currentSubscription.User.email" disabled />
        </div>
        <div class="p-field">
          <label for="editSubPlan">Plan</label>
          <Dropdown id="editSubPlan" v-model="currentSubscription.Plan.id" :options="adminStore.plans" optionLabel="name" optionValue="id" />
        </div>
        <div class="p-field">
          <label for="editSubStatus">Status</label>
          <Dropdown id="editSubStatus" v-model="currentSubscription.status" :options="['active', 'cancelled', 'expired']" />
        </div>
        <div class="p-field">
          <label for="editSubNextBilling">Next Billing Date</label>
          <Calendar id="editSubNextBilling" v-model="currentSubscription.next_billing_date" dateFormat="yy-mm-dd" />
        </div>
        <template #footer>
          <Button label="Cancel" icon="pi pi-times" class="p-button-text" @click="showEditSubscriptionDialog = false" />
          <Button label="Save" icon="pi pi-check" class="p-button-primary" @click="saveSubscription" />
        </template>
      </Dialog>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useToast } from 'primevue/usetoast'
import { useConfirm } from 'primevue/useconfirm'
import { useAdminStore } from '@/store/admin'
import DashboardLayout from '@/layouts/DashboardLayout.vue'

const adminStore = useAdminStore()
const toast = useToast()
const confirm = useConfirm()

const showEditSubscriptionDialog = ref(false)
const currentSubscription = ref(null)

onMounted(() => {
  adminStore.fetchSubscriptions()
  adminStore.fetchPlans() // Needed for dropdown
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

const onPageChange = (event) => {
  adminStore.fetchSubscriptions(event.page + 1, event.rows)
}

const editSubscription = (subscription) => {
  currentSubscription.value = { ...subscription, User: { ...subscription.User }, Plan: { ...subscription.Plan } }
  showEditSubscriptionDialog.value = true
}

const saveSubscription = async () => {
  // In a real app, you'd call an API to update subscription
  console.log('Saving subscription:', currentSubscription.value)
  toast.add({ severity: 'success', summary: 'Success', detail: 'Subscription saved successfully!', life: 3000 })
  showEditSubscriptionDialog.value = false
  await adminStore.fetchSubscriptions(adminStore.pagination.page, adminStore.pagination.limit) // Refresh data
}

const confirmCancelSubscription = (event, subscription) => {
  confirm.require({
    target: event.currentTarget,
    message: `Are you sure you want to cancel subscription for ${subscription.User.email}?`,
    icon: 'pi pi-exclamation-triangle',
    acceptClass: 'p-button-danger',
    accept: () => {
      // In a real app, call API to cancel subscription
      toast.add({ severity: 'success', summary: 'Confirmed', detail: 'Subscription cancelled!', life: 3000 })
      adminStore.fetchSubscriptions(adminStore.pagination.page, adminStore.pagination.limit) // Refresh data
    },
    reject: () => {
      toast.add({ severity: 'info', summary: 'Rejected', detail: 'Cancellation aborted', life: 3000 })
    }
  })
}
</script>

<style scoped>
.admin-subscriptions-page {
  padding: 2rem;
}

.admin-subscriptions-page h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
  color: var(--text-color);
}

.admin-subscriptions-page .subtitle {
  font-size: 1.1rem;
  margin-bottom: 2rem;
  color: var(--text-color-secondary);
}

.loading-state {
  text-align: center;
  padding: 2rem;
}

.subscriptions-table {
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
