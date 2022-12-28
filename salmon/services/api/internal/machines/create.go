package machines

import (
	"context"
	"strings"

	"github.com/eskpil/salmon/pkg/models"
	"github.com/eskpil/salmon/services/api/database"

	interfaceService "github.com/eskpil/salmon/services/api/internal/interfaces"
)

func Create(ctx context.Context, machine models.Machine) error {
	machine.Groups = make([]string, 1)

	_, err := database.Conn.Exec(ctx, "INSERT INTO machines VALUES(?,?,?,?,?,?)", machine.Id, machine.Name, machine.Host, machine.Hostname, strings.Join(machine.Groups, ";;-_-;;"), machine.Fqdn)

	if err != nil {
		return err
	}

	for _, iface := range machine.Interfaces {
		if err = interfaceService.Create(ctx, machine.Id, iface); err != nil {
			return err
		}
	}

	return nil
}
