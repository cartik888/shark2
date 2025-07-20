<template>
  <PublicLayout>
    <div class="contact-page section-padding text-center">
      <section class="contact-section container">
        <h1>Contact Us</h1>
        <p class="subtitle">We'd love to hear from you! Please fill out the form below or reach out to us directly.</p>

        <div class="contact-content">
          <div class="contact-form-card">
            <Card>
              <template #content>
                <form @submit.prevent="submitForm" class="p-fluid form-spacing">
                  <div class="p-field">
                    <label for="name">Name</label>
                    <InputText id="name" v-model="form.name" required />
                  </div>
                  <div class="p-field">
                    <label for="email">Email</label>
                    <InputText id="email" type="email" v-model="form.email" required />
                  </div>
                  <div class="p-field">
                    <label for="subject">Subject</label>
                    <InputText id="subject" v-model="form.subject" required />
                  </div>
                  <div class="p-field">
                    <label for="message">Message</label>
                    <Textarea id="message" v-model="form.message" rows="5" required />
                  </div>
                  <Button type="submit" label="Send Message" class="p-button-primary full-width-button" />
                </form>
              </template>
            </Card>
          </div>

          <div class="contact-info-card">
            <Card>
              <template #title>
                <h2>Our Information</h2>
              </template>
              <template #content>
                <p class="info-item"><i class="pi pi-map-marker"></i> 123 SaaS Street, Tech City, 90210</p>
                <p class="info-item"><i class="pi pi-phone"></i> +1 (555) 123-4567</p>
                <p class="info-item"><i class="pi pi-envelope"></i> support@saasplatform.com</p>
                <Divider class="info-divider" />
                <h3>Business Hours</h3>
                <p class="info-item">Monday - Friday: 9:00 AM - 5:00 PM (PST)</p>
                <p class="info-item">Saturday - Sunday: Closed</p>
              </template>
            </Card>
          </div>
        </div>
      </section>
    </div>
  </PublicLayout>
</template>

<script setup>
import { ref } from 'vue'
import { useToast } from 'primevue/usetoast'
import PublicLayout from '@/layouts/PublicLayout.vue'

const toast = useToast()

const form = ref({
  name: '',
  email: '',
  subject: '',
  message: ''
})

const submitForm = () => {
  console.log('Contact form submitted:', form.value)
  toast.add({ severity: 'success', summary: 'Message Sent', detail: 'We will get back to you shortly!', life: 3000 })
  
  form.value = {
    name: '',
    email: '',
    subject: '',
    message: ''
  }
}
</script>

<style scoped>
.contact-page {
  background: var(--orinix-gradient-bg);
  color: var(--text-color);
}

.contact-section h1 {
  color: white;
  margin-bottom: 1rem;
}

.contact-section .subtitle {
  margin-bottom: 3rem;
}

.contact-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
  max-width: 1000px;
  margin: 0 auto;
  text-align: left; /* Override global text-center for content */
}

.contact-form-card .p-card,
.contact-info-card .p-card {
  padding: 2rem; /* Consistent padding for cards */
}

.form-spacing {
  display: flex;
  flex-direction: column;
  gap: 1.5rem; /* Space between form fields */
}

.p-field label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: bold;
  color: var(--text-color);
}

.full-width-button {
  width: 100%;
  padding: 0.75rem 1.5rem; /* Adjust button padding */
  font-size: 1.1rem;
}

.contact-info-card h2 {
  font-size: 2rem; /* Adjust title size for info card */
  margin-bottom: 1.5rem;
  color: var(--text-color);
}

.contact-info-card h3 {
  font-size: 1.5rem; /* Adjust subtitle size for info card */
  margin-top: 1.5rem;
  margin-bottom: 1rem;
  color: var(--text-color);
}

.info-item {
  margin-bottom: 1rem;
  display: flex;
  align-items: center;
  gap: 0.75rem;
  color: var(--text-color-secondary);
  font-size: 1.1rem;
}

.info-item i {
  color: var(--orinix-accent-blue); /* Icon color */
  font-size: 1.3rem;
}

.info-divider {
  margin: 2rem 0;
  border-color: var(--orinix-border-color);
}

@media (max-width: 1024px) {
  .contact-content {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 768px) {
  .contact-section h1 {
    font-size: 2.5rem;
  }
  .contact-section .subtitle {
    font-size: 1rem;
  }
  .contact-form-card .p-card,
  .contact-info-card .p-card {
    padding: 1.5rem;
  }
  .contact-info-card h2 {
    font-size: 1.8rem;
  }
  .contact-info-card h3 {
    font-size: 1.3rem;
  }
  .info-item {
    font-size: 1rem;
  }
}
</style>
