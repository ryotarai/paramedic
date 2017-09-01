// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package servicecatalogiface provides an interface to enable mocking the AWS Service Catalog service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package servicecatalogiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/servicecatalog"
)

// ServiceCatalogAPI provides an interface to enable mocking the
// servicecatalog.ServiceCatalog service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Service Catalog.
//    func myFunc(svc servicecatalogiface.ServiceCatalogAPI) bool {
//        // Make svc.AcceptPortfolioShare request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := servicecatalog.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockServiceCatalogClient struct {
//        servicecatalogiface.ServiceCatalogAPI
//    }
//    func (m *mockServiceCatalogClient) AcceptPortfolioShare(input *servicecatalog.AcceptPortfolioShareInput) (*servicecatalog.AcceptPortfolioShareOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockServiceCatalogClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ServiceCatalogAPI interface {
	AcceptPortfolioShare(*servicecatalog.AcceptPortfolioShareInput) (*servicecatalog.AcceptPortfolioShareOutput, error)
	AcceptPortfolioShareWithContext(aws.Context, *servicecatalog.AcceptPortfolioShareInput, ...request.Option) (*servicecatalog.AcceptPortfolioShareOutput, error)
	AcceptPortfolioShareRequest(*servicecatalog.AcceptPortfolioShareInput) (*request.Request, *servicecatalog.AcceptPortfolioShareOutput)

	AssociatePrincipalWithPortfolio(*servicecatalog.AssociatePrincipalWithPortfolioInput) (*servicecatalog.AssociatePrincipalWithPortfolioOutput, error)
	AssociatePrincipalWithPortfolioWithContext(aws.Context, *servicecatalog.AssociatePrincipalWithPortfolioInput, ...request.Option) (*servicecatalog.AssociatePrincipalWithPortfolioOutput, error)
	AssociatePrincipalWithPortfolioRequest(*servicecatalog.AssociatePrincipalWithPortfolioInput) (*request.Request, *servicecatalog.AssociatePrincipalWithPortfolioOutput)

	AssociateProductWithPortfolio(*servicecatalog.AssociateProductWithPortfolioInput) (*servicecatalog.AssociateProductWithPortfolioOutput, error)
	AssociateProductWithPortfolioWithContext(aws.Context, *servicecatalog.AssociateProductWithPortfolioInput, ...request.Option) (*servicecatalog.AssociateProductWithPortfolioOutput, error)
	AssociateProductWithPortfolioRequest(*servicecatalog.AssociateProductWithPortfolioInput) (*request.Request, *servicecatalog.AssociateProductWithPortfolioOutput)

	AssociateTagOptionWithResource(*servicecatalog.AssociateTagOptionWithResourceInput) (*servicecatalog.AssociateTagOptionWithResourceOutput, error)
	AssociateTagOptionWithResourceWithContext(aws.Context, *servicecatalog.AssociateTagOptionWithResourceInput, ...request.Option) (*servicecatalog.AssociateTagOptionWithResourceOutput, error)
	AssociateTagOptionWithResourceRequest(*servicecatalog.AssociateTagOptionWithResourceInput) (*request.Request, *servicecatalog.AssociateTagOptionWithResourceOutput)

	CreateConstraint(*servicecatalog.CreateConstraintInput) (*servicecatalog.CreateConstraintOutput, error)
	CreateConstraintWithContext(aws.Context, *servicecatalog.CreateConstraintInput, ...request.Option) (*servicecatalog.CreateConstraintOutput, error)
	CreateConstraintRequest(*servicecatalog.CreateConstraintInput) (*request.Request, *servicecatalog.CreateConstraintOutput)

	CreatePortfolio(*servicecatalog.CreatePortfolioInput) (*servicecatalog.CreatePortfolioOutput, error)
	CreatePortfolioWithContext(aws.Context, *servicecatalog.CreatePortfolioInput, ...request.Option) (*servicecatalog.CreatePortfolioOutput, error)
	CreatePortfolioRequest(*servicecatalog.CreatePortfolioInput) (*request.Request, *servicecatalog.CreatePortfolioOutput)

	CreatePortfolioShare(*servicecatalog.CreatePortfolioShareInput) (*servicecatalog.CreatePortfolioShareOutput, error)
	CreatePortfolioShareWithContext(aws.Context, *servicecatalog.CreatePortfolioShareInput, ...request.Option) (*servicecatalog.CreatePortfolioShareOutput, error)
	CreatePortfolioShareRequest(*servicecatalog.CreatePortfolioShareInput) (*request.Request, *servicecatalog.CreatePortfolioShareOutput)

	CreateProduct(*servicecatalog.CreateProductInput) (*servicecatalog.CreateProductOutput, error)
	CreateProductWithContext(aws.Context, *servicecatalog.CreateProductInput, ...request.Option) (*servicecatalog.CreateProductOutput, error)
	CreateProductRequest(*servicecatalog.CreateProductInput) (*request.Request, *servicecatalog.CreateProductOutput)

	CreateProvisioningArtifact(*servicecatalog.CreateProvisioningArtifactInput) (*servicecatalog.CreateProvisioningArtifactOutput, error)
	CreateProvisioningArtifactWithContext(aws.Context, *servicecatalog.CreateProvisioningArtifactInput, ...request.Option) (*servicecatalog.CreateProvisioningArtifactOutput, error)
	CreateProvisioningArtifactRequest(*servicecatalog.CreateProvisioningArtifactInput) (*request.Request, *servicecatalog.CreateProvisioningArtifactOutput)

	CreateTagOption(*servicecatalog.CreateTagOptionInput) (*servicecatalog.CreateTagOptionOutput, error)
	CreateTagOptionWithContext(aws.Context, *servicecatalog.CreateTagOptionInput, ...request.Option) (*servicecatalog.CreateTagOptionOutput, error)
	CreateTagOptionRequest(*servicecatalog.CreateTagOptionInput) (*request.Request, *servicecatalog.CreateTagOptionOutput)

	DeleteConstraint(*servicecatalog.DeleteConstraintInput) (*servicecatalog.DeleteConstraintOutput, error)
	DeleteConstraintWithContext(aws.Context, *servicecatalog.DeleteConstraintInput, ...request.Option) (*servicecatalog.DeleteConstraintOutput, error)
	DeleteConstraintRequest(*servicecatalog.DeleteConstraintInput) (*request.Request, *servicecatalog.DeleteConstraintOutput)

	DeletePortfolio(*servicecatalog.DeletePortfolioInput) (*servicecatalog.DeletePortfolioOutput, error)
	DeletePortfolioWithContext(aws.Context, *servicecatalog.DeletePortfolioInput, ...request.Option) (*servicecatalog.DeletePortfolioOutput, error)
	DeletePortfolioRequest(*servicecatalog.DeletePortfolioInput) (*request.Request, *servicecatalog.DeletePortfolioOutput)

	DeletePortfolioShare(*servicecatalog.DeletePortfolioShareInput) (*servicecatalog.DeletePortfolioShareOutput, error)
	DeletePortfolioShareWithContext(aws.Context, *servicecatalog.DeletePortfolioShareInput, ...request.Option) (*servicecatalog.DeletePortfolioShareOutput, error)
	DeletePortfolioShareRequest(*servicecatalog.DeletePortfolioShareInput) (*request.Request, *servicecatalog.DeletePortfolioShareOutput)

	DeleteProduct(*servicecatalog.DeleteProductInput) (*servicecatalog.DeleteProductOutput, error)
	DeleteProductWithContext(aws.Context, *servicecatalog.DeleteProductInput, ...request.Option) (*servicecatalog.DeleteProductOutput, error)
	DeleteProductRequest(*servicecatalog.DeleteProductInput) (*request.Request, *servicecatalog.DeleteProductOutput)

	DeleteProvisioningArtifact(*servicecatalog.DeleteProvisioningArtifactInput) (*servicecatalog.DeleteProvisioningArtifactOutput, error)
	DeleteProvisioningArtifactWithContext(aws.Context, *servicecatalog.DeleteProvisioningArtifactInput, ...request.Option) (*servicecatalog.DeleteProvisioningArtifactOutput, error)
	DeleteProvisioningArtifactRequest(*servicecatalog.DeleteProvisioningArtifactInput) (*request.Request, *servicecatalog.DeleteProvisioningArtifactOutput)

	DescribeConstraint(*servicecatalog.DescribeConstraintInput) (*servicecatalog.DescribeConstraintOutput, error)
	DescribeConstraintWithContext(aws.Context, *servicecatalog.DescribeConstraintInput, ...request.Option) (*servicecatalog.DescribeConstraintOutput, error)
	DescribeConstraintRequest(*servicecatalog.DescribeConstraintInput) (*request.Request, *servicecatalog.DescribeConstraintOutput)

	DescribePortfolio(*servicecatalog.DescribePortfolioInput) (*servicecatalog.DescribePortfolioOutput, error)
	DescribePortfolioWithContext(aws.Context, *servicecatalog.DescribePortfolioInput, ...request.Option) (*servicecatalog.DescribePortfolioOutput, error)
	DescribePortfolioRequest(*servicecatalog.DescribePortfolioInput) (*request.Request, *servicecatalog.DescribePortfolioOutput)

	DescribeProduct(*servicecatalog.DescribeProductInput) (*servicecatalog.DescribeProductOutput, error)
	DescribeProductWithContext(aws.Context, *servicecatalog.DescribeProductInput, ...request.Option) (*servicecatalog.DescribeProductOutput, error)
	DescribeProductRequest(*servicecatalog.DescribeProductInput) (*request.Request, *servicecatalog.DescribeProductOutput)

	DescribeProductAsAdmin(*servicecatalog.DescribeProductAsAdminInput) (*servicecatalog.DescribeProductAsAdminOutput, error)
	DescribeProductAsAdminWithContext(aws.Context, *servicecatalog.DescribeProductAsAdminInput, ...request.Option) (*servicecatalog.DescribeProductAsAdminOutput, error)
	DescribeProductAsAdminRequest(*servicecatalog.DescribeProductAsAdminInput) (*request.Request, *servicecatalog.DescribeProductAsAdminOutput)

	DescribeProductView(*servicecatalog.DescribeProductViewInput) (*servicecatalog.DescribeProductViewOutput, error)
	DescribeProductViewWithContext(aws.Context, *servicecatalog.DescribeProductViewInput, ...request.Option) (*servicecatalog.DescribeProductViewOutput, error)
	DescribeProductViewRequest(*servicecatalog.DescribeProductViewInput) (*request.Request, *servicecatalog.DescribeProductViewOutput)

	DescribeProvisionedProduct(*servicecatalog.DescribeProvisionedProductInput) (*servicecatalog.DescribeProvisionedProductOutput, error)
	DescribeProvisionedProductWithContext(aws.Context, *servicecatalog.DescribeProvisionedProductInput, ...request.Option) (*servicecatalog.DescribeProvisionedProductOutput, error)
	DescribeProvisionedProductRequest(*servicecatalog.DescribeProvisionedProductInput) (*request.Request, *servicecatalog.DescribeProvisionedProductOutput)

	DescribeProvisioningArtifact(*servicecatalog.DescribeProvisioningArtifactInput) (*servicecatalog.DescribeProvisioningArtifactOutput, error)
	DescribeProvisioningArtifactWithContext(aws.Context, *servicecatalog.DescribeProvisioningArtifactInput, ...request.Option) (*servicecatalog.DescribeProvisioningArtifactOutput, error)
	DescribeProvisioningArtifactRequest(*servicecatalog.DescribeProvisioningArtifactInput) (*request.Request, *servicecatalog.DescribeProvisioningArtifactOutput)

	DescribeProvisioningParameters(*servicecatalog.DescribeProvisioningParametersInput) (*servicecatalog.DescribeProvisioningParametersOutput, error)
	DescribeProvisioningParametersWithContext(aws.Context, *servicecatalog.DescribeProvisioningParametersInput, ...request.Option) (*servicecatalog.DescribeProvisioningParametersOutput, error)
	DescribeProvisioningParametersRequest(*servicecatalog.DescribeProvisioningParametersInput) (*request.Request, *servicecatalog.DescribeProvisioningParametersOutput)

	DescribeRecord(*servicecatalog.DescribeRecordInput) (*servicecatalog.DescribeRecordOutput, error)
	DescribeRecordWithContext(aws.Context, *servicecatalog.DescribeRecordInput, ...request.Option) (*servicecatalog.DescribeRecordOutput, error)
	DescribeRecordRequest(*servicecatalog.DescribeRecordInput) (*request.Request, *servicecatalog.DescribeRecordOutput)

	DescribeTagOption(*servicecatalog.DescribeTagOptionInput) (*servicecatalog.DescribeTagOptionOutput, error)
	DescribeTagOptionWithContext(aws.Context, *servicecatalog.DescribeTagOptionInput, ...request.Option) (*servicecatalog.DescribeTagOptionOutput, error)
	DescribeTagOptionRequest(*servicecatalog.DescribeTagOptionInput) (*request.Request, *servicecatalog.DescribeTagOptionOutput)

	DisassociatePrincipalFromPortfolio(*servicecatalog.DisassociatePrincipalFromPortfolioInput) (*servicecatalog.DisassociatePrincipalFromPortfolioOutput, error)
	DisassociatePrincipalFromPortfolioWithContext(aws.Context, *servicecatalog.DisassociatePrincipalFromPortfolioInput, ...request.Option) (*servicecatalog.DisassociatePrincipalFromPortfolioOutput, error)
	DisassociatePrincipalFromPortfolioRequest(*servicecatalog.DisassociatePrincipalFromPortfolioInput) (*request.Request, *servicecatalog.DisassociatePrincipalFromPortfolioOutput)

	DisassociateProductFromPortfolio(*servicecatalog.DisassociateProductFromPortfolioInput) (*servicecatalog.DisassociateProductFromPortfolioOutput, error)
	DisassociateProductFromPortfolioWithContext(aws.Context, *servicecatalog.DisassociateProductFromPortfolioInput, ...request.Option) (*servicecatalog.DisassociateProductFromPortfolioOutput, error)
	DisassociateProductFromPortfolioRequest(*servicecatalog.DisassociateProductFromPortfolioInput) (*request.Request, *servicecatalog.DisassociateProductFromPortfolioOutput)

	DisassociateTagOptionFromResource(*servicecatalog.DisassociateTagOptionFromResourceInput) (*servicecatalog.DisassociateTagOptionFromResourceOutput, error)
	DisassociateTagOptionFromResourceWithContext(aws.Context, *servicecatalog.DisassociateTagOptionFromResourceInput, ...request.Option) (*servicecatalog.DisassociateTagOptionFromResourceOutput, error)
	DisassociateTagOptionFromResourceRequest(*servicecatalog.DisassociateTagOptionFromResourceInput) (*request.Request, *servicecatalog.DisassociateTagOptionFromResourceOutput)

	ListAcceptedPortfolioShares(*servicecatalog.ListAcceptedPortfolioSharesInput) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error)
	ListAcceptedPortfolioSharesWithContext(aws.Context, *servicecatalog.ListAcceptedPortfolioSharesInput, ...request.Option) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error)
	ListAcceptedPortfolioSharesRequest(*servicecatalog.ListAcceptedPortfolioSharesInput) (*request.Request, *servicecatalog.ListAcceptedPortfolioSharesOutput)

	ListConstraintsForPortfolio(*servicecatalog.ListConstraintsForPortfolioInput) (*servicecatalog.ListConstraintsForPortfolioOutput, error)
	ListConstraintsForPortfolioWithContext(aws.Context, *servicecatalog.ListConstraintsForPortfolioInput, ...request.Option) (*servicecatalog.ListConstraintsForPortfolioOutput, error)
	ListConstraintsForPortfolioRequest(*servicecatalog.ListConstraintsForPortfolioInput) (*request.Request, *servicecatalog.ListConstraintsForPortfolioOutput)

	ListLaunchPaths(*servicecatalog.ListLaunchPathsInput) (*servicecatalog.ListLaunchPathsOutput, error)
	ListLaunchPathsWithContext(aws.Context, *servicecatalog.ListLaunchPathsInput, ...request.Option) (*servicecatalog.ListLaunchPathsOutput, error)
	ListLaunchPathsRequest(*servicecatalog.ListLaunchPathsInput) (*request.Request, *servicecatalog.ListLaunchPathsOutput)

	ListPortfolioAccess(*servicecatalog.ListPortfolioAccessInput) (*servicecatalog.ListPortfolioAccessOutput, error)
	ListPortfolioAccessWithContext(aws.Context, *servicecatalog.ListPortfolioAccessInput, ...request.Option) (*servicecatalog.ListPortfolioAccessOutput, error)
	ListPortfolioAccessRequest(*servicecatalog.ListPortfolioAccessInput) (*request.Request, *servicecatalog.ListPortfolioAccessOutput)

	ListPortfolios(*servicecatalog.ListPortfoliosInput) (*servicecatalog.ListPortfoliosOutput, error)
	ListPortfoliosWithContext(aws.Context, *servicecatalog.ListPortfoliosInput, ...request.Option) (*servicecatalog.ListPortfoliosOutput, error)
	ListPortfoliosRequest(*servicecatalog.ListPortfoliosInput) (*request.Request, *servicecatalog.ListPortfoliosOutput)

	ListPortfoliosForProduct(*servicecatalog.ListPortfoliosForProductInput) (*servicecatalog.ListPortfoliosForProductOutput, error)
	ListPortfoliosForProductWithContext(aws.Context, *servicecatalog.ListPortfoliosForProductInput, ...request.Option) (*servicecatalog.ListPortfoliosForProductOutput, error)
	ListPortfoliosForProductRequest(*servicecatalog.ListPortfoliosForProductInput) (*request.Request, *servicecatalog.ListPortfoliosForProductOutput)

	ListPrincipalsForPortfolio(*servicecatalog.ListPrincipalsForPortfolioInput) (*servicecatalog.ListPrincipalsForPortfolioOutput, error)
	ListPrincipalsForPortfolioWithContext(aws.Context, *servicecatalog.ListPrincipalsForPortfolioInput, ...request.Option) (*servicecatalog.ListPrincipalsForPortfolioOutput, error)
	ListPrincipalsForPortfolioRequest(*servicecatalog.ListPrincipalsForPortfolioInput) (*request.Request, *servicecatalog.ListPrincipalsForPortfolioOutput)

	ListProvisioningArtifacts(*servicecatalog.ListProvisioningArtifactsInput) (*servicecatalog.ListProvisioningArtifactsOutput, error)
	ListProvisioningArtifactsWithContext(aws.Context, *servicecatalog.ListProvisioningArtifactsInput, ...request.Option) (*servicecatalog.ListProvisioningArtifactsOutput, error)
	ListProvisioningArtifactsRequest(*servicecatalog.ListProvisioningArtifactsInput) (*request.Request, *servicecatalog.ListProvisioningArtifactsOutput)

	ListRecordHistory(*servicecatalog.ListRecordHistoryInput) (*servicecatalog.ListRecordHistoryOutput, error)
	ListRecordHistoryWithContext(aws.Context, *servicecatalog.ListRecordHistoryInput, ...request.Option) (*servicecatalog.ListRecordHistoryOutput, error)
	ListRecordHistoryRequest(*servicecatalog.ListRecordHistoryInput) (*request.Request, *servicecatalog.ListRecordHistoryOutput)

	ListResourcesForTagOption(*servicecatalog.ListResourcesForTagOptionInput) (*servicecatalog.ListResourcesForTagOptionOutput, error)
	ListResourcesForTagOptionWithContext(aws.Context, *servicecatalog.ListResourcesForTagOptionInput, ...request.Option) (*servicecatalog.ListResourcesForTagOptionOutput, error)
	ListResourcesForTagOptionRequest(*servicecatalog.ListResourcesForTagOptionInput) (*request.Request, *servicecatalog.ListResourcesForTagOptionOutput)

	ListResourcesForTagOptionPages(*servicecatalog.ListResourcesForTagOptionInput, func(*servicecatalog.ListResourcesForTagOptionOutput, bool) bool) error
	ListResourcesForTagOptionPagesWithContext(aws.Context, *servicecatalog.ListResourcesForTagOptionInput, func(*servicecatalog.ListResourcesForTagOptionOutput, bool) bool, ...request.Option) error

	ListTagOptions(*servicecatalog.ListTagOptionsInput) (*servicecatalog.ListTagOptionsOutput, error)
	ListTagOptionsWithContext(aws.Context, *servicecatalog.ListTagOptionsInput, ...request.Option) (*servicecatalog.ListTagOptionsOutput, error)
	ListTagOptionsRequest(*servicecatalog.ListTagOptionsInput) (*request.Request, *servicecatalog.ListTagOptionsOutput)

	ListTagOptionsPages(*servicecatalog.ListTagOptionsInput, func(*servicecatalog.ListTagOptionsOutput, bool) bool) error
	ListTagOptionsPagesWithContext(aws.Context, *servicecatalog.ListTagOptionsInput, func(*servicecatalog.ListTagOptionsOutput, bool) bool, ...request.Option) error

	ProvisionProduct(*servicecatalog.ProvisionProductInput) (*servicecatalog.ProvisionProductOutput, error)
	ProvisionProductWithContext(aws.Context, *servicecatalog.ProvisionProductInput, ...request.Option) (*servicecatalog.ProvisionProductOutput, error)
	ProvisionProductRequest(*servicecatalog.ProvisionProductInput) (*request.Request, *servicecatalog.ProvisionProductOutput)

	RejectPortfolioShare(*servicecatalog.RejectPortfolioShareInput) (*servicecatalog.RejectPortfolioShareOutput, error)
	RejectPortfolioShareWithContext(aws.Context, *servicecatalog.RejectPortfolioShareInput, ...request.Option) (*servicecatalog.RejectPortfolioShareOutput, error)
	RejectPortfolioShareRequest(*servicecatalog.RejectPortfolioShareInput) (*request.Request, *servicecatalog.RejectPortfolioShareOutput)

	ScanProvisionedProducts(*servicecatalog.ScanProvisionedProductsInput) (*servicecatalog.ScanProvisionedProductsOutput, error)
	ScanProvisionedProductsWithContext(aws.Context, *servicecatalog.ScanProvisionedProductsInput, ...request.Option) (*servicecatalog.ScanProvisionedProductsOutput, error)
	ScanProvisionedProductsRequest(*servicecatalog.ScanProvisionedProductsInput) (*request.Request, *servicecatalog.ScanProvisionedProductsOutput)

	SearchProducts(*servicecatalog.SearchProductsInput) (*servicecatalog.SearchProductsOutput, error)
	SearchProductsWithContext(aws.Context, *servicecatalog.SearchProductsInput, ...request.Option) (*servicecatalog.SearchProductsOutput, error)
	SearchProductsRequest(*servicecatalog.SearchProductsInput) (*request.Request, *servicecatalog.SearchProductsOutput)

	SearchProductsAsAdmin(*servicecatalog.SearchProductsAsAdminInput) (*servicecatalog.SearchProductsAsAdminOutput, error)
	SearchProductsAsAdminWithContext(aws.Context, *servicecatalog.SearchProductsAsAdminInput, ...request.Option) (*servicecatalog.SearchProductsAsAdminOutput, error)
	SearchProductsAsAdminRequest(*servicecatalog.SearchProductsAsAdminInput) (*request.Request, *servicecatalog.SearchProductsAsAdminOutput)

	TerminateProvisionedProduct(*servicecatalog.TerminateProvisionedProductInput) (*servicecatalog.TerminateProvisionedProductOutput, error)
	TerminateProvisionedProductWithContext(aws.Context, *servicecatalog.TerminateProvisionedProductInput, ...request.Option) (*servicecatalog.TerminateProvisionedProductOutput, error)
	TerminateProvisionedProductRequest(*servicecatalog.TerminateProvisionedProductInput) (*request.Request, *servicecatalog.TerminateProvisionedProductOutput)

	UpdateConstraint(*servicecatalog.UpdateConstraintInput) (*servicecatalog.UpdateConstraintOutput, error)
	UpdateConstraintWithContext(aws.Context, *servicecatalog.UpdateConstraintInput, ...request.Option) (*servicecatalog.UpdateConstraintOutput, error)
	UpdateConstraintRequest(*servicecatalog.UpdateConstraintInput) (*request.Request, *servicecatalog.UpdateConstraintOutput)

	UpdatePortfolio(*servicecatalog.UpdatePortfolioInput) (*servicecatalog.UpdatePortfolioOutput, error)
	UpdatePortfolioWithContext(aws.Context, *servicecatalog.UpdatePortfolioInput, ...request.Option) (*servicecatalog.UpdatePortfolioOutput, error)
	UpdatePortfolioRequest(*servicecatalog.UpdatePortfolioInput) (*request.Request, *servicecatalog.UpdatePortfolioOutput)

	UpdateProduct(*servicecatalog.UpdateProductInput) (*servicecatalog.UpdateProductOutput, error)
	UpdateProductWithContext(aws.Context, *servicecatalog.UpdateProductInput, ...request.Option) (*servicecatalog.UpdateProductOutput, error)
	UpdateProductRequest(*servicecatalog.UpdateProductInput) (*request.Request, *servicecatalog.UpdateProductOutput)

	UpdateProvisionedProduct(*servicecatalog.UpdateProvisionedProductInput) (*servicecatalog.UpdateProvisionedProductOutput, error)
	UpdateProvisionedProductWithContext(aws.Context, *servicecatalog.UpdateProvisionedProductInput, ...request.Option) (*servicecatalog.UpdateProvisionedProductOutput, error)
	UpdateProvisionedProductRequest(*servicecatalog.UpdateProvisionedProductInput) (*request.Request, *servicecatalog.UpdateProvisionedProductOutput)

	UpdateProvisioningArtifact(*servicecatalog.UpdateProvisioningArtifactInput) (*servicecatalog.UpdateProvisioningArtifactOutput, error)
	UpdateProvisioningArtifactWithContext(aws.Context, *servicecatalog.UpdateProvisioningArtifactInput, ...request.Option) (*servicecatalog.UpdateProvisioningArtifactOutput, error)
	UpdateProvisioningArtifactRequest(*servicecatalog.UpdateProvisioningArtifactInput) (*request.Request, *servicecatalog.UpdateProvisioningArtifactOutput)

	UpdateTagOption(*servicecatalog.UpdateTagOptionInput) (*servicecatalog.UpdateTagOptionOutput, error)
	UpdateTagOptionWithContext(aws.Context, *servicecatalog.UpdateTagOptionInput, ...request.Option) (*servicecatalog.UpdateTagOptionOutput, error)
	UpdateTagOptionRequest(*servicecatalog.UpdateTagOptionInput) (*request.Request, *servicecatalog.UpdateTagOptionOutput)
}

var _ ServiceCatalogAPI = (*servicecatalog.ServiceCatalog)(nil)
