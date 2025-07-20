<template>
  <PublicLayout>
    <div class="auth-page">
      <Card class="auth-card">
        <template #title>Register</template>
        <template #content>
          <form @submit.prevent="handleRegister" class="p-fluid">
            <div class="p-field">
              <label for="email">Email</label>
              <InputText id="email" type="email" v-model="email" required />
            </div>
            <div class="p-field">
              <label for="password">Password</label>
              <Password id="password" v-model="password" toggleMask required />
            </div>
            <div class="p-field">
              <label for="confirmPassword">Confirm Password</label>
              <Password id="confirmPassword" v-model="confirmPassword" toggleMask :feedback="false" required />
            </div>
            <Button type="submit" label="Register" :loading="authStore.loading" class="p-button-primary" />
          </form>
          <div class="auth-links">
            <router-link to="/login">Already have an account? Login</router-link>
          </div>
        </template>
      </Card>
    </div>
  </PublicLayout>
</template>

<script setup>
import { ref } from 'vue'
import { useToast } from 'primevue/usetoast'
import { useAuthStore } from '@/store/auth'
import PublicLayout from '@/layouts/PublicLayout.vue'

const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const authStore = useAuthStore()
const toast = useToast()

const handleRegister = async () => {
  if (password.value !== confirmPassword.value) {
    toast.add({ severity: 'error', summary: 'Validation Error', detail: 'Passwords do not match', life: 3000 })
    return
  }
  try {
    await authStore.register(email.value, password.value)
    toast.add({ severity: 'success', summary: 'Success', detail: 'Registration successful!', life: 3000 })
  } catch (error) {
    toast.add({ severity: 'error', summary: 'Registration Failed', detail: error.message || 'An error occurred', life: 3000 })
  }
}
</script>

<style scoped>
.auth-page {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: calc(100vh - 120px); /* Adjust for header/footer */
  background: var(--surface-ground);
}

.auth-card {
  width: 100%;
  max-width: 400px;
  padding: 2rem;
  text-align: center;
}

.auth-card .p-card-title {
  margin-bottom: 2rem;
  color: var(--text-color);
}

.p-field {
  margin-bottom: 1.5rem;
  text-align: left;
}

.p-field label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: bold;
  color: var(--text-color);
}

.auth-links {
  margin-top: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.auth-links a {
  color: var(--primary-color);
  text-decoration: none;
}

.auth-links a:hover {
  text-decoration: underline;
}
</style>
