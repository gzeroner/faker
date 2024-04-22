package faker

import "github.com/gzeroner/faker/provider/en_US"

type Gender int8

const (
	GenderUndefined Gender = iota
	GenderMale
	GenderFemale
)

func (f *Faker) Name(gender ...Gender) string {
	return f.FirstName(gender...) + " " + f.LastName()
}

func (f *Faker) FirstName(gender ...Gender) string {
	g := f.getGender(gender)
	var names []string
	switch g {
	case GenderMale:
		names = en_US.GetFirstNameMale()
	case GenderFemale:
		names = en_US.GetFirstNameFemale()
	default:
		male := en_US.GetFirstNameMale()
		female := en_US.GetFirstNameFemale()
		names = append(male, female...)
	}
	return names[f.RandomBetween(0, len(names)-1)]
}

func (f *Faker) LastName() string {
	names := en_US.GetLastName()
	return names[f.RandomBetween(0, len(names)-1)]
}

func (f *Faker) getGender(genders []Gender) Gender {
	if genders == nil {
		return GenderUndefined
	}

	switch len(genders) {
	case 1:
		return genders[0]
	default:
		return GenderUndefined
	}
}
