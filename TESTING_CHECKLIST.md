# Admin Dashboard APIs - Testing Checklist

## ✅ Pre-Testing Setup

- [ ] Go backend is running: `go run main.go`
- [ ] MySQL database is running
- [ ] Database contains sample data (or has real data)
- [ ] Admin user account exists in database
- [ ] JWT authentication is working
- [ ] Postman is installed (or use cURL/Insomnia)

---

## 🔐 Authentication Testing

### Login Test
- [ ] POST to `/api/v1/auth/login` with admin credentials
- [ ] Response status is 200
- [ ] Response contains `data.token` (JWT token)
- [ ] Token is not empty
- [ ] Copy token for use in dashboard API tests

### Token Validation
- [ ] Try dashboard API without token → Should return 401
- [ ] Try dashboard API with invalid token → Should return 401
- [ ] Try dashboard API with expired token → Should return 401
- [ ] Try dashboard API with valid token → Should return 200

---

## 📊 Total Users API Testing

**Endpoint**: `GET /api/v1/admin/dashboard/total-users`

### Success Scenarios
- [ ] Request succeeds (200 status)
- [ ] Response has `status: "success"`
- [ ] Response has `message` field
- [ ] Response has `data` object
- [ ] `data.total_users` is a number ≥ 0
- [ ] `data.active_users` is a number ≥ 0
- [ ] `data.inactive_users` is a number ≥ 0
- [ ] `total_users` = `active_users` + `inactive_users`

### Data Validation
- [ ] `total_users` matches database count
- [ ] `active_users` = users with `is_active=true AND is_blocked=false`
- [ ] `inactive_users` includes blocked users
- [ ] Number increases after creating new user

### Error Handling
- [ ] Missing auth token → 401 response
- [ ] Non-admin user → 403 response
- [ ] Database error handled → 500 response

### Performance
- [ ] Response time < 500ms
- [ ] Response size < 1KB

---

## 💳 Active Subscriptions API Testing

**Endpoint**: `GET /api/v1/admin/dashboard/active-subscriptions`

### Success Scenarios
- [ ] Request succeeds (200 status)
- [ ] Response has all required fields
- [ ] `data.active_subscriptions` is a number ≥ 0
- [ ] `data.total_subscriptions` is a number ≥ 0
- [ ] `data.paying_customers` is a number ≥ 0
- [ ] `data.total_revenue` is a number ≥ 0
- [ ] `active_subscriptions` ≤ `total_subscriptions`

### Data Validation
- [ ] `active_subscriptions` matches count of `status='active'`
- [ ] `paying_customers` = active subscriptions count
- [ ] `total_revenue` > 0 if active subscriptions exist
- [ ] Revenue includes only completed payments
- [ ] Revenue calculation is correct (sum of amounts)

### Data Consistency
- [ ] After new payment → `total_revenue` increases
- [ ] After subscription cancel → `active_subscriptions` decreases
- [ ] Refunded amounts don't affect revenue

### Error Handling
- [ ] Missing auth token → 401 response
- [ ] Non-admin user → 403 response
- [ ] Database error handled → 500 response

### Performance
- [ ] Response time < 500ms
- [ ] Works with large payment datasets

---

## 💰 Monthly Revenue API Testing

**Endpoint**: `GET /api/v1/admin/dashboard/monthly-revenue`

### Success Scenarios
- [ ] Request succeeds (200 status)
- [ ] Response has all required fields
- [ ] `data.current_month` is formatted as "YYYY-MM"
- [ ] `data.current_revenue` is a number
- [ ] `data.previous_month` is formatted as "YYYY-MM"
- [ ] `data.previous_revenue` is a number
- [ ] `data.revenue_difference` is calculated correctly
- [ ] `data.percentage_change` is calculated correctly
- [ ] `data.trend` is either "up" or "down"

### Current Month Calculation
- [ ] Uses current calendar month (not sliding window)
- [ ] Starts from 1st of month
- [ ] Ends at start of next month (exclusive)
- [ ] Only includes completed payments

### Trend Logic
- [ ] `trend: "up"` when `current_revenue` > `previous_revenue`
- [ ] `trend: "down"` when `current_revenue` ≤ `previous_revenue`
- [ ] `percentage_change` is positive when trend is "up"
- [ ] `percentage_change` is negative when trend is "down"

### Edge Cases
- [ ] When `previous_revenue` = 0 → `percentage_change` = 100 (if current > 0)
- [ ] When both revenues = 0 → `percentage_change` = 0
- [ ] Handles first month of data correctly
- [ ] Works on first day of month
- [ ] Works on last day of month

### Data Validation
- [ ] Revenue only includes `status='completed'` payments
- [ ] Refunded amounts not included
- [ ] Correct month is queried
- [ ] Previous month is one month before current

### Error Handling
- [ ] Missing auth token → 401 response
- [ ] Non-admin user → 403 response
- [ ] Database error handled → 500 response

### Performance
- [ ] Response time < 500ms
- [ ] Efficient date range queries

---

## 📈 User Growth API Testing

**Endpoint**: `GET /api/v1/admin/dashboard/user-growth`

### Success Scenarios
- [ ] Request succeeds (200 status)
- [ ] Response has all required fields
- [ ] `data.total_users` is a number > 0
- [ ] `data.current_month_new` is a number ≥ 0
- [ ] `data.growth_by_month` is an array
- [ ] Array contains objects with `month` and `new_users`
- [ ] `data.last_updated` is a timestamp

### Data Validation
- [ ] `total_users` matches actual user count
- [ ] `current_month_new` = users registered in current month
- [ ] Each array item has `month` in "YYYY-MM" format
- [ ] Each array item has `new_users` ≥ 0
- [ ] Array is sorted by month (oldest to newest)
- [ ] Array contains up to 12 months of data
- [ ] No gaps in months with data

### 12-Month Coverage
- [ ] Returns last 12 months of data (or less if < 12 months old)
- [ ] Most recent data is current month
- [ ] Oldest data is from 12 months ago

### Current Month Data
- [ ] `current_month_new` matches last item in array
- [ ] Includes all users registered from 1st to today
- [ ] Updates when new user registers

### Edge Cases
- [ ] First month with no users doesn't cause errors
- [ ] Works when < 12 months of data exist
- [ ] Handles year-end transitions correctly
- [ ] Months with 0 registrations shown with `new_users: 0`

### Time-Based Testing
- [ ] Test on 1st of month (includes 0 days of registration)
- [ ] Test mid-month (partial month data)
- [ ] Test end of month (full month data)

### Error Handling
- [ ] Missing auth token → 401 response
- [ ] Non-admin user → 403 response
- [ ] Database error handled → 500 response

### Performance
- [ ] Response time < 500ms
- [ ] Efficient grouping queries

---

## 🔄 Integration Testing

### Coordinated Data Testing
- [ ] Add new user → `total_users` increases
- [ ] Add payment → `total_revenue` increases
- [ ] Activate subscription → `active_subscriptions` increases
- [ ] Cancel subscription → `active_subscriptions` decreases
- [ ] Month-end transition → `current_month_new` resets

### Concurrent Requests
- [ ] Multiple simultaneous requests don't cause errors
- [ ] No race conditions in data calculations
- [ ] All requests return consistent data

### Cross-API Consistency
- [ ] `active_subscriptions` from API 2 matches subscription count
- [ ] Revenue from API 3 matches sum from API 2
- [ ] User count from API 1 >= users in API 4 growth data

---

## 🧪 Load Testing

### Single Request Performance
- [ ] Each API response < 500ms under normal load
- [ ] Each API response < 100ms with cached data

### Multiple Requests
- [ ] 10 concurrent requests → all respond < 1s
- [ ] 100 concurrent requests → no timeouts
- [ ] Dashboard loads all 4 metrics in < 2s

### Database Connection Pool
- [ ] No connection pool exhaustion errors
- [ ] Clean connection closure after query

---

## 🐛 Debugging Checklist

### If Tests Fail

**Query Returns 0/Wrong Values**
- [ ] Verify database has sample data
- [ ] Check SQL query manually in database client
- [ ] Verify date filters are correct
- [ ] Check payment status values
- [ ] Verify subscription status values

**Response is Slow**
- [ ] Check database indexes exist
- [ ] Check database for large table scans
- [ ] Monitor CPU/memory usage
- [ ] Check for N+1 queries
- [ ] Add query caching if needed

**Authentication Fails**
- [ ] Verify JWT token is valid
- [ ] Check token hasn't expired
- [ ] Verify user is admin in database
- [ ] Check Authorization header format

**Data Inconsistencies**
- [ ] Check for timezone mismatches
- [ ] Verify payment status values
- [ ] Check date range boundaries
- [ ] Verify relationship integrity

---

## 📝 Test Report Template

```
Admin Dashboard API Test Report
================================
Date: [DATE]
Tester: [NAME]
Version: [VERSION]

Total Users API
- [ ] ✅ PASS / ❌ FAIL
  Issues: [if any]
  
Active Subscriptions API
- [ ] ✅ PASS / ❌ FAIL
  Issues: [if any]
  
Monthly Revenue API
- [ ] ✅ PASS / ❌ FAIL
  Issues: [if any]
  
User Growth API
- [ ] ✅ PASS / ❌ FAIL
  Issues: [if any]

Overall: ✅ READY FOR PRODUCTION / ❌ NEEDS FIXES

Notes:
[Any additional notes]
```

---

## 🚀 Production Deployment Checklist

Before deploying to production:

- [ ] All tests pass
- [ ] Error handling covers all edge cases
- [ ] Database backups created
- [ ] Indexes optimized
- [ ] Authentication tokens valid
- [ ] CORS properly configured
- [ ] Rate limiting implemented (if needed)
- [ ] Logging enabled
- [ ] Monitoring alerts set up
- [ ] Documentation complete
- [ ] Admin users verified
- [ ] Sample data cleaned up (if needed)

---

**Testing Guide Version**: 1.0
**Last Updated**: January 25, 2026
