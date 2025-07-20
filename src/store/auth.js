import { defineStore } from "pinia"
import { ref, computed } from "vue"
import { jwtDecode } from "jwt-decode"
import { authApi } from "@/api/auth"
import router from "@/router"

export const useAuthStore = defineStore("auth", () => {
  const token = ref(null)
  const user = ref(null)
  const loading = ref(false)

  const isAuthenticated = computed(() => !!token.value)
  const isAdmin = computed(() => user.value?.role === "admin")
  const isUser = computed(() => user.value?.role === "user")

  const initializeAuth = () => {
    const storedToken = localStorage.getItem("token")
    if (storedToken) {
      try {
        const decoded = jwtDecode(storedToken)
        if (decoded.exp * 1000 > Date.now()) {
          token.value = storedToken
          user.value = {
            id: decoded.user_id,
            email: decoded.email,
            role: decoded.role,
          }
        } else {
          logout()
        }
      } catch (error) {
        logout()
      }
    }
  }

  const login = async (email, password) => {
    loading.value = true
    try {
      const response = await authApi.login({ email, password })
      const { token: authToken, user: userData } = response.data.data

      token.value = authToken
      user.value = userData
      localStorage.setItem("token", authToken)

      // Redirect based on role
      if (userData.role === "admin") {
        router.push("/admin")
      } else {
        router.push("/dashboard")
      }

      return response.data
    } catch (error) {
      throw error.response?.data || error
    } finally {
      loading.value = false
    }
  }

  const register = async (email, password) => {
    loading.value = true
    try {
      const response = await authApi.register({ email, password })
      const { token: authToken, user: userData } = response.data.data

      token.value = authToken
      user.value = userData
      localStorage.setItem("token", authToken)

      router.push("/dashboard")
      return response.data
    } catch (error) {
      throw error.response?.data || error
    } finally {
      loading.value = false
    }
  }

  const logout = () => {
    token.value = null
    user.value = null
    localStorage.removeItem("token")
    router.push("/login")
  }

  return {
    token,
    user,
    loading,
    isAuthenticated,
    isAdmin,
    isUser,
    initializeAuth,
    login,
    register,
    logout,
  }
})
