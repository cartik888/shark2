import { getApiClient } from "@/lib/apiClient"

export const plansApi = {
  getPublicPlans: () => {
    const api = getApiClient() // Get client when function is called
    return api.get("/plans/public")
  },
}
