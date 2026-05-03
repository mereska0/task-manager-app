# Task Manager App
**THIS PROJECT WAS BUILT FOR EDUCATIONAL PURPOSES**

A simple web-based task manager built with Go, PostgreSQL, and a vanilla JavaScript frontend.

## Overview

This repository contains a small educational project that provides:

- REST API for task management (`GET`, `POST`, `PUT`, `DELETE`)
- PostgreSQL persistence for tasks
- Browser UI served from static files
- Middleware for logging and error recovery

## Features

- Add new tasks
- Toggle task completion
- Edit task text
- Delete tasks
- Automatically creates the `tasks` table on startup

## Architecture

- `main.go` — application entrypoint, router setup, middleware, and static file serving
- `handler/manager.go` — HTTP handlers for task endpoints
- `service/manager.go` — business logic and database operations
- `model/manager.go` — task model definition
- `storage/postgres.go` — PostgreSQL connection and schema initialization
- `app.js` — frontend logic for task CRUD operations
- `index.html` / `style.css` — UI markup and styles

## Requirements

- Go 1.26+
- PostgreSQL


## Setup

1. Start PostgreSQL and ensure it is accessible.
2. Create the `tmanager` database if it does not already exist:

```bash
createdb -h localhost -p 5432 -U postgres tmanager
```

3. Run the application from the project root:

```bash
go run ./...
```

The server listens on `http://localhost:8080`.

## Usage

- Open `http://localhost:8080` in your browser.
- Use the form to add tasks.
- Click the checkmark to mark tasks completed.
- Edit tasks using the edit button.
- Delete tasks with the trash button.

## API Endpoints

- `GET /tasks/` — list all tasks
- `POST /tasks/` — add a task; body: `{ "task": "Task text" }`
- `PUT /tasks/{id}` — update a task; body: `{ "changes": true }` or `{ "changes": "new text" }`
- `DELETE /tasks/{id}` — delete a task

## Notes

- Static frontend files are served from the project root.
- The database table is created automatically on startup if missing.
- Error responses are returned as JSON.
