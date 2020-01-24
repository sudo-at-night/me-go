# About
My personal website's SSR code lives here. This server is used to return meta tags in the initial HTML which is crucial for tools like Open Graph.

# Run container
```bash
# Run dev container
./scripts/reload-container.sh

# Run production container
./scripts/reload-container.sh --production
```

# Build application
```bash
# Build the application into binary executable called "ssr"
./scripts/build.sh
```