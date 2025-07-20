import { defineStore } from "pinia"
import { ref } from "vue"
import { plansApi } from "@/api/plans"

export const usePlansStore = defineStore("plans", () => {
  const plans = ref([])
  const loading = ref(false)

  const fetchPublicPlans = async () => {
    loading.value = true
    try {
      const response = await plansApi.getPublicPlans()
      plans.value = response.data.data
    } catch (error) {
      console.error("Failed to fetch plans:", error)
    } finally {
      loading.value = false
    }
  }

  return {
    plans,
    loading,
    fetchPublicPlans,
  }
})
