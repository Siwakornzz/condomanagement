# Backend Service

This is the backend service for **CondoManagement**, written in **Go**.

## Architecture

This project follows **Hexagonal Architecture** (also known as **Ports and Adapters**). The core idea is to separate the business logic from external dependencies, making it easier to test, maintain, and extend.

### Core Logic
At the heart of the system, we have the core business logic, which includes:
- **Entities**: The core business objects (e.g., `User`, `Role`)
- **Use Cases**: The business operations (e.g., `CreateUser`, `GetUser`)

### Ports
**Ports** define the interfaces through which the core logic communicates with the outside world. For example:
- **UserRepositoryPort**: Interface for storing and retrieving users from a database
- **CreateUserPort**: Interface for creating users within the system

### Adapters
**Adapters** implement the Ports and connect the core logic to external systems. For example:
- **HTTP Adapter**: Handles incoming HTTP requests and converts them into use case invocations
- **Database Adapter (GORM)**: Implements the `UserRepositoryPort` interface and interacts with the database
- **Cache Adapter (Redis)**: Implements caching logic for frequently accessed data

By following the Hexagonal Architecture, we ensure that the core business logic remains decoupled from the infrastructure and external services.



## 📂 Project Structure

```
└── 📁backend
    └── 📁cmd
        └── main.go
    └── 📁internal
        └── 📁adapter
            └── 📁cache
                └── 📁redis
            └── 📁database
                └── 📁gorm
            └── 📁http
            └── readme.md
        └── 📁config
            └── config.go
            └── env.example
        └── 📁core
            └── 📁role
            └── 📁user
        └── 📁middleware
        └── 📁utils
    └── 📁pkg
        └── 📁logging
        └── 📁metrics
        └── 📁tracing
    └── .env
    └── docker-compose.yml
    └── Dockerfile
    └── go.mod
    └── go.sum
    └── readme.md
```

## 🚀 Getting Started

### 1️⃣ Clone the Repository

```sh
git clone https://github.com/your-repo.git
cd backend
```

### 2️⃣ Setup Environment ( cp is an command for copying files or directories.)

```sh
cp internal/config/env.example internal/config/.env
```

### 3️⃣ Run Backend with Docker

```sh
docker-compose --env-file .env up --build
```

### 4️⃣ Run Locally with Go

```sh
go build -o cmd/main cmd/main.go
./cmd/main
```

### 🛠️ External Services Used

| Service    | Purpose                                  |
| :--------- | :--------------------------------------- |
| Fiber      | Web framework for building APIs          |
| PostgreSQL | Stores user and application data         |
| Swagger    | API documentation tool                   |
| Redis      | Caching system for improving performance |
| Prometheus | Metrics monitoring                       |
| Jaeger     | Distributed tracing                      |

### 📌 Deployment Steps

| Method | Command                                          |
| :----- | :----------------------------------------------- |
| Docker | `docker-compose --env-file .env up --build`      |
| Local  | `go build -o cmd/main cmd/main.go && ./cmd/main` |

### 📚 Documentation

- [API Documentation](http://localhost:8080/swagger/index.html)
- [Prometheus Metrics](http://localhost:9090)

```
    _--_     _--_    _--_     _--_     _--_     _--_     _--_     _--_
   ( () )___( () )  ( () )___( () )   ( () )___( () )   ( () )___( () )
    \           /    \           /     \           /     \           /
     (  ' _ `  )      (  ' _ `  )       (  ' _ `  )       (  ' _ `  )
      \  ___  /        \  ___  /         \  ___  /         \  ___  /
    .__( `-' )          ( `-' )           ( `-' )        .__( `-' )  ___
   / !  `---' \      _--'`---_          .--`---'\       /   /`---'`-'   \
  /  \         !    /         \___     /        _>\    /   /          ._/ 
 !   /\        !   /   /       !  \   /  /-___-'   ) /'   /.-----\___/     / -)
 !   !_\______/\. (   <        !__/ /'  (        _/  \___//          `----'   |
  \    \       ! \ \   \      /\    \___/`------' )       \            ______/
   \___/   )  /__/  \--/   \ /  \  ._    \      `<         `--_____----'
     \    /   !       `.    )-   \/  ) ___>-_     \   /-\    \    /
     /   !   /         !   !  `.    / /      `-_   `-/  /    !   !
    !   /__ /___       /  /__   \__/ (  \---__/ `-_    /     /  /__
    (______)____)     (______)        \__)         `-_/     (______)
```