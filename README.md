# Gin Digital Garden 🌱

A modern, lightweight digital garden (backend) built with Go and Gin. Plant your thoughts, watch them grow.

## ✨ Features

- 🚀 Blazing fast performance with Go and Gin
- 📝 Markdown-based content management
- 🔍 SEO-friendly URL slugs
- 🗑️ Soft delete support
- 🔄 Automatic content updates
- 🎨 Clean and modern API design

## 🛠️ Tech Stack

- Go 1.20+
- Gin Web Framework
- GORM for database operations
- MySQL for data storage
- Redis for caching

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/alucpro/gin-digital-garden.git

# Install dependencies
go mod download

# Set up your environment
cp .env.example .env

# Run the application
go run main.go
```

## 📚 API Endpoints

- `GET /posts` - List all posts
- `GET /posts/:slug` - Get a specific post
- `POST /posts` - Create a new post
- `PUT /posts/:slug` - Update a post
- `DELETE /posts/:slug` - Delete a post

## 🤝 Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Made with ❤️ by [Your Name]
