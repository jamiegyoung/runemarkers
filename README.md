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

To install the dependencies, run:

```bash
go mod download
```

To start the project locally for development, run:

```bash
go run ./cmd/server/
```

To build the project, run:

```bash
go run ./cmd/build/
```

Open `http://localhost:8080` with your browser to see the result.


### Requirements

- Go >= 1.22

### Scripts

- `./cmd/server/` — Starts the application in development mode at `http://localhost:8080`.
- `./cmd/build/` — Creates an optimized production build of your application.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for more information.

## Contributors

<a href="https://github.com/jamiegyoung/runemarkers/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=jamiegyoung/runemarkers" />
</a>

Made with [contrib.rocks](https://contrib.rocks).

