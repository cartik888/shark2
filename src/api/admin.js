import { getApiClient } from "@/lib/apiClient"

export const adminApi = {
  getUsers: (page = 1, limit = 10) => {
    const api = getApiClient() // Get client when function is called
    return api.get(`/admin/users?page=${page}&limit=${limit}`)
  },
  getPlans: () => {
    const api = getApiClient() // Get client when function is called
    return api.get("/admin/plans")
  },
  getSubscriptions: (page = 1, limit = 10) => {
    const api = getApiClient() // Get client when function is called
    return api.get(`/admin/subscriptions?page=${page}&limit=${limit}`)
  },
  getPayments: (page = 1, limit = 10) => {
    const api = getApiClient() // Get client when function is called
    return api.get(`/admin/payments?page=${page}&limit=${limit}`)
  },
  getInvoices: (page = 1, limit = 10) => {
    const api = getApiClient() // Get client when function is called
    return api.get(`/admin/invoices?page=${page}&limit=${limit}`)
  },
  createCredKey: (data) => {
    const api = getApiClient()
    return api.post("/admin/cred-keys", data)
  },
  getAllCredKeys: () => {
    const api = getApiClient()
    return api.get("/admin/cred-keys")
  },
  updateCredKey: (id, data) => {
    const api = getApiClient()
    return api.put(`/admin/cred-keys/${id}`, data)
  },
  blockUser: (id, isBlocked) => {
    const api = getApiClient()
    return api.patch(`/admin/user/${id}/block`, { is_blocked: isBlocked })
  },
  getUserDetails: (id) => {
    const api = getApiClient()
    return api.get(`/admin/user/${id}/details`)
  },
  getKeyLogForUser: (userId) => {
    const api = getApiClient()
    return api.get(`/admin/key-log/${userId}`)
  },
}
