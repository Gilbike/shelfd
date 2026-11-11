![shelfd][cover_url]

<div align="center">
<strong>100% local, lightweight personal library</strong>

<hr/>

![GitHub Actions Workflow Status][workflow_shield_url]
![GitHub Release][release_shield_url]
![License][license_shield_url]

</div>

> [!WARNING]
> This is a 0.x project under heavy development. Functionality and design are subject to change.

## Getting started

### Requirements

The application is lightweight and compiles into a single binary with the frontend embedded.

- **Binary size:** ~12 MB
- **Memory footprint:** ~50 MB RAM at runtime

### Using Docker Compose

```yaml
services:
  shelfd:
    image: ghcr.io/gilbike/shelfd:latest
    container_name: shelfd
    restart: unless-stopped
    volumes:
      - shelfd_data:/app/data
    ports:
      - "8080:8080"
    environment: # See .env.example
      SHELFD_PORT: 8080
      SHELFD_LOG_FORMAT: json

volumes:
  shelfd_data:
```

### Using Docker

```bash
docker run --name shelfd -d -p 8080:8080 -v shelfd_data:/app/data ghcr.io/gilbike/shelfd:latest
```

### Building the image yourself

```bash
docker build -t shelfd:latest .
docker run --name shelfd -d -p 8080:8080 -v shelfd_data:/app/data shelfd:latest
```

### Standalone Binary

Download the latest executable for your platfrom from [Releases][releases_url] page.

```bash
# Linux (amd64)
./shelfd-linux-amd64

# Linux (arm64)
./shelfd-linux-arm64

# Window
./shelfd-windows-amd64.exe
```

## Configuration

Every configuration option and it's default value can be found in the [example env][example_env_url] file.

## Roadmap

- ✅ CRUD Book management
- 🏗️ Collections
- 🏗️ Progress tracking
- 🏗️ Customization
- 🏗️ i18n and themes
- 🏗️ Reviews and ratings
- 🏗️ Per instance social feed
- 🏗️ Mobile app

## Contributing

Everything you need to know is located in the [CONTRIBUTING.md][contributing_text_url] file.

## License

The project is licensed under the [MIT license][license_text_url]

[cover_url]: .github/assets/cover.jpg
[example_env_url]: https://github.com/Gilbike/shelfd/blob/master/.env.example
[workflow_shield_url]: https://img.shields.io/github/actions/workflow/status/Gilbike/shelfd/ci.yaml?style=flat-square
[release_shield_url]: https://img.shields.io/github/v/release/Gilbike/shelfd?display_name=tag&style=flat-square&link=https%3A%2F%2Fgithub.com%2FGilbike%2Fshelfd%2Freleases
[license_shield_url]: https://img.shields.io/github/license/Gilbike/shelfd?style=flat-square
[license_text_url]: https://github.com/Gilbike/shelfd/blob/master/LICENSE
[contributing_text_url]: https://github.com/Gilbike/shelfd/blob/master/CONTRIBUTING.md
[releases_url]: https://github.com/Gilbike/shelfd/releases
