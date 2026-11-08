![shelfd][cover_url]

100% local, lightweight personal library

> [!WARNING]
> This is a 0.x project under heavy development. Functionality and design are subject to change.

## Getting started

### Building the image yourself

```bash
docker build -t shelfd:latest .
docker run --name shelfd -d -p 8080:8080 -v shelfd-data:/app/data shelfd:latest
```

## Roadmap

- ✅ CRUD Book management
- 🏗️ Collections
- 🏗️ Progress tracking
- 🏗️ Customization
- 🏗️ i18n and themes
- 🏗️ Reviews and ratings
- 🏗️ Per instance social feed
- 🏗️ Mobile app

[cover_url]: https://raw.githubusercontent.com/Gilbike/shelfd/refs/heads/master/.github/assets/cover.jpg
