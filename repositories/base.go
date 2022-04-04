package repositories

import (
	"strings"

	"github.com/jinzhu/gorm"
)

func makeFilterFunc(query interface{}, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	filterFunc := func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
	return filterFunc
}

func makePatternMatchingFilter(keyword string) string {
	var builder strings.Builder
	builder.WriteString("%")
	builder.WriteString(keyword)
	builder.WriteString("%")
	return builder.String()
}
