package client

import (
	"context"
	"time"

	detailPb "github.com/dn-github/details/pb"
	reviewPb "github.com/dn-github/reviews/pb"
)

// GetDetails calls details service and returns price, genre and error
func GetDetails(c detailPb.DetailServiceClient, book string) (int64, string, error) {
	detailRequest := detailPb.Book{
		Name: book,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	res, err := c.Details(ctx, &detailRequest)
	if err != nil {
		return 0, "", err
	}

	return res.Price, res.Genre, nil
}

// GetReviews calls reviews service and returns review and rating and error
func GetReviews(c reviewPb.ReviewServiceClient, book string) (string, int64, error) {
	reviewRequest := reviewPb.Book{
		Name: book,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	res, err := c.Reviews(ctx, &reviewRequest)
	if err != nil {
		return "", 0, err
	}

	return res.Review, res.Rating, nil
}
