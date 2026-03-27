**This repo is a refractored version of my CRUD API, converted to clean architecture.**




**Backend CRUD Application (Go + Gin + GORM)**


This project is a backend CRUD application built using Go, Gin, and GORM. The goal of this project was not just to make a basic app work, but to understand how real backend applications are structured and how different components work together.

I focused on building the project step by step and improving it over time instead of rushing through it.

**What this project includes:**


1.User registration and login

2.Password hashing and authentication logic

3.CRUD operations using GORM

4.Proper database models and relationships

5.Structured logging

6.Environment variable configuration using .env

7.Clean and organized project structure


**Tech Stack**


1.Go (Golang)

2.Gin (HTTP framework)

3.GORM (ORM for database operations)

4.PostgreSQL

5.Zap Logger

7.dotenv

**Architecture**

This project follows the Clean Architecture pattern to keep the code modular, testable, and easy to maintain as the application grows.

The application is divided into clear layers, where each layer has a single responsibility and depends only on the layer below it.

Handlers  →  Services  →  Repository  →  Database

for testing purposes