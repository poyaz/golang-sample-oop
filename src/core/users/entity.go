package users

type GenderEnum uint8

const (
	Male GenderEnum = iota + 1
	Female
)

type Users struct {
	Id     uint64
	Name   string
	Age    uint8
	Gender GenderEnum
}
