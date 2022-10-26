<p align="center">
  <a href="https://runemarkers.net">
    <picture>
      <img src="./public/logo-128-background.png" height="128">
    </picture>
    <h1 align="center">RuneMarkers.net</h1>
  </a>
</p>

## Tiles Contribution

If you wish to contribute to tiles, follow the instructions below.

- Fork the repo
- Create a new branch with a name like so `tiles/<tile name>` e.g. `tiles/vorkath-melee`
- Add the tiles to the 'tiles' folder following the format used in other tiles (look at abyssal sire for an example)
- Commit the changes and push to the repo
- Open a pull request with a description of the changes
- Wait for the pull request to be accepted

## Development

To start the project locally, run:

```bash
yarn dev
```

Open `http://localhost:3000` with your browser to see the result.

### Requirements

- Node.js >= 12.22.0
- Yarn 1 (Classic)

### Scripts

- `yarn dev` — Starts the application in development mode at `http://localhost:3000`.
- `yarn build` — Creates an optimized production build of your application.
- `yarn start` — Starts the application in production mode.
- `yarn type-check` — Validate code using TypeScript compiler.
- `yarn lint` — Runs ESLint for all files in the `src` directory.
- `yarn format` — Runs Prettier for all files in the `src` directory.
- `yarn commit` — Run commitizen. Alternative to `git commit`.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for more information.
