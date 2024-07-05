package enums

type ProductTypeEnum string

const (
	HUB ProductTypeEnum = "Hub"
	ORQ ProductTypeEnum = "Orchestrator"
	STD ProductTypeEnum = "Studio"
	MAG ProductTypeEnum = "Manager"
)
