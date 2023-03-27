package repository

import (
	"api-beer-challenge/database"
	"api-beer-challenge/internal/entity"
	"api-beer-challenge/internal/model"
	"api-beer-challenge/settings"
	"context"
	"log"
	"testing"
	"time"
)

type testCaseBeer struct {
	Description string
	Input       model.BeerInput
	Output      entity.Beer
	Outputs     []entity.Beer
	ExpectedErr error
}

var testCasesBeer map[string]testCaseBeer

var repo Repository

func setUpTestCases() {
	createdAt := time.Unix(time.Now().Unix(), 0).UTC()

	testCasesBeer = map[string]testCaseBeer{
		"find_none_beers": {
			Description: "find current beers",
			Outputs:     []entity.Beer{},
			ExpectedErr: nil,
		},
		"insert_beer": {
			Description: "insert new beer",
			Input: model.BeerInput{
				Name:      "Corona",
				Brewery:   "Grupo Modelo",
				Country:   "Mexico",
				Price:     25.00,
				Currency:  "MXN",
				CreatedAt: createdAt,
				UpdatedAt: createdAt,
			},
			Output: entity.Beer{
				ID:        1,
				Name:      "Corona",
				Brewery:   "Grupo Modelo",
				Country:   "Mexico",
				Price:     25.00,
				Currency:  "MXN",
				CreatedAt: createdAt,
				UpdatedAt: createdAt,
			},
			ExpectedErr: nil,
		},
		"find_beer_by_id": {
			Description: "find beer by id",
			Output: entity.Beer{
				ID:        1,
				Name:      "Corona",
				Brewery:   "Grupo Modelo",
				Country:   "Mexico",
				Price:     25.00,
				Currency:  "MXN",
				CreatedAt: createdAt,
				UpdatedAt: createdAt,
			},
			ExpectedErr: nil,
		},
		"find_beer_boxprice": {
			Description: "find beer boxprice",
			ExpectedErr: nil,
		},
	}
}

func TestBeer(t *testing.T) {
	setUpTestCases()

	ctx := context.Background()
	nSettings, err := settings.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.GetConnection(ctx, nSettings)
	if err != nil {
		log.Fatal(err)
	}

	repo = New(db)
	err = repo.RestartTable(ctx, "./../../database/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	findCurrentBeers(ctx, t)
	insertBeer(ctx, t)
	findBeerByID(ctx, t)
	findBeerBoxPrice(ctx, t)
}

func findCurrentBeers(ctx context.Context, t *testing.T) {
	testCase := testCasesBeer["find_none_beers"]

	t.Run(testCase.Description, func(t *testing.T) {
		bb, err := repo.FindBeers(ctx)
		if err != nil {
			t.Fatalf("expected error %v, got %v", testCase.ExpectedErr, err)
		}

		size := len(testCase.Outputs)
		current := len(bb)
		if size != current {
			t.Fatalf("expected %d beers, got %d beers", size, current)
		}
	})
}

func insertBeer(ctx context.Context, t *testing.T) {
	testCase := testCasesBeer["insert_beer"]

	t.Run(testCase.Description, func(t *testing.T) {
		b, err := repo.InsertBeer(ctx, &testCase.Input)
		if err != nil {
			t.Fatalf("expected error %v, got %v", testCase.ExpectedErr, err)
		}

		assertBeer(&testCase.Output, b, t)
	})
}

func findBeerByID(ctx context.Context, t *testing.T) {
	testCase := testCasesBeer["find_beer_by_id"]

	t.Run(testCase.Description, func(t *testing.T) {
		var ID uint64 = 1
		b, err := repo.FindBeerByID(ctx, ID)
		if err != nil {
			t.Fatalf("expected error %v, got %v", testCase.ExpectedErr, err)
		}

		assertBeer(&testCase.Output, b, t)
	})
}

func findBeerBoxPrice(ctx context.Context, t *testing.T) {
	testCase := testCasesBeer["find_beer_boxprice"]

	t.Run(testCase.Description, func(t *testing.T) {
		var ID, quantity uint64 = 1, 6
		currency := "USD"

		price, err := repo.FindBoxPriceBeer(ctx, ID, quantity, currency)
		if err != nil {
			t.Fatalf("expected error %v, got %v", testCase.ExpectedErr, err)
		}

		if price == 0 {
			t.Fatalf("expected price != 0, got %v", price)
		}
	})
}

func assertBeer(want, got *entity.Beer, t testing.TB) {
	if want.ID != got.ID {
		t.Fatalf("expected ID %d, got ID %d", want.ID, got.ID)
	}

	if want.Name != got.Name {
		t.Fatalf("expected Name %s, got Name %s", want.Name, got.Name)
	}

	if want.Brewery != got.Brewery {
		t.Fatalf("expected Brewery %s, got Brewery %s", want.Brewery, got.Brewery)
	}

	if want.Country != got.Country {
		t.Fatalf("expected Country %s, got Country %s", want.Country, got.Country)
	}

	if want.Price != got.Price {
		t.Fatalf("expected Price %.2f, got Price %.2f", want.Price, got.Price)
	}

	if want.Currency != got.Currency {
		t.Fatalf("expected Currency %s, got Currency %s", want.Currency, got.Currency)
	}

	if !want.CreatedAt.Equal(got.CreatedAt) {
		t.Fatalf("expected CreatedAt %s, got CreatedAt %s", want.CreatedAt, got.CreatedAt)
	}

	if !want.UpdatedAt.Equal(got.UpdatedAt) {
		t.Fatalf("expected UpdatedAt %s, got UpdatedAt %s", want.UpdatedAt, got.UpdatedAt)
	}
}
