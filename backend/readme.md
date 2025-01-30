# Backend Service

This is the backend service for **CondoManagement**, written in **Go**.

### Hexagonal Architecture (Ports and Adapters)

This project follows **Hexagonal Architecture** (also known as **Ports and Adapters**). The core idea is to separate the business logic from external dependencies, making it easier to test, maintain, and extend.

#### Core

The core logic contains the business rules and domain logic, independent of external systems. This is located in the `internal` folder.

#### Ports

**Ports** define the interfaces for communication between the core and external systems. For example:

- `UserRepositoryPort`: Interface to manage user data
- `CreateUserPort`: Interface for creating users

#### Adapters

**Adapters** are responsible for implementing the Ports and enabling communication with external systems, such as:

- **HTTP Adapter**: Handles incoming requests and maps them to the appropriate use case
- **Database Adapter (GORM)**: Implements `UserRepositoryPort` and interacts with the database
- **Cache Adapter (Redis)**: Implements caching mechanisms using Redis

### Architecture Diagram

Below is a high-level diagram of how the **Hexagonal Architecture** is implemented in this project:

![Hexagonal Architecture](https://github.com/Siwakornzz/condomanagement/tree/blob/main/assets/images/hexagonal-architecture.png)

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
