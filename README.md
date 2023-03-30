# Shelt

## Develop:

To develop shelt locally, you need:

- [NodeJS](https://nodejs.org/en/download)
- [Go](https://go.dev/learn/)

**To run the dev environment locally, all three components need to run:**


*Strapi (Port 1337)*
- `cd ./shelt-db`
- `npm install` (when running for the first time)
- run via VSCode F5 or `npm run develop`

*Go (Port 8080)*
- `cd ./shelt-backend`
- `go mod tidy` (when running for the first time)
- run via VSCode F5 or `go run .`

**Rendering Route Examples (localhost:8080...)**

|||
|-|-|
| render animal[3] with template[1] of all animals | /animal-view?animal-id=3&template-id=1 |
| render list with list-template[2] of all animals | /animal-collection-view?template-id=2 |
| ..of animals with Tags "tag1" or "tag2" | /animal-collection-view?template-id=2&tags=tag1,tag2 |
