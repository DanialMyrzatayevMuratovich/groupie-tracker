package controller

import (
	"groupie-tracker/models"
	"testing"
)

var testArtists = []models.Artist{
	{ID: 1, Name: "Queen", Members: []string{"Freddie Mercury", "Brian May", "Roger Taylor", "John Deacon"}, CreationDate: 1970, FirstAlbum: "Queen"},
	{ID: 2, Name: "Gorillaz", Members: []string{"Damon Albarn", "Jamie Hewlett"}, CreationDate: 1998, FirstAlbum: "Gorillaz"},
	{ID: 3, Name: "Travis Scott", Members: []string{"Jacques Bermon Webster II"}, CreationDate: 2008, FirstAlbum: "Owl Pharaoh"},
	{ID: 4, Name: "Foo Fighters", Members: []string{"Dave Grohl", "Taylor Hawkins", "Nate Mendel"}, CreationDate: 1994, FirstAlbum: "Foo Fighters"},
}

func TestFilterArtists_ByName(t *testing.T) {
	got := filterArtists(testArtists, "queen")
	if len(got) != 1 || got[0].Name != "Queen" {
		t.Errorf("expected [Queen], got %v", got)
	}
}

func TestFilterArtists_ByMember(t *testing.T) {
	got := filterArtists(testArtists, "dave grohl")
	if len(got) != 1 || got[0].Name != "Foo Fighters" {
		t.Errorf("expected [Foo Fighters], got %v", got)
	}
}

func TestFilterArtists_ByCreationYear(t *testing.T) {
	got := filterArtists(testArtists, "1998")
	if len(got) != 1 || got[0].Name != "Gorillaz" {
		t.Errorf("expected [Gorillaz], got %v", got)
	}
}

func TestFilterArtists_ByFirstAlbum(t *testing.T) {
	got := filterArtists(testArtists, "owl pharaoh")
	if len(got) != 1 || got[0].Name != "Travis Scott" {
		t.Errorf("expected [Travis Scott], got %v", got)
	}
}

func TestFilterArtists_CaseInsensitive(t *testing.T) {
	for _, q := range []string{"QUEEN", "queen", "Queen", "QuEeN"} {
		got := filterArtists(testArtists, q)
		if len(got) != 1 || got[0].Name != "Queen" {
			t.Errorf("query %q: expected [Queen], got %v", q, got)
		}
	}
}

func TestFilterArtists_NoMatch(t *testing.T) {
	got := filterArtists(testArtists, "beatles")
	if len(got) != 0 {
		t.Errorf("expected 0 results, got %d", len(got))
	}
}

func TestFilterArtists_EmptyQuery_ReturnsAll(t *testing.T) {
	got := filterArtists(testArtists, "")
	if len(got) != len(testArtists) {
		t.Errorf("expected all %d artists, got %d", len(testArtists), len(got))
	}
}

func TestMatchesQuery(t *testing.T) {
	a := models.Artist{
		Name:         "Foo Fighters",
		Members:      []string{"Dave Grohl"},
		CreationDate: 1994,
		FirstAlbum:   "Foo Fighters",
	}

	cases := []struct {
		query string
		want  bool
	}{
		{"foo fighters", true},
		{"dave grohl", true},
		{"1994", true},
		{"beatles", false},
		{"2000", false},
	}

	for _, c := range cases {
		if got := matchesQuery(a, c.query); got != c.want {
			t.Errorf("matchesQuery(%q) = %v, want %v", c.query, got, c.want)
		}
	}
}
