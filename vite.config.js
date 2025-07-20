import { defineConfig } from "vite"
import vue from "@vitejs/plugin-vue"
import { resolve } from "path"

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      "@": resolve(__dirname, "src"),
    },
  },
  server: {
    port: 3000,
  },
  optimizeDeps: {
    include: [
      "vue",
      "vue-router",
      "pinia",
      "axios",
      "jwt-decode",
      "primevue/config",
      "primevue/toastservice",
      "primevue/confirmationservice",
      // Include commonly used PrimeVue components if you notice slow startup
      "primevue/button",
      "primevue/inputtext",
      "primevue/password",
      "primevue/card",
      "primevue/datatable",
      "primevue/column",
      "primevue/toast",
      "primevue/confirmdialog",
      "primevue/sidebar",
      "primevue/menu",
      "primevue/avatar",
      "primevue/badge",
      "primevue/dropdown",
      "primevue/dialog",
      "primevue/textarea",
      "primevue/calendar",
      "primevue/chip",
      "primevue/tag",
      "primevue/fileupload",
      "primevue/progressbar",
      "primevue/divider",
      "primevue/checkbox",
      "primevue/radiobutton",
      "primevue/inputnumber",
      "primevue/confirmpopup",
      "primevue/paginator",
      "@primevue/themes/aura",
    ],
  },
})
