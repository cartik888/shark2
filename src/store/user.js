import { apiClient } from "@/lib/apiClient"

const state = {
  profile: null,
  subscription: null,
  payments: [],
  loading: false,
}

const mutations = {
  SET_PROFILE(state, profile) {
    state.profile = profile
  },
  SET_SUBSCRIPTION(state, subscription) {
    state.subscription = subscription
  },
  SET_PAYMENTS(state, payments) {
    state.payments = payments
  },
  SET_LOADING(state, loading) {
    state.loading = loading
  },
}

const actions = {
  async fetchProfile({ commit }) {
    try {
      commit("SET_LOADING", true)
      const response = await apiClient.get("/me")
      commit("SET_PROFILE", response.data.data)
      return response.data
    } catch (error) {
      throw error
    } finally {
      commit("SET_LOADING", false)
    }
  },

  async updateProfile({ commit }, profileData) {
    try {
      commit("SET_LOADING", true)
      const response = await apiClient.put("/profile", profileData)
      commit("SET_PROFILE", response.data.data)
      return response.data
    } catch (error) {
      throw error
    } finally {
      commit("SET_LOADING", false)
    }
  },

  async fetchSubscription({ commit }) {
    try {
      const response = await apiClient.get("/subscriptions/my")
      commit("SET_SUBSCRIPTION", response.data.data)
      return response.data
    } catch (error) {
      throw error
    }
  },

  async processCheckout({ commit }, checkoutData) {
    try {
      commit("SET_LOADING", true)
      const response = await apiClient.post("/payments/checkout", checkoutData)
      return response.data
    } catch (error) {
      throw error
    } finally {
      commit("SET_LOADING", false)
    }
  },

  async fetchPayments({ commit }) {
    try {
      const response = await apiClient.get("/payments/my")
      commit("SET_PAYMENTS", response.data.data)
      return response.data
    } catch (error) {
      throw error
    }
  },
}

const getters = {
  hasActiveSubscription: (state) => {
    return state.subscription && state.subscription.subscription.status === "active"
  },

  subscriptionKey: (state) => {
    return state.subscription?.assigned_key?.dummy_key || null
  },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
}
