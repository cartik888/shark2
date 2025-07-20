// This file will hold the single, initialized Axios instance
let apiClient = null

export const setApiClient = (client) => {
  apiClient = client
}

export const getApiClient = () => {
  if (!apiClient) {
    // This should ideally not happen if setup is correct in main.js
    console.error("API client not initialized. Ensure createApiClient is called and setApiClient is used in main.js.")
    // You might want to throw an error or return a dummy client in production
    throw new Error("API client not initialized.")
  }
  return apiClient
}
