<p align="center">
      <img src="https://github.com/gaelgoth/swiss-deals-frontend/raw/main/src/assets/images/logos/logo.png" alt="swiss-deal-logo" width="60px" height="auto">

</p>

<h1 align="center">
      Swiss Deals API - Online swiss shop deals aggregation
</h1>

## Introduction

This API aggregates the deals of the day on the following online sites:

- Digitec
- Galaxus
- Qoqa
  - Qbeer
  - Qlock
  - Qoqa
  - Qsport
  - Qwine
- ...

## Installation ⚒️

> Compatible with Golang 1.19

Run development server with `Makefile`

   ```bash
   # auto reload on file modification 
   make watch
   ```

OR

Run development server with `Dockerfile`

```bash
# auto reload on file modification with air
docker-compose -f docker-compose.dev.yml up   
```

## Documentation 📖

- Swagger UI: <https://deal-backend-production.up.railway.app/swagger/index.html>

## Contributing 🦸

 Feel free to contribute via Pull Request.

## Changelog 📆

Please refer to the [CHANGELOG](CHANGELOG.md) file.

## Credits 🙏

- [Fiber](https://github.com/gofiber/fiber)
- [Colly](https://github.com/gocolly/colly)
