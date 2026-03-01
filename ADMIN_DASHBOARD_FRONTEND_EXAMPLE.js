// Example implementation guide for Admin Dashboard APIs

// Frontend/Vue.js Example - Dashboard Component

// 1. Get Total Users
async function getTotalUsers() {
  const response = await fetch('http://localhost:8080/api/v1/admin/dashboard/total-users', {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
      'Content-Type': 'application/json'
    }
  });
  
  const data = await response.json();
  return data.data;
  // Response:
  // {
  //   "total_users": 1250,
  //   "active_users": 980,
  //   "inactive_users": 270
  // }
}

// 2. Get Active Subscriptions
async function getActiveSubscriptions() {
  const response = await fetch('http://localhost:8080/api/v1/admin/dashboard/active-subscriptions', {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
      'Content-Type': 'application/json'
    }
  });
  
  const data = await response.json();
  return data.data;
  // Response:
  // {
  //   "active_subscriptions": 520,
  //   "total_subscriptions": 850,
  //   "paying_customers": 520,
  //   "total_revenue": 52500.00
  // }
}

// 3. Get Monthly Revenue
async function getMonthlyRevenue() {
  const response = await fetch('http://localhost:8080/api/v1/admin/dashboard/monthly-revenue', {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
      'Content-Type': 'application/json'
    }
  });
  
  const data = await response.json();
  return data.data;
  // Response:
  // {
  //   "current_month": "2026-01",
  //   "current_revenue": 15750.50,
  //   "previous_month": "2025-12",
  //   "previous_revenue": 12300.00,
  //   "revenue_difference": 3450.50,
  //   "percentage_change": 28.05,
  //   "trend": "up"
  // }
}

// 4. Get User Growth
async function getUserGrowth() {
  const response = await fetch('http://localhost:8080/api/v1/admin/dashboard/user-growth', {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
      'Content-Type': 'application/json'
    }
  });
  
  const data = await response.json();
  return data.data;
  // Response:
  // {
  //   "total_users": 1250,
  //   "current_month_new": 85,
  //   "growth_by_month": [
  //     { "month": "2025-02", "new_users": 45 },
  //     { "month": "2025-03", "new_users": 62 },
  //     { "month": "2026-01", "new_users": 85 }
  //   ],
  //   "last_updated": "2026-01-25T10:30:00Z"
  // }
}

// Vue.js Component Example
// Save as: src/views/admin/AdminDashboard.vue

/*
<template>
  <div class="admin-dashboard">
    <div class="dashboard-header">
      <h1>Admin Dashboard</h1>
      <span class="last-updated">Last updated: {{ lastUpdated }}</span>
    </div>

    <div class="metrics-grid">
      <!-- Total Users Card -->
      <div class="metric-card">
        <div class="metric-icon users-icon">👥</div>
        <div class="metric-content">
          <h3>Total Active Users</h3>
          <div class="metric-value">{{ totalUsers.active_users }}</div>
          <div class="metric-subtext">
            {{ totalUsers.total_users }} total users
          </div>
        </div>
      </div>

      <!-- Active Subscriptions Card -->
      <div class="metric-card">
        <div class="metric-icon subscriptions-icon">📊</div>
        <div class="metric-content">
          <h3>Paying Customers</h3>
          <div class="metric-value">{{ activeSubscriptions.paying_customers }}</div>
          <div class="metric-subtext">
            Revenue: ${{ activeSubscriptions.total_revenue.toFixed(2) }}
          </div>
        </div>
      </div>

      <!-- Monthly Revenue Card -->
      <div class="metric-card">
        <div class="metric-icon revenue-icon">💰</div>
        <div class="metric-content">
          <h3>Monthly Revenue</h3>
          <div class="metric-value">${{ monthlyRevenue.current_revenue.toFixed(2) }}</div>
          <div class="metric-subtext" :class="monthlyRevenue.trend === 'up' ? 'positive' : 'negative'">
            <span v-if="monthlyRevenue.trend === 'up'">↑</span>
            <span v-else>↓</span>
            {{ Math.abs(monthlyRevenue.percentage_change).toFixed(2) }}% vs last month
          </div>
        </div>
      </div>

      <!-- New Users This Month Card -->
      <div class="metric-card">
        <div class="metric-icon growth-icon">📈</div>
        <div class="metric-content">
          <h3>New Users This Month</h3>
          <div class="metric-value">{{ userGrowth.current_month_new }}</div>
          <div class="metric-subtext">
            Growth tracking started
          </div>
        </div>
      </div>
    </div>

    <!-- User Growth Chart -->
    <div class="chart-container">
      <h2>User Growth - Last 12 Months</h2>
      <canvas id="userGrowthChart"></canvas>
    </div>

    <!-- Revenue Trend Card -->
    <div class="revenue-trend-card">
      <h2>Revenue Comparison</h2>
      <div class="revenue-comparison">
        <div class="revenue-item">
          <span class="month">{{ monthlyRevenue.current_month }}</span>
          <div class="revenue-bar" :style="{ width: (monthlyRevenue.current_revenue / maxRevenue * 100) + '%' }"></div>
          <span class="amount">${{ monthlyRevenue.current_revenue.toFixed(2) }}</span>
        </div>
        <div class="revenue-item">
          <span class="month">{{ monthlyRevenue.previous_month }}</span>
          <div class="revenue-bar previous" :style="{ width: (monthlyRevenue.previous_revenue / maxRevenue * 100) + '%' }"></div>
          <span class="amount">${{ monthlyRevenue.previous_revenue.toFixed(2) }}</span>
        </div>
      </div>
      <div class="revenue-change" :class="monthlyRevenue.trend">
        <p v-if="monthlyRevenue.trend === 'up'">
          ✓ Revenue increased by ${{ monthlyRevenue.revenue_difference.toFixed(2) }}
        </p>
        <p v-else>
          ✗ Revenue decreased by ${{ Math.abs(monthlyRevenue.revenue_difference).toFixed(2) }}
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import { Chart } from 'chart.js';

export default {
  name: 'AdminDashboard',
  data() {
    return {
      totalUsers: {},
      activeSubscriptions: {},
      monthlyRevenue: {},
      userGrowth: {},
      lastUpdated: new Date().toLocaleString(),
      maxRevenue: 0,
      refreshInterval: null
    };
  },
  methods: {
    async loadDashboardData() {
      try {
        const [users, subs, revenue, growth] = await Promise.all([
          getTotalUsers(),
          getActiveSubscriptions(),
          getMonthlyRevenue(),
          getUserGrowth()
        ]);

        this.totalUsers = users;
        this.activeSubscriptions = subs;
        this.monthlyRevenue = revenue;
        this.userGrowth = growth;
        
        this.maxRevenue = Math.max(revenue.current_revenue, revenue.previous_revenue);
        this.lastUpdated = new Date().toLocaleString();
        
        this.drawUserGrowthChart();
      } catch (error) {
        console.error('Error loading dashboard data:', error);
      }
    },
    drawUserGrowthChart() {
      const ctx = document.getElementById('userGrowthChart');
      if (!ctx) return;

      const labels = this.userGrowth.growth_by_month.map(item => item.month);
      const data = this.userGrowth.growth_by_month.map(item => item.new_users);

      new Chart(ctx, {
        type: 'line',
        data: {
          labels: labels,
          datasets: [{
            label: 'New Users',
            data: data,
            borderColor: 'rgb(75, 192, 192)',
            backgroundColor: 'rgba(75, 192, 192, 0.1)',
            tension: 0.1
          }]
        },
        options: {
          responsive: true,
          plugins: {
            title: {
              display: true,
              text: 'User Growth Over 12 Months'
            }
          }
        }
      });
    }
  },
  mounted() {
    this.loadDashboardData();
    // Refresh data every 5 minutes
    this.refreshInterval = setInterval(() => {
      this.loadDashboardData();
    }, 5 * 60 * 1000);
  },
  beforeUnmount() {
    if (this.refreshInterval) {
      clearInterval(this.refreshInterval);
    }
  }
};
</script>

<style scoped>
.admin-dashboard {
  padding: 20px;
  background-color: #f5f5f5;
}

.dashboard-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
}

.dashboard-header h1 {
  font-size: 28px;
  color: #333;
}

.last-updated {
  font-size: 12px;
  color: #999;
}

.metrics-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

.metric-card {
  background: white;
  border-radius: 8px;
  padding: 20px;
  display: flex;
  gap: 15px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s;
}

.metric-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
}

.metric-icon {
  font-size: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 60px;
  height: 60px;
  background-color: #f0f0f0;
  border-radius: 8px;
}

.metric-content h3 {
  color: #999;
  font-size: 12px;
  margin: 0 0 10px 0;
  text-transform: uppercase;
}

.metric-value {
  font-size: 24px;
  font-weight: bold;
  color: #333;
  margin-bottom: 5px;
}

.metric-subtext {
  font-size: 12px;
  color: #999;
}

.metric-subtext.positive {
  color: #4caf50;
}

.metric-subtext.negative {
  color: #f44336;
}

.chart-container {
  background: white;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 30px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.revenue-trend-card {
  background: white;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.revenue-comparison {
  margin: 20px 0;
}

.revenue-item {
  display: flex;
  align-items: center;
  margin-bottom: 15px;
  gap: 15px;
}

.revenue-item .month {
  width: 60px;
  font-weight: bold;
}

.revenue-bar {
  height: 25px;
  background: linear-gradient(90deg, #4caf50, #81c784);
  border-radius: 4px;
  flex: 1;
  min-width: 100px;
}

.revenue-bar.previous {
  background: linear-gradient(90deg, #90caf9, #64b5f6);
}

.revenue-item .amount {
  width: 100px;
  text-align: right;
  font-weight: bold;
}

.revenue-change {
  margin-top: 20px;
  padding: 15px;
  border-radius: 4px;
  text-align: center;
}

.revenue-change.up {
  background-color: #c8e6c9;
  color: #2e7d32;
}

.revenue-change.down {
  background-color: #ffcdd2;
  color: #c62828;
}
*/
