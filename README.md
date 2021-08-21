# web2pdf: any web page to PDF document

Print web pages to PDF documents using Chromium-based browsers.

## Installation instructions

```
go get -u github.com/erdnaxe/web2pdf
```

This will get and compile web2pdf. By default the executable will be in
`~/go/bin/web2pdf`.

## Usage examples

### Print all unread entries from Miniflux

[Miniflux](https://miniflux.app/) is a feed reader. Using the API, we can
print each unread entry to a PDF document to generate an offline newspaper.
Replace `username` and `miniflux.example.com` with your Miniflux username
and host.

```bash
curl -su username "https://miniflux.example.com/v1/entries?status=unread" | jq -r '.entries[] | "web2pdf --output=\(.id).pdf \"\(.url)\" # \(.feed.category.title) - \(.feed.title) - \(.title)"'
```
