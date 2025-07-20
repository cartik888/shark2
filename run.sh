#!/bin/bash

echo "🦈 Starting Shark SaaS Backend Setup..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or higher."
    exit 1
fi

# Check if MySQL is running
if ! command -v mysql &> /dev/null; then
    echo "⚠️  MySQL client not found. Please ensure MySQL server is running."
fi

# Install dependencies
echo "📦 Installing dependencies..."
go mod download

# Create database if it doesn't exist
echo "🗄️  Setting up database..."
mysql -u root -p -e "CREATE DATABASE IF NOT EXISTS shark_saas;"

# Run the application
echo "🚀 Starting Shark SaaS Backend..."
echo "📊 Admin Panel will be available at: http://localhost:8080"
echo "🔑 Default Admin: admin@shark.com / admin123"
echo ""

go run main.go
