<template>
  <DashboardLayout>
    <div class="subscription-page container">
      <h1>My Subscription</h1>
      <p class="subtitle">Manage your current plan and view details.</p>

      <div v-if="userStore.loading" class="loading-state">
        <ProgressBar mode="indeterminate" style="height: 6px"></ProgressBar>
        <p>Loading subscription details...</p>
      </div>

      <div v-else-if="userStore.subscription" class="subscription-details">
        <Card>
          <template #title>Current Plan: {{ userStore.subscription.Plan?.name || 'N/A' }}</template>
          <template #content>
            <div class="p-grid p-nogutter">
              <div class="p-col-12 p-md-6">
                <p><strong>Status:</strong> <Tag :value="userStore.subscription.status" :severity="getSeverity(userStore.subscription.status)" /></p>
                <p><strong>Price:</strong> ${{ userStore.subscription.Plan?.price || 'N/A' }} / {{ userStore.subscription.Plan?.interval || 'N/A' }}</p>
                <p><strong>Start Date:</strong> {{ formatDate(userStore.subscription.start_date) }}</p>
                <p v-if="userStore.subscription.next_billing_date"><strong>Next Billing:</strong> {{ formatDate(userStore.subscription.next_billing_date) }}</p>
                <p v-if="userStore.subscription.end_date"><strong>End Date:</strong> {{ formatDate(userStore.subscription.end_date) }}</p>
                <p v-if="userStore.subscription.cancelled_at"><strong>Cancelled On:</strong> {{ formatDate(userStore.subscription.cancelled_at) }}</p>
                <p v-if="userStore.subscription.dummy_key"><strong>Your Access Key:</strong> <Chip :label="userStore.subscription.dummy_key" icon="pi pi-key" /></p>
              </div>
              <div class="p-col-12 p-md-6">
                <p><strong>Features:</strong></p>
                <ul class="features-list">
                  <li v-for="(feature, index) in userStore.subscription.Plan?.features?.split(',')" :key="index">
                    <i class="pi pi-check-circle" style="color: var(--primary-color)"></i> {{ feature.trim() }}
                  </li>
                </ul>
              </div>
            </div>
            <Divider />
            <div class="action-buttons">
              <Button label="Change Plan" class="p-button-secondary" @click="showChangePlanDialog = true" />
              <Button label="Cancel Subscription" class="p-button-danger" @click="confirmCancel($event)" />
            </div>
          </template>
        </Card>
      </div>

      <div v-else class="no-subscription-state">
        <Card>
          <template #content>
            <p>You do not have an active subscription.</p>
            <Button label="View Plans" class="p-button-primary" @click="router.push('/pricing')" />
          </template>
        </Card>
      </div>

      <!-- Change Plan Dialog -->
      <Dialog v-model:visible="showChangePlanDialog" header="Change Your Plan" :modal="true" class="p-fluid">
        <div class="p-grid p-nogutter">
          <div v-for="plan in plansStore.plans" :key="plan.id" class="p-col-12 p-md-6 p-lg-4">
            <Card class="plan-option-card" :class="{ 'selected-plan': selectedPlan && selectedPlan.id === plan.id }" @click="selectedPlan = plan">
              <template #title>{{ plan.name }}</template>
              <template #subtitle>${{ plan.price }} / {{ plan.interval }}</template>
              <template #content>
                <ul class="features-list">
                  <li v-for="(feature, index) in plan.features.split(',')" :key="index">
                    <i class="pi pi-check-circle" style="color: var(--primary-color)"></i> {{ feature.trim() }}
                  </li>
                </ul>
              </template>
            </Card>
          </div>
        </div>
        <template #footer>
          <Button label="Cancel" icon="pi pi-times" class="p-button-text" @click="showChangePlanDialog = false" />
          <Button label="Confirm Change" icon="pi pi-check" class="p-button-primary" @click="confirmChangePlan" :disabled="!selectedPlan || (selectedPlan && selectedPlan.id === userStore.subscription?.Plan?.id)" />
        </template>
      </Dialog>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from 'primevue/usetoast'
import { useConfirm } from 'primevue/useconfirm'
import { useUserStore } from '@/store/user'
import { usePlansStore } from '@/store/plans'
import DashboardLayout from '@/layouts/DashboardLayout.vue'

const userStore = useUserStore()
const plansStore = usePlansStore()
const router = useRouter()
const toast = useToast()
const confirm = useConfirm()

const showChangePlanDialog = ref(false)
const selectedPlan = ref(null)

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

const confirmCancel = (event) => {
  confirm.require({
    target: event.currentTarget,
    message: 'Are you sure you want to cancel your subscription? This action cannot be undone.',
    icon: 'pi pi-exclamation-triangle',
    acceptClass: 'p-button-danger',
    accept: () => {
      // Simulate cancellation
      toast.add({ severity: 'success', summary: 'Confirmed', detail: 'Subscription cancelled!', life: 3000 })
      // In a real app, call backend API to cancel
      userStore.fetchSubscription() // Refresh data
    },
    reject: () => {
      toast.add({ severity: 'info', summary: 'Rejected', detail: 'Cancellation aborted', life: 3000 })
    }
  })
}

const confirmChangePlan = () => {
  if (selectedPlan.value) {
    // Simulate plan change
    toast.add({ severity: 'success', summary: 'Plan Changed', detail: `Your plan has been changed to ${selectedPlan.value.name}!`, life: 3000 })
    // In a real app, call backend API to change plan
    showChangePlanDialog.value = false
    userStore.fetchSubscription() // Refresh data
  }
}

onMounted(() => {
  userStore.fetchSubscription()
  plansStore.fetchPublicPlans()
})
</script>

<style scoped>
.subscription-page {
  padding: 2rem;
}

.subscription-page h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
  color: var(--text-color);
}

.subscription-page .subtitle {
  font-size: 1.1rem;
  margin-bottom: 2rem;
  color: var(--text-color-secondary);
}

.loading-state {
  text-align: center;
  padding: 2rem;
}

.subscription-details .p-card {
  text-align: left;
}

.subscription-details .p-card-title {
  color: var(--primary-color);
}

.subscription-details p {
  margin-bottom: 0.75rem;
  color: var(--text-color-secondary);
}

.features-list {
  list-style: none;
  padding: 0;
  margin-top: 0.5rem;
}

.features-list li {
  margin-bottom: 0.5rem;
  color: var(--text-color);
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.features-list li i {
  font-size: 1.1rem;
}

.action-buttons {
  margin-top: 1.5rem;
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
}

.no-subscription-state {
  text-align: center;
  padding: 2rem;
}

.plan-option-card {
  cursor: pointer;
  border: 2px solid transparent;
  transition: all 0.2s ease-in-out;
}

.plan-option-card:hover {
  border-color: var(--primary-200);
}

.plan-option-card.selected-plan {
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px var(--primary-color);
}

.plan-option-card .p-card-title {
  color: var(--text-color);
}

.plan-option-card .p-card-subtitle {
  color: var(--text-color-secondary);
}
</style>
