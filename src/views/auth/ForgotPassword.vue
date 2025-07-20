<template>
  <PublicLayout>
    <div class="auth-page">
      <Card class="auth-card">
        <template #title>Forgot Password</template>
        <template #content>
          <form @submit.prevent="handleForgotPassword" class="p-fluid">
            <div class="p-field">
              <label for="email">Email</label>
              <InputText id="email" type="email" v-model="email" required />
            </div>
            <Button type="submit" label="Reset Password" :loading="loading" class="p-button-primary" />
          </form>
          <div class="auth-links">
            <router-link to="/login">Back to Login</router-link>
          </div>
        </template>
      </Card>
    </div>
  </PublicLayout>
</template>

<script setup>
import { ref } from 'vue'
import { useToast } from 'primevue/usetoast'
import PublicLayout from '@/layouts/PublicLayout.vue'

const email = ref('')
const loading = ref(false)
const toast = useToast()

const handleForgotPassword = async () => {
  loading.value = true
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500))
    console.log('Password reset requested for:', email.value)
    toast.add({ severity: 'success', summary: 'Success', detail: 'Password reset link sent to your email!', life: 3000 })
  } catch (error) {
    toast.add({ severity: 'error', summary: 'Error', detail: 'Failed to send reset link. Please try again.', life: 3000 })
  } finally {
    loading.value = false
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
