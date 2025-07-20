<template>
  <DashboardLayout>
    <div class="admin-users-page container">
      <h1>Manage Users</h1>
      <p class="subtitle">View and manage all registered users.</p>

      <div v-if="adminStore.loading" class="loading-state">
        <ProgressBar mode="indeterminate" style="height: 6px"></ProgressBar>
        <p>Loading users...</p>
      </div>

      <div v-else class="users-table">
        <DataTable :value="adminStore.users" paginator :rows="adminStore.pagination.limit" :totalRecords="adminStore.pagination.total"
                   @page="onPageChange" tableStyle="min-width: 50rem">
          <Column field="id" header="ID"></Column>
          <Column field="email" header="Email"></Column>
          <Column field="role" header="Role"></Column>
          <Column field="is_active" header="Active">
            <template #body="slotProps">
              <i :class="slotProps.data.is_active ? 'pi pi-check-circle' : 'pi pi-times-circle'"
                 :style="{ color: slotProps.data.is_active ? 'var(--green-500)' : 'var(--red-500)' }"></i>
            </template>
          </Column>
          <Column field="is_blocked" header="Blocked">
            <template #body="slotProps">
              <i :class="slotProps.data.is_blocked ? 'pi pi-lock' : 'pi pi-unlock'"
                 :style="{ color: slotProps.data.is_blocked ? 'var(--red-500)' : 'var(--green-500)' }"></i>
            </template>
          </Column>
          <Column field="created_at" header="Joined Date">
            <template #body="slotProps">
              {{ formatDate(slotProps.data.created_at) }}
            </template>
          </Column>
          <Column header="Actions">
            <template #body="slotProps">
              <Button icon="pi pi-info-circle" class="p-button-text p-button-sm" @click="viewUserDetails(slotProps.data)" v-tooltip.top="'View Details'" />
              <Button 
                :icon="slotProps.data.is_blocked ? 'pi pi-unlock' : 'pi pi-lock'" 
                :class="slotProps.data.is_blocked ? 'p-button-text p-button-sm p-button-success' : 'p-button-text p-button-sm p-button-warning'" 
                @click="confirmBlockUser($event, slotProps.data)" 
                v-tooltip.top="slotProps.data.is_blocked ? 'Unblock User' : 'Block User'"
              />
              <Button icon="pi pi-pencil" class="p-button-text p-button-sm" @click="editUser(slotProps.data)" v-tooltip.top="'Edit User'" />
              <Button icon="pi pi-trash" class="p-button-text p-button-sm p-button-danger" @click="confirmDeleteUser($event, slotProps.data)" v-tooltip.top="'Delete User'" />
            </template>
          </Column>
        </DataTable>
      </div>

      <!-- Edit User Dialog -->
      <Dialog v-model:visible="showEditUserDialog" :header="editMode === 'add' ? 'Add New User' : 'Edit User'" :modal="true" class="p-fluid">
        <div class="p-field">
          <label for="editEmail">Email</label>
          <InputText id="editEmail" v-model="currentUser.email" :disabled="editMode === 'edit'" />
        </div>
        <div class="p-field">
          <label for="editRole">Role</label>
          <Dropdown id="editRole" v-model="currentUser.role" :options="['user', 'admin']" />
        </div>
        <div class="p-field">
          <label for="editIsActive">Active</label>
          <Checkbox id="editIsActive" v-model="currentUser.is_active" :binary="true" />
        </div>
        <template #footer>
          <Button label="Cancel" icon="pi pi-times" class="p-button-text" @click="showEditUserDialog = false" />
          <Button label="Save" icon="pi pi-check" class="p-button-primary" @click="saveUser" />
        </template>
      </Dialog>

      <!-- User Details Dialog -->
      <Dialog v-model:visible="showUserDetailsDialog" header="User Details" :modal="true" class="p-fluid user-details-dialog">
        <div v-if="adminStore.userDetails">
          <h3>User Information</h3>
          <p><strong>ID:</strong> {{ adminStore.userDetails.id }}</p>
          <p><strong>Email:</strong> {{ adminStore.userDetails.email }}</p>
          <p><strong>Role:</strong> {{ adminStore.userDetails.role }}</p>
          <p><strong>First Name:</strong> {{ adminStore.userDetails.first_name || 'N/A' }}</p>
          <p><strong>Last Name:</strong> {{ adminStore.userDetails.last_name || 'N/A' }}</p>
          <p><strong>Active:</strong> <i :class="adminStore.userDetails.is_active ? 'pi pi-check-circle' : 'pi pi-times-circle'"
                 :style="{ color: adminStore.userDetails.is_active ? 'var(--green-500)' : 'var(--red-500)' }"></i></p>
          <p><strong>Blocked:</strong> <i :class="adminStore.userDetails.is_blocked ? 'pi pi-lock' : 'pi pi-unlock'"
                 :style="{ color: adminStore.userDetails.is_blocked ? 'var(--red-500)' : 'var(--green-500)' }"></i></p>
          <p><strong>Joined:</strong> {{ formatDate(adminStore.userDetails.created_at) }}</p>

          <Divider />

          <h3>Assigned Keys (Subscription Keys)</h3>
          <div v-if="adminStore.userDetails.SubscriptionKeys && adminStore.userDetails.SubscriptionKeys.length > 0">
            <DataTable :value="adminStore.userDetails.SubscriptionKeys" class="p-datatable-sm" :rows="5" paginator>
              <Column field="dummy_key" header="Dummy Key"></Column>
              <Column field="original_key.key_value" header="Original Key"></Column>
              <Column field="is_used" header="Used">
                <template #body="slotProps">
                  <i :class="slotProps.data.is_used ? 'pi pi-check-circle' : 'pi pi-times-circle'"
                     :style="{ color: slotProps.data.is_used ? 'var(--green-500)' : 'var(--red-500)' }"></i>
                </template>
              </Column>
              <Column field="assigned_at" header="Assigned At">
                <template #body="slotProps">
                  {{ formatDate(slotProps.data.assigned_at) }}
                </template>
              </Column>
            </DataTable>
          </div>
          <div v-else>
            <p>No subscription keys assigned to this user.</p>
          </div>
        </div>
        <div v-else>
          <p>Loading user details...</p>
        </div>
        <template #footer>
          <Button label="Close" icon="pi pi-times" class="p-button-text" @click="showUserDetailsDialog = false" />
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
import Tooltip from 'primevue/tooltip'; // Import Tooltip

// Register Tooltip directive globally or locally if not already
// app.directive('tooltip', Tooltip); // In main.js

const adminStore = useAdminStore()
const toast = useToast()
const confirm = useConfirm()

const showEditUserDialog = ref(false)
const currentUser = ref(null)
const editMode = ref('edit') // 'edit' or 'add'

const showUserDetailsDialog = ref(false)

onMounted(() => {
  adminStore.fetchUsers()
})

const formatDate = (dateString) => {
  if (!dateString) return 'N/A'
  const options = { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit' }
  return new Date(dateString).toLocaleDateString(undefined, options)
}

const onPageChange = (event) => {
  adminStore.fetchUsers(event.page + 1, event.rows)
}

const editUser = (user) => {
  currentUser.value = { ...user }
  editMode.value = 'edit'
  showEditUserDialog.value = true
}

const saveUser = async () => {
  // In a real app, you'd call an API to update/create user
  console.log('Saving user:', currentUser.value)
  toast.add({ severity: 'success', summary: 'Success', detail: 'User saved successfully!', life: 3000 })
  showEditUserDialog.value = false
  await adminStore.fetchUsers(adminStore.pagination.page, adminStore.pagination.limit) // Refresh data
}

const confirmDeleteUser = (event, user) => {
  confirm.require({
    target: event.currentTarget,
    message: `Are you sure you want to delete user ${user.email}?`,
    icon: 'pi pi-exclamation-triangle',
    acceptClass: 'p-button-danger',
    accept: () => {
      // In a real app, call API to delete user
      toast.add({ severity: 'success', summary: 'Confirmed', detail: 'User deleted!', life: 3000 })
      adminStore.fetchUsers(adminStore.pagination.page, adminStore.pagination.limit) // Refresh data
    },
    reject: () => {
      toast.add({ severity: 'info', summary: 'Rejected', detail: 'Deletion aborted', life: 3000 })
    }
  })
}

const viewUserDetails = async (user) => {
  await adminStore.fetchUserDetails(user.id)
  showUserDetailsDialog.value = true
}

const confirmBlockUser = (event, user) => {
  const action = user.is_blocked ? 'unblock' : 'block';
  confirm.require({
    target: event.currentTarget,
    message: `Are you sure you want to ${action} user ${user.email}?`,
    icon: 'pi pi-exclamation-triangle',
    acceptClass: user.is_blocked ? 'p-button-success' : 'p-button-warning',
    accept: async () => {
      try {
        await adminStore.blockUser(user.id, !user.is_blocked);
        toast.add({ severity: 'success', summary: 'Confirmed', detail: `User ${action}ed successfully!`, life: 3000 });
        await adminStore.fetchUsers(adminStore.pagination.page, adminStore.pagination.limit); // Refresh data
      } catch (error) {
        toast.add({ severity: 'error', summary: 'Operation Failed', detail: error.message || 'An error occurred', life: 3000 });
      }
    },
    reject: () => {
      toast.add({ severity: 'info', summary: 'Rejected', detail: `User ${action} aborted`, life: 3000 });
    }
  });
};
</script>

<style scoped>
.admin-users-page {
  padding: 2rem;
}

.admin-users-page h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
  color: var(--text-color);
}

.admin-users-page .subtitle {
  font-size: 1.1rem;
  margin-bottom: 2rem;
  color: var(--text-color-secondary);
}

.loading-state {
  text-align: center;
  padding: 2rem;
}

.users-table {
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

.user-details-dialog .p-dialog-content p {
  margin-bottom: 0.75rem;
  color: var(--text-color-secondary);
}

.user-details-dialog .p-dialog-content h3 {
  margin-top: 1.5rem;
  margin-bottom: 1rem;
  color: var(--text-color);
}
</style>
