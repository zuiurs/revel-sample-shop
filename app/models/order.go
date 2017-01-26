package models

import (
	"strconv"
	"strings"
	"time"
)

type Order struct {
	ID        uint64     `json:"id" gorm:"primary_key"`
	OrderID   uint64     `json:"order_id"`
	ItemID    uint64     `json:"item_id"`
	Num       int64      `json:"num"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// format: &<id>#<num>
// e.g.) &1#5&4#1&...
func ParseCart(s string) (map[int]int, error) {
	cartElems := strings.Split(s, "&")[1:]

	elemList := make(map[int]int)

	for _, v := range cartElems {
		e := strings.Split(v, "#")

		id, err := strconv.Atoi(e[0])
		if err != nil {
			return nil, err
		}

		num, err := strconv.Atoi(e[1])
		if err != nil {
			return nil, err
		}

		elemList[id] = num
	}

	return elemList, nil
}
