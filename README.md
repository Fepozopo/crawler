# Web Crawler

This is a simple web crawler built in Go. It crawls a specified website, extracts all internal links, and generates a report listing the pages and the number of internal links found to each page.

## Prerequisites

- Go installed on your machine.
- Access to the terminal or command line.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Fepozopo/crawler
   cd crawler
   ```

2. Build the project:
   ```bash
   go build
   ```

## Usage

Run the web crawler with the following command:

```bash
./crawler <website> [concurrency] [maxPages]
```

- `<website>`: The URL of the website you want to crawl.
- `[concurrency]`: (Optional) The number of concurrent requests. Default is 4.
- `[maxPages]`: (Optional) The maximum number of pages to crawl. Default is unlimited (0).

### Example

```bash
./crawler http://example.com 3 25
```

This command will crawl `http://example.com` with a concurrency of 3 requests at a time, and stop after 25 pages.

## Features

- Extracts and normalizes URLs from `<a>` tags.
- Handles relative URLs based on the provided base URL.
- Concurrency control for efficient crawling.
- Limits the crawl to internal links within the same domain.
- Generates a report listing pages and their internal link counts.

## Testing

To run the tests, use the following command:

```bash
go test ./...
```
