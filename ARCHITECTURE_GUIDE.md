# Admin Dashboard APIs - Architecture & Data Flow

## System Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                     FRONTEND (Vue.js/React)                      │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │  Admin Dashboard Component                               │  │
│  │  - Metric Cards (Users, Subscriptions, Revenue, Growth)  │  │
│  │  - Charts & Graphs                                       │  │
│  │  - Real-time Updates                                     │  │
│  └──────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
                                 ▼
                    HTTP GET (with JWT Token)
                                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                    API GATEWAY / ROUTER                          │
│  /api/v1/admin/dashboard/                                        │
│  ├── total-users                                                 │
│  ├── active-subscriptions                                        │
│  ├── monthly-revenue                                             │
│  └── user-growth                                                 │
└─────────────────────────────────────────────────────────────────┘
                                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                    MIDDLEWARE CHAIN                              │
│  1. AuthMiddleware    (Verify JWT Token)                         │
│  2. AdminMiddleware   (Check Admin Role)                         │
└─────────────────────────────────────────────────────────────────┘
                                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                   CONTROLLER LAYER                               │
│  AdminController                                                 │
│  ├── GetTotalUsers()                                             │
│  ├── GetActiveSubscriptions()                                    │
│  ├── GetMonthlyRevenue()                                         │
│  └── GetUserGrowth()                                             │
└─────────────────────────────────────────────────────────────────┘
                                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                   DATABASE LAYER (MySQL)                         │
│  ├── Users Table         (id, email, created_at, is_blocked)   │
│  ├── Subscriptions       (id, user_id, status, created_at)     │
│  ├── Payments            (id, user_id, amount, status, date)   │
│  └── Subscription Keys   (id, is_used, assigned_at)            │
└─────────────────────────────────────────────────────────────────┘
```

---

## Request-Response Flow

### 1. Total Users API
```
REQUEST:
GET /api/v1/admin/dashboard/total-users
Authorization: Bearer eyJhbGc...

FLOW:
1. AuthMiddleware → Validate JWT
2. AdminMiddleware → Check admin role
3. GetTotalUsers() in AdminController
   ├── COUNT(*) FROM users
   ├── COUNT(*) WHERE is_active=true AND is_blocked=false
   └── COUNT(*) WHERE is_active=false OR is_blocked=true
4. Return aggregated results

RESPONSE:
{
  "status": "success",
  "message": "Total users retrieved successfully",
  "data": {
    "total_users": 1250,
    "active_users": 980,
    "inactive_users": 270
  }
}
```

### 2. Active Subscriptions API
```
REQUEST:
GET /api/v1/admin/dashboard/active-subscriptions
Authorization: Bearer eyJhbGc...

FLOW:
1. AuthMiddleware → Validate JWT
2. AdminMiddleware → Check admin role
3. GetActiveSubscriptions() in AdminController
   ├── COUNT(*) FROM subscriptions
   ├── COUNT(*) WHERE status='active'
   ├── JOIN payments ON subscriptions.id = payments.subscription_id
   ├── SUM(payments.amount) WHERE status='completed' AND subscription.status='active'
   └── Return paying customers = active subscriptions

RESPONSE:
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

### 3. Monthly Revenue API
```
REQUEST:
GET /api/v1/admin/dashboard/monthly-revenue
Authorization: Bearer eyJhbGc...

FLOW:
1. AuthMiddleware → Validate JWT
2. AdminMiddleware → Check admin role
3. GetMonthlyRevenue() in AdminController
   ├── Calculate: current_month_start = 2026-01-01
   ├── Calculate: current_month_end = 2026-02-01
   ├── SUM(amount) WHERE status='completed' 
   │   AND created_at >= current_month_start 
   │   AND created_at < current_month_end
   ├── Repeat for previous month
   └── Calculate percentage change & trend

RESPONSE:
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

### 4. User Growth API
```
REQUEST:
GET /api/v1/admin/dashboard/user-growth
Authorization: Bearer eyJhbGc...

FLOW:
1. AuthMiddleware → Validate JWT
2. AdminMiddleware → Check admin role
3. GetUserGrowth() in AdminController
   ├── Query: DATE_TRUNC('month', created_at) as month
   ├── GROUP BY month for last 12 months
   ├── COUNT(*) as new_users per month
   ├── ORDER BY month ASC
   └── Current month new users = TODAY'S MONTH registration count

RESPONSE:
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

## Database Query Patterns

### Pattern 1: Simple Aggregation
```sql
-- Get total users
SELECT COUNT(*) as total_users FROM users;

-- Get active users
SELECT COUNT(*) as active_users 
FROM users 
WHERE is_active = true AND is_blocked = false;
```

### Pattern 2: Conditional Aggregation
```sql
-- Get subscriptions with status filter
SELECT COUNT(*) as active_subscriptions 
FROM subscriptions 
WHERE status = 'active';
```

### Pattern 3: Join with Sum
```sql
-- Revenue from active paying customers
SELECT SUM(payments.amount) as total_revenue
FROM payments
JOIN subscriptions ON payments.subscription_id = subscriptions.id
WHERE payments.status = 'completed' 
AND subscriptions.status = 'active';
```

### Pattern 4: Date Range Filtering
```sql
-- Monthly revenue
SELECT SUM(amount) as revenue
FROM payments
WHERE status = 'completed'
AND created_at >= '2026-01-01'
AND created_at < '2026-02-01';
```

### Pattern 5: Group by Date
```sql
-- User registrations by month
SELECT DATE_TRUNC('month', created_at) as month, COUNT(*) as new_users
FROM users
WHERE created_at >= NOW() - INTERVAL '12 months'
GROUP BY DATE_TRUNC('month', created_at)
ORDER BY month ASC;
```

---

## Data Consistency & Relationships

```
Users
  ├── id (PK)
  ├── email
  ├── created_at
  ├── is_active
  ├── is_blocked
  └── Relationships:
      ├── → Subscriptions (1:N)
      ├── → Payments (1:N)
      └── → UserProfile (1:1)

Subscriptions
  ├── id (PK)
  ├── user_id (FK → Users)
  ├── status (active/inactive/cancelled/expired)
  ├── created_at
  └── Relationships:
      ├── → Users (N:1)
      ├── → Payments (1:N)
      └── → Plans (N:1)

Payments
  ├── id (PK)
  ├── user_id (FK → Users)
  ├── subscription_id (FK → Subscriptions)
  ├── amount
  ├── status (pending/completed/failed/refunded)
  ├── created_at
  └── Relationships:
      ├── → Users (N:1)
      ├── → Subscriptions (N:1)
      └── → Invoices (1:1)

SubscriptionKeys
  ├── id (PK)
  ├── assigned_to_user_id (FK → Users)
  ├── is_used
  ├── assigned_at
  └── Relationships:
      └── → Users (N:1)
```

---

## Performance Optimization Strategy

```
┌─────────────────────────────────────────┐
│     Frontend Request Received           │
└────────────────┬────────────────────────┘
                 ▼
        ┌────────────────────┐
        │ Check Cache Layer? │
        └─────┬──────────────┘
              │
        ┌─────▼──────────────────────────┐
    YES │ Return Cached Data (Fast)      │
        └───────────────────────────────┘
              │
        ┌─────▼──────────────────────────┐
        │ NO - Query Database            │
        └─────┬──────────────────────────┘
              │
        ┌─────▼──────────────────────────┐
        │ Optimizations Applied:         │
        │ ✓ Indexed Columns              │
        │ ✓ Aggregation at DB Level      │
        │ ✓ Minimal Data Transfer        │
        │ ✓ Single Query per Metric      │
        └─────┬──────────────────────────┘
              │
        ┌─────▼──────────────────────────┐
        │ Cache Result (5-15 min TTL)    │
        └─────┬──────────────────────────┘
              │
        ┌─────▼──────────────────────────┐
        │ Return to Frontend             │
        │ (JSON Response)                │
        └────────────────────────────────┘
```

---

## Error Handling Flow

```
Frontend Request
       ▼
┌─────────────────────────────────────────┐
│ 1. Missing Authorization Header?        │
│    YES → 401 Unauthorized               │
│    NO  → Continue                       │
└────────────────┬────────────────────────┘
                 ▼
┌─────────────────────────────────────────┐
│ 2. Invalid JWT Token?                   │
│    YES → 401 Unauthorized               │
│    NO  → Continue                       │
└────────────────┬────────────────────────┘
                 ▼
┌─────────────────────────────────────────┐
│ 3. User is Admin?                       │
│    YES → Continue                       │
│    NO  → 403 Forbidden                  │
└────────────────┬────────────────────────┘
                 ▼
┌─────────────────────────────────────────┐
│ 4. Database Query Error?                │
│    YES → 500 Internal Server Error      │
│    NO  → Continue                       │
└────────────────┬────────────────────────┘
                 ▼
┌─────────────────────────────────────────┐
│ 5. Return Success (200)                 │
│    with aggregated metrics              │
└─────────────────────────────────────────┘
```

---

## Caching Strategy (Optional Enhancement)

```go
// Pseudo-code for caching
type CacheManager struct {
    Cache redis.Client
    TTL   time.Duration
}

// For Monthly Revenue (Cache for 1 hour)
func (cm *CacheManager) GetMonthlyRevenue() {
    key := fmt.Sprintf("revenue:%s", time.Now().Format("2006-01"))
    
    // Try cache first
    if cached, _ := cm.Cache.Get(key); cached != nil {
        return cached  // Fast path
    }
    
    // Query database
    data := queryDatabase()
    
    // Store in cache with 1 hour TTL
    cm.Cache.Set(key, data, 1*time.Hour)
    
    return data
}

// For User Growth (Cache for 6 hours - rarely changes)
func (cm *CacheManager) GetUserGrowth() {
    key := "growth:12months"
    
    if cached, _ := cm.Cache.Get(key); cached != nil {
        return cached
    }
    
    data := queryDatabase()
    cm.Cache.Set(key, data, 6*time.Hour)
    
    return data
}
```

---

## Scalability Considerations

### Current State (Single DB)
```
Request → Server → Single MySQL DB
```

### Future: Multi-DB Setup
```
Request → Load Balancer 
         → API Server 1 ┐
         → API Server 2 ├─→ Read Replica
         → API Server 3 ┘
              ↓
    Primary DB (Writes)
              ↓
    Read Replica (Reads)
```

### Optimization Layers
1. **Database Level**: Indexes, partitioning
2. **Cache Level**: Redis for aggregated metrics
3. **API Level**: Response caching headers
4. **Frontend Level**: Client-side caching

---

**Documentation Version**: 1.0
**Last Updated**: January 25, 2026
