package unit

import (
	. "github.com/onsi/gomega"
	"github.com/OnpreeyaMi/test-se.git/entity"
	"testing"
	"github.com/asaskevich/govalidator"
)
func TestStudentValidation(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`กรณีข้อมูลครบทุก field`, func(t *testing.T) {
		student := entity.Students{
			Fullname: "John Doe",
			Age:      20,
			Email:    "john.doe@gmail.com",
			GPA:      3.05,
		}
		ok, err := govalidator.ValidateStruct(student)

		g.Expect(err).To(BeNil())
		g.Expect(ok).To(BeTrue())
	})

	t.Run(`กรณี Fullname เป็นค่าว่าง`, func(t *testing.T) {
		student := entity.Students{
			Fullname: "",
			Age:      20,
			Email:    "john.doe@gmail.com",
			GPA:      3.05,
		}
		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Fullname is required"))
	})

	t.Run((`กรณี Age น้อยกว่า 18`), func(t *testing.T) {
		student := entity.Students{
			Fullname: "John Doe",
			Age:      17,
			Email:    "john.doe@gmail.com",
			GPA:      3.05,
		}
		ok, err := govalidator.ValidateStruct(student)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Age must be at least 18"))
	})

	t.Run((`กรณี Email ไม่ถูกต้อง`), func(t *testing.T) {
		student := entity.Students{
			Fullname: "John Doe",
			Age:      20,
			Email:    "john.doegmail.com",
			GPA:      3.05,
		}
		ok, err := govalidator.ValidateStruct(student)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is invalid"))
	})

	t.Run((`กรณี GPA ไม่ถูกต้อง`), func(t *testing.T) {
		student := entity.Students{
			Fullname: "John Doe",
			Age:      20,
			Email:    "john.doe@gmail.com",
			GPA:      4.50,
		}
		ok, err := govalidator.ValidateStruct(student)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("GPA must be between 0.00 and 4.00"))
	})
}