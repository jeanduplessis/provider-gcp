/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	domainmapping "github.com/upbound/provider-gcp/internal/controller/cloudrun/domainmapping"
	service "github.com/upbound/provider-gcp/internal/controller/cloudrun/service"
	serviceiammember "github.com/upbound/provider-gcp/internal/controller/cloudrun/serviceiammember"
	v2job "github.com/upbound/provider-gcp/internal/controller/cloudrun/v2job"
	v2service "github.com/upbound/provider-gcp/internal/controller/cloudrun/v2service"
)

// Setup_cloudrun creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudrun(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domainmapping.Setup,
		service.Setup,
		serviceiammember.Setup,
		v2job.Setup,
		v2service.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
