# Groupie Tracker

A web application that consumes the [Groupie Trackers API](https://groupietrackers.herokuapp.com/api) and displays information about musical artists, their members, and concert tour dates.

## Features

- Browse all artists with images and founding year
- View artist details: members, first album, tour dates by location
- Live search by artist name, member, first album, or creation year
- Concurrent API fetching via goroutines and channels
- In-memory cache with automatic background refresh

## Stack

- **Backend:** Go (standard library only)
- **Frontend:** HTML, CSS, vanilla JS

## Run

```bash
git clone https://01.tomorrow-school.ai/git/dmyrzata/groupie-tracker.git
cd groupie-tracker
go run cmd/main.go
```

Open [http://localhost:3000](http://localhost:3000)

## Test

```bash
go test ./...
```

## Authors

- **dmyrzata**
- **akeles**
