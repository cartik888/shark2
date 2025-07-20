<template>
  <PublicLayout>
    <div class="about-page section-padding text-center">
      <section class="about-section container">
        <h1>About Our SaaS Platform</h1>
        <p class="subtitle max-w-4xl mx-auto mb-4">
          Welcome to {{ appName }}, your partner in achieving unparalleled productivity and efficiency.
          We are dedicated to providing innovative software solutions that empower businesses
          and individuals to streamline their operations and focus on what truly matters.
        </p>
        <p class="subtitle max-w-4xl mx-auto">
          Founded in 2024, our mission is to simplify complex tasks through intuitive design
          and powerful features. We believe that great software should be accessible, reliable,
          and constantly evolving to meet the demands of a dynamic world.
        </p>
      </section>

      <section class="mission-section container section-margin-top">
        <h2>Our Mission</h2>
        <p class="mission-text max-w-3xl mx-auto">
          To empower businesses and individuals with cutting-edge, user-friendly tools that
          drive growth, foster collaboration, and unlock their full potential.
        </p>
      </section>

      <section class="values-section container section-margin-top">
        <h2>Our Values</h2>
        <div class="value-grid">
          <Card class="value-card">
            <template #title><i class="pi pi-lightbulb"></i><br/>Innovation</template>
            <template #content>
              <p>We constantly seek new ways to improve and deliver groundbreaking solutions.</p>
            </template>
          </Card>
          <Card class="value-card">
            <template #title><i class="pi pi-shield"></i><br/>Integrity</template>
            <template #content>
              <p>We operate with honesty, transparency, and a strong ethical compass.</p>
            </template>
          </Card>
          <Card class="value-card">
            <template #title><i class="pi pi-users"></i><br/>Customer Focus</template>
            <template #content>
              <p>Our users are at the heart of everything we do. Their success is our success.</p>
            </template>
          </Card>
          <Card class="value-card">
            <template #title><i class="pi pi-star"></i><br/>Excellence</template>
            <template #content>
              <p>We strive for the highest quality in our products and services.</p>
            </template>
          </Card>
        </div>
      </section>
    </div>
  </PublicLayout>
</template>

<script setup>
import PublicLayout from '@/layouts/PublicLayout.vue'

const appName = import.meta.env.VITE_APP_NAME || 'Orinix'
</script>

<style scoped>
.about-page {
  background: var(--orinix-gradient-bg);
  color: var(--text-color);
}

.about-section h1 {
  color: white;
  margin-bottom: 1.5rem;
}

.about-section .subtitle {
  line-height: 1.6;
}

.section-margin-top {
  margin-top: 4rem;
}

.mission-section h2, .values-section h2 {
  color: white;
  margin-bottom: 2rem;
}

.mission-text {
  font-size: 1.2rem;
  color: var(--text-color); /* Slightly lighter for mission statement */
}

.value-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
  margin-top: 2rem;
}

.value-card {
  padding: 2rem;
  text-align: center;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.value-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
}

.value-card :deep(.p-card-title) {
  font-size: 1.5rem;
  color: var(--text-color);
  margin-bottom: 1rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.75rem;
}

.value-card :deep(.p-card-title i) {
  color: var(--orinix-accent-blue);
  font-size: 3rem;
}

.value-card :deep(.p-card-content) p {
  color: var(--text-color-secondary);
}

@media (max-width: 1024px) {
  .value-grid {
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  }
}

@media (max-width: 768px) {
  .about-section h1, .mission-section h2, .values-section h2 {
    font-size: 2.5rem;
  }
  .about-section .subtitle, .mission-text {
    font-size: 1rem;
  }
  .value-grid {
    grid-template-columns: 1fr;
  }
  .value-card {
    padding: 1.5rem;
  }
  .value-card :deep(.p-card-title) {
    font-size: 1.3rem;
  }
  .value-card :deep(.p-card-title i) {
    font-size: 2.5rem;
  }
}
</style>
