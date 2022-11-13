package records

import (
	"context"
	"fmt"
	"strings"

	"github.com/eskpil/cod/trout/database"
	zoneService "github.com/eskpil/cod/trout/internal/zones"
	log "github.com/sirupsen/logrus"
)

// This function checks if the fqdn of the record trying to be created in the respective zone actually belongs in the zone
// if the zone fqdn cod. a record fqdn with test.cod. would fall through but a record with fqdn test.codtest. would fail
// in that zone.
func checkFqdnZoneOwnership(subjectFqdn string, zoneFqdn string, limit int) (error, bool) {
	if 0 >= limit {
		return nil, false
	}

	parts := strings.Split(subjectFqdn, ".")

	if len(parts)-1 == 0 {
		return nil, false
	}

	result := strings.TrimPrefix(subjectFqdn, fmt.Sprintf("%s.", parts[0]))

	if strings.Compare(result, zoneFqdn) == 0 {
		return nil, true
	} else {
		if _, ok := checkFqdnZoneOwnership(strings.Join(parts[1:], "."), zoneFqdn, limit-1); !ok {
			return fmt.Errorf("Fqdn: %s does not belong in zone with fqdn: %s\n", subjectFqdn, zoneFqdn), false
		}
	}

	return nil, false
}

func Create(ctx context.Context, record database.Record) error {
	zone, err := zoneService.GetById(ctx, record.ZoneId)
	if err != nil {
		log.Errorf("Could not look up zone: %s by id: %v\n", record.ZoneId, err)
		return fmt.Errorf("Zone: %s does not exists.", record.ZoneId)
	}

	if err, _ := checkFqdnZoneOwnership(record.Fqdn, zone.Fqdn, 256); err != nil {
		return err
	}

	_, err = getCollection().InsertOne(ctx, record)
	return nil
}
