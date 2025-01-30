# Backend Service

This is the backend service for **CondoManagement**, written in **Go**.

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