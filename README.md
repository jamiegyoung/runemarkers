<p align="center">
  <a href="https://runemarkers.net">
    <img src="./assets/logo-256-background.png" height="256">
  </a>
  <h1 align="center">RuneMarkers</h1>
</p>

<p align="center">
<a href="https://runemarkers.net">https://runemarkers.net</a>
</p>

## Tiles

You can either contribute tiles directly by following the guide below, or create an issue with the `Tiles Request` template

### Tiles Contribution

If you wish to contribute to tiles, follow the instructions below.

- Fork the repo
- Create a new branch with a name like so `tiles/<tile name>` e.g. `tiles/vorkath-melee`
- Add the tiles to the 'entities' folder following the format used in other tiles (look at abyssal sire for an example)
- Commit the changes and push to the repo
- Open a pull request with a description of the changes
- Wait for the pull request to be accepted

## Development

### Using Docker (recommended)

To start the project locally using Docker, run:

```bash
docker-compose up
# or
docker compose up
# or if you want to run in the background
docker compose up -d
```

This will run a development server at `http://localhost:8080` with hot reloading enabled.

### Locally

To start the project locally, run:

```bash
pnpm dev
```

Open `http://localhost:8080` with your browser to see the result.


### Requirements

- Go >= 1.22

### Scripts

- `pnpm dev` — Starts the application in development mode at `http://localhost:3000`.
- `pnpm build` — Creates an optimized production build of your application.
- `pnpm start` — Starts the application in production mode.
- `pnpm type-check` — Validate code using TypeScript compiler.
- `pnpm lint` — Runs ESLint for all files in the `src` directory.
- `pnpm format` — Runs Prettier for all files in the `src` directory.
- `pnpm commit` — Run commitizen. Alternative to `git commit`.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for more information.

## Contributors

<a href="https://github.com/jamiegyoung/runemarkers/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=jamiegyoung/runemarkers" />
</a>

Made with [contrib.rocks](https://contrib.rocks).





# runemarkers-go
A rewrite of RuneMarkers using a custom SSG in go

## Development

To start the development server run the following:

`go run ./cmd/server`

This will start a server on port `8080` and rebuild when any changes are made

## Build

To build the project run the following:

`go run ./cmd/build`
