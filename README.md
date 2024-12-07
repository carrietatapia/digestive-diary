# Digestive Diary

**Digestive Diary** is a comprehensive health and wellness tracking system that helps users log and monitor their meals, bowel movements, symptoms, exercises, and overall well-being. It connects the dots between daily activities and health, providing insights into patterns that may impact digestive and general health.

## Features
- **Meal Tracking**: Log meals, beverages, medications, and symptoms with optional weather and exercise details.
- **Digestive Health Monitoring**: Record stool consistency (based on the Bristol Stool Scale), color, and other observations.
- **Activity Logging**: Track exercises, intensity, and duration.
- **Menstrual Cycle Tracking**: Log menstrual phases and associated symptoms (for female users).
- **Weather Tracking**: Record environmental conditions during activities or health events.
- **User-Friendly API**: Manage data with CRUD operations for all entities.

---

## Installation

### Prerequisites
- **Go**: Make sure you have Go installed (version 1.18 or higher recommended).
- **PostgreSQL**: Ensure PostgreSQL is installed and running.
- **Swagger UI**: For exploring the API documentation.

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/digestive-diary.git
   cd digestive-diary
   ```

2. Configure the .env file with your database credentials:
   ```bash
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=digestive_diary
   ```

3. Run database migrations to set up the schema:

   ```bash
   go run cmd/migrate.go
   ```
4. Install dependencies:

   ```bash
   go mod tidy
   ```

5. Run the application:
   ```bash
   go run main.go
   ```
6. Access the Swagger API documentation: Navigate to http://localhost:8080/swagger/index.html in your browser.

7. (Optional) Run tests to verify the setup:

   ```bash
   go test ./...
   ```
