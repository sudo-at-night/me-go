# About
My personal website's SSR code lives here. This server is used to return meta tags in the initial HTML which is crucial for tools like Open Graph.

> :warning: **DEPRECATED:** My personal website is now located at [sudo-at-night/me-next](https://github.com/sudo-at-night/me-next).  
> This repository is no longer in use.

# Run container
```bash
# Run dev container
docker-compose -f docker/compose.dev.yml up -d

# Run production container
docker-compose -f docker/compose.prod.yml up -d
```

# Build application
```bash
# Build the application into binary executable called "ssr"
./scripts/build-go.sh
```

# Build image
```bash
# Build docker image for later reuse
./scripts/build-image.sh
```
