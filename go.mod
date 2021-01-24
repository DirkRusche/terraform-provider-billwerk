module terraform-provider-billwerk

go 1.15

replace billwerk-go => ../billwerk-go

require (
	billwerk-go v0.0.0-00010101000000-000000000000
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.4.1
)
