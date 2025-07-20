import axios from "axios"
import { useAuthStore } from "@/store/auth"

const API_URL = import.meta.env.VITE_API_URL

// This function will create and configure the Axios instance
// It takes the Pinia instance to ensure useAuthStore can be called correctly within interceptors
export const createApiClient = (piniaInstance) => {
  const api = axios.create({
    baseURL: API_URL,
    headers: {
      "Content-Type": "application/json",
    },
  })

  // Request interceptor to add auth token
  api.interceptors.request.use(
    (config) => {
      const authStore = useAuthStore(piniaInstance) // Get store instance using the provided piniaInstance
      const token = authStore.token
      if (token) {
        config.headers.Authorization = `Bearer ${token}`
      }
      return config
    },
    (error) => {
      return Promise.reject(error)
    },
  )

  // Response interceptor to handle auth errors
  api.interceptors.response.use(
    (response) => response,
    (error) => {
      const authStore = useAuthStore(piniaInstance) // Get store instance using the provided piniaInstance
      if (error.response?.status === 401) {
        authStore.logout()
      }
      return Promise.reject(error)
    },
  )

  return api
}

// We no longer export a default 'api' instance directly from here.
// Other API modules will import the 'api' instance from a globally set variable via src/lib/apiClient.js
