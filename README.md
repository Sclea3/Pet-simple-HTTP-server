# Pet-project: Simple HTTP (maybe gRPC in the future) Server with Docker

## Overview

Implement a simple HTTP server in Go for managing users. The main logic of the application includes user creation and task management, such as entering a referral code, subscribing to Telegram or Twitter, and receiving points as a reward. You can customize rewards for tasks and add other tasks based on your imagination.

## Required Functionality

### Core Features

- **Authorization Middleware** using Access tokens for all endpoints (e.g., JWT)
- **HTTP API** implementation:
    - `GET /users/{id}/status` – Fetch all available information about a user
    - `GET /users/leaderboard` – Get a leaderboard of top users with the highest balance
    - `POST /users/{id}/task/complete` – Mark a task as completed
    - `POST /users/{id}/referrer` – Input a referral code (could be another user’s id)
- **Data Storage** using PostgreSQL for all user-related information (with migration support via `golang-migrate`)
- **Docker Compose** for project setup and execution

### Additional Requirements

- Proper **error handling** (e.g., invalid data, non-existent users, etc.)
- Follow **SOLID principles** for clean and maintainable code

## Project Structure

The project should follow a clean, standard Go project layout:

### `/cmd`
- Contains the main application entry point (`main.go`)

### `/internal`
- Internal code and libraries that won’t be used outside the application

### `/pkg`
- Not used in this project since it's not expected to be shared externally

### `/api`
- OpenAPI/Swagger specifications, JSON schema, protocol definitions

### `/configs`
- Configuration files and default settings

### `/build`
- Contains CI/CD configurations (`/build/ci`), Docker setup (`/build/package`)

### `/deploy`
- Setup files for Kubernetes and docker-compose

### `/test`
- Unit tests for the application

## Example Source Layout

This layout is based on the **[Go Project Layout](https://github.com/golang-standards/project-layout)**