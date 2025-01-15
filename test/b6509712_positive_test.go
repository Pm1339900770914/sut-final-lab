package unit

import (
	"sut-final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestAllValid(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("All field valid", func(t *testing.T) {
		product := entity.Product{
			Name:  "Poom",
			Price: 100,
			SKU:   "B11111",
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
