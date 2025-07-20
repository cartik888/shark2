import { createApp } from "vue"
import { createPinia } from "pinia"
import PrimeVue from "primevue/config"
import Aura from "@primevue/themes/aura"
import ToastService from "primevue/toastservice"
import ConfirmationService from "primevue/confirmationservice"

import App from "./App.vue"
import router from "./router"
import { createApiClient } from "./api"
import { setApiClient } from "./lib/apiClient"

// Import the new global stylesheet
import "./assets/style.css"

// PrimeVue Components (keep as is)
import Button from "primevue/button"
import InputText from "primevue/inputtext"
import Password from "primevue/password"
import Card from "primevue/card"
import DataTable from "primevue/datatable"
import Column from "primevue/column"
import Toast from "primevue/toast"
import ConfirmDialog from "primevue/confirmdialog"
import Sidebar from "primevue/sidebar"
import Menu from "primevue/menu"
import Avatar from "primevue/avatar"
import Badge from "primevue/badge"
import Dropdown from "primevue/dropdown"
import Dialog from "primevue/dialog"
import Textarea from "primevue/textarea"
import Calendar from "primevue/calendar"
import Chip from "primevue/chip"
import Tag from "primevue/tag"
import FileUpload from "primevue/fileupload"
import ProgressBar from "primevue/progressbar"
import Divider from "primevue/divider"
import Checkbox from "primevue/checkbox"
import RadioButton from "primevue/radiobutton"
import InputNumber from "primevue/inputnumber"
import ConfirmPopup from "primevue/confirmpopup"
import Paginator from "primevue/paginator"

// PrimeVue CSS (these imports are now handled by style.css, but keeping them here won't hurt)
// import "primevue/resources/primevue.min.css"
// import "primeicons/primeicons.css"

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
const initializedApiClient = createApiClient(pinia)
setApiClient(initializedApiClient)

app.use(router)
app.use(PrimeVue, {
  theme: {
    preset: Aura,
    options: {
      darkModeSelector: ".dark-mode",
    },
  },
})
app.use(ToastService)
app.use(ConfirmationService)

// Register PrimeVue components globally (keep as is)
app.component("Button", Button)
app.component("InputText", InputText)
app.component("Password", Password)
app.component("Card", Card)
app.component("DataTable", DataTable)
app.component("Column", Column)
app.component("Toast", Toast)
app.component("ConfirmDialog", ConfirmDialog)
app.component("Sidebar", Sidebar)
app.component("Menu", Menu)
app.component("Avatar", Avatar)
app.component("Badge", Badge)
app.component("Dropdown", Dropdown)
app.component("Dialog", Dialog)
app.component("Textarea", Textarea)
app.component("Calendar", Calendar)
app.component("Chip", Chip)
app.component("Tag", Tag)
app.component("FileUpload", FileUpload)
app.component("ProgressBar", ProgressBar)
app.component("Divider", Divider)
app.component("Checkbox", Checkbox)
app.component("RadioButton", RadioButton)
app.component("InputNumber", InputNumber)
app.component("ConfirmPopup", ConfirmPopup)
app.component("Paginator", Paginator)

app.mount("#app")
