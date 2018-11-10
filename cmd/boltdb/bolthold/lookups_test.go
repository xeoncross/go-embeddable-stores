package lookups

import (
	"math/rand"
	"os"
	"testing"
	"time"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/timshannon/bolthold"
	"go.etcd.io/bbolt"
)

type Employee struct {
	ID        string `boltholdKey:"ID"`
	FirstName string
	LastName  string
	Division  string
	Bio       string
	Age       int
	Created   time.Time
}

func newEmployee() *Employee {
	return &Employee{
		FirstName: randomdata.FirstName(randomdata.Male),
		LastName:  randomdata.LastName(),
		Division:  randomdata.Day(),
		Bio:       randomdata.Paragraph(),
		Age:       rand.Intn(100), // 0 is ok for our test
		Created:   time.Now(),
	}
}

func TestLookups(t *testing.T) {
	_ = &Employee{}
}

func BenchmarkLookups(b *testing.B) {

	filename := "test.db"

	// Remove before test instead of after so we can review the file
	os.Remove(filename)

	// Remove db after testing
	// defer os.Remove(filename)

	s, err := bolthold.Open(filename, 0666, nil)
	if err != nil {
		b.Error(err)
	}

	// Batch insert records
	err = s.Bolt().Update(func(tx *bbolt.Tx) error {
		for i := 0; i < 50000; i++ {
			err = s.TxInsert(tx, bolthold.NextSequence(), newEmployee())
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()

	var result []Employee

	for i := 0; i < b.N; i++ {
		_ = s.Find(&result, bolthold.Where("Age").Lt(10))
	}

}
