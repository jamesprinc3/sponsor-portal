# DoCSoc Sponsor Portal

## Dependencies
 - Go 1.8+
 - [Glide](https://github.com/Masterminds/glide#install)
 - [Migrate](https://github.com/mattes/migrate#cli-usage)
 - Docker
 - npm
 - [yarn](https://yarnpkg.com/en/docs/install)

## Build & Run
 - **Ensure you have the dependencies listed above**
 - `make install` to install npm and go packages
 - `docker-compose up -d` and `make setup` to start docker and migrate/seed the db
 
 - `make client` to build the front-end assets for production
 - `make client-dev` to build the front-end assets for development
 - `make watch` to build the front-end assets for development and watch for changes (recommended)

 - `make server` to build the server
 - `make run` to start the server

## Example sponsor

 - Email: `ada@sponsor.com`
 - Password: `password`
