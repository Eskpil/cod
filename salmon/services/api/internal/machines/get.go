package machines

import (
	"context"
	"strings"

	"github.com/eskpil/salmon/pkg/models"
	"github.com/eskpil/salmon/services/api/database"

	interfaceService "github.com/eskpil/salmon/services/api/internal/interfaces"
)

func GetById(ctx context.Context, id string) (models.Machine, error) {
	row := database.Conn.ExecRow(ctx, "SELECT id, name, host, hostname, groups, fqdn FROM machines WHERE id = ?", id)

	machine := models.Machine{}

	groups := ""
	if err := row.Scan(&machine.Id, &machine.Name, &machine.Host, &machine.Hostname, &groups, &machine.Fqdn); err != nil {
		return machine, err
	}

	machine.Groups = strings.Split(groups, ";;-_-;;")

	// Get all interfaces by the machiens id
	interfaces, err := interfaceService.GetAll(ctx, machine.Id)

	if err != nil {
		return machine, err
	}

	machine.Interfaces = interfaces

	return machine, nil
}

func GetAll(ctx context.Context) ([]models.Machine, error) {
	machines := []models.Machine{}

	rows, err := database.Conn.ExecRows(ctx, "SELECT id, name, host, hostname, groups, fqdn FROM machines")
	if err != nil {
		return machines, err
	}

	for rows.Next() {
		machine := models.Machine{}
		groups := ""

		if err := rows.Scan(&machine.Id, &machine.Name, &machine.Host, &machine.Hostname, &groups, &machine.Fqdn); err != nil {
			return machines, err
		}

		// Get all interfaces by the machiens id
		interfaces, err := interfaceService.GetAll(ctx, machine.Id)

		if err != nil {
			return machines, err
		}

		machine.Groups = strings.Split(groups, ";;-_-;;")
		machine.Interfaces = interfaces
		machines = append(machines, machine)
	}

	return machines, nil
}
