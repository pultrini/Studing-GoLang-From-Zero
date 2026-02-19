# 🐹 Studying GoLang From Zero
This repository documents my journey learning Go (Golang), following the curriculum developed by Stephany Henrique Batista. The goal is to move from fundamental syntax to building a production-ready REST API with modern architecture and security standards.
## 📌 About the Studie
**Focus**: High-performance systems, simplicity, and industry-standard tools.
**Core Goal**: Transitioning from basic syntax (Go 1.19+) to complex implementations including Concurrency, Unit Testing, and Identity Providers.

## 🚀 Learning Roadmap
### 1. Language & Syntax (The Foundation)
Understanding the core mechanics and the simplicity of Go's design.
- [X] Fundamentals: Variables, Types, and Functions (func).
- [ ] Data Structures: struct, Interface, and Generics.
- [ ] Memory Management: Pointers and robust Error Handling.
- [ ] Ecosystem: Dependency management with go mod.
- [ ] Concurrency: 
  - Managing state with sync.WaitGroup and sync.Mutex.Communication via Channels and Goroutines.
### 2. Practical Project: Professional REST API
Building a scalable application using the Domain/Service/Infrastructure architectural pattern.
- [ ] Web Layer: Implementing the Chi Router for RESTful services.
- [ ] Persistence: Deep dive into PostgreSQL.
- [ ] Security: Authentication and Authorization using Keycloak.
- [ ] Messaging: Handling Queues (Async Calls).
- [ ] Quality Assurance: Extensive Unit Testing using Testify.
### 🛠️ Tech Stack
| Category | Technology |
| :--- | :--- |
| **Language** | Go (v1.26+) |
| **Framework** | Chi (REST API) |
| **Database** | PostgreSQL |
| **Identity** | Keycloak |
| **Testing** | Testify |

### 📂 Repository StructureBash.
├── 01 - Basic/       
├── 02 - Functions/ 
├── ...
├── 0n - Final Studies/
├── api-project/        # Main project implementation
│   ├── cmd/            # Entry points
│   ├── internal/       # Domain, Services, and Infrastructure (Clean Arch)
│   └── tests/          # Unit and integration tests
└── README.md


✨ __"Go is not meant to innovate; it is meant to be productive"__