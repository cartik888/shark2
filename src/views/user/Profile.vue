<template>
  <DashboardLayout>
    <div class="profile-page container">
      <h1>My Profile</h1>
      <p class="subtitle">Update your personal information.</p>

      <div v-if="userStore.loading" class="loading-state">
        <ProgressBar mode="indeterminate" style="height: 6px"></ProgressBar>
        <p>Loading profile...</p>
      </div>

      <div v-else class="profile-form">
        <Card>
          <template #content>
            <form @submit.prevent="saveProfile" class="p-fluid">
              <div class="p-field">
                <label for="email">Email (Read-only)</label>
                <InputText id="email" v-model="profileForm.email" disabled />
              </div>
              <div class="p-field">
                <label for="firstName">First Name</label>
                <InputText id="firstName" v-model="profileForm.first_name" />
              </div>
              <div class="p-field">
                <label for="lastName">Last Name</label>
                <InputText id="lastName" v-model="profileForm.last_name" />
              </div>
              <div class="p-field">
                <label for="phone">Phone</label>
                <InputText id="phone" v-model="profileForm.phone" />
              </div>
              <div class="p-field">
                <label for="company">Company</label>
                <InputText id="company" v-model="profileForm.company" />
              </div>
              <Button type="submit" label="Save Profile" :loading="userStore.loading" class="p-button-primary" />
            </form>
          </template>
        </Card>
      </div>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useToast } from 'primevue/usetoast'
import { useUserStore } from '@/store/user'
import { useAuthStore } from '@/store/auth'
import DashboardLayout from '@/layouts/DashboardLayout.vue'

const userStore = useUserStore()
const authStore = useAuthStore()
const toast = useToast()

const profileForm = ref({
  email: '',
  first_name: '',
  last_name: '',
  phone: '',
  company: ''
})

onMounted(async () => {
  await userStore.fetchProfile()
  if (userStore.profile) {
    profileForm.value = { ...userStore.profile, email: authStore.user?.email }
  }
})

// Watch for changes in userStore.profile and update form
watch(() => userStore.profile, (newProfile) => {
  if (newProfile) {
    profileForm.value = { ...newProfile, email: authStore.user?.email }
  }
}, { deep: true })

const saveProfile = async () => {
  try {
    await userStore.updateProfile(profileForm.value)
    toast.add({ severity: 'success', summary: 'Success', detail: 'Profile updated successfully!', life: 3000 })
  } catch (error) {
    toast.add({ severity: 'error', summary: 'Update Failed', detail: error.message || 'An error occurred', life: 3000 })
  }
}
</script>

<style scoped>
.profile-page {
  padding: 2rem;
}

.profile-page h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
  color: var(--text-color);
}

.profile-page .subtitle {
  font-size: 1.1rem;
  margin-bottom: 2rem;
  color: var(--text-color-secondary);
}

.loading-state {
  text-align: center;
  padding: 2rem;
}

.profile-form .p-card {
  max-width: 600px;
  margin: 0 auto;
  text-align: left;
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
