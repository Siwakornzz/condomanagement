# Project Name CondoManagement

## Overview

This project is a [condo management service]. It is structured into different folders for infrastructure, backend, and configuration management.

## Folder Structure

```
└── 📁condomanagement
    └── 📁.vscode
        └── launch.json
    └── 📁backend
        └── .env
        └── 📁cmd
            └── main.go
        └── docker-compose.yml
        └── Dockerfile
        └── go.mod
        └── go.sum
        └── 📁internal
            └── 📁adapter
                └── 📁cache
                    └── 📁redis
                └── 📁database
                    └── 📁gorm
                └── 📁http
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
        └── readme.md
    └── 📁infra
        └── docker-compose.yml
        └── readme.md
    └── readme.md
```

- **`.vscode/`** - VS Code configuration files.
- **`infra/`** - Infrastructure-related files, such as Docker Compose.
- **`backend/`** - Backend services written in Go.

## Getting Started

1. Clone the repository:

   ```sh
   git clone https://github.com/Siwakornzz/condomanagement
   cd backend
   ```

2. Set up the environment:

   ```sh
   // cp is a command to copy files from one location to another
   cp infra/docker-compose.yml.example infra/docker-compose.yml
   cp backend/cmd/main.go.example backend/cmd/main.go
   cp backend/internal/config/env.example backend/internal/config/env
   ```

3. Build and run the backend services:
   ```sh
   cd infra
   docker-compose up -d
   cd ../backend
   go build -o cmd/main cmd/main.go
   ./cmd/main
   ```

```
                                              ,           ^'^  _
                                              )               (_) ^'^
         _/\_                    .---------. ((        ^'^
         (('>                    )`'`'`'`'`( ||                 ^'^
    _    /^|                    /`'`'`'`'`'`\||           ^'^
    =>--/__|m---               /`'`'`'`'`'`'`\|
         ^^           ,,,,,,, /`'`'`'`'`'`'`'`\      ,
                     .-------.`|`````````````|`  .   )
                    / .^. .^. \|  ,^^, ,^^,  |  / \ ((
                   /  |_| |_|  \  |__| |__|  | /,-,\||
        _         /_____________\ |")| |  |  |/ |_| \|
       (")         |  __   __  |  '==' '=='  /_______\     _
      (' ')        | /  \ /  \ |   _______   |,^, ,^,|    (")
       \  \        | |--| |--| |  ((--.--))  ||_| |_||   (' ')
     _  ^^^ _      | |__| |("| |  ||  |  ||  |,-, ,-,|   /  /
   ,' ',  ,' ',    |           |  ||  |  ||  ||_| |_||   ^^^
.,,|xxx|,.|xxx|,.,,'==========='==''=='==''=='=======',,....,,,,.,

```
