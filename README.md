# social-network-common-service
## Project Summary
This is project about social network that allows users to share content, images, and emotions, and have real-time communication capabilities, while ensuring high performance, security, and scalability using the microservices architecture.

#### Technologies:
- Back-end:
  - Language: Go.
  - Frameworks/Platforms: Gin-Gonic, gRPC, Swagger, JWT, Google-Wire, SQLX, Redis, Zap, WebSocket.
  - Database: MariaDB.
- Front-end:
  - Language: JavaScript.
  - Frameworks/Platforms: React, Tailwind CSS, FireBase.

## The project includes repositories
- [common-service](https://github.com/nhutHao02/social-network-common-service)
- [user-service](https://github.com/nhutHao02/social-network-user-service)
- [tweet-service](https://github.com/nhutHao02/social-network-tweet-service)
- [chat-service](https://github.com/nhutHao02/social-network-chat-service)
- [notification-service](https://github.com/nhutHao02/social-network-notification-service)
- [Front-end-service (in progress)](https://github.com/nhutHao02/)

## This service
This is the service that provides ultils, validators, and middleware for services.

## Project structure
```
.
├── go.mod
├── go.sum
├── main.go
├── middleware
│   └── middleware.go
├── model
│   ├── paging.go
│   └── response.go
├── rabbitmq
│   └── rabbitmq.go
├── README.md
├── request
│   └── request.go
├── utils
│   ├── constanst
│   │   └── constants.go
│   ├── error
│   │   └── error.go
│   ├── logger
│   │   └── logger.go
│   └── token
│       └── token.go
└── validation
    └── validation.go
```

