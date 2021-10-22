package education

import "fmt"

type Course struct {
	Id          uint64
	Title       string
	Description string
}

func (c Course) String() string {
	return fmt.Sprintf("Id: %d Title: %s Description: %s", c.Id, c.Title, c.Description)
}
