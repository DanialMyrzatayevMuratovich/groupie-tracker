package controller

import (
	"groupie-tracker/api"
	"groupie-tracker/models"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ArtistFull combines artist profile with concert data.
type ArtistFull struct {
	Artist    models.Artist
	Locations []string
	Dates     map[string][]string
}

// dataCache holds the last successful API response with an expiry.
type dataCache struct {
	mu        sync.RWMutex
	artists   []models.Artist
	locations []models.Location
	dates     []models.Date
	relations []models.Relation
	expiresAt time.Time
}

var cache dataCache

const cacheTTL = 5 * time.Minute

// apiResult carries a value or error from a goroutine.
type apiResult[T any] struct {
	value T
	err   error
}

// fetchAsync launches fn in a goroutine and returns a buffered channel with the result.
func fetchAsync[T any](fn func() (T, error)) <-chan apiResult[T] {
	ch := make(chan apiResult[T], 1)
	go func() {
		v, e := fn()
		ch <- apiResult[T]{v, e}
	}()
	return ch
}

// getAllData returns all API data, using cached values when fresh.
func getAllData() ([]models.Artist, []models.Location, []models.Date, []models.Relation, error) {
	cache.mu.RLock()
	if time.Now().Before(cache.expiresAt) {
		a, l, d, r := cache.artists, cache.locations, cache.dates, cache.relations
		cache.mu.RUnlock()
		return a, l, d, r, nil
	}
	cache.mu.RUnlock()

	// Fetch all four endpoints concurrently via goroutines and channels.
	artistsCh := fetchAsync(api.FetchArtists)
	locationsCh := fetchAsync(api.FetchLocations)
	datesCh := fetchAsync(api.FetchDates)
	relationsCh := fetchAsync(api.FetchRelations)

	ar, lr, dr, rr := <-artistsCh, <-locationsCh, <-datesCh, <-relationsCh

	for _, err := range []error{ar.err, lr.err, dr.err, rr.err} {
		if err != nil {
			return nil, nil, nil, nil, err
		}
	}

	cache.mu.Lock()
	cache.artists, cache.locations, cache.dates, cache.relations = ar.value, lr.value, dr.value, rr.value
	cache.expiresAt = time.Now().Add(cacheTTL)
	cache.mu.Unlock()

	return ar.value, lr.value, dr.value, rr.value, nil
}

// WarmCache pre-fetches all data into the cache. Safe to call concurrently.
func WarmCache() error {
	_, _, _, _, err := getAllData()
	return err
}

// GetArtists returns all artists.
func GetArtists() ([]models.Artist, error) {
	artists, _, _, _, err := getAllData()
	return artists, err
}

// GetArtistByID returns combined artist data, or nil if not found.
func GetArtistByID(id int) (*ArtistFull, error) {
	artists, locations, _, relations, err := getAllData()
	if err != nil {
		return nil, err
	}

	var artist *models.Artist
	for i := range artists {
		if artists[i].ID == id {
			artist = &artists[i]
			break
		}
	}
	if artist == nil {
		return nil, nil
	}

	result := &ArtistFull{Artist: *artist}

	for i := range locations {
		if locations[i].ID == id {
			result.Locations = locations[i].Locations
			break
		}
	}
	for i := range relations {
		if relations[i].ID == id {
			result.Dates = relations[i].DatesLocations
			break
		}
	}

	return result, nil
}

// matchesQuery checks whether an artist matches a lowercase query string.
func matchesQuery(a models.Artist, q string) bool {
	if strings.Contains(strings.ToLower(a.Name), q) {
		return true
	}
	if strings.Contains(strings.ToLower(a.FirstAlbum), q) {
		return true
	}
	if strings.Contains(strconv.Itoa(a.CreationDate), q) {
		return true
	}
	for _, m := range a.Members {
		if strings.Contains(strings.ToLower(m), q) {
			return true
		}
	}
	return false
}

// filterArtists is a pure function — testable without hitting the API.
func filterArtists(artists []models.Artist, query string) []models.Artist {
	q := strings.ToLower(query)
	var result []models.Artist
	for _, a := range artists {
		if matchesQuery(a, q) {
			result = append(result, a)
		}
	}
	return result
}

// SearchArtists returns artists matching by name, member, first album, or creation year.
func SearchArtists(query string) ([]models.Artist, error) {
	artists, _, _, _, err := getAllData()
	if err != nil {
		return nil, err
	}
	return filterArtists(artists, query), nil
}
