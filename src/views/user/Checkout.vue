<template>
  <div class="min-h-screen bg-gray-50 py-12">
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="bg-white rounded-lg shadow-lg overflow-hidden">
        <div class="px-6 py-8">
          <h1 class="text-3xl font-bold text-gray-900 mb-8">Complete Your Purchase</h1>
          
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
            <!-- Plan Summary -->
            <div class="bg-gray-50 rounded-lg p-6">
              <h2 class="text-xl font-semibold text-gray-900 mb-4">Plan Summary</h2>
              <div v-if="selectedPlan" class="space-y-4">
                <div class="flex justify-between items-center">
                  <span class="text-gray-600">Plan:</span>
                  <span class="font-medium">{{ selectedPlan.name }}</span>
                </div>
                <div class="flex justify-between items-center">
                  <span class="text-gray-600">Price:</span>
                  <span class="font-medium">${{ selectedPlan.price }}/{{ selectedPlan.interval }}</span>
                </div>
                <div class="flex justify-between items-center">
                  <span class="text-gray-600">Currency:</span>
                  <span class="font-medium">{{ selectedPlan.currency }}</span>
                </div>
                <div class="border-t pt-4">
                  <div class="flex justify-between items-center text-lg font-semibold">
                    <span>Total:</span>
                    <span>${{ selectedPlan.price }}</span>
                  </div>
                </div>
                <div class="mt-4">
                  <h3 class="font-medium text-gray-900 mb-2">Features:</h3>
                  <p class="text-gray-600 text-sm">{{ selectedPlan.features }}</p>
                </div>
              </div>
            </div>

            <!-- Payment Form -->
            <div>
              <h2 class="text-xl font-semibold text-gray-900 mb-4">Payment Information</h2>
              <form @submit.prevent="processPayment" class="space-y-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">
                    Payment Method
                  </label>
                  <select 
                    v-model="paymentForm.paymentMethod" 
                    required
                    class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  >
                    <option value="">Select Payment Method</option>
                    <option value="credit_card">Credit Card</option>
                    <option value="debit_card">Debit Card</option>
                    <option value="paypal">PayPal</option>
                    <option value="bank_transfer">Bank Transfer</option>
                  </select>
                </div>

                <div v-if="paymentForm.paymentMethod === 'credit_card' || paymentForm.paymentMethod === 'debit_card'">
                  <div class="grid grid-cols-1 gap-4">
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">
                        Card Number
                      </label>
                      <input 
                        type="text" 
                        placeholder="1234 5678 9012 3456"
                        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                        required
                      />
                    </div>
                    <div class="grid grid-cols-2 gap-4">
                      <div>
                        <label class="block text-sm font-medium text-gray-700 mb-2">
                          Expiry Date
                        </label>
                        <input 
                          type="text" 
                          placeholder="MM/YY"
                          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                          required
                        />
                      </div>
                      <div>
                        <label class="block text-sm font-medium text-gray-700 mb-2">
                          CVV
                        </label>
                        <input 
                          type="text" 
                          placeholder="123"
                          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                          required
                        />
                      </div>
                    </div>
                  </div>
                </div>

                <div class="border-t pt-6">
                  <button 
                    type="submit" 
                    :disabled="loading"
                    class="w-full bg-blue-600 text-white py-3 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
                  >
                    <span v-if="loading">Processing...</span>
                    <span v-else>Complete Payment - ${{ selectedPlan?.price }}</span>
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Success Modal -->
    <div v-if="showSuccessModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-8 max-w-md w-full mx-4">
        <div class="text-center">
          <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-green-100 mb-4">
            <svg class="h-6 w-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
          </div>
          <h3 class="text-lg font-medium text-gray-900 mb-2">Payment Successful!</h3>
          <p class="text-gray-600 mb-6">Your subscription is now active. Here's your activation key:</p>
          
          <div class="bg-gray-100 rounded-lg p-4 mb-6">
            <div class="flex items-center justify-between">
              <code class="text-sm font-mono text-gray-800">{{ activationKey }}</code>
              <button 
                @click="copyKey"
                class="ml-2 px-3 py-1 bg-blue-600 text-white text-sm rounded hover:bg-blue-700"
              >
                {{ copied ? 'Copied!' : 'Copy' }}
              </button>
            </div>
          </div>
          
          <p class="text-sm text-gray-500 mb-6">
            Save this key safely. You'll need it to activate your desktop application.
          </p>
          
          <button 
            @click="goToDashboard"
            class="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700"
          >
            Go to Dashboard
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
  name: 'Checkout',
  data() {
    return {
      selectedPlan: null,
      paymentForm: {
        paymentMethod: ''
      },
      loading: false,
      showSuccessModal: false,
      activationKey: '',
      copied: false
    }
  },
  computed: {
    ...mapState('plans', ['plans'])
  },
  async mounted() {
    const planId = this.$route.query.plan
    if (planId) {
      await this.fetchPlans()
      this.selectedPlan = this.plans.find(plan => plan.id == planId)
    }
    
    if (!this.selectedPlan) {
      this.$router.push('/pricing')
    }
  },
  methods: {
    ...mapActions('plans', ['fetchPlans']),
    ...mapActions('user', ['processCheckout']),
    
    async processPayment() {
      if (!this.paymentForm.paymentMethod) {
        alert('Please select a payment method')
        return
      }

      this.loading = true
      
      try {
        const response = await this.processCheckout({
          plan_id: this.selectedPlan.id,
          payment_method: this.paymentForm.paymentMethod
        })
        
        this.activationKey = response.data.dummy_key
        this.showSuccessModal = true
      } catch (error) {
        alert(error.response?.data?.error || 'Payment failed. Please try again.')
      } finally {
        this.loading = false
      }
    },
    
    copyKey() {
      navigator.clipboard.writeText(this.activationKey)
      this.copied = true
      setTimeout(() => {
        this.copied = false
      }, 2000)
    },
    
    goToDashboard() {
      this.$router.push('/dashboard')
    }
  }
}
</script>
