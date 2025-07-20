<template>
  <PublicLayout>
    <div class="pricing-page">
      <section class="pricing-hero section-padding text-center">
        <div class="container">
          <h1>Pricing Plans</h1>
          <p class="pricing-subtitle">Coca landing page UI Kit. no credit card required. All plans come with a free, 14 day trial of our Premium features.</p>
          <div class="pricing-toggle">
            <Button label="Monthly" :class="{ 'p-button-primary': isMonthly, 'p-button-secondary': !isMonthly }" @click="isMonthly = true" />
            <Button label="Yearly" :class="{ 'p-button-primary': !isMonthly, 'p-button-secondary': isMonthly }" @click="isMonthly = false" />
          </div>
        </div>
      </section>

      <section class="pricing-plans-section section-padding">
        <div class="container">
          <div class="plan-grid">
            <Card v-for="plan in plansStore.plans" :key="plan.id" class="plan-card" :class="{ 'most-popular': plan.name === 'Pro' }">
              <template #header v-if="plan.name === 'Pro'">
                <div class="popular-tag">MOST POPULAR</div>
              </template>
              <template #title>{{ plan.name }}</template>
              <template #content>
                <div class="price-display">
                  <span class="price">${{ isMonthly ? plan.price : (plan.price * 10).toFixed(0) }}</span>
                  <span class="interval">/{{ isMonthly ? 'month' : 'year' }}</span>
                </div>
                <Divider class="plan-divider" />
                <ul class="features-list">
                  <li v-for="(feature, index) in plan.features.split(',')" :key="index">
                    <i class="pi pi-check-circle" style="color: var(--orinix-accent-blue)"></i> {{ feature.trim() }}
                  </li>
                </ul>
              </template>
              <template #footer>
                <Button label="Choose plan" class="p-button-primary" @click="selectPlan(plan)" />
              </template>
            </Card>
          </div>
        </div>
      </section>

      <!-- FAQ Section (re-using from Home.vue for consistency) -->
      <section class="faq-section section-padding">
        <div class="container text-center">
          <h2>Frequently Asked Question</h2>
          <p class="faq-subtitle">Create custom landing pages with Omega that converts.</p>
          <div class="faq-grid">
            <div class="faq-item">
              <h3>What's gonna be your question?</h3>
              <p>Create custom landing pages with Omega that converts more visitors than any website. With lots of unique blocks, you can easily build a page without any design or custom coding, with Omega that converts more visitors than any website.</p>
            </div>
            <div class="faq-item">
              <h3>What's gonna be your question?</h3>
              <p>Create custom landing pages with Omega that converts more visitors than any website. With lots of unique blocks, you can easily build a page without any design or custom coding, with Omega that converts more visitors than any website.</p>
            </div>
            <div class="faq-item">
              <h3>What's gonna be your question?</h3>
              <p>Create custom landing pages with Omega that converts more visitors than any website. With lots of unique blocks, you can easily build a page without any design or custom coding, with Omega that converts more visitors than any website.</p>
            </div>
            <div class="faq-item">
              <h3>What's gonna be your question?</h3>
              <p>Create custom landing pages with Omega that converts more visitors than any website. With lots of unique blocks, you can easily build a page without any design or custom coding, with Omega that converts more visitors than any website.</p>
            </div>
          </div>
          <p class="faq-contact">Didn't find your answer? <router-link to="/contact">Contact us here</router-link></p>
        </div>
      </section>
    </div>
  </PublicLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import { usePlansStore } from '@/store/plans'
import PublicLayout from '@/layouts/PublicLayout.vue'

const plansStore = usePlansStore()
const authStore = useAuthStore()
const router = useRouter()
const isMonthly = ref(true) // Default to monthly pricing

const selectPlan = (plan) => {
  if (!authStore.isAuthenticated) {
    // Store selected plan in localStorage and redirect to register
    localStorage.setItem('selectedPlan', JSON.stringify(plan))
    router.push('/register')
  } else {
    // Redirect to checkout page with plan ID
    router.push(`/checkout?plan=${plan.id}`)
  }
}

onMounted(() => {
  plansStore.fetchPublicPlans()
})
</script>

<style scoped>
/* Scoped styles for Pricing.vue */
.pricing-page {
  background: var(--orinix-gradient-bg);
  color: var(--text-color);
}

.pricing-hero {
  background: var(--orinix-dark-blue); /* Solid dark blue for hero */
  padding-top: 6rem;
  padding-bottom: 4rem;
}

.pricing-hero h1 {
  color: white;
  margin-bottom: 1rem;
}

.pricing-subtitle {
  font-size: 1.1rem;
  margin-bottom: 3rem;
  color: var(--orinix-text-secondary);
  max-width: 700px;
  margin-left: auto;
  margin-right: auto;
}

.pricing-toggle {
  display: inline-flex;
  border-radius: 50px;
  overflow: hidden;
  background: var(--orinix-card-bg);
  border: 1px solid var(--orinix-border-color);
}

.pricing-toggle .p-button {
  border-radius: 50px; /* Ensure buttons inside toggle are also pill-shaped */
  padding: 0.75rem 1.5rem;
  font-weight: 600;
  background: transparent;
  color: var(--text-color-secondary);
  border: none;
  box-shadow: none;
}

.pricing-toggle .p-button.p-button-primary {
  background: var(--orinix-accent-blue);
  color: white;
}

.pricing-toggle .p-button.p-button-secondary {
  background: transparent;
  color: var(--text-color-secondary);
}

.pricing-plans-section {
  background: var(--orinix-gradient-bg);
}

.plan-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  justify-content: center;
  align-items: stretch; /* Ensure cards stretch to same height */
}

.plan-card {
  text-align: center;
  padding: 2rem;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  position: relative;
  overflow: hidden;
}

.plan-card.most-popular {
  border: 2px solid var(--orinix-accent-blue);
  box-shadow: 0 0 0 3px var(--orinix-accent-blue), var(--surface-shadow);
}

.popular-tag {
  position: absolute;
  top: 0;
  left: 50%;
  transform: translate(-50%, -50%);
  background: #ffc107; /* Yellow color from image */
  color: var(--orinix-dark-blue);
  padding: 0.5rem 1.5rem;
  border-radius: 50px;
  font-weight: bold;
  font-size: 0.9rem;
  white-space: nowrap;
  box-shadow: 0 5px 15px rgba(255, 193, 7, 0.3);
}

.plan-card .p-card-title {
  font-size: 1.8rem;
  color: white;
  margin-bottom: 1rem;
}

.price-display {
  margin-top: 1rem;
  margin-bottom: 1.5rem;
}

.price {
  font-size: 3rem;
  font-weight: bold;
  color: var(--orinix-accent-blue);
}

.interval {
  font-size: 1.2rem;
  color: var(--orinix-text-secondary);
  margin-left: 0.5rem;
}

.plan-divider {
  margin: 1.5rem 0;
  border-color: var(--orinix-border-color);
}

.features-list {
  list-style: none;
  padding: 0;
  margin-top: 1rem;
  text-align: left;
  flex-grow: 1; /* Allow list to take available space */
}

.features-list li {
  margin-bottom: 0.75rem;
  color: var(--orinix-text-light);
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.features-list li i {
  font-size: 1.1rem;
}

.plan-card .p-card-footer {
  margin-top: 2rem; /* Space before button */
}

/* FAQ Section (re-used from Home.vue) */
.faq-section {
  background: var(--orinix-dark-blue); /* Different background for contrast */
}

.faq-section h2 {
  color: white;
}

.faq-subtitle {
  color: var(--orinix-text-secondary);
}

.faq-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 2rem;
  text-align: left;
}

.faq-item {
  background: var(--orinix-card-bg);
  padding: 2rem;
  border-radius: var(--border-radius);
  box-shadow: var(--surface-shadow);
  border: 1px solid var(--orinix-border-color);
}

.faq-item h3 {
  color: white;
  margin-bottom: 1rem;
}

.faq-item p {
  color: var(--orinix-text-secondary);
}

.faq-contact {
  margin-top: 3rem;
  font-size: 1.1rem;
  color: var(--orinix-text-secondary);
}

.faq-contact a {
  color: var(--orinix-accent-blue);
  font-weight: 600;
}

/* Responsive Adjustments for Pricing Page */
@media (max-width: 1024px) {
  .plan-grid {
    grid-template-columns: 1fr;
  }
  .faq-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 768px) {
  .pricing-hero h1 {
    font-size: 2.5rem;
  }
  .pricing-subtitle {
    font-size: 1rem;
  }
  .pricing-toggle {
    flex-direction: column;
    width: 100%;
  }
  .pricing-toggle .p-button {
    width: 100%;
    border-radius: 0;
  }
  .pricing-toggle .p-button:first-child {
    border-top-left-radius: 50px;
    border-top-right-radius: 50px;
  }
  .pricing-toggle .p-button:last-child {
    border-bottom-left-radius: 50px;
    border-bottom-right-radius: 50px;
  }
}

@media (max-width: 480px) {
  .pricing-hero h1 {
    font-size: 2rem;
  }
  .price {
    font-size: 2.5rem;
  }
}
</style>
