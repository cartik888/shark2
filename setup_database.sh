#!/bin/bash

echo "🚀 Setting up MySQL Database for Shark SaaS Backend..."

# Check if MySQL is installed
if ! command -v mysql &> /dev/null; then
    echo "❌ MySQL is not installed. Please install MySQL first."
    echo "Download from: https://dev.mysql.com/downloads/mysql/"
    exit 1
fi

# Check if MySQL service is running
if ! pgrep -x "mysqld" > /dev/null; then
    echo "❌ MySQL service is not running. Please start MySQL service."
    echo "Windows: net start mysql"
    echo "macOS: brew services start mysql"
    echo "Linux: sudo systemctl start mysql"
    exit 1
fi

echo "✅ MySQL is installed and running"

# Prompt for MySQL root password
echo "Please enter your MySQL root password:"
read -s ROOT_PASSWORD

# Test MySQL connection
mysql -u root -p$ROOT_PASSWORD -e "SELECT 1;" > /dev/null 2>&1
if [ $? -ne 0 ]; then
    echo "❌ Failed to connect to MySQL. Please check your root password."
    exit 1
fi

echo "✅ MySQL connection successful"

# Run the setup script
echo "🔧 Creating database and user..."
mysql -u root -p$ROOT_PASSWORD < scripts/setup_mysql.sql

if [ $? -eq 0 ]; then
    echo "✅ Database setup completed successfully!"
    echo ""
    echo "📋 Database Details:"
    echo "   Database: shark_saas"
    echo "   User: shark_user"
    echo "   Password: shark_password_2024"
    echo "   Host: localhost"
    echo "   Port: 3306"
    echo ""
    echo "🚀 You can now run: go run main.go"
else
    echo "❌ Database setup failed. Please check the error messages above."
    exit 1
fi
