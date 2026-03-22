# EmailN

Project to manage email campaigns.

## Prerequisites

- Go 1.25+

## Installation

```bash
go mod download
```

## Run Tests

```bash
go test ./...
```

## Project Structure

```
internal/
├── contract/          # DTOs and contracts
│   └── NewCampaign.go
└── domain/
    └── campaign/      # Business logic
        ├── campaign.go
        ├── repository.go
        └── service.go
```

## Features

- Create email campaigns with name, content, and contact list
- Validation for required fields
- Repository pattern for data persistence (interface)
