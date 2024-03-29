package worker

import (
	"context"
	"sync"

	"github.com/hm-edu/pki-service/ent"
	"github.com/hm-edu/pki-service/ent/certificate"
	"github.com/hm-edu/pki-service/ent/domain"

	pkiHelper "github.com/hm-edu/pki-service/pkg/helper"
	"github.com/hm-edu/sectigo-client/sectigo"
	"github.com/hm-edu/sectigo-client/sectigo/ssl"

	"go.uber.org/zap"
)

// Syncer holds the sectigo client and the database instance.
type Syncer struct {
	Client *sectigo.Client
	Db     *ent.Client
}

// SyncAllCertificates downloads all available information from the Sectigo API and stores it in the database.
func (s *Syncer) SyncAllCertificates() {
	logger := zap.L()
	ctx := context.Background()
	certs, certificates, err := s.Client.SslService.List(&ssl.ListSSLRequest{Size: 200})

	if err != nil {
		logger.Fatal("Error while listing certificates", zap.Error(err))
		return
	}
	offset := 0
	for {
		var wg sync.WaitGroup
		details := []*ssl.Details{}
		for _, cert := range *certs {
			wg.Add(1)
			go func(cert ssl.ListItem) {
				defer wg.Done()
				item, err := s.Client.SslService.Details(cert.SslID)
				if err != nil {
					logger.Error("Error while getting certificate details", zap.Error(err), zap.Int("id", cert.SslID))
					return
				}
				// In the the requested time is empty due to the ACME issuance.
				if item.Requested == nil {
					cert, err := s.Client.SslService.Collect(item.SslID, "x509CO")
					if err != nil {
						logger.Error("Error while collecting certificate", zap.Error(err), zap.Int("id", item.SslID))
						return
					}
					certs, err := pkiHelper.ParseCertificates([]byte(*cert))
					if err != nil {
						logger.Error("Error while parsing certificate", zap.Error(err), zap.Int("id", item.SslID))
						return
					}
					item.Requested = &ssl.JSONDate{Time: certs[0].NotBefore}
				}
				details = append(details, item)
			}(cert)
		}
		wg.Wait()
		logger.Info("Got certificate details for certificates", zap.Int("count", len(details)))
		for _, item := range details {
			if item.SerialNumber == "" {
				continue
			}
			logger.Info("Updating certificate", zap.Int("id", item.SslID), zap.String("serial", item.SerialNumber))
			sans := []string{item.CommonName}

			for _, domain := range item.SubjectAlternativeNames {
				if domain != item.CommonName {
					sans = append(sans, domain)
				}
			}
			ids := []int{}

			for _, fqdn := range sans {
				id, err := s.Db.Domain.Create().SetFqdn(fqdn).OnConflictColumns(domain.FieldFqdn).Ignore().ID(ctx)

				if err != nil {
					logger.Error("Error while creating domain", zap.Error(err))
					continue
				}
				ids = append(ids, id)
			}

			creator := s.Db.Certificate.Create().SetCommonName(item.CommonName).SetSslId(item.SslID).SetNotAfter(item.Expires.Time).SetSerial(pkiHelper.NormalizeSerial(item.SerialNumber))

			if item.Requested != nil {
				creator.SetNotBefore(item.Requested.Time)
			}
			creator.SetStatus(certificate.Status(item.Status))
			id, err := creator.OnConflictColumns(certificate.FieldSerial).UpdateNewValues().ID(ctx)
			if err != nil {
				logger.Error("Error while creating certificate", zap.Error(err))
			}

			_, err = s.Db.Certificate.UpdateOneID(id).ClearDomains().AddDomainIDs(ids...).Save(ctx)
			if err != nil {
				logger.Error("Error while creating certificate", zap.Error(err))
			}
		}
		certificates -= len(*certs)
		if certificates <= 0 {
			break
		}
		offset += len(*certs)
		certs, _, err = s.Client.SslService.List(&ssl.ListSSLRequest{Size: 200, Position: offset})

		if err != nil {
			logger.Fatal("Error while listing certificates", zap.Error(err))
			return
		}
	}

}
