<template>
  <DashboardLayout>
    <div class="admin-cred-keys-page container">
      <h1>Manage Credential Keys</h1>
      <p class="subtitle">Create and manage activation, trial, and partner keys.</p>

      <div class="action-bar">
        <Button label="Add New Key" icon="pi pi-plus" class="p-button-primary" @click="addNewKey" />
      </div>

      <div v-if="adminStore.loading" class="loading-state">
        <ProgressBar mode="indeterminate" style="height: 6px"></ProgressBar>
        <p>Loading credential keys...</p>
      </div>

      <div v-else class="cred-keys-table">
        <DataTable :value="adminStore.credKeys" tableStyle="min-width: 50rem">
          <Column field="id" header="ID"></Column>
          <Column field="key_value" header="Key Value"></Column>
          <Column field="key_type" header="Key Type"></Column>
          <Column field="is_active" header="Active">
            <template #body="slotProps">
              <i :class="slotProps.data.is_active ? 'pi pi-check-circle' : 'pi pi-times-circle'"
                 :style="{ color: slotProps.data.is_active ? 'var(--green-500)' : 'var(--red-500)' }"></i>
            </template>
          </Column>
          <Column field="created_at" header="Created At">
            <template #body="slotProps">
              {{ formatDate(slotProps.data.created_at) }}
            </template>
          </Column>
          <Column header="Actions">
            <template #body="slotProps">
              <Button icon="pi pi-pencil" class="p-button-text p-button-sm" @click="editKey(slotProps.data)" />
            </template>
          </Column>
        </DataTable>
      </div>

      <!-- Edit/Add Credential Key Dialog -->
      <Dialog v-model:visible="showEditKeyDialog" :header="editMode === 'add' ? 'Add New Credential Key' : 'Edit Credential Key'" :modal="true" class="p-fluid">
        <div class="p-field">
          <label for="keyValue">Key Value</label>
          <InputText id="keyValue" v-model="currentKey.key_value" required />
        </div>
        <div class="p-field">
          <label for="keyType">Key Type</label>
          <Dropdown id="keyType" v-model="currentKey.key_type" :options="['activation', 'trial', 'partner']" required />
        </div>
        <div class="p-field-checkbox">
          <Checkbox id="isActive" v-model="currentKey.is_active" :binary="true" />
          <label for="isActive">Active</label>
        </div>
        <template #footer>
          <Button label="Cancel" icon="pi pi-times" class="p-button-text" @click="showEditKeyDialog = false" />
          <Button label="Save" icon="pi pi-check" class="p-button-primary" @click="saveKey" />
        </template>
      </Dialog>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useToast } from 'primevue/usetoast'
import { useAdminStore } from '@/store/admin'
import DashboardLayout from '@/layouts/DashboardLayout.vue'

const adminStore = useAdminStore()
const toast = useToast()

const showEditKeyDialog = ref(false)
const currentKey = ref(null)
const editMode = ref('add') // 'add' or 'edit'

onMounted(() => {
  adminStore.fetchCredKeys()
})

const formatDate = (dateString) => {
  if (!dateString) return 'N/A'
  const options = { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit' }
  return new Date(dateString).toLocaleDateString(undefined, options)
}

const addNewKey = () => {
  currentKey.value = {
    key_value: '',
    key_type: 'activation',
    is_active: true,
  }
  editMode.value = 'add'
  showEditKeyDialog.value = true
}

const editKey = (key) => {
  currentKey.value = { ...key }
  editMode.value = 'edit'
  showEditKeyDialog.value = true
}

const saveKey = async () => {
  try {
    if (editMode.value === 'add') {
      await adminStore.createCredKey(currentKey.value)
      toast.add({ severity: 'success', summary: 'Success', detail: 'Credential key created successfully!', life: 3000 })
    } else {
      await adminStore.updateCredKey(currentKey.value.id, currentKey.value)
      toast.add({ severity: 'success', summary: 'Success', detail: 'Credential key updated successfully!', life: 3000 })
    }
    showEditKeyDialog.value = false
    await adminStore.fetchCredKeys() // Refresh data
  } catch (error) {
    toast.add({ severity: 'error', summary: 'Operation Failed', detail: error.message || 'An error occurred', life: 3000 })
  }
}
</script>

<style scoped>
.admin-cred-keys-page {
  padding: 2rem;
}

.admin-cred-keys-page h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
  color: var(--text-color);
}

.admin-cred-keys-page .subtitle {
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

.cred-keys-table {
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
