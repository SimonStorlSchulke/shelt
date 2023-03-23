# Shelt

## Develop:

To develop shelt locally, you need:

- [NodeJS](https://nodejs.org/en/download)
- [Go](https://go.dev/learn/)
- angular `npm install -g @angular/cli`


**To run the dev environment locally, all three components need to run:**


*Strapi (Port 1337)*
- `cd ./shelt-db`
- `npm install` (when running for the first time)
- `npm run develop`

*Go (Port 8080)*
- `cd ./shelt-backend`
- `go mod tidy` (when running for the first time)


*Angular (Port 4200)*
- `cd ./shelt-frontend`
- `npm install` (when running for the first time)
- `ng serve --open` <- opens a browser window woth the frontend
