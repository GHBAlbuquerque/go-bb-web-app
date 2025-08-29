# go-bb-web-app
## Bed & Breakfast

A full-stack web application built with Go that allows guests to book rooms at a Bed & Breakfast and enables administrators to manage reservations through a secure back-office interface.

## Key Features
• Showcase the property
• Allow for booking a room for one or more nights
• Check a room’s availability
• Book the room
• Notify guest, and notify property owner

## Project Structure
main.go - main file, starts server
routes.go - router with all paths in the app

## Libs used

- [chi router](https://github.com/go-chi/chi/v5)
- [scs session management](https://github.com/alexedwards/scs/v2)
- [nosurf cookie management](https://github.com/justinas/nosurf )

## Running the project
go run cmd/web/*.go        

