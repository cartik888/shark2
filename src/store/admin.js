import { defineStore } from "pinia"
import { ref } from "vue"
import { adminApi } from "@/api/admin"

export const useAdminStore = defineStore("admin", () => {
  const users = ref([])
  const plans = ref([])
  const subscriptions = ref([])
  const payments = ref([])
  const invoices = ref([])
  const loading = ref(false)
  const pagination = ref({
    page: 1,
    limit: 10,
    total: 0,
  })

  const credKeys = ref([])
  const userDetails = ref(null)
  const userKeyLog = ref([])

  const fetchUsers = async (page = 1, limit = 10) => {
    loading.value = true
    try {
      const response = await adminApi.getUsers(page, limit)
      users.value = response.data.data.users
      pagination.value = response.data.data.pagination
    } catch (error) {
      console.error("Failed to fetch users:", error)
    } finally {
      loading.value = false
    }
  }

  const fetchPlans = async () => {
    loading.value = true
    try {
      const response = await adminApi.getPlans()
      plans.value = response.data.data
    } catch (error) {
      console.error("Failed to fetch plans:", error)
    } finally {
      loading.value = false
    }
  }

  const fetchSubscriptions = async (page = 1, limit = 10) => {
    loading.value = true
    try {
      const response = await adminApi.getSubscriptions(page, limit)
      subscriptions.value = response.data.data.subscriptions
      pagination.value = response.data.data.pagination
    } catch (error) {
      console.error("Failed to fetch subscriptions:", error)
    } finally {
      loading.value = false
    }
  }

  const fetchPayments = async (page = 1, limit = 10) => {
    loading.value = true
    try {
      const response = await adminApi.getPayments(page, limit)
      payments.value = response.data.data.payments
      pagination.value = response.data.data.pagination
    } catch (error) {
      console.error("Failed to fetch payments:", error)
    } finally {
      loading.value = false
    }
  }

  return {
    users,
    plans,
    subscriptions,
    payments,
    invoices,
    loading,
    pagination,
    fetchUsers,
    fetchPlans,
    fetchSubscriptions,
    fetchPayments,
    credKeys,
    userDetails,
    userKeyLog,
    fetchCredKeys: async () => {
      loading.value = true
      try {
        const response = await adminApi.getAllCredKeys()
        credKeys.value = response.data.data
      } catch (error) {
        console.error("Failed to fetch credential keys:", error)
      } finally {
        loading.value = false
      }
    },
    createCredKey: async (data) => {
      loading.value = true
      try {
        const response = await adminApi.createCredKey(data)
        return response.data
      } catch (error) {
        throw error.response?.data || error
      } finally {
        loading.value = false
      }
    },
    updateCredKey: async (id, data) => {
      loading.value = true
      try {
        const response = await adminApi.updateCredKey(id, data)
        return response.data
      } catch (error) {
        throw error.response?.data || error
      } finally {
        loading.value = false
      }
    },
    blockUser: async (id, isBlocked) => {
      loading.value = true
      try {
        const response = await adminApi.blockUser(id, isBlocked)
        return response.data
      } catch (error) {
        throw error.response?.data || error
      } finally {
        loading.value = false
      }
    },
    fetchUserDetails: async (id) => {
      loading.value = true
      try {
        const response = await adminApi.getUserDetails(id)
        userDetails.value = response.data.data
      } catch (error) {
        console.error("Failed to fetch user details:", error)
      } finally {
        loading.value = false
      }
    },
    fetchKeyLogForUser: async (userId) => {
      loading.value = true
      try {
        const response = await adminApi.getKeyLogForUser(userId)
        userKeyLog.value = response.data.data
      } catch (error) {
        console.error("Failed to fetch key log for user:", error)
      } finally {
        loading.value = false
      }
    },
  }
})
