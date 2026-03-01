# Admin Dashboard APIs - Quick Integration Guide

## 📋 What Was Created

You now have **4 professional admin dashboard APIs** ready to use:

| API | Endpoint | Purpose |
|-----|----------|---------|
| **Total Users** | `GET /api/v1/admin/dashboard/total-users` | Active platform users count |
| **Active Subscriptions** | `GET /api/v1/admin/dashboard/active-subscriptions` | Paying customers & revenue |
| **Monthly Revenue** | `GET /api/v1/admin/dashboard/monthly-revenue` | Revenue comparison (current vs last month) |
| **User Growth** | `GET /api/v1/admin/dashboard/user-growth` | New registrations by month (12 months) |

---

## 🚀 Quick Start

### Step 1: Start Backend
```bash
cd "d:\go-gin-saa-s-backend (5)"
go run main.go
```

### Step 2: Import Postman Collection
1. Open **Postman**
2. Click **Import** → **Upload Files**
3. Select **admin_dashboard_postman.json**
4. Click **Import**

### Step 3: Set Auth Token
1. Login endpoint: `POST /api/v1/auth/login`
   ```json
   {
     "email": "admin@example.com",
     "password": "your_password"
   }
   ```
2. Copy the JWT token from response
3. In Postman, set variable: `auth_token` = your_token
4. Now test all dashboard endpoints

---

## 📊 API Response Formats

### 1️⃣ Total Users
```json
GET /api/v1/admin/dashboard/total-users

{
  "status": "success",
  "data": {
    "total_users": 1250,
    "active_users": 980,
    "inactive_users": 270
  }
}
```

### 2️⃣ Active Subscriptions
```json
GET /api/v1/admin/dashboard/active-subscriptions

{
  "status": "success",
  "data": {
    "active_subscriptions": 520,
    "total_subscriptions": 850,
    "paying_customers": 520,
    "total_revenue": 52500.00
  }
}
```

### 3️⃣ Monthly Revenue
```json
GET /api/v1/admin/dashboard/monthly-revenue

{
  "status": "success",
  "data": {
    "current_month": "2026-01",
    "current_revenue": 15750.50,
    "previous_month": "2025-12",
    "previous_revenue": 12300.00,
    "revenue_difference": 3450.50,
    "percentage_change": 28.05,
    "trend": "up"
  }
}
```

### 4️⃣ User Growth
```json
GET /api/v1/admin/dashboard/user-growth

{
  "status": "success",
  "data": {
    "total_users": 1250,
    "current_month_new": 85,
    "growth_by_month": [
      {"month": "2025-02", "new_users": 45},
      {"month": "2025-03", "new_users": 62},
      {"month": "2026-01", "new_users": 85}
    ],
    "last_updated": "2026-01-25T10:30:00Z"
  }
}
```

---

## 💻 Frontend Integration (Vue.js Example)

### Install Axios (if not already installed)
```bash
npm install axios
```

### Create Composable/Store
```javascript
// src/api/admin.js

import axios from 'axios';

const API_BASE = 'http://localhost:8080/api/v1/admin/dashboard';

export const adminDashboardAPI = {
  getTotalUsers() {
    return axios.get(`${API_BASE}/total-users`);
  },
  
  getActiveSubscriptions() {
    return axios.get(`${API_BASE}/active-subscriptions`);
  },
  
  getMonthlyRevenue() {
    return axios.get(`${API_BASE}/monthly-revenue`);
  },
  
  getUserGrowth() {
    return axios.get(`${API_BASE}/user-growth`);
  }
};
```

### Use in Vue Component
```vue
<template>
  <div class="dashboard">
    <div class="metrics-grid">
      <MetricCard 
        title="Total Active Users"
        :value="totalUsers.active_users"
        icon="👥"
      />
      <MetricCard 
        title="Paying Customers"
        :value="activeSubscriptions.paying_customers"
        icon="💳"
      />
      <MetricCard 
        title="Monthly Revenue"
        :value="'$' + monthlyRevenue.current_revenue.toFixed(2)"
        icon="💰"
        :trend="monthlyRevenue.trend"
        :change="monthlyRevenue.percentage_change"
      />
      <MetricCard 
        title="New Users This Month"
        :value="userGrowth.current_month_new"
        icon="📈"
      />
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { adminDashboardAPI } from '@/api/admin';

export default {
  setup() {
    const totalUsers = ref({});
    const activeSubscriptions = ref({});
    const monthlyRevenue = ref({});
    const userGrowth = ref({});

    const loadDashboard = async () => {
      try {
        const [users, subs, revenue, growth] = await Promise.all([
          adminDashboardAPI.getTotalUsers(),
          adminDashboardAPI.getActiveSubscriptions(),
          adminDashboardAPI.getMonthlyRevenue(),
          adminDashboardAPI.getUserGrowth()
        ]);

        totalUsers.value = users.data.data;
        activeSubscriptions.value = subs.data.data;
        monthlyRevenue.value = revenue.data.data;
        userGrowth.value = growth.data.data;
      } catch (error) {
        console.error('Error loading dashboard:', error);
      }
    };

    onMounted(() => {
      loadDashboard();
      // Refresh every 5 minutes
      setInterval(loadDashboard, 5 * 60 * 1000);
    });

    return {
      totalUsers,
      activeSubscriptions,
      monthlyRevenue,
      userGrowth
    };
  }
};
</script>
```

---

## 🔒 Security Notes

✅ **All endpoints require:**
- Valid JWT token in Authorization header
- Admin role/privileges
- Bearer token format: `Authorization: Bearer <token>`

❌ **Will return 401 if:**
- Token is missing
- Token is invalid/expired
- Token belongs to non-admin user

❌ **Will return 403 if:**
- User is authenticated but not an admin

---

## 🧪 Testing with cURL

### Get Auth Token
```bash
curl -X POST "http://localhost:8080/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"password123"}'
```

### Test Total Users
```bash
curl -X GET "http://localhost:8080/api/v1/admin/dashboard/total-users" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json"
```

### Test Monthly Revenue
```bash
curl -X GET "http://localhost:8080/api/v1/admin/dashboard/monthly-revenue" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json"
```

### Test User Growth
```bash
curl -X GET "http://localhost:8080/api/v1/admin/dashboard/user-growth" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json"
```

---

## 📁 Files Reference

| File | Purpose |
|------|---------|
| `controllers/admin_controller.go` | Backend API implementations |
| `routes/routes.go` | Route definitions |
| `ADMIN_DASHBOARD_API.md` | Complete API documentation |
| `ADMIN_DASHBOARD_FRONTEND_EXAMPLE.js` | Vue.js component examples |
| `admin_dashboard_postman.json` | Postman collection for testing |
| `IMPLEMENTATION_SUMMARY.md` | Full implementation details |

---

## ⚡ Performance Tips

1. **Caching**: Add Redis caching for monthly revenue (data doesn't change within a month)
   ```go
   // Pseudo-code
   cacheKey := fmt.Sprintf("revenue:%s", currentMonth.Format("2006-01"))
   if cached := cache.Get(cacheKey); cached != nil {
       return cached
   }
   ```

2. **Database Indexing**: Ensure these columns are indexed:
   - `users.created_at`
   - `payments.created_at`
   - `payments.status`
   - `subscriptions.status`

3. **Query Optimization**: Queries use `SELECT` with aggregation, no N+1 queries

---

## 🐛 Troubleshooting

### "Unauthorized" Error
- Check JWT token is valid
- Verify token hasn't expired
- Ensure Bearer prefix is included

### "Forbidden" Error
- Verify user is an admin
- Check `is_admin` flag in user record

### Empty/Zero Results
- Verify test data exists in database
- Check date ranges match your data
- Ensure payments have correct status

### Slow Response
- Check database indexes
- Consider adding caching layer
- Use query optimization techniques

---

## 📞 Support

For questions or issues:
1. Check `ADMIN_DASHBOARD_API.md` for full documentation
2. Review `ADMIN_DASHBOARD_FRONTEND_EXAMPLE.js` for implementation details
3. Test with Postman collection first
4. Check database contains sample data

---

**Status**: ✅ Production Ready
**Last Updated**: January 25, 2026
**Version**: 1.0
