package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// IngestLicense is the resolver for the ingestLicense field.
func (r *mutationResolver) IngestLicense(ctx context.Context, license *model.IDorLicenseInput) (string, error) {
	if license.LicenseInput != nil {
		if err := validateLicenseInput(license.LicenseInput); err != nil {
			return "", err
		}
	}
	return r.Backend.IngestLicense(ctx, license)
}

// IngestLicenses is the resolver for the ingestLicenses field.
func (r *mutationResolver) IngestLicenses(ctx context.Context, licenses []*model.IDorLicenseInput) ([]string, error) {
	for _, l := range licenses {
		if l.LicenseInput != nil {
			if err := validateLicenseInput(l.LicenseInput); err != nil {
				return nil, err
			}
		}
	}
	return r.Backend.IngestLicenses(ctx, licenses)
}

// Licenses is the resolver for the licenses field.
func (r *queryResolver) Licenses(ctx context.Context, licenseSpec model.LicenseSpec) ([]*model.License, error) {
	return r.Backend.Licenses(ctx, &licenseSpec)
}

// LicenseList is the resolver for the licenseList field.
func (r *queryResolver) LicenseList(ctx context.Context, licenseSpec model.LicenseSpec, after *string, first *int) (*model.LicenseConnection, error) {
	return r.Backend.LicenseList(ctx, licenseSpec, after, first)
}
