package grpc

import (
	"context"

	"github.com/hm-edu/domain-rest-interface/ent"
	"github.com/hm-edu/domain-rest-interface/pkg/store"
	pb "github.com/hm-edu/portal-apis"
	"github.com/hm-edu/portal-common/helper"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type domainAPIServer struct {
	pb.UnimplementedDomainServiceServer
	store  *store.DomainStore
	logger *zap.Logger
	tracer trace.Tracer
	admins []string
}

func newDomainAPIServer(store *store.DomainStore, logger *zap.Logger, admins []string) *domainAPIServer {
	tracer := otel.GetTracerProvider().Tracer("domains")
	return &domainAPIServer{store: store, logger: logger, tracer: tracer, admins: admins}
}

func (api *domainAPIServer) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest) (*pb.CheckPermissionResponse, error) {

	ctx, span := api.tracer.Start(ctx, "CheckPermission")
	defer span.End()
	span.SetAttributes(attribute.String("user", req.User), attribute.StringSlice("domains", req.Domains))
	domains, err := api.store.ListDomains(ctx, req.User, true, false)
	log := otelzap.New(api.logger.With(zap.String("user", req.User), zap.Strings("domains", req.Domains)))
	if err != nil {
		return nil, err
	}
	log.Ctx(ctx).Info("Checking permissions", zap.String("user", req.User), zap.Strings("domains", req.Domains))
	permissions := helper.Map(req.Domains, func(t string) *pb.Permission {
		if helper.Any(domains, func(d *ent.Domain) bool { return d.Fqdn == t }) {
			log.Ctx(ctx).Info("Permission granted", zap.String("user", req.User), zap.String("domain", t))
			return &pb.Permission{Domain: t, Granted: true}
		}
		log.Ctx(ctx).Info("Permission denied", zap.String("user", req.User), zap.String("domain", t))
		return &pb.Permission{Domain: t, Granted: false}
	})
	log.Ctx(ctx).Info("Checked permissions", zap.String("user", req.User), zap.Any("permissions", permissions))
	resp := pb.CheckPermissionResponse{Permissions: permissions}

	return &resp, nil
}

func (api *domainAPIServer) CheckRegistration(ctx context.Context, req *pb.CheckRegistrationRequest) (*pb.CheckRegistrationResponse, error) {

	ctx, span := api.tracer.Start(ctx, "CheckRegistration")
	defer span.End()

	api.logger.Info("Checking registrations domains", zap.Strings("domains", req.Domains))
	domains, err := api.store.ListAllDomains(ctx, true)
	if err != nil {
		span.RecordError(err)
		api.logger.Error("Checking registrations failed", zap.Strings("domains", req.Domains), zap.Error(err))
		return nil, err
	}

	missing := helper.Where(req.Domains, func(t string) bool {
		return !helper.Any(domains, func(d string) bool {
			return d == t
		})
	})
	api.logger.Info("Checked registrations", zap.Strings("domains", req.Domains), zap.Strings("missing", missing))
	return &pb.CheckRegistrationResponse{Missing: missing}, nil
}

func (api *domainAPIServer) ListDomains(ctx context.Context, req *pb.ListDomainsRequest) (*pb.ListDomainsResponse, error) {
	ctx, span := api.tracer.Start(ctx, "ListDomains")
	defer span.End()
	api.logger.Info("Listing domains", zap.String("user", req.User))
	domains, err := api.store.ListDomains(ctx, req.User, req.Approved, helper.Contains(api.admins, req.User))
	if err != nil {
		span.RecordError(err)
		api.logger.Error("Listing domains failed", zap.String("user", req.User), zap.Error(err))
		return nil, err
	}
	api.logger.Debug("Listed domains", zap.String("user", req.User), zap.Any("domains", domains))
	resp := pb.ListDomainsResponse{Domains: helper.Map(domains, func(t *ent.Domain) string { return t.Fqdn })}
	return &resp, nil
}
