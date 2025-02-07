# Geass
web crawler for you

## Usage

### Instalization manual
```bash
git clone github/osamikoyo/geass
cd geass
task run
```
### Docker
```bash
git clone github/osamikoyo/geass
task docker-build
task docker-run
```

### Handler
```bash
curl "localhost:PORT_IN_CONFIG/get/content?url=EXEMPLE_URL"
```
### output:
```json
{
  "url": "https://example.com",
  "title": "Example Domain",
  "meta_description": "This domain is for use in illustrative examples in documents.",
  "content": {
    "headings": ["Example Domain", "This domain is for use..."],
    "text": "This domain is for use in illustrative examples in documents...",
    "images": [
      {
        "src": "https://example.com/image.png",
        "alt": "Example Image"
      }
    ]
  },
  "links": [
    {
      "text": "More information...",
      "href": "https://www.iana.org/domains/example"
    }
  ],
  "technical": {
    "status_code": 200,
    "content_type": "text/html",
    "last_modified": "2023-10-01T12:00:00Z"
  },
  "metadata": {
    "language": "en",
    "canonical": "https://example.com",
    "robots": "index, follow"
  }
}
```