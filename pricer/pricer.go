package pricer

import (
	"github.com/doyyan/ECS/bucket"
)

type pricer interface {
	Price(bucket bucket.Bucket)
}
