<template>
  <DashboardLayout>
    <div class="admin-plans-page container">
      <h1>Manage Plans</h1>
      <p class="subtitle">Create, edit, or deactivate subscription plans.</p>

      <div class="action-bar">
        <Button label="Add New Plan" icon="pi pi-plus" class="p-button-primary" @click="addNewPlan" />
      </div>

      <div v-if="adminStore.loading" class="loading-state">
        <ProgressBar mode="indeterminate" style="height: 6px"></ProgressBar>
        <p>Loading plans...</p>
      </div>

      <div v-else class="plans-table">
        <DataTable :value="adminStore.plans" tableStyle="min-width: 50rem">
          <Column field="id" header="ID"></Column>
          <Column field="name" header="Name"></Column>
          <Column field="price" header="Price">
            <template #body="slotProps">
              ${{ slotProps.data.price.toFixed(2) }}
            </template>
          </Column>
          <Column field="interval" header="Interval"></Column>
          <Column field="is_active" header="Active">
            <template #body="slotProps">
              <i :class="slotProps.data.is_active ? 'pi pi-check-circle' : 'pi pi-times-circle'"
                 :style="{ color: slotProps.data.is_active ? 'var(--green-500)' : 'var(--red-500)' }"></i>
            </template>
          </Column>
          <Column header="Actions">
            <template #body="slotProps">
              <Button icon="pi pi-pencil" class="p-button-text p-button-sm" @click="editPlan(slotProps.data)" />
              <Button icon="pi pi-trash" class="p-button-text p-button-sm p-button-danger" @click="confirmDeletePlan($event, slotProps.data)" />
            </template>
          </Column>
        </DataTable>
      </div>

      <!-- Edit/Add Plan Dialog -->
      <Dialog v-model:visible="showEditPlanDialog" :header="editMode === 'add' ? 'Add New Plan' : 'Edit Plan'" :modal="true" class="p-fluid">
        <div class="p-field">
          <label for="planName">Plan Name</label>
          <InputText id="planName" v-model="currentPlan.name" required />
        </div>
        <div class="p-field">
          <label for="planDescription">Description</label>
          <Textarea id="planDescription" v-model="currentPlan.description" rows="3" />
        </div>
        <div class="p-field">
          <label for="planPrice">Price</label>
          <InputNumber id="planPrice" v-model="currentPlan.price" mode="currency" currency="USD" locale="en-US" :min="0" required />
        </div>
        <div class="p-field">
          <label for="planInterval">Billing Interval</label>
          <Dropdown id="planInterval" v-model="currentPlan.interval" :options="['monthly', 'yearly']" required />
        </div>
        <div class="p-field">
          <label for="planFeatures">Features (comma-separated)</label>
          <Textarea id="planFeatures" v-model="currentPlan.features" rows="3" />
        </div>
        <div class="p-field-checkbox">
          <Checkbox id="planIsActive" v-model="currentPlan.is_active" :binary="true" />
          <label for="planIsActive">Active Plan</label>
        </div>
        <template #footer>
          <Button label="Cancel" icon="pi pi-times" class="p-button-text" @click="showEditPlanDialog = false" />
          <Button label="Save" icon="pi pi-check" class="p-button-primary" @click="savePlan" />
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

const showEditPlanDialog = ref(false)
const currentPlan = ref(null)
const editMode = ref('add') // 'add' or 'edit'

onMounted(() => {
  adminStore.fetchPlans()
})

const addNewPlan = () => {
  currentPlan.value = {
    name: '',
    description: '',
    price: 0,
    currency: 'USD',
    interval: 'monthly',
    features: '',
    is_active: true,
  }
  editMode.value = 'add'
  showEditPlanDialog.value = true
}

const editPlan = (plan) => {
  currentPlan.value = { ...plan }
  editMode.value = 'edit'
  showEditPlanDialog.value = true
}

const savePlan = async () => {
  // In a real app, you'd call an API to create/update plan
  console.log('Saving plan:', currentPlan.value)
  toast.add({ severity: 'success', summary: 'Success', detail: 'Plan saved successfully!', life: 3000 })
  showEditPlanDialog.value = false
  await adminStore.fetchPlans() // Refresh data
}

const confirmDeletePlan = (event, plan) => {
  confirm.require({
    target: event.currentTarget,
    message: `Are you sure you want to delete plan "${plan.name}"?`,
    icon: 'pi pi-exclamation-triangle',
    acceptClass: 'p-button-danger',
    accept: () => {
      // In a real app, call API to delete plan
      toast.add({ severity: 'success', summary: 'Confirmed', detail: 'Plan deleted!', life: 3000 })
      adminStore.fetchPlans() // Refresh data
    },
    reject: () => {
      toast.add({ severity: 'info', summary: 'Rejected', detail: 'Deletion aborted', life: 3000 })
    }
  })
}
</script>

<style scoped>
.admin-plans-page {
  padding: 2rem;
}

.admin-plans-page h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
  color: var(--text-color);
}

.admin-plans-page .subtitle {
  font-size: 1.1rem;
  margin-bottom: 2rem;
  color: var(--text-color-secondary);
}

.action-bar {
  margin-bottom: 2rem;
  text-align: right;
}

.loading-state {
  text-align: center;
  padding: 2rem;
}

.plans-table {
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

.p-field-checkbox {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 1.5rem;
}
</style>
