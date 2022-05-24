package config

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Data Base Client", func() {

	Context("data base", func() {
		AfterEach(func() {
			os.Clearenv()

		})

		When("Map To Database TimeStamp ", func() {
			It("should return  OK", func() {

				timeStamp := time.Now()
				nilTimeStamp := time.Time{}
				sqlTimeStamp := NewDatabaseTimeStamp(timeStamp)
				sqlNullTimeStamp := NewDatabaseTimeStamp(nilTimeStamp)

				Expect(reflect.TypeOf(sql.NullTime{})).Should(Equal(reflect.TypeOf(sqlNullTimeStamp)))
				Expect(reflect.TypeOf(time.Time{})).Should(Equal(reflect.TypeOf(sqlTimeStamp)))
			})
		})
		When("Map To Database String", func() {
			It("should return  OK", func() {
				text := "garex"
				nilText := ""
				textExpected := "{ false}"
				sqlText := NewDatabaseString(text)
				sqlNullText := NewDatabaseString(nilText)

				Expect(textExpected).Should(Equal(fmt.Sprintf("%v", sqlNullText)))
				Expect(text).Should(Equal(fmt.Sprintf("%v", sqlText)))
			})
		})

	})

})
