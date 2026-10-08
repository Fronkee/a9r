# A9R

A fast terminal UI for managing AWS resources.

## Features

- EC2 instance detail/json viewer
- S3 bucket detail/json viewer
- Live search filtering
- Auto refresh
- Vim-style navigation
- AWS multi profile support
- Multi-region support
- Auto region select from profile
- Responsive terminal UI

## Preview

(screenshot later)

## Installation

### Clone

git clone ...

### Run

go run .

## Requirements

- Go 1.24+
- AWS credentials configured

## AWS Config Example

~/.aws/config

```ini
[default]
region = ap-southeast-1

[profile movie-uat]
region = us-east-1

[profile music-uat]
region = ap-south-1
```

Selecting a profile auto-selects its `region`.
You can still change the region manually.

## Controls

| Key | Action |
|-----|--------|
| TAB | switch focus |
| j | json |
| d | detail |
| / | search |
| r | refresh |
| a | auto refresh |
| q | quit |

## Architecture

configs/
models/
services/
ui/

## Future Roadmap

- IAM
- VPC
- EKS
- ASG
- CloudWatch
- Fuzzy search

## Tech Stack

- Go
- tview
- AWS SDK v2

## License

MIT
