# E-Corp

E-Corp is a tool to bootstrap new projects with a predefined structure and configuration. It helps developers quickly set up a new project with some opinionated defaults in mind.

The goal is after you run the initialization script, you will have a fully functional project structure with all the necessary files and folders in place, so you can start developing your application and focusing on business logic right away without worrying about setting up the basic structure.

## Initialize a new project

You can initialize a new project by running the following command in your terminal:

```bash
wget -qO- https://raw.githubusercontent.com/anonychun/ecorp/refs/heads/main/new.sh | bash -s <project-name>
```

Replace `<project-name>` with the desired name for your new project. This command will create a new directory with the specified project name and set up the necessary files and folders.

The new directory will be created in the current working directory with the basename of the project name you provided. For example, if you run the command with `github.com/anonychun/verification-api`, a new directory named `verification-api` will be created.

## Project structure

The initialized project will have the following structure:

- **`bin`** - Scripts for various development and deployment tasks.
- **`cmd`** - Main application entry points and CLI commands.
  - **`server`** - HTTP server application.
  - **`db`** - Database management CLI.
  - **`generate`** - Code generation utilities.
- **`migrations`** - Database migration files.
- **`internal`** - Internal application code.
  - **`api`** - HTTP API utilities.
  - **`app`** - Application layer with business logic (use cases and handlers).
  - **`bootstrap`** - Coordination of application dependencies.
  - **`config`** - Configuration management with environment variable loading.
  - **`consts`** - Application constants.
  - **`current`** - Context utilities for request-scoped data.
  - **`db`** - Database layer.
  - **`entity`** - Database models and business entities.
  - **`middleware`** - HTTP middleware.
  - **`repository`** - Data access layer with database operations.
  - **`server`** - HTTP server setup and routing configuration.

## Usage

After initializing a new project, navigate to the project directory:

```bash
cd <project-name>
```

Fill the environment variables in the `.env` file as needed.

### Development

To start the development server with hot-reloading, run:

```bash
./bin/dev
```

This will start the server and watch for file changes, automatically restarting the server when changes are detected.

### Generate code

To generate code you can use the generator command in the `cmd/generate` package.

#### Migration

To create a new database migration, run:

```bash
go run cmd/generate/main.go migration <migration-name> <migration-type>
```

The `<migration-type>` can be either `sql` or `go`, the default is `sql`.

#### App

To generate a new application component, run:

```bash
go run cmd/generate/main.go app <component-name>
```

E.g., `go run cmd/generate/main.go app api/v1/user`.

#### Repository

To generate a new repository, run:

```bash
go run cmd/generate/main.go repository <repository-name>
```

E.g., `go run cmd/generate/main.go repository user`.

### Database

You can manage your database using the provided CLI commands.

#### Create database

To create the database, run:

```bash
go run cmd/db/main.go create
```

#### Drop database

To drop the database, run:

```bash
go run cmd/db/main.go drop
```

#### Migrate database

To migrate the database, run:

```bash
go run cmd/db/main.go migrate
```

#### Rollback database

To rollback the last applied migration, run:

```bash
go run cmd/db/main.go rollback
```

#### Seed database

To seed the database with initial data, run:

```bash
go run cmd/db/main.go seed
```

#### Setup database

To set up the database (create, migrate, and seed), run:

```bash
go run cmd/db/main.go setup
```

#### Reset database

To reset the database (drop, create, migrate, and seed), run:

```bash
go run cmd/db/main.go reset
```

### Server

To start the HTTP server, run:

```bash
go run cmd/server/main.go start
```

### Transaction

To execute a function within a database transaction in the use case layer, you can use the `repository.Transaction` function. Here's an example:

```go
func (u *Usecase) Delete(ctx context.Context, req DeleteRequest) error {
	exists, err := u.repository.Admin.ExistsById(ctx, req.Id)
	if err != nil {
		return err
	}

	if !exists {
		return consts.ErrAdminNotFound
	}

	return u.repository.Transaction(ctx, func(ctx context.Context) error {
		err := u.repository.AdminSession.DeleteAllByAdminId(ctx, req.Id)
		if err != nil {
			return err
		}

		return u.repository.Admin.DeleteById(ctx, req.Id)
	})
}
```

### Current

Current is a package that provides utilities for managing request-scoped data using context. It allows you to set and get values associated with the current request, such as user information or request ID.

Here's an example of how to use the `current` package:

```go
// Setting a value in the context
user := &entity.User{ID: 1, Name: "Achun"}
ctx = current.SetUser(ctx, user)

// Getting a value from the context
user := current.User(ctx)
```

## Starter kit

E-Corp comes with default starter kit to help you get started quickly.

- **Admin management**: Includes admin authentication and CRUD operations.

You can explore the generated code, adapt it to your project's requirements, and use it to become familiar with the structure and conventions.
