package database

type Record struct {
	Id     string      `json:"id" bson:"_id"`
	ZoneId string      `json:"-" bson:"zone_id"`
	Fqdn   string      `json:"fqdn" bson:"fqdn"`
	Type   uint16      `json:"type" bson:"type"`
	Value  interface{} `json:"value" bson:"value"`
	Ttl    uint32      `json:"ttl" bson:"ttl"`
}

func init() {}
