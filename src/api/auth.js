import { getApiClient } from "@/lib/apiClient"

export const authApi = {
  login: (data) => {
    const api = getApiClient() // Get client when function is called
    return api.post("/auth/login", data)
  },
  register: (data) => {
    const api = getApiClient() // Get client when function is called
    return api.post("/auth/register", data)
  },
  getMe: () => {
    const api = getApiClient() // Get client when function is called
    return api.get("/me")
  },
}
