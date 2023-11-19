# Vocabulary

Vocabulary is a web application designed for language learners. It allows users to store and manage vocabularies learned during their classes. Users can add, update, or delete vocabularies and associate them with categories for better organization.

## Features

- Add new vocabularies with the following fields:
    - **Words**: The vocabulary word(s).
    - **Translation**: The translation of the word(s).
    - **Used in Phrase**: An example of how the vocabulary is used in a sentence.
    - **Explanation**: Additional details or notes about the vocabulary.
    - **Categories**: Associate vocabularies with categories (users can create or search categories dynamically).
- View a list of vocabularies on the main screen:
    - Each row displays the vocabulary and an example sentence.
    - Options to update or delete each vocabulary.
- Clean architecture design:
    - Business logic is isolated.
    - Repository interface for database communication.
- Easily configurable with multiple ` + "`main.go`" + ` files demonstrating integration with various database systems, loggers, and servers.
- Quick deployment with Docker Compose.

## Installation

### Prerequisites

- Docker and Docker Compose installed on your machine.
- Golang and Node.js (optional, for local development).

### Quick Start

1. Clone the repository:
   ```bash
   git clone https://github.com/axeltaglia/vocabulary.git
   cd vocabulary
2. Start the application using Docker Compose:
   ```bash
   docker-compose up -d

Access the application in your browser at http://localhost:<configured-port>.

## Project Structure

### Backend
The backend is built using Golang with clean architecture principles. The application uses Gorm for database interactions and supports the following main functionalities:

- Vocabulary management (VocabularyEntity)
- Category management
- Database communication through a repository interface (gormRepository).

### Frontend
The frontend is built using React.js and provides an intuitive interface for:

- Adding vocabularies via a form in a popup.
- Viewing, updating, and deleting vocabularies in a table format.
- Searching and managing categories dynamically.

### Database
The database is managed using PostgreSQL and includes the following entities:

- Vocabulary: Represents the vocabulary data.
- Category: Represents categories associated with vocabularies.

### Configuration
The application uses a conf.json file for configurations, including:

- Database credentials
- API port

```bash
{
  "DbConfig": {
    "Host": "localhost",
    "Port": 5432,
    "DbName": "vocabulary",
    "User": "admin",
    "Password": "password"
  },
  "ApiPort": 8080
}
```

### Architecture Highlights
- Clean Architecture: The business logic is decoupled from external systems.
- Repository Pattern: Abstracts database operations via interfaces.
- Configurable Main Files: Demonstrates flexibility in connecting with different systems.
- Dockerized Deployment: Simplifies the setup and deployment process.
- 