package repository

import (
	"api-beer-challenge/database"
	"api-beer-challenge/internal/entity"
	"api-beer-challenge/internal/model"
	"api-beer-challenge/settings"
	"context"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"
)

type testCaseBeer struct {
	Description string
	Input       model.InputBeer
	InputU      model.InputUBeer
	Output      entity.Beer
	Outputs     []entity.Beer
	ExpectedErr error
}

var testCasesBeer map[string]testCaseBeer

var repo *repository

func setUpTestCases() {
	createdAt := time.Unix(time.Now().Unix(), 0).UTC()
	updatedAt := createdAt.Add(3 * time.Second)

	testCasesBeer = map[string]testCaseBeer{
		"find_none_beers": {
			Description: "find current beers",
			Outputs:     []entity.Beer{},
			ExpectedErr: nil,
		},
		"insert_beer": {
			Description: "insert new beer",
			Input: model.InputBeer{
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
		"update_beer_by_id": {
			Description: "update beer by id",
			InputU: model.InputUBeer{
				Name:      model.InputU{Value: "Estrella", Valid: true},
				Brewery:   model.InputU{Value: "BeerHouse", Valid: true},
				Country:   model.InputU{Value: "Mexico", Valid: true},
				Price:     model.InputU{Value: 20.00, Valid: true},
				Currency:  model.InputU{Value: "MXN", Valid: true},
				UpdatedAt: model.InputU{Value: updatedAt, Valid: true},
			},
			Output: entity.Beer{
				ID:        1,
				Name:      "Estrella",
				Brewery:   "BeerHouse",
				Country:   "Mexico",
				Price:     20.00,
				Currency:  "MXN",
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
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
		"find_beer_not_found": {
			Description: "find beer not found",
			ExpectedErr: nil,
		},
		"find_beer_boxprice": {
			Description: "find beer boxprice",
			ExpectedErr: nil,
		},
		"delete_beer_by_id": {
			Description: "delete beer by id",
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
	findBeerBoxPriceFake(ctx, t)
	updateBeerByID(ctx, t)
	deleteBeerByID(ctx, t)
	findBeerNotFound(ctx, t)
}

func TestIsEqualsNotDeletedItems(t *testing.T) {
	ok := repo.EqualsNotDeletedItems(nil)
	if ok {
		t.Fatalf("expected equals false, got %v", ok)
	}

	ok = repo.EqualsNotDeletedItems(ErrNoneEntityDeleted)
	if !ok {
		t.Fatalf("expected equals true, got %v", ok)
	}
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
			if errors.Is(err, ErrRequestInvalid) || errors.Is(err, context.DeadlineExceeded) {
				fmt.Println("API ERROR: ", err)
				return
			}

			t.Fatalf("expected error %v, got %v", testCase.ExpectedErr, err)
		}

		if price == 0 {
			t.Fatalf("expected price != 0, got %v", price)
		}
	})
}

func findBeerBoxPriceFake(ctx context.Context, t *testing.T) {
	testCase := testCasesBeer["find_beer_boxprice"]

	t.Run(testCase.Description, func(t *testing.T) {
		var ID, quantity uint64 = 1, 6
		currency := "USD"

		price, err := repo.FindBoxPriceBeerFake(ctx, ID, quantity, currency)
		if err != nil {
			t.Fatalf("expected error %v, got %v", testCase.ExpectedErr, err)
		}

		if price == 0 {
			t.Fatalf("expected price != 0, got %v", price)
		}
	})
}

func findBeerNotFound(ctx context.Context, t *testing.T) {
	testCase := testCasesBeer["find_beer_not_found"]

	t.Run(testCase.Description, func(t *testing.T) {
		var ID uint64 = 1
		b, err := repo.FindBeerByID(ctx, ID)
		if err != nil {
			t.Fatalf("expected error %v, got %v", testCase.ExpectedErr, err)
		}

		if b != nil {
			t.Fatalf("expected beer nil, got %v", b)
		}
	})
}

func updateBeerByID(ctx context.Context, t *testing.T) {
	testCase := testCasesBeer["update_beer_by_id"]

	t.Run(testCase.Description, func(t *testing.T) {
		var ID uint64 = 1
		b, err := repo.UpdateBeerByID(ctx, ID, &testCase.InputU)
		if err != nil {
			t.Fatalf("expected error %v, got %v", testCase.ExpectedErr, err)
		}

		assertBeer(&testCase.Output, b, t)
	})
}

func deleteBeerByID(ctx context.Context, t *testing.T) {
	testCase := testCasesBeer["delete_beer_by_id"]

	t.Run(testCase.Description, func(t *testing.T) {
		var ID uint64 = 1
		err := repo.DeleteBeerByID(ctx, ID)
		if err != nil {
			t.Fatalf("expected error %v, got %v", testCase.ExpectedErr, err)
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
