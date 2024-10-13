# Bravo Platform - Backend Service

Welcome to the **Bravo Platform** backend service! This repository contains the backend system that powers the Bravo platform—an innovative educational solution designed for schools and universities.

## 🌟 Introduction

Bravo is a comprehensive education platform aimed at revolutionizing the learning experience for both students and teachers. It provides tools that make learning more engaging and accessible, helping to create a better educational environment.

## 🎯 Our Mission

Our mission is to deliver a robust application that enhances the learning journey by making it easier for students to access educational content and for teachers to manage their curriculum. We strive to improve the overall experience of education through technology, creating a seamless platform for all.

## 🛠️ Development Guidelines

Here are the key commands to help you develop and maintain the backend service:

### Running the Service
```bash
go run main.go # to run the service
```

```bash
make migrate_up # to up the apply migration record
```

```bash
make migrate_down # to rollback the migration record
```

```bash
make migrate_force # to force update the previous migration record (Not recommend, please use create_migration instead)
```

```bash
make commit # to commit your code change
```

```bash
make push # to push your change on your current Git branch
```


