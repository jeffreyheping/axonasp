# Install and Run AxonASP with Docker

## Overview
This page explains how to install and run the Docker distribution of G3Pix AxonASP.

Use the official container image:
- `ghcr.io/guimaraeslucas/axonasp`

## Syntax
Pull image:

```bash
docker pull ghcr.io/guimaraeslucas/axonasp
```

Run container:

```bash
docker run --name axonasp -p 8801:8801 ghcr.io/guimaraeslucas/axonasp
```

## Parameters and Arguments
Common Docker command arguments used in AxonASP deployment:
- `--name`:
  - Type: String
  - Required: No
  - Purpose: Assigns a fixed container name.
- `-p host_port:container_port`:
  - Type: Port mapping
  - Required: Yes for host access
  - Purpose: Exposes AxonASP HTTP server port to the host.
- `-d`:
  - Type: Flag
  - Required: No
  - Purpose: Runs container in detached mode.
- `-v host_path:container_path`:
  - Type: Volume mapping
  - Required: Recommended
  - Purpose: Persists and customizes `www`, `config`, and other runtime data.
- `--restart unless-stopped`:
  - Type: Restart policy
  - Required: No
  - Purpose: Keeps service available after host reboot.

## Return Values
- `docker pull` returns success when the image is downloaded and available locally.
- `docker run` returns a running container ID or an immediate error message if configuration is invalid.
- Exit code `0` indicates command success.

## Remarks
- Expose port `8801` for HTTP mode unless your configuration changes server port.
- Mount your `www` and `config` directories if you need custom applications and configuration.
- For production deployment, use a reverse proxy in front of the container.
- You can stop and remove the container safely without removing the image.

## Code Example
Basic installation and run:

```bash
# Pull latest image from GHCR
docker pull ghcr.io/guimaraeslucas/axonasp

# Run AxonASP in detached mode with host port mapping
docker run -d --name axonasp -p 8801:8801 --restart unless-stopped ghcr.io/guimaraeslucas/axonasp

# Check container status
docker ps
```

Run with mounted application and config directories:

```bash
docker run -d \
  --name axonasp \
  -p 8801:8801 \
  -v /opt/axonasp/www:/opt/axonasp/www \
  -v /opt/axonasp/config:/opt/axonasp/config \
  --restart unless-stopped \
  ghcr.io/guimaraeslucas/axonasp
```