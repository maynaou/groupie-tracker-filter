package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchArtistsAndRelationsData() error {
	// Fetch artists data
	artistsResp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return fmt.Errorf("error fetching artists data: %w", err)
	}
	defer artistsResp.Body.Close()

	data, err := io.ReadAll(artistsResp.Body)
	if err != nil {
		return fmt.Errorf("error reading artists response: %w", err)
	}

	err = json.Unmarshal(data, &ApiObjects)
	if err != nil {
		return fmt.Errorf("error unmarshalling artists data: %w", err)
	}
	// Fetch relations data for each artist and store it in the artist object
	for i := range ApiObjects {
		artist := &ApiObjects[i]
		if artist.RelationsURL != "" {
			relationResp, err := http.Get(artist.RelationsURL)
			if err != nil {
				return fmt.Errorf("error fetching relations from %s: %w", artist.RelationsURL, err)
			}
			defer relationResp.Body.Close()

			relationData, err := io.ReadAll(relationResp.Body)
			if err != nil {
				return fmt.Errorf("error reading relations response from %s: %w", artist.RelationsURL, err)
			}

			err = json.Unmarshal(relationData, &artist.Relation)
			if err != nil {
				return fmt.Errorf("error unmarshalling relations data from %s: %w", artist.RelationsURL, err)
			}
		}
	}

	return nil
}
