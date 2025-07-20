@echo off
echo 🚀 Setting up MySQL Database for Shark SaaS Backend...

REM Check if MySQL is installed
mysql --version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ MySQL is not installed. Please install MySQL first.
    echo Download from: https://dev.mysql.com/downloads/mysql/
    pause
    exit /b 1
)

echo ✅ MySQL is installed

REM Prompt for MySQL root password
set /p ROOT_PASSWORD=Please enter your MySQL root password: 

REM Test MySQL connection
mysql -u root -p%ROOT_PASSWORD% -e "SELECT 1;" >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Failed to connect to MySQL. Please check your root password.
    pause
    exit /b 1
)

echo ✅ MySQL connection successful

REM Run the setup script
echo 🔧 Creating database and user...
mysql -u root -p%ROOT_PASSWORD% < scripts/setup_mysql.sql

if %errorlevel% equ 0 (
    echo ✅ Database setup completed successfully!
    echo.
    echo 📋 Database Details:
    echo    Database: shark_saas
    echo    User: shark_user
    echo    Password: shark_password_2024
    echo    Host: localhost
    echo    Port: 3306
    echo.
    echo 🚀 You can now run: go run main.go
) else (
    echo ❌ Database setup failed. Please check the error messages above.
)

pause
