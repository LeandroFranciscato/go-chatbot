# go-chatbot

![Build](https://github.com/LeandroFranciscato/go-chatbot/actions/workflows/go.yml/badge.svg)

This project aims to create a chatbot that can interact with users in a meaningful and engaging way.

## Prerequisites

This project has a few prerequisites:

- Go 1.18+: [official Go website](https://golang.org/dl/).

- MongoDB: This project uses MongoDB as its primary database.[official MongoDB website](https://www.mongodb.com/try/download/community).

## Project Goals

The main goal of this project is to develop a chatbot that can understand and respond to user inputs.  
The chatbot will be specifically designed to facilitate the review process after a client confirms a delivery.  
It will interact with the client, asking for feedback and rating, and then store this information.

## Getting Started

To get started with this project, clone the repository and install the necessary dependencies:

```bash
git clone https://github.com/LeandroFranciscato/go-chatbot
cd go-chatbot
go mod tidy
go run main.go
```

Set the environment variable `MIGRATE` to have a initial mock data for tests purpose.

e.g.:

```bash
MIGRATE=1 go run main.go
```

## Tools

### build and run

```bash
make run
```

### generating mocks

```bash
make mocks
```

### genarating testing report

```bash
make test
```

Then open `coverage.html`.
