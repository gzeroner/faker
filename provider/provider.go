package provider

type Provider interface {
	Address
	Person
}

type Address interface {
	GetCityPrefix() []string
	GetCitySuffix() []string
	GetBuildingNumber() []string
	GetStreetSuffix() []string
	GetPostcode() []string
	GetState() []string
	GetStateAbbr() []string
	GetCountry() []string
}

type Person interface {
	GetFirstNameMale() []string
	GetFirstNameFemale() []string
	GetLastName() []string
}

type Company interface {
	GetJobTitleFormat() []string
	GetCompanySuffix() []string
}

type Text interface {
	GetText() string
}
