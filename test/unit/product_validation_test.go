package unit

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/Yutthana27/lab05/model"
)

func TestProductValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	// กรณีข้อมูลครบทุก field → ต้องผ่าน
	t.Run("Valid product", func(t *testing.T) {
		product := model.Product{
			Name:        "Coke Zero",
			Price:       25.0,
			Stock:       10,
			Description: "Sugar-free drink",
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	// กรณี Name เป็นค่าว่าง 
	t.Run("Name is required", func(t *testing.T) {
		product := model.Product{
			Name:        "", // ผิดตรงนี้
			Price:       25.0,
			Stock:       10,
			Description: "Sugar-free drink",
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Name is required"))
	})

	// กรณี Price น้อยกว่าเท่ากับ 0 
	t.Run("Price must be greater than 0", func(t *testing.T) {
		product := model.Product{
			Name:        "Coke Zero",
			Price:       -1, // ผิดตรงนี้
			Stock:       10,
			Description: "Sugar-free drink",
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Price must be greater than 0"))
	})

	// กรณี Stock น้อยกว่า 0 
	t.Run("Stock cannot be negative", func(t *testing.T) {
		product := model.Product{
			Name:        "Coke Zero",
			Price:       25.0,
			Stock:       -1, // ผิดตรงนี้
			Description: "Sugar-free drink",
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Stock cannot be nagative"))
	})

	// กรณี Description น้อยกว่า 10 ตัวอักษร
	t.Run("Description too short", func(t *testing.T) {
		product := model.Product{
			Name:        "Coke Zero",
			Price:       25.0,
			Stock:       10,
			Description: "Too short", // ผิดตรงนี้ 
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Description must be at least 10 characters"))
	})
}
