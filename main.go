package main

import (
	"context"
	"log"
	"net/http"

	"github.com/blukai/gqlgenc-omitempty-example/gen"
	"github.com/sanity-io/litter"
)

func strPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func main() {
	genClient := gen.NewClient(http.DefaultClient,
		"http://localhost:8080/graphql")
	ctx := context.Background()

	added, err := genClient.AddPlanet(ctx, []*gen.AddPlanetInput{
		{Name: "Earth", Moons: []*gen.MoonRef{{Name: strPtr("Moon")}}},
		{Name: "Jupiter", Moons: []*gen.MoonRef{{Name: strPtr("Callisto")}}},
	})
	if err != nil {
		log.Fatalf("could not add: %s\n", err)
	}

	_, err = genClient.UpdatePlanet(ctx, gen.UpdatePlanetInput{
		Filter: &gen.PlanetFilter{
			ID: []string{added.AddPlanet.Planet[0].ID},
		},
		Remove: &gen.PlanetPatch{
			Moons: []*gen.MoonRef{
				{ID: &added.AddPlanet.Planet[0].Moons[0].ID},
			},
		},
	})
	if err != nil {
		log.Fatalf("could not update: %s\n", err)
	}

	planets, err := genClient.QueryPlanet(ctx)
	if err != nil {
		log.Fatalf("could not query: %s\n", err)
	}

	litter.Dump(planets.QueryPlanet)
}
