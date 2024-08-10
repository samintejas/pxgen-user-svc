# Project organisation


## Directory Descriptions

### `/cmd`
Contains the main entry points for the application.
- **`/app`**: Main application executable.
- **`/worker`**: Secondary service or worker executable.

### `/internal`
Holds the core logic of the application.
- **`/config`**: Configuration management.
- **`/db`**: Database interactions.
- **`/models`**: Data structures or models.
- **`/services`**: Business logic and service layer.
- **`/utils`**: Utility functions used across the application.

### `/pkg`
Contains reusable code and libraries.
- **`/api`**: API definitions and handlers.
- **`/cache`**: Caching mechanisms.

### `/scripts`
Deployment, setup, and management scripts.

### `/docs`
Documentation related to the project.
- **`architecture.md`**: Project architecture documentation.

### `/tests`
Contains tests for the application.
- **`/integration`**: Integration tests.
- **`/unit`**: Unit tests.

### `go.mod` & `go.sum`
Go module files for dependency management.
