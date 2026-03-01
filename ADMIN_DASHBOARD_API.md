# Admin Dashboard APIs

## Overview
This document describes the admin dashboard APIs for retrieving key metrics and analytics about the SaaS platform.

## Base URL
```
http://localhost:8080/api/v1/admin
```

## Authentication
All endpoints require:
- **Authorization Header**: `Bearer {jwt_token}`
- **Admin Role**: User must have admin privileges

## Endpoints

### 1. Get Total Users
**Endpoint**: `GET /dashboard/total-users`

**Description**: Retrieve total number of platform users, including active and inactive users.

**Headers**:
```json
{
  "Authorization": "Bearer {jwt_token}"
}
```

**Response**:
```json
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

**Status Codes**:
- `200 OK` - Successfully retrieved user statistics
- `401 Unauthorized` - Missing or invalid authentication
- `403 Forbidden` - User does not have admin privileges
- `500 Internal Server Error` - Database error

---

### 2. Get Active Subscriptions
**Endpoint**: `GET /dashboard/active-subscriptions`

**Description**: Retrieve active subscription and paying customer information.

**Headers**:
```json
{
  "Authorization": "Bearer {jwt_token}"
}
```

**Response**:
```json
{
  "status": "success",
  "message": "Active subscriptions retrieved successfully",
  "data": {
    "active_subscriptions": 520,
    "total_subscriptions": 850,
    "paying_customers": 520,
    "total_revenue": 52500.00
  }
}
```

**Status Codes**:
- `200 OK` - Successfully retrieved subscription statistics
- `401 Unauthorized` - Missing or invalid authentication
- `403 Forbidden` - User does not have admin privileges
- `500 Internal Server Error` - Database error

---

### 3. Get Monthly Revenue
**Endpoint**: `GET /dashboard/monthly-revenue`

**Description**: Retrieve current month's revenue and comparison with previous month.

**Headers**:
```json
{
  "Authorization": "Bearer {jwt_token}"
}
```

**Response**:
```json
{
  "status": "success",
  "message": "Monthly revenue retrieved successfully",
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

**Fields Description**:
- `current_month`: Current month in YYYY-MM format
- `current_revenue`: Total revenue for current month
- `previous_month`: Previous month in YYYY-MM format
- `previous_revenue`: Total revenue for previous month
- `revenue_difference`: Absolute difference in revenue
- `percentage_change`: Percentage change (positive = increase)
- `trend`: "up" or "down" based on comparison

**Status Codes**:
- `200 OK` - Successfully retrieved monthly revenue
- `401 Unauthorized` - Missing or invalid authentication
- `403 Forbidden` - User does not have admin privileges
- `500 Internal Server Error` - Database error

---

### 4. Get User Growth
**Endpoint**: `GET /dashboard/user-growth`

**Description**: Retrieve new user registrations by month for the last 12 months.

**Headers**:
```json
{
  "Authorization": "Bearer {jwt_token}"
}
```

**Response**:
```json
{
  "status": "success",
  "message": "User growth data retrieved successfully",
  "data": {
    "total_users": 1250,
    "current_month_new": 85,
    "growth_by_month": [
      {
        "month": "2025-02",
        "new_users": 45
      },
      {
        "month": "2025-03",
        "new_users": 62
      },
      {
        "month": "2025-04",
        "new_users": 78
      },
      {
        "month": "2026-01",
        "new_users": 85
      }
    ],
    "last_updated": "2026-01-25T10:30:00Z"
  }
}
```

**Fields Description**:
- `total_users`: Total number of users in the platform
- `current_month_new`: Number of new users registered in the current month
- `growth_by_month`: Array of user growth data for each month
  - `month`: Month in YYYY-MM format
  - `new_users`: Number of new users registered in that month
- `last_updated`: Timestamp of when the data was retrieved

**Status Codes**:
- `200 OK` - Successfully retrieved user growth data
- `401 Unauthorized` - Missing or invalid authentication
- `403 Forbidden` - User does not have admin privileges
- `500 Internal Server Error` - Database error

---

## Usage Examples

### Example 1: Get Total Users (cURL)
```bash
curl -X GET "http://localhost:8080/api/v1/admin/dashboard/total-users" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

### Example 2: Get Active Subscriptions (cURL)
```bash
curl -X GET "http://localhost:8080/api/v1/admin/dashboard/active-subscriptions" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

### Example 3: Get Monthly Revenue (JavaScript/Fetch)
```javascript
const response = await fetch('http://localhost:8080/api/v1/admin/dashboard/monthly-revenue', {
  method: 'GET',
  headers: {
    'Authorization': `Bearer ${authToken}`,
    'Content-Type': 'application/json'
  }
});

const data = await response.json();
console.log(data.data);
```

### Example 4: Get User Growth (JavaScript/Fetch)
```javascript
const response = await fetch('http://localhost:8080/api/v1/admin/dashboard/user-growth', {
  method: 'GET',
  headers: {
    'Authorization': `Bearer ${authToken}`,
    'Content-Type': 'application/json'
  }
});

const data = await response.json();
console.log('Total Users:', data.data.total_users);
console.log('Growth Data:', data.data.growth_by_month);
```

---

## Postman Collection

### Import Instructions
1. Open Postman
2. Click "Import"
3. Select "Raw Text" or paste the following JSON

```json
{
  "info": {
    "name": "Admin Dashboard APIs",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Get Total Users",
      "request": {
        "method": "GET",
        "header": [
          {
            "key": "Authorization",
            "value": "Bearer {{auth_token}}"
          }
        ],
        "url": {
          "raw": "{{base_url}}/admin/dashboard/total-users",
          "host": ["{{base_url}}"],
          "path": ["admin", "dashboard", "total-users"]
        }
      }
    },
    {
      "name": "Get Active Subscriptions",
      "request": {
        "method": "GET",
        "header": [
          {
            "key": "Authorization",
            "value": "Bearer {{auth_token}}"
          }
        ],
        "url": {
          "raw": "{{base_url}}/admin/dashboard/active-subscriptions",
          "host": ["{{base_url}}"],
          "path": ["admin", "dashboard", "active-subscriptions"]
        }
      }
    },
    {
      "name": "Get Monthly Revenue",
      "request": {
        "method": "GET",
        "header": [
          {
            "key": "Authorization",
            "value": "Bearer {{auth_token}}"
          }
        ],
        "url": {
          "raw": "{{base_url}}/admin/dashboard/monthly-revenue",
          "host": ["{{base_url}}"],
          "path": ["admin", "dashboard", "monthly-revenue"]
        }
      }
    },
    {
      "name": "Get User Growth",
      "request": {
        "method": "GET",
        "header": [
          {
            "key": "Authorization",
            "value": "Bearer {{auth_token}}"
          }
        ],
        "url": {
          "raw": "{{base_url}}/admin/dashboard/user-growth",
          "host": ["{{base_url}}"],
          "path": ["admin", "dashboard", "user-growth"]
        }
      }
    }
  ],
  "variable": [
    {
      "key": "base_url",
      "value": "http://localhost:8080/api/v1"
    },
    {
      "key": "auth_token",
      "value": ""
    }
  ]
}
```

---

## Data Aggregation Notes

1. **Total Users**: Counts all users in the system with breakdown by active/inactive status
2. **Active Subscriptions**: Counts subscriptions with "active" status and calculates total revenue from active paying customers
3. **Monthly Revenue**: Aggregates all completed payments within calendar months for comparison
4. **User Growth**: Groups user creation dates by month for the past 12 months

---

## Error Handling

All endpoints return standardized error responses:

### Unauthorized Error (401)
```json
{
  "status": "error",
  "message": "Unauthorized",
  "error": "Missing or invalid authentication token"
}
```

### Forbidden Error (403)
```json
{
  "status": "error",
  "message": "Forbidden",
  "error": "Admin privileges required"
}
```

### Server Error (500)
```json
{
  "status": "error",
  "message": "Failed to retrieve data",
  "error": "Database connection error"
}
```

---

## Rate Limiting
Currently, no rate limiting is implemented. Consider adding rate limiting in production:
- Suggested: 100 requests per 15 minutes per admin user
- Implement using middleware in `middlewares/ratelimit.go`

---

## Best Practices

1. **Caching**: Consider caching monthly revenue and user growth data since they don't change frequently
2. **Pagination**: For large datasets, consider adding pagination parameters
3. **Date Range Filtering**: Add optional `start_date` and `end_date` query parameters for custom ranges
4. **CSV Export**: Consider adding export functionality for reports

---

## Future Enhancements

- [ ] Add date range filtering for all metrics
- [ ] Add CSV/PDF export functionality
- [ ] Add real-time dashboard updates using WebSockets
- [ ] Add custom report generation
- [ ] Add anomaly detection for revenue changes
- [ ] Add user activity heatmaps
- [ ] Add churn rate calculation
- [ ] Add customer lifetime value (CLV) metrics
