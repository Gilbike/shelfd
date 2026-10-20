# Shelfd

> [!WARNING]
> This is a 0.x project under heavy development. Functionality and design are subject to change.

## Getting started

### Building the image yourself

```bash
docker build -t shelfd:latest .
docker run --name shelfd -d -p 8080:8080 -v shelfd-data:/app/data shelfd:latest
```
