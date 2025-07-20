# WebImg

A secure image hosting and management API service built with Go, featuring user authentication and cloud storage integration.

## Features

- 🔐 **User Authentication** - Secure signup/login with JWT tokens
- 📸 **Image Upload** - Upload images with authentication protection
- 🖼️ **Image Retrieval** - Access your uploaded images securely
- ☁️ **Cloud Storage** - MinIO object storage integration
- 🛡️ **Secure** - Protected endpoints with middleware authentication
- 🚀 **Fast** - Built with Gin web framework for high performance

## Tech Stack

- **Backend**: Go 1.23+ with Gin web framework
- **Database**: MySQL with GORM ORM
- **Storage**: MinIO object storage
- **Authentication**: JWT tokens
- **Password Security**: bcrypt hashing

## Prerequisites

- Go 1.23 or higher
- MySQL database
- MinIO server (or compatible S3 storage)

## Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd webimg
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   Create a `.env` file in the root directory:
   ```env
   PORT=8080
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=webimg
   JWT_SECRET=your_jwt_secret_key
   MINIO_ENDPOINT=localhost:9000
   MINIO_ACCESS_KEY=your_minio_access_key
   MINIO_SECRET_KEY=your_minio_secret_key
   MINIO_BUCKET=webimg-bucket
   MINIO_USE_SSL=false
   ```

4. **Set up the database**
   - Create a MySQL database named `webimg`
   - The application will automatically create the required tables on startup

5. **Set up MinIO**
   - Install and start MinIO server
   - Create a bucket named `webimg-bucket` (or as specified in your config)

6. **Build and run**
   ```bash
   go build -o webimg
   ./webimg
   ```

## API Endpoints

### Authentication

#### Sign Up
```http
POST /signup
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

#### Login
```http
POST /login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

**Response**: Returns JWT token for authenticated requests

#### Validate Token
```http
GET /validate
Authorization: Bearer <jwt_token>
```

### Image Management

#### Upload Image
```http
POST /upload
Authorization: Bearer <jwt_token>
Content-Type: multipart/form-data

file: <image_file>
```

**Response**: Returns image filename/identifier

#### Get Image
```http
GET /image/:filename
Authorization: Bearer <jwt_token>
```

**Response**: Returns the image file

## Usage Examples

### Using cURL

1. **Sign up a new user**
   ```bash
   curl -X POST http://localhost:8080/signup \
     -H "Content-Type: application/json" \
     -d '{"email": "test@example.com", "password": "password123"}'
   ```

2. **Login**
   ```bash
   curl -X POST http://localhost:8080/login \
     -H "Content-Type: application/json" \
     -d '{"email": "test@example.com", "password": "password123"}'
   ```

3. **Upload an image**
   ```bash
   curl -X POST http://localhost:8080/upload \
     -H "Authorization: Bearer <your_jwt_token>" \
     -F "file=@/path/to/your/image.jpg"
   ```

4. **Retrieve an image**
   ```bash
   curl -X GET http://localhost:8080/image/filename.jpg \
     -H "Authorization: Bearer <your_jwt_token>" \
     --output downloaded_image.jpg
   ```

## Project Structure

```
webimg/
├── main.go              # Application entry point
├── go.mod               # Go module dependencies
├── config/              # Configuration management
├── controllers/         # Request handlers
│   ├── usersController.go
│   └── imagesController.go
├── models/              # Data models
│   └── userModel.go
├── middleware/          # Authentication middleware
├── initializers/        # Database and MinIO setup
└── dto/                 # Data transfer objects
```

## Configuration

The application uses environment variables for configuration. Key settings include:

- `PORT`: Server port (default: 8080)
- `DB_*`: Database connection parameters
- `JWT_SECRET`: Secret key for JWT token signing
- `MINIO_*`: MinIO/S3 storage configuration

## Security Features

- **Password Hashing**: User passwords are hashed using bcrypt
- **JWT Authentication**: Secure token-based authentication
- **Protected Routes**: Image upload/retrieval requires authentication
- **Unique Email**: Email addresses must be unique across users

## Development

### Running in Development Mode
```bash
go run main.go
```

### Building for Production
```bash
go build -o webimg
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

If you encounter any issues or have questions, please open an issue on the repository.
