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

	/*
		filename := "test1.db"

		// Remove before test instead of after so we can review the file
		os.Remove(filename)

		// Remove db after testing
		defer os.Remove(filename)

		s, err := bolthold.Open(filename, 0666, nil)
		if err != nil {
			t.Error(err)
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
			t.Error(err)
		}

		result, err := s.FindAggregate(&Employee{}, nil, "Division")

		for i := range result {
			var division string
			employee := &Employee{}

			result[i].Group(&division)
			result[i].Min("Created", employee)

			fmt.Printf("The most senior employee in the %s division is %s.\n",
				division, employee.FirstName+" "+employee.LastName)
		}
	*/

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

	// Show time in "seconds" because "5756004313 ns" is hard to read
	timer := func(name string, f func(*testing.B)) func(*testing.B) {
		return func(b *testing.B) {
			start := time.Now()
			f(b)
			b.Logf("%f seconds\n", time.Since(start).Seconds())
		}
	}

	b.Run("LessThan", timer("LessThan", func(b *testing.B) {
		var result []Employee

		for i := 0; i < b.N; i++ {
			_ = s.Find(&result, bolthold.Where("Age").Lt(10))
		}
	}))

	b.Run("ComplexClause", timer("LessThan", func(b *testing.B) {
		var result []Employee

		for i := 0; i < b.N; i++ {
			_ = s.Find(&result, bolthold.Where("FirstName").Eq("John").And("Age").Lt(10).Or(
				bolthold.Where("Division").In("Saturday", "Sunday"),
			).SortBy("LastName"))
		}
	}))

	b.Run("Division", timer("DivisionAggregate", func(b *testing.B) {

		result, err := s.FindAggregate(&Employee{}, nil, "Division")

		if err != nil {
			b.Error(err)
		}

		for i := range result {
			var division string
			employee := &Employee{}

			result[i].Group(&division)
			result[i].Min("Created", employee)
			_ = employee

			// fmt.Printf("The most senior employee in the %s division is %s.\n",
			// 	division, employee.FirstName+" "+employee.LastName)
		}

	}))

	b.Run("Second", func(b *testing.B) {
		time.Sleep(time.Second * 20)
	})

}
