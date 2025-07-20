<template>
  <PublicLayout>
    <div class="home-page">
      <!-- Hero Section -->
      <section class="hero-section section-padding">
        <div class="container hero-content">
          <h1>Integrate AI For A Competitive Analysis For Your Business <span class="star-icon">⭐</span></h1>
          <p class="hero-subtitle">
            Elevate your advertising game with creatives that consistently outperform your competitors, setting your brand apart.
          </p>
          <div class="hero-actions">
            <Button label="Join The Waitlist" class="p-button-primary p-button-lg" @click="router.push('/register')" />
            <Button label="Know More" class="p-button-secondary p-button-lg" @click="router.push('/about')" />
          </div>
        </div>
      </section>

      <!-- Dashboard Mockup Section -->
      <section class="dashboard-mockup-section">
        <div class="container">
          <img src="/images/orinix-dashboard.png" alt="Orinix Dashboard" class="dashboard-image" />
        </div>
      </section>

      <!-- Features Section -->
      <section class="features-section section-padding">
        <div class="container">
          <h2>Help You Find The Best Analysis For Your Business</h2>
          <div class="feature-list-grid">
            <div class="feature-item">
              <div class="feature-icon"><i class="pi pi-check-circle"></i></div>
              <div class="feature-text">
                <h3>Trusted and Accurate</h3>
                <p>Amet minim mollit non deserunt ullamco est sit dolor do amet sint. Velit officia consequat duis</p>
              </div>
            </div>
            <div class="feature-item">
              <div class="feature-icon"><i class="pi pi-building"></i></div>
              <div class="feature-text">
                <h3>Hired By Top Company</h3>
                <p>Amet minim mollit non deserunt ullamco est sit dolor do amet sint. Velit officia consequat duis</p>
              </div>
            </div>
            <div class="feature-item">
              <div class="feature-icon"><i class="pi pi-briefcase"></i></div>
              <div class="feature-text">
                <h3>Various Categories</h3>
                <p>Amet minim mollit non deserunt ullamco est sit dolor do amet sint. Velit officia consequat duis</p>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Video Section -->
      <section class="video-section section-padding">
        <div class="container text-center">
          <h2>Why Orinix would be your best fit?</h2>
          <p class="video-subtitle">Watch this 1 min video to learn about Orinix.</p>
          <div class="video-player-placeholder">
            <img src="/placeholder.svg?height=400&width=800" alt="Video Placeholder" class="video-thumbnail" />
            <Button icon="pi pi-play" class="p-button-rounded p-button-lg p-button-primary video-play-button" />
          </div>
        </div>
      </section>

      <!-- Benefits Grid Section -->
      <section class="benefits-section section-padding">
        <div class="container">
          <div class="benefits-grid">
            <div class="benefit-item">
              <div class="benefit-icon"><i class="pi pi-mobile"></i></div>
              <h3>Fully Responsive</h3>
              <p>With lots of unique blocks, you can easily build a page without coding.</p>
            </div>
            <div class="benefit-item">
              <div class="benefit-icon"><i class="pi pi-th-large"></i></div>
              <h3>Multiple Layouts</h3>
              <p>With lots of unique blocks, you can easily build a page without coding.</p>
            </div>
            <div class="benefit-item">
              <div class="benefit-icon"><i class="pi pi-bolt"></i></div>
              <h3>Faster Loading</h3>
              <p>With lots of unique blocks, you can easily build a page without coding.</p>
            </div>
            <div class="benefit-item">
              <div class="benefit-icon"><i class="pi pi-comments"></i></div>
              <h3>Super Support</h3>
              <p>With lots of unique blocks, you can easily build a page without coding.</p>
            </div>
            <div class="benefit-item">
              <div class="benefit-icon"><i class="pi pi-book"></i></div>
              <h3>Rich Documentation</h3>
              <p>With lots of unique blocks, you can easily build a page without coding.</p>
            </div>
            <div class="benefit-item">
              <div class="benefit-icon"><i class="pi pi-sync"></i></div>
              <h3>Lifetime Updates</h3>
              <p>With lots of unique blocks, you can easily build a page without coding.</p>
            </div>
          </div>
        </div>
      </section>

      <!-- Useful Software Section -->
      <section class="software-section section-padding">
        <div class="container text-center">
          <h2>Useful software that We assist.</h2>
          <div class="software-logo-grid">
            <img src="/placeholder.svg?height=60&width=60" alt="Software Logo 1" />
            <img src="/placeholder.svg?height=60&width=60" alt="Software Logo 2" />
            <img src="/placeholder.svg?height=60&width=60" alt="Software Logo 3" />
            <img src="/placeholder.svg?height=60&width=60" alt="Software Logo 4" />
            <img src="/placeholder.svg?height=60&width=60" alt="Software Logo 5" />
            <img src="/placeholder.svg?height=60&width=60" alt="Software Logo 6" />
            <img src="/placeholder.svg?height=60&width=60" alt="Software Logo 7" />
            <img src="/placeholder.svg?height=60&width=60" alt="Software Logo 8" />
          </div>
        </div>
      </section>

      <!-- Testimonials Section -->
      <section class="testimonials-section section-padding">
        <div class="container text-center">
          <h2>What people are saying about Circle</h2>
          <div class="testimonial-grid">
            <Card class="testimonial-card">
              <template #content>
                <div class="testimonial-header">
                  <Avatar label="LA" shape="circle" class="testimonial-avatar" />
                  <div class="testimonial-info">
                    <h4>Leslie Alexander</h4>
                    <span class="p-text-secondary">Community</span>
                  </div>
                </div>
                <p class="testimonial-text">
                  Circle is being used for my project, and the team has been very helpful. thanks, are there any new Tools you've tried this week?
                </p>
                <div class="star-rating">
                  <i class="pi pi-star-fill"></i><i class="pi pi-star-fill"></i><i class="pi pi-star-fill"></i><i class="pi pi-star-fill"></i><i class="pi pi-star-fill"></i>
                </div>
              </template>
            </Card>
            <Card class="testimonial-card">
              <template #content>
                <div class="testimonial-header">
                  <Avatar label="DR" shape="circle" class="testimonial-avatar" />
                  <div class="testimonial-info">
                    <h4>Dianne Russell</h4>
                    <span class="p-text-secondary">Community</span>
                  </div>
                </div>
                <p class="testimonial-text">
                  Amet minim mollit non deserunt ullamco est sit dolor do amet sint. Velit officia consequat duis enim velit. mollit.
                </p>
                <div class="star-rating">
                  <i class="pi pi-star-fill"></i><i class="pi pi-star-fill"></i><i class="pi pi-star-fill"></i><i class="pi pi-star-fill"></i><i class="pi pi-star-fill"></i>
                </div>
              </template>
            </Card>
            <Card class="testimonial-card">
              <template #content>
                <div class="testimonial-header">
                  <Avatar label="DS" shape="circle" class="testimonial-avatar" />
                  <div class="testimonial-info">
                    <h4>Darrell Steward</h4>
                    <span class="p-text-secondary">Community</span>
                  </div>
                </div>
                <p class="testimonial-text">
                  Making your own mobile app. Thank you, with the new technologies!
                </p>
                <div class="star-rating">
                  <i class="pi pi-star-fill"></i><i class="pi pi-star-fill"></i><i class="pi pi-star-fill"></i><i class="pi pi-star-fill"></i><i class="pi pi-star-fill"></i>
                </div>
              </template>
            </Card>
          </div>
        </div>
      </section>

      <!-- FAQ Section -->
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

      <!-- Bottom CTA Section -->
      <section class="bottom-cta-section section-padding">
        <div class="container cta-card-wrapper">
          <Card class="cta-card">
            <template #content>
              <div class="cta-content-grid">
                <div class="cta-text-content">
                  <h2>Explore Free Version now!</h2>
                  <p>Search all the open positions on the web. Get your own personalized dashboard for lifetime.</p>
                  <div class="cta-buttons">
                    <Button label="Join Waitlist" class="p-button-primary" @click="router.push('/register')" />
                    <Button label="Contact" class="p-button-secondary" @click="router.push('/contact')" />
                  </div>
                </div>
                <div class="cta-image">
                  <img src="/placeholder.svg?height=300&width=400" alt="Illustration" />
                </div>
              </div>
            </template>
          </Card>
        </div>
      </section>
    </div>
  </PublicLayout>
</template>

<script setup>
import PublicLayout from '@/layouts/PublicLayout.vue'
import { useRouter } from 'vue-router'

const router = useRouter()
</script>

<style scoped>
/* Scoped styles for Home.vue */
.home-page {
  background: var(--orinix-gradient-bg);
  color: var(--text-color);
}

/* Hero Section */
.hero-section {
  text-align: center;
  padding-top: 8rem;
  padding-bottom: 8rem;
  position: relative;
  overflow: hidden;
}

.hero-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: url('/placeholder.svg?height=100&width=100') repeat; /* Placeholder for abstract pattern */
  opacity: 0.1;
  z-index: 0;
}

.hero-content {
  position: relative;
  z-index: 1;
}

.hero-content h1 {
  font-size: 4.5rem;
  max-width: 900px;
  margin: 0 auto 1.5rem auto;
  color: white;
}

.star-icon {
  display: inline-block;
  margin-left: 10px;
  font-size: 3.5rem;
  vertical-align: middle;
}

.hero-subtitle {
  font-size: 1.3rem;
  max-width: 700px;
  margin: 0 auto 3rem auto;
  color: var(--orinix-text-light);
}

.hero-actions {
  display: flex;
  justify-content: center;
  gap: 1.5rem;
}

/* Dashboard Mockup Section */
.dashboard-mockup-section {
  padding-top: 0;
  padding-bottom: 6rem;
  background: var(--orinix-gradient-bg); /* Ensure background matches hero */
  position: relative;
  z-index: 1;
}

.dashboard-image {
  width: 100%;
  max-width: 1000px;
  height: auto;
  display: block;
  margin: 0 auto;
  border-radius: 15px;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.5);
}

/* Features Section */
.features-section {
  background: var(--orinix-dark-blue);
  text-align: center;
}

.features-section h2 {
  margin-bottom: 4rem;
  color: white;
}

.feature-list-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 3rem;
  text-align: left;
}

.feature-item {
  display: flex;
  align-items: flex-start;
  gap: 1.5rem;
}

.feature-icon {
  background: var(--orinix-accent-blue);
  border-radius: 50%;
  width: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  box-shadow: 0 5px 15px rgba(0, 123, 255, 0.3);
}

.feature-icon i {
  font-size: 2rem;
  color: white;
}

.feature-text h3 {
  color: white;
  margin-bottom: 0.5rem;
}

.feature-text p {
  color: var(--orinix-text-secondary);
}

/* Video Section */
.video-section {
  background: var(--orinix-gradient-bg);
}

.video-section h2 {
  color: white;
  margin-bottom: 1rem;
}

.video-subtitle {
  font-size: 1.1rem;
  margin-bottom: 3rem;
  color: var(--orinix-text-secondary);
}

.video-player-placeholder {
  position: relative;
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
  border-radius: 15px;
  overflow: hidden;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
}

.video-thumbnail {
  width: 100%;
  height: auto;
  display: block;
}

.video-play-button {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 80px;
  height: 80px;
  font-size: 2.5rem;
  background: rgba(0, 123, 255, 0.8);
  border: 2px solid white;
  color: white;
  cursor: pointer;
  transition: all 0.3s ease;
}

.video-play-button:hover {
  background: var(--orinix-accent-blue);
  transform: translate(-50%, -50%) scale(1.05);
}

/* Benefits Section */
.benefits-section {
  background: var(--orinix-dark-blue);
  text-align: center;
}

.benefits-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 3rem;
}

.benefit-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}

.benefit-icon {
  background: var(--orinix-card-bg); /* Darker background for icon */
  border-radius: 50%;
  width: 80px;
  height: 80px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 1.5rem;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
}

.benefit-icon i {
  font-size: 2.5rem;
  color: var(--orinix-accent-blue);
}

.benefit-item h3 {
  color: white;
  margin-bottom: 0.75rem;
}

.benefit-item p {
  color: var(--orinix-text-secondary);
}

/* Software Section */
.software-section {
  background: var(--orinix-gradient-bg);
}

.software-section h2 {
  color: white;
  margin-bottom: 3rem;
}

.software-logo-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(100px, 1fr));
  gap: 2rem;
  justify-items: center;
  align-items: center;
}

.software-logo-grid img {
  max-width: 100px;
  height: auto;
  filter: brightness(0) invert(1); /* Ensures logos turn white regardless of original color */
  opacity: 0.7;
  transition: opacity 0.3s ease;
}

.software-logo-grid img:hover {
  opacity: 1;
}

/* Testimonials Section */
.testimonials-section {
  background: var(--orinix-dark-blue);
  text-align: center;
}

.testimonials-section h2 {
  color: white;
  margin-bottom: 4rem;
}

.testimonial-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.testimonial-card {
  text-align: left;
  padding: 1.5rem;
}

.testimonial-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.testimonial-avatar {
  background: var(--orinix-accent-blue);
  color: white;
  font-weight: bold;
  width: 50px;
  height: 50px;
  font-size: 1.2rem;
}

.testimonial-info h4 {
  color: white;
  margin-bottom: 0.25rem;
}

.testimonial-info .p-text-secondary {
  font-size: 0.9rem;
  color: var(--orinix-text-secondary);
}

.testimonial-text {
  font-style: italic;
  color: var(--orinix-text-light);
  margin-bottom: 1.5rem;
}

.star-rating i {
  color: gold;
  font-size: 1.2rem;
  margin-right: 0.2rem;
}

/* FAQ Section */
.faq-section {
  background: var(--orinix-gradient-bg);
}

.faq-section h2 {
  color: white;
  margin-bottom: 1rem;
}

.faq-subtitle {
  font-size: 1.1rem;
  margin-bottom: 3rem;
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

/* Bottom CTA Section */
.bottom-cta-section {
  background: var(--orinix-dark-blue);
  padding-top: 0; /* Adjust padding as the card has its own */
  padding-bottom: 6rem;
}

.cta-card-wrapper {
  margin-top: -4rem; /* Overlap with previous section */
  position: relative;
  z-index: 2;
}

.cta-card {
  padding: 3rem;
  background: var(--orinix-gradient-bg-light); /* Lighter gradient for CTA card */
  border: none; /* Remove default card border */
  box-shadow: 0 15px 40px rgba(0, 0, 0, 0.6);
  overflow: hidden; /* For image positioning */
}

.cta-content-grid {
  display: grid;
  grid-template-columns: 2fr 1fr;
  align-items: center;
  gap: 2rem;
  text-align: left;
}

.cta-text-content h2 {
  font-size: 3rem;
  color: white;
  margin-bottom: 1rem;
}

.cta-text-content p {
  font-size: 1.1rem;
  color: var(--orinix-text-light);
  margin-bottom: 2rem;
}

.cta-buttons {
  display: flex;
  gap: 1rem;
}

.cta-image {
  text-align: right;
}

.cta-image img {
  max-width: 100%;
  height: auto;
  display: block;
  margin-left: auto;
}

/* Responsive Adjustments for Home Page */
@media (max-width: 1024px) {
  .hero-content h1 {
    font-size: 3.5rem;
  }
  .star-icon {
    font-size: 2.5rem;
  }
  .feature-list-grid, .benefits-grid, .testimonial-grid, .faq-grid {
    grid-template-columns: 1fr;
  }
  .cta-content-grid {
    grid-template-columns: 1fr;
    text-align: center;
  }
  .cta-image {
    text-align: center;
  }
  .cta-image img {
    margin: 0 auto;
  }
  .cta-buttons {
    justify-content: center;
  }
}

@media (max-width: 768px) {
  .hero-section {
    padding-top: 6rem;
    padding-bottom: 4rem;
  }
  .hero-content h1 {
    font-size: 2.8rem;
  }
  .star-icon {
    font-size: 2rem;
  }
  .hero-subtitle {
    font-size: 1rem;
  }
  .section-padding {
    padding: 4rem 0;
  }
  .features-section h2, .video-section h2, .benefits-section h2, .software-section h2, .testimonials-section h2, .faq-section h2, .cta-text-content h2 {
    font-size: 2rem;
  }
  .feature-item {
    flex-direction: column;
    align-items: center;
    text-align: center;
  }
  .feature-icon {
    margin-bottom: 1rem;
  }
  .faq-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 480px) {
  .hero-content h1 {
    font-size: 2rem;
  }
  .star-icon {
    font-size: 1.5rem;
  }
  .hero-actions {
    flex-direction: column;
    gap: 1rem;
  }
  .p-button.p-button-lg {
    width: 100%;
  }
  .cta-text-content h2 {
    font-size: 1.8rem;
  }
}
</style>
