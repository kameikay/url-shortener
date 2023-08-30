# Welcome to the URL Shortner 🪄

![Repository Banner](images/gitbg.png)

## Overview 💡

This is a code challange by [Devgym](https://app.devgym.com.br/) to create a URL shortner.

## How to play with docker-compose and air 🐳 🎮

To play the app with docker-compose, you can clone this repository and run the following commands:

```bash
# Clone the repository
git clone

# Enter the repository folder
cd url-shortner

# Run the Makefile to run the docker-compose with air
make air
```

## Requirements 

In this challenge you must create a server that shortens urls and redirects.

-- Create an http server that contains two endpoints:
- POST / - takes a url and returns a unique code
- GET /:code - use the code to redirect to the original url

-- The code is a unique code, the same url sent multiple times generates different codes
-- The code is up to 6 characters long

## Study Suggestions

If you're eager to dive into the technologies that power this app, I recommend exploring the following:

Go (Golang): Golang is an excellent language for building fast and efficient server-side applications. Its simplicity and performance make it a great choice for this project.

PostgreSQL: For a robust and scalable database solution, consider delving into PostgreSQL. It offers advanced features and reliability, perfect for storing the data behind this app.
