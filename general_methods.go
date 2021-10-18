package goladok3

import (
	"context"
	"errors"
	"fmt"
)

// IsLadokPermissionsSufficent compare ladok permissions with p
func (c *Client) IsLadokPermissionsSufficent(ctx context.Context, ps Permissions) (Permissions, error) {
	egna, _, err := c.KataloginformationService.GetAnvandarbehorighetEgna(ctx)
	if err != nil {
		return nil, err
	}

	if len(egna.Anvandarbehorighet) < 1 {
		return nil, errors.New("missing Användarbehörigheter")
	}

	cfg := &GetBehorighetsprofilerCfg{
		UID: egna.UID,
	}
	profil, _, err := c.KataloginformationService.GetBehorighetsprofil(ctx, cfg)
	if err != nil {
		return nil, err
	}

	missingPermission := Permissions{}
	if missingPermission == nil {
		fmt.Println("is nil")
	}

	for pk, pv := range ps {
		for _, behorighet := range profil.Behorighetsprofiler {
			for _, s := range behorighet.Systemaktiviteter {
				if s.ID == pk {
					if pv == s.Rattighetsniva {
						continue
					}
				}
			}
		}
		missingPermission[pk] = pv
	}
	return missingPermission, nil
}

func (c *Client) environment() (string, error) {
	switch c.certificate.Subject.OrganizationalUnit[1] {
	case envIntTestAPI:
		return envIntTestAPI, nil
	case envProdAPI:
		return envProdAPI, nil
	case envTestAPI:
		return envTestAPI, nil
	default:
		return "", ErrNoEnvFound
	}
}
