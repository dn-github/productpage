package server

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	detailPb "github.com/dn-github/details/pb"
	"github.com/dn-github/productpage/client"
	"github.com/dn-github/productpage/pb"
	reviewPb "github.com/dn-github/reviews/pb"
)

type productPageImpl struct {
	reviewClient reviewPb.ReviewServiceClient
	detailClient detailPb.DetailServiceClient
}

func NewProductPageImpl() *productPageImpl {
	reviewConn, err := grpc.Dial("reviews:3001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}

	detailConn, err := grpc.Dial("details:3002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}

	return &productPageImpl{
		reviewClient: reviewPb.NewReviewServiceClient(reviewConn),
		detailClient: detailPb.NewDetailServiceClient(detailConn),
	}
}

// Ratings ...
func (r *productPageImpl) Product(ctx context.Context, book *pb.Book) (*pb.Description, error) {
	review, rating, err := client.GetReviews(r.reviewClient, book.Name)
	if err != nil {
		log.Fatalf(err.Error())
		review = "review service not available"
		rating = 0
	}
	price, genre, err := client.GetDetails(r.detailClient, book.Name)
	if err != nil {
		log.Fatalf(err.Error())
		price = 0
		genre = "None"
	}
	logMessage := fmt.Sprintf("Book %s Review %s Rating %d Price %d Genre %s", book.Name, review, rating, price, genre)
	log.Println(logMessage)
	return &pb.Description{
		Review: review,
		Rating: rating,
		Price:  price,
		Genre:  genre,
	}, nil
}
