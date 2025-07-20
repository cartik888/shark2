import { getApiClient } from "@/lib/apiClient"

export const userApi = {
  getProfile: () => {
    const api = getApiClient()
    return api.get("/me")
  },
  updateProfile: (data) => {
    const api = getApiClient()
    return api.put("/profile", data)
  },
  getMySubscription: () => {
    const api = getApiClient()
    return api.get("/subscriptions/my")
  },
  checkout: (data) => {
    const api = getApiClient()
    return api.post("/payments/checkout", data)
  },
  getMyPayments: () => {
    const api = getApiClient()
    return api.get("/payments/my")
  },
}
