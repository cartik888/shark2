# 🦈 Shark SaaS Backend

A complete Go-based SaaS backend built with Gin framework, featuring user management, subscription handling, payment processing, and desktop app key validation.

## 🚀 Features

- **User Authentication & Authorization** - JWT-based auth with role management
- **Subscription Management** - Multiple plans with flexible billing
- **Payment Processing** - Stripe integration with payment tracking
- **Desktop App Integration** - Key validation system for desktop applications
- **Admin Dashboard** - Complete admin panel for user and subscription management
- **Credential Keys** - Activation keys, trial keys, and partner keys
- **Support System** - Built-in ticketing system
- **RESTful API** - Clean API design with proper error handling

## 🛠️ Tech Stack

- **Backend**: Go 1.21+ with Gin framework
- **Database**: MySQL with GORM ORM
- **Authentication**: JWT tokens
- **Password Hashing**: bcrypt
- **CORS**: gin-contrib/cors
- **Environment**: godotenv

## 📦 Installation

### Prerequisites

- Go 1.21 or higher
- MySQL 5.7 or higher
- Git

### Quick Start

1. **Clone the repository**
   \`\`\`bash
   git clone <repository-url>
   cd subscription-saas-backend
   \`\`\`

2. **Install dependencies**
   \`\`\`bash
   go mod download
   \`\`\`

3. **Create MySQL database**
   \`\`\`sql
   CREATE DATABASE shark_saas;
   \`\`\`

4. **Configure environment**
   \`\`\`bash
   cp .env.example .env
   # Edit .env with your database credentials
   \`\`\`

5. **Run the application**
   \`\`\`bash
   go run main.go
   \`\`\`

The server will start on `http://localhost:8080`

### Using Docker

\`\`\`bash
# Build and run with Docker Compose
docker-compose up -d

# Or build manually
docker build -t shark-saas-backend .
docker run -p 8080:8080 shark-saas-backend
\`\`\`

## 🔧 Configuration

Update the `.env` file with your configuration:

\`\`\`env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_mysql_password
DB_NAME=shark_saas

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key

# Server Configuration
PORT=8080
GIN_MODE=debug
\`\`\`

## 📚 API Documentation

### Authentication Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/auth/login` | User login |
| POST | `/api/v1/auth/register` | User registration |
| GET | `/api/v1/auth/profile` | Get user profile |
| PUT | `/api/v1/auth/profile` | Update user profile |

### Subscription Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/plans` | Get all active plans |
| GET | `/api/v1/plans/:id` | Get specific plan |
| POST | `/api/v1/user/buy-plan` | Purchase a plan |
| GET | `/api/v1/user/subscriptions` | Get user subscriptions |

### Key Validation Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/keys/validate` | Validate any key type |
| POST | `/api/v1/keys/validate-cred` | Validate credential key |
| POST | `/api/v1/keys/validate-subscription` | Validate subscription key |

### Admin Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/admin/dashboard` | Admin dashboard stats |
| GET | `/api/v1/admin/users` | Get all users |
| POST | `/api/v1/admin/users/:id/block` | Block user |
| GET | `/api/v1/admin/cred-keys` | Get credential keys |
| POST | `/api/v1/admin/cred-keys` | Create credential key |

## 🔑 Default Admin Account

- **Email**: admin@shark.com
- **Password**: admin123

## 🏗️ Project Structure

\`\`\`
subscription-saas-backend/
├── config/
│   └── database.go          # Database configuration
├── controllers/
│   ├── auth_controller.go   # Authentication logic
│   ├── user_controller.go   # User management
│   ├── admin_controller.go  # Admin operations
│   └── ...
├── middlewares/
│   └── auth.go             # JWT middleware
├── models/
│   ├── user.go             # User models
│   ├── subscription.go     # Subscription models
│   └── ...
├── routes/
│   └── routes.go           # API routes
├── utils/
│   ├── jwt.go              # JWT utilities
│   ├── response.go         # Response helpers
│   └── key_generator.go    # Key generation
├── main.go                 # Application entry point
├── .env                    # Environment variables
├── docker-compose.yml      # Docker configuration
└── README.md
\`\`\`

## 🧪 Testing

### Manual Testing with curl

\`\`\`bash
# Register a new user
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123","first_name":"Test","last_name":"User"}'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# Get plans
curl -X GET http://localhost:8080/api/v1/plans

# Validate a key
curl -X POST http://localhost:8080/api/v1/keys/validate \
  -H "Content-Type: application/json" \
  -d '{"key_value":"SHARK-ACTIVATION-2024-001"}'
\`\`\`

## 🔒 Security Features

- **JWT Authentication** - Secure token-based authentication
- **Password Hashing** - bcrypt for secure password storage
- **CORS Protection** - Configurable CORS policies
- **Input Validation** - Request validation and sanitization
- **Role-based Access** - Admin and user role separation

## 📈 Monitoring & Logging

- **Health Check** - `/health` endpoint for monitoring
- **Request Logging** - Gin's built-in logging middleware
- **Error Handling** - Centralized error response system

## 🚀 Deployment

### Production Deployment

1. **Build the application**
   \`\`\`bash
   go build -o shark-saas-backend main.go
   \`\`\`

2. **Set production environment**
   \`\`\`bash
   export GIN_MODE=release
   \`\`\`

3. **Run with process manager**
   \`\`\`bash
   # Using systemd, PM2, or similar
   ./shark-saas-backend
   \`\`\`

### Docker Deployment

\`\`\`bash
# Production build
docker build -t shark-saas-backend:latest .

# Deploy with docker-compose
docker-compose -f docker-compose.prod.yml up -d
\`\`\`

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

For support and questions:

- Create an issue in the repository
- Contact: support@shark-technologies.com
- Documentation: [API Docs](docs/api.md)

## 🔄 Changelog

### v1.0.0
- Initial release
- User authentication and authorization
- Subscription management
- Payment processing
- Admin dashboard
- Desktop app key validation
- Support system

---
Public APIs (/api/v1/):

POST   /api/v1/auth/login
POST   /api/v1/auth/register
GET    /api/v1/plans
GET    /api/v1/plans/:id
POST   /api/v1/keys/validate
POST   /api/v1/keys/validate-cred
POST   /api/v1/keys/validate-subscription
Protected APIs (require authentication, /api/v1/):

GET    /api/v1/auth/profile
PUT    /api/v1/auth/profile
GET    /api/v1/user/dashboard
GET    /api/v1/user/subscriptions
GET    /api/v1/user/payments
POST   /api/v1/user/buy-plan
GET    /api/v1/user/plan-usage
GET    /api/v1/payments
GET    /api/v1/payments/:id
POST   /api/v1/payment-methods
GET    /api/v1/payment-methods
Admin APIs (require admin authentication, /api/v1/admin/):

GET    /api/v1/admin/dashboard
GET    /api/v1/admin/users
GET    /api/v1/admin/users/:id
POST   /api/v1/admin/users/:id/block
POST   /api/v1/admin/users/:id/unblock
GET    /api/v1/admin/cred-keys
POST   /api/v1/admin/cred-keys
DELETE /api/v1/admin/cred-keys/:id
GET    /api/v1/admin/subscriptions
GET    /api/v1/admin/payments
GET    /api/v1/admin/plans
POST   /api/v1/admin/plans
PUT    /api/v1/admin/plans/:id
DELETE /api/v1/admin/plans/:id
GET    /api/v1/admin/support
Health Check:

GET    /health
**Built with ❤️ by Shark Technologies**
#   s h a r k 2  
 