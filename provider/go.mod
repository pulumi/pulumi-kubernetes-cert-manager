module github.com/joeduffy/pulumi-chart-cert-manager

go 1.15

require (
	github.com/joeduffy/pulumi-go-helmbase v0.0.10
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi-kubernetes/sdk/v3 v3.7.3
	github.com/pulumi/pulumi/pkg/v3 v3.12.0
	github.com/pulumi/pulumi/sdk/v3 v3.13.2
)

replace github.com/joeduffy/pulumi-go-helmbase v0.0.10 => ../../pulumi-go-helmbase

replace github.com/pulumi/pulumi/sdk/v3 v3.13.2 => ../../../pulumi/pulumi/sdk
