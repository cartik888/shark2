<template>
  <div class="public-layout">
    <header class="header">
      <nav class="navbar container">
        <div class="nav-brand">
          <router-link to="/" class="brand-link">
            <img src="/images/orinix-logo.svg" alt="Orinix Logo" />
            {{ appName }}
          </router-link>
        </div>
        
        <div class="nav-menu">
          <router-link to="/" class="nav-link">Homepage</router-link>
          <router-link to="/features" class="nav-link">Features</router-link>
          <router-link to="/use-cases" class="nav-link">Use Cases</router-link>
          <router-link to="/pricing" class="nav-link">Pricing</router-link>
          <router-link to="/blog" class="nav-link">Blog</router-link>
        </div>
        
        <div class="nav-actions">
          <Button 
            label="Join The Waitlist" 
            class="p-button-primary"
            @click="router.push('/register')"
          />
        </div>
      </nav>
    </header>
    
    <main class="main-content">
      <slot />
    </main>
    
    <footer class="footer">
      <div class="container">
        <div class="footer-grid">
          <div class="footer-brand">
            <router-link to="/" class="brand-link">
              <img src="/images/orinix-logo.svg" alt="Orinix Logo" />
              {{ appName }}
            </router-link>
            <p>&copy; 2024 Orinix. All rights reserved.</p>
            <div class="social-icons">
              <a href="#" aria-label="Twitter"><i class="pi pi-twitter"></i></a>
              <a href="#" aria-label="Facebook"><i class="pi pi-facebook"></i></a>
              <a href="#" aria-label="LinkedIn"><i class="pi pi-linkedin"></i></a>
            </div>
          </div>

          <div class="footer-column">
            <h4>Product</h4>
            <ul>
              <li><a href="#">Landingpage</a></li>
              <li><a href="#">Features</a></li>
              <li><a href="#">Documentation</a></li>
              <li><a href="#">Referral Program</a></li>
              <li><a href="#">Pricing</a></li>
            </ul>
          </div>

          <div class="footer-column">
            <h4>Services</h4>
            <ul>
              <li><a href="#">Documentation</a></li>
              <li><a href="#">Design</a></li>
              <li><a href="#">Themes</a></li>
              <li><a href="#">Illustrations</a></li>
              <li><a href="#">UI Kit</a></li>
            </ul>
          </div>

          <div class="footer-column">
            <h4>Company</h4>
            <ul>
              <li><a href="#">About</a></li>
              <li><a href="#">Terms</a></li>
              <li><a href="#">Privacy Policy</a></li>
              <li><a href="#">Careers</a></li>
            </ul>
          </div>

          <div class="footer-column">
            <h4>More</h4>
            <ul>
              <li><a href="#">Documentation</a></li>
              <li><a href="#">License</a></li>
              <li><a href="#">Changelog</a></li>
            </ul>
          </div>
        </div>
        <p class="copyright"></p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()
const appName = import.meta.env.VITE_APP_NAME || 'Orinix'
</script>

<style scoped>
/* Scoped styles for PublicLayout.vue */
/* These styles are specific to this component and override global styles if necessary */

.public-layout {
  background: var(--orinix-gradient-bg);
}

.header {
  background: var(--orinix-dark-blue);
}

.brand-link {
  color: var(--text-color);
}

.nav-link {
  color: var(--text-color-secondary);
}

.nav-link:hover {
  color: var(--text-color);
}

.footer {
  background: var(--orinix-dark-blue);
}

.footer-brand .brand-link {
  color: var(--text-color);
}

.footer-column h4 {
  color: var(--text-color);
}

.footer-column ul li a {
  color: var(--text-color-secondary);
}

.footer-column ul li a:hover {
  color: var(--text-color);
}

.copyright {
  color: var(--text-color-secondary);
}

/* Responsive adjustments for PublicLayout */
@media (max-width: 768px) {
  .navbar {
    flex-direction: column;
    align-items: flex-start;
  }
  .nav-menu {
    display: none; /* Hide for mobile, could add a toggle button */
  }
  .nav-actions {
    width: 100%;
    justify-content: flex-start;
  }
  .footer-grid {
    grid-template-columns: 1fr;
    text-align: center;
  }
  .footer-brand {
    align-items: center;
  }
  .social-icons {
    justify-content: center;
  }
}
</style>
