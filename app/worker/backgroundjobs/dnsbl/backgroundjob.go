package job

import (
	"context"
	"fmt"
	"time"

	"github.com/sujithshajee/dnsbl/app/dnsbl"
	"github.com/sujithshajee/dnsbl/app/ent"
	"github.com/sujithshajee/dnsbl/app/ent/ip"
)

const checkTime = 24 * time.Hour

type DNSBLTask struct {
	cl    *ent.Client
	dnsbl dnsbl.DNSBL
}

func NewDNSBLJob(cl *ent.Client, dnsbl dnsbl.DNSBL) *DNSBLTask {
	return &DNSBLTask{
		cl:    cl,
		dnsbl: dnsbl,
	}
}

func (s *DNSBLTask) Execute(ctx context.Context, ipaddr string) error {
	tx, err := s.cl.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}
	defer func() { _ = tx.Rollback() }()

	ipe, err := tx.IP.Query().Where(
		ip.IPAddressEQ(ipaddr),
	).Only(ctx)
	switch {
	case ent.IsNotFound(err):
		ipe, err = tx.IP.Create().
			SetIPAddress(ipaddr).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("ceating new ip: %w", err)
		}
	case err != nil:
		return fmt.Errorf("ip lookup: %w", err)
	default:
		nextAllowed := ipe.UpdatedAt.Add(checkTime)
		if time.Now().Before(nextAllowed) {
			_ = tx.Rollback()
			return fmt.Errorf("please wait before checking the IP address again")
		}
		_, err = ipe.Update().SetUpdatedAt(time.Now()).Save(ctx)
		if err != nil {
			return fmt.Errorf("ip updated at: %w", err)
		}
	}

	query, err := tx.AppQuery.Create().
		SetIpaddress(ipe).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("query update: %w", err)
	}

	resp, err := s.dnsbl.Query(ctx, ipaddr)
	if err != nil {
		return fmt.Errorf("job execution: %w", err)
	}

	for _, code := range resp.ResponseCodes {
		_, err := tx.AppResponse.Create().
			SetQuery(query).
			SetResponsecode(string(code)).
			SetCodedescription("spamhaus response").
			Save(ctx)
		if err != nil {
			return fmt.Errorf("saving response: %w", err)
		}
	}

	return tx.Commit()
}
