<h1 align="center">Gotyl</h1>
<p align="center"><em></em></p>

<p align="center">
  Generate high-quality music cards with cover art, lyrics, and clean visual themes.
</p>

<p align="center">
  
  <a href="https://github.com/Chewawi/gotyl/graphs/commit-activity">
    <img alt="Maintenance"
      src="https://img.shields.io/badge/Maintained-Yes-6C3BAA?style=for-the-badge&logo=undertale&logoColor=E9E3F5&labelColor=1E1A2B" /></a>

  <a href="https://github.com/Chewawi/gotyl/stargazers">
    <img alt="Stars"
      src="https://img.shields.io/github/stars/Chewawi/gotyl?style=for-the-badge&logo=github&logoColor=E9E3F5&labelColor=1E1A2B&color=6C3BAA" /></a>

  <a href="https://github.com/Chewawi/gotyl/issues">
    <img alt="Issues"
      src="https://img.shields.io/github/issues/Chewawi/gotyl?style=for-the-badge&logo=github&logoColor=E9E3F5&labelColor=1E1A2B&color=6C3BAA" /></a>

  <a href="https://go.dev/">
  <img alt="Go version"
    src="https://img.shields.io/badge/Go-latest-6C3BAA?style=for-the-badge&logo=go&logoColor=E9E3F5&labelColor=1E1A2B" /></a>

  <a href="https://www.gnu.org/licenses/agpl-3.0.html">
  <img alt="License"
    src="https://img.shields.io/badge/License-AGPL--3.0-6C3BAA?style=for-the-badge&logo=gnu&logoColor=E9E3F5&labelColor=1E1A2B" /></a>
    
</p>


---

## Overview

**Gotyl** is a Go-based renderer and API for generating stylized music cards from
Spotify tracks and albums.

It focuses on:

- deterministic rendering  
- performance and caching  
- clean typography  
- minimal but expressive layouts  

Gotyl exposes both a **CLI** and an **HTTP API**.  
It is intended to be run as a tool or service — **not** consumed as a Go library.

---

## Features

- High-resolution poster rendering
- Album cover composition
- Lyrics layout with font fallback
- Spotify scannable codes
- Light and dark themes
- Aggressive in-memory caching
- CLI tool and HTTP API

---

## Quick start

Build the CLI:

```bash
go build ./cmd/gostyl
````

Run an example:

```bash
go run ./cmd/gostyl --query "PDA" --theme Light --out assets/examples/card_pda.png
```

---

## Example output

<p align="center">
  <img src="assets/examples/card_pda.png" alt="Light Theme Card" width="400" />
  <img src="assets/examples/card_dark.png" alt="Dark Theme Card" width="400" />
</p>

---

## Environment

Gotyl requires Spotify API credentials.

Copy the example environment file:

```bash
cp .env.example .env
```

Edit `.env` and provide the required values.

---

## API

Gotyl can also run as an HTTP service.

Start the API in development mode:

```bash
mage Dev
```

The API allows generating cards via query parameters such as:

* track search
* theme selection
* output format (PNG / JPEG)
* optional Spotify scannable codes

The API surface is intentionally small and evolving.

---

## Development (Mage)

This repository includes a `magefile.go` with common development tasks.

Install Mage:

```bash
go install github.com/magefile/mage@latest
```

Available targets:

* `mage` or `mage Build` — build binaries
* `mage Run` — run the API
* `mage Dev` — fast development mode
* `mage Test` — run tests
* `mage Clean` — clean build artifacts

Example:

```bash
cp .env.example .env
mage Dev
```

---

## Contributing

Contributions are welcome.

Good areas to contribute:

* new layouts or themes
* rendering optimizations
* API improvements
* typography and font handling

</br>

---

</br>

<p align="center">
Made with 💜 by Kodeko Studios<br>
<sub>
Inspired by (and a rewrite of) the original Python project
<a href="https://github.com/TrueMyst/BeatPrints/">BeatPrints</a>
</sub>
</p>