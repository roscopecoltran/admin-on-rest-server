package seeds

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/azumads/faker"
)

var Fake *faker.Faker

var (
	Root, _ = os.Getwd()
)

func init() {
	Fake, _ = faker.New("en")
	Fake.Rand = rand.New(rand.NewSource(42))
	rand.Seed(time.Now().UnixNano())

	filepaths, _ := filepath.Glob("shared/data/seeds/*.yml")
	if err := configor.Load(&Seeds, filepaths...); err != nil {
		panic(err)
	}
}
