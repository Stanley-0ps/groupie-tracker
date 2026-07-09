## Groupie Tracker 

This project is a web application that tracks musical artists, their locations, and dates of events. It serves data from an external API and displays artist information using Mapbox GL JS integration for locations.

#### How to use

```bash
git clone https://github.com/Stanley-0ps/groupie-tracker.git
```

```bash
go run .
```

The project contains two main components:

- **The server**: responsible for fetching data from the api and serving the web application.
- **The API**: responsible for fetching and serving artist-related data from the external source.

> Setting ports (to use a specific port):

```bash
export PORT=<port number>
export APIPORT=<api port number>"
```

#### Mapbox access token

The artist page now uses Mapbox GL JS to render tour locations. Create a free Mapbox account, generate a **Default public token**, then either place it in `.env` (automatically sourced by every `make` target via `set -a`) or export it manually before running the app:

```bash
MAPBOX_ACCESS_TOKEN=<your-public-token>   # inside .env
# or
export MAPBOX_ACCESS_TOKEN=<your-public-token>
```

Without a token the embedded map is replaced with a short warning message.

#### SerpAPI key

City images are pulled from SerpAPI. Add your SerpAPI token to `.env` (or export it) so the image fetcher can call the API:

```bash
SERPAPI_KEY=<your-serp-api-key>   # inside .env
# or
export SERPAPI_KEY=<your-serp-api-key>
```

When the key is missing, location image lookups are skipped and fall back to placeholders.

#### Mapbox Integration
This project integrates Mapbox GL JS to display artist locations on an interactive map. The locations are fetched from an external API and are dynamically displayed on the map as markers, representing various places where the artists have performed or will perform in the future.
