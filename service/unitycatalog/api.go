// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Metastore Assignments, Account Metastores, Account Storage Credentials, Catalogs, External Locations, Functions, Grants, Metastores, Providers, Recipient Activation, Recipients, Schemas, Shares, Storage Credentials, Table Constraints, Tables, etc.
package unitycatalog

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewAccountMetastoreAssignments(client *client.DatabricksClient) *AccountMetastoreAssignmentsAPI {
	return &AccountMetastoreAssignmentsAPI{
		impl: &accountMetastoreAssignmentsImpl{
			client: client,
		},
	}
}

// These APIs manage metastore assignments to a workspace.
type AccountMetastoreAssignmentsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(AccountMetastoreAssignmentsService)
	impl AccountMetastoreAssignmentsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *AccountMetastoreAssignmentsAPI) WithImpl(impl AccountMetastoreAssignmentsService) *AccountMetastoreAssignmentsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level AccountMetastoreAssignments API implementation
func (a *AccountMetastoreAssignmentsAPI) Impl() AccountMetastoreAssignmentsService {
	return a.impl
}

// Assigns a workspace to a metastore.
//
// Creates an assignment to a metastore for a workspace
func (a *AccountMetastoreAssignmentsAPI) Create(ctx context.Context, request CreateMetastoreAssignment) (*MetastoreAssignment, error) {
	return a.impl.Create(ctx, request)
}

// Delete a metastore assignment.
//
// Deletes a metastore assignment to a workspace, leaving the workspace with no
// metastore.
func (a *AccountMetastoreAssignmentsAPI) Delete(ctx context.Context, request DeleteAccountMetastoreAssignmentRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a metastore assignment.
//
// Deletes a metastore assignment to a workspace, leaving the workspace with no
// metastore.
func (a *AccountMetastoreAssignmentsAPI) DeleteByWorkspaceIdAndMetastoreId(ctx context.Context, workspaceId int64, metastoreId string) error {
	return a.impl.Delete(ctx, DeleteAccountMetastoreAssignmentRequest{
		WorkspaceId: workspaceId,
		MetastoreId: metastoreId,
	})
}

// Gets the metastore assignment for a workspace.
//
// Gets the metastore assignment, if any, for the workspace specified by ID. If
// the workspace is assigned a metastore, the mappig will be returned. If no
// metastore is assigned to the workspace, the assignment will not be found and
// a 404 returned.
func (a *AccountMetastoreAssignmentsAPI) Get(ctx context.Context, request GetAccountMetastoreAssignmentRequest) (*MetastoreAssignment, error) {
	return a.impl.Get(ctx, request)
}

// Gets the metastore assignment for a workspace.
//
// Gets the metastore assignment, if any, for the workspace specified by ID. If
// the workspace is assigned a metastore, the mappig will be returned. If no
// metastore is assigned to the workspace, the assignment will not be found and
// a 404 returned.
func (a *AccountMetastoreAssignmentsAPI) GetByWorkspaceId(ctx context.Context, workspaceId int64) (*MetastoreAssignment, error) {
	return a.impl.Get(ctx, GetAccountMetastoreAssignmentRequest{
		WorkspaceId: workspaceId,
	})
}

// Get all workspaces assigned to a metastore.
//
// Gets a list of all Databricks workspace IDs that have been assigned to given
// metastore.
func (a *AccountMetastoreAssignmentsAPI) List(ctx context.Context, request ListAccountMetastoreAssignmentsRequest) ([]MetastoreAssignment, error) {
	return a.impl.List(ctx, request)
}

// Get all workspaces assigned to a metastore.
//
// Gets a list of all Databricks workspace IDs that have been assigned to given
// metastore.
func (a *AccountMetastoreAssignmentsAPI) ListByMetastoreId(ctx context.Context, metastoreId string) ([]MetastoreAssignment, error) {
	return a.impl.List(ctx, ListAccountMetastoreAssignmentsRequest{
		MetastoreId: metastoreId,
	})
}

// Updates a metastore assignment to a workspaces.
//
// Updates an assignment to a metastore for a workspace. Currently, only the
// default catalog may be updated
func (a *AccountMetastoreAssignmentsAPI) Update(ctx context.Context, request UpdateMetastoreAssignment) (*MetastoreAssignment, error) {
	return a.impl.Update(ctx, request)
}

func NewAccountMetastores(client *client.DatabricksClient) *AccountMetastoresAPI {
	return &AccountMetastoresAPI{
		impl: &accountMetastoresImpl{
			client: client,
		},
	}
}

// These APIs manage Unity Catalog metastores for an account. A metastore
// contains catalogs that can be associated with workspaces
type AccountMetastoresAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(AccountMetastoresService)
	impl AccountMetastoresService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *AccountMetastoresAPI) WithImpl(impl AccountMetastoresService) *AccountMetastoresAPI {
	a.impl = impl
	return a
}

// Impl returns low-level AccountMetastores API implementation
func (a *AccountMetastoresAPI) Impl() AccountMetastoresService {
	return a.impl
}

// Create metastore.
//
// Creates a Unity Catalog metastore.
func (a *AccountMetastoresAPI) Create(ctx context.Context, request CreateMetastore) (*MetastoreInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete a metastore.
//
// Deletes a Databricks Unity Catalog metastore for an account, both specified
// by ID.
func (a *AccountMetastoresAPI) Delete(ctx context.Context, request DeleteAccountMetastoreRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a metastore.
//
// Deletes a Databricks Unity Catalog metastore for an account, both specified
// by ID.
func (a *AccountMetastoresAPI) DeleteByMetastoreId(ctx context.Context, metastoreId string) error {
	return a.impl.Delete(ctx, DeleteAccountMetastoreRequest{
		MetastoreId: metastoreId,
	})
}

// Get a metastore.
//
// Gets a Databricks Unity Catalog metastore from an account, both specified by
// ID.
func (a *AccountMetastoresAPI) Get(ctx context.Context, request GetAccountMetastoreRequest) (*MetastoreInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a metastore.
//
// Gets a Databricks Unity Catalog metastore from an account, both specified by
// ID.
func (a *AccountMetastoresAPI) GetByMetastoreId(ctx context.Context, metastoreId string) (*MetastoreInfo, error) {
	return a.impl.Get(ctx, GetAccountMetastoreRequest{
		MetastoreId: metastoreId,
	})
}

// Get all metastores associated with an account.
//
// Gets all Unity Catalog metastores associated with an account specified by ID.
func (a *AccountMetastoresAPI) List(ctx context.Context) (*ListMetastoresResponse, error) {
	return a.impl.List(ctx)
}

// Update a metastore.
//
// Updates an existing Unity Catalog metastore.
func (a *AccountMetastoresAPI) Update(ctx context.Context, request UpdateMetastore) (*MetastoreInfo, error) {
	return a.impl.Update(ctx, request)
}

func NewAccountStorageCredentials(client *client.DatabricksClient) *AccountStorageCredentialsAPI {
	return &AccountStorageCredentialsAPI{
		impl: &accountStorageCredentialsImpl{
			client: client,
		},
	}
}

// These APIs manage storage credentials for a particular metastore.
type AccountStorageCredentialsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(AccountStorageCredentialsService)
	impl AccountStorageCredentialsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *AccountStorageCredentialsAPI) WithImpl(impl AccountStorageCredentialsService) *AccountStorageCredentialsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level AccountStorageCredentials API implementation
func (a *AccountStorageCredentialsAPI) Impl() AccountStorageCredentialsService {
	return a.impl
}

// Create a storage credential.
//
// Creates a new storage credential. The request object is specific to the
// cloud:
//
// * **AwsIamRole** for AWS credentials * **AzureServicePrincipal** for Azure
// credentials * **GcpServiceAcountKey** for GCP credentials.
//
// The caller must be a metastore admin and have the
// **CREATE_STORAGE_CREDENTIAL** privilege on the metastore.
func (a *AccountStorageCredentialsAPI) Create(ctx context.Context, request CreateStorageCredential) (*StorageCredentialInfo, error) {
	return a.impl.Create(ctx, request)
}

// Gets the named storage credential.
//
// Gets a storage credential from the metastore. The caller must be a metastore
// admin, the owner of the storage credential, or have a level of privilege on
// the storage credential.
func (a *AccountStorageCredentialsAPI) Get(ctx context.Context, request GetAccountStorageCredentialRequest) (*StorageCredentialInfo, error) {
	return a.impl.Get(ctx, request)
}

// Gets the named storage credential.
//
// Gets a storage credential from the metastore. The caller must be a metastore
// admin, the owner of the storage credential, or have a level of privilege on
// the storage credential.
func (a *AccountStorageCredentialsAPI) GetByMetastoreIdAndStorageCredentialName(ctx context.Context, metastoreId string, storageCredentialName string) (*StorageCredentialInfo, error) {
	return a.impl.Get(ctx, GetAccountStorageCredentialRequest{
		MetastoreId:           metastoreId,
		StorageCredentialName: storageCredentialName,
	})
}

// Get all storage credentials assigned to a metastore.
//
// Gets a list of all storage credentials that have been assigned to given
// metastore.
func (a *AccountStorageCredentialsAPI) List(ctx context.Context, request ListAccountStorageCredentialsRequest) ([]StorageCredentialInfo, error) {
	return a.impl.List(ctx, request)
}

// Get all storage credentials assigned to a metastore.
//
// Gets a list of all storage credentials that have been assigned to given
// metastore.
func (a *AccountStorageCredentialsAPI) ListByMetastoreId(ctx context.Context, metastoreId string) ([]StorageCredentialInfo, error) {
	return a.impl.List(ctx, ListAccountStorageCredentialsRequest{
		MetastoreId: metastoreId,
	})
}

func NewCatalogs(client *client.DatabricksClient) *CatalogsAPI {
	return &CatalogsAPI{
		impl: &catalogsImpl{
			client: client,
		},
	}
}

// A catalog is the first layer of Unity Catalog’s three-level namespace.
// It’s used to organize your data assets. Users can see all catalogs on which
// they have been assigned the USE_CATALOG data permission.
//
// In Unity Catalog, admins and data stewards manage users and their access to
// data centrally across all of the workspaces in a Databricks account. Users in
// different workspaces can share access to the same data, depending on
// privileges granted centrally in Unity Catalog.
type CatalogsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(CatalogsService)
	impl CatalogsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *CatalogsAPI) WithImpl(impl CatalogsService) *CatalogsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Catalogs API implementation
func (a *CatalogsAPI) Impl() CatalogsService {
	return a.impl
}

// Create a catalog.
//
// Creates a new catalog instance in the parent metastore if the caller is a
// metastore admin or has the **CREATE_CATALOG** privilege.
func (a *CatalogsAPI) Create(ctx context.Context, request CreateCatalog) (*CatalogInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete a catalog.
//
// Deletes the catalog that matches the supplied name. The caller must be a
// metastore admin or the owner of the catalog.
func (a *CatalogsAPI) Delete(ctx context.Context, request DeleteCatalogRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a catalog.
//
// Deletes the catalog that matches the supplied name. The caller must be a
// metastore admin or the owner of the catalog.
func (a *CatalogsAPI) DeleteByName(ctx context.Context, name string) error {
	return a.impl.Delete(ctx, DeleteCatalogRequest{
		Name: name,
	})
}

// Get a catalog.
//
// Gets the specified catalog in a metastore. The caller must be a metastore
// admin, the owner of the catalog, or a user that has the **USE_CATALOG**
// privilege set for their account.
func (a *CatalogsAPI) Get(ctx context.Context, request GetCatalogRequest) (*CatalogInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a catalog.
//
// Gets the specified catalog in a metastore. The caller must be a metastore
// admin, the owner of the catalog, or a user that has the **USE_CATALOG**
// privilege set for their account.
func (a *CatalogsAPI) GetByName(ctx context.Context, name string) (*CatalogInfo, error) {
	return a.impl.Get(ctx, GetCatalogRequest{
		Name: name,
	})
}

// List catalogs.
//
// Gets an array of catalogs in the metastore. If the caller is the metastore
// admin, all catalogs will be retrieved. Otherwise, only catalogs owned by the
// caller (or for which the caller has the **USE_CATALOG** privilege) will be
// retrieved. There is no guarantee of a specific ordering of the elements in
// the array.
//
// This method is generated by Databricks SDK Code Generator.
func (a *CatalogsAPI) ListAll(ctx context.Context) ([]CatalogInfo, error) {
	response, err := a.impl.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.Catalogs, nil
}

// Update a catalog.
//
// Updates the catalog that matches the supplied name. The caller must be either
// the owner of the catalog, or a metastore admin (when changing the owner field
// of the catalog).
func (a *CatalogsAPI) Update(ctx context.Context, request UpdateCatalog) (*CatalogInfo, error) {
	return a.impl.Update(ctx, request)
}

func NewExternalLocations(client *client.DatabricksClient) *ExternalLocationsAPI {
	return &ExternalLocationsAPI{
		impl: &externalLocationsImpl{
			client: client,
		},
	}
}

// An external location is an object that combines a cloud storage path with a
// storage credential that authorizes access to the cloud storage path. Each
// external location is subject to Unity Catalog access-control policies that
// control which users and groups can access the credential. If a user does not
// have access to an external location in Unity Catalog, the request fails and
// Unity Catalog does not attempt to authenticate to your cloud tenant on the
// user’s behalf.
//
// Databricks recommends using external locations rather than using storage
// credentials directly.
//
// To create external locations, you must be a metastore admin or a user with
// the **CREATE_EXTERNAL_LOCATION** privilege.
type ExternalLocationsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ExternalLocationsService)
	impl ExternalLocationsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ExternalLocationsAPI) WithImpl(impl ExternalLocationsService) *ExternalLocationsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level ExternalLocations API implementation
func (a *ExternalLocationsAPI) Impl() ExternalLocationsService {
	return a.impl
}

// Create an external location.
//
// Creates a new external location entry in the metastore. The caller must be a
// metastore admin or have the **CREATE_EXTERNAL_LOCATION** privilege on both
// the metastore and the associated storage credential.
func (a *ExternalLocationsAPI) Create(ctx context.Context, request CreateExternalLocation) (*ExternalLocationInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete an external location.
//
// Deletes the specified external location from the metastore. The caller must
// be the owner of the external location.
func (a *ExternalLocationsAPI) Delete(ctx context.Context, request DeleteExternalLocationRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete an external location.
//
// Deletes the specified external location from the metastore. The caller must
// be the owner of the external location.
func (a *ExternalLocationsAPI) DeleteByName(ctx context.Context, name string) error {
	return a.impl.Delete(ctx, DeleteExternalLocationRequest{
		Name: name,
	})
}

// Get an external location.
//
// Gets an external location from the metastore. The caller must be either a
// metastore admin, the owner of the external location, or a user that has some
// privilege on the external location.
func (a *ExternalLocationsAPI) Get(ctx context.Context, request GetExternalLocationRequest) (*ExternalLocationInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get an external location.
//
// Gets an external location from the metastore. The caller must be either a
// metastore admin, the owner of the external location, or a user that has some
// privilege on the external location.
func (a *ExternalLocationsAPI) GetByName(ctx context.Context, name string) (*ExternalLocationInfo, error) {
	return a.impl.Get(ctx, GetExternalLocationRequest{
		Name: name,
	})
}

// List external locations.
//
// Gets an array of external locations (__ExternalLocationInfo__ objects) from
// the metastore. The caller must be a metastore admin, the owner of the
// external location, or a user that has some privilege on the external
// location. There is no guarantee of a specific ordering of the elements in the
// array.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ExternalLocationsAPI) ListAll(ctx context.Context) ([]ExternalLocationInfo, error) {
	response, err := a.impl.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.ExternalLocations, nil
}

// Update an external location.
//
// Updates an external location in the metastore. The caller must be the owner
// of the external location, or be a metastore admin. In the second case, the
// admin can only update the name of the external location.
func (a *ExternalLocationsAPI) Update(ctx context.Context, request UpdateExternalLocation) (*ExternalLocationInfo, error) {
	return a.impl.Update(ctx, request)
}

func NewFunctions(client *client.DatabricksClient) *FunctionsAPI {
	return &FunctionsAPI{
		impl: &functionsImpl{
			client: client,
		},
	}
}

// Functions implement User-Defined Functions (UDFs) in Unity Catalog.
//
// The function implementation can be any SQL expression or Query, and it can be
// invoked wherever a table reference is allowed in a query. In Unity Catalog, a
// function resides at the same level as a table, so it can be referenced with
// the form __catalog_name__.__schema_name__.__function_name__.
type FunctionsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(FunctionsService)
	impl FunctionsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *FunctionsAPI) WithImpl(impl FunctionsService) *FunctionsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Functions API implementation
func (a *FunctionsAPI) Impl() FunctionsService {
	return a.impl
}

// Create a function.
//
// # Creates a new function
//
// The user must have the following permissions in order for the function to be
// created: - **USE_CATALOG** on the function's parent catalog - **USE_SCHEMA**
// and **CREATE_FUNCTION** on the function's parent schema
func (a *FunctionsAPI) Create(ctx context.Context, request CreateFunction) (*FunctionInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete a function.
//
// Deletes the function that matches the supplied name. For the deletion to
// succeed, the user must satisfy one of the following conditions: - Is the
// owner of the function's parent catalog - Is the owner of the function's
// parent schema and have the **USE_CATALOG** privilege on its parent catalog -
// Is the owner of the function itself and have both the **USE_CATALOG**
// privilege on its parent catalog and the **USE_SCHEMA** privilege on its
// parent schema
func (a *FunctionsAPI) Delete(ctx context.Context, request DeleteFunctionRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a function.
//
// Deletes the function that matches the supplied name. For the deletion to
// succeed, the user must satisfy one of the following conditions: - Is the
// owner of the function's parent catalog - Is the owner of the function's
// parent schema and have the **USE_CATALOG** privilege on its parent catalog -
// Is the owner of the function itself and have both the **USE_CATALOG**
// privilege on its parent catalog and the **USE_SCHEMA** privilege on its
// parent schema
func (a *FunctionsAPI) DeleteByName(ctx context.Context, name string) error {
	return a.impl.Delete(ctx, DeleteFunctionRequest{
		Name: name,
	})
}

// Get a function.
//
// Gets a function from within a parent catalog and schema. For the fetch to
// succeed, the user must satisfy one of the following requirements: - Is a
// metastore admin - Is an owner of the function's parent catalog - Have the
// **USE_CATALOG** privilege on the function's parent catalog and be the owner
// of the function - Have the **USE_CATALOG** privilege on the function's parent
// catalog, the **USE_SCHEMA** privilege on the function's parent schema, and
// the **EXECUTE** privilege on the function itself
func (a *FunctionsAPI) Get(ctx context.Context, request GetFunctionRequest) (*FunctionInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a function.
//
// Gets a function from within a parent catalog and schema. For the fetch to
// succeed, the user must satisfy one of the following requirements: - Is a
// metastore admin - Is an owner of the function's parent catalog - Have the
// **USE_CATALOG** privilege on the function's parent catalog and be the owner
// of the function - Have the **USE_CATALOG** privilege on the function's parent
// catalog, the **USE_SCHEMA** privilege on the function's parent schema, and
// the **EXECUTE** privilege on the function itself
func (a *FunctionsAPI) GetByName(ctx context.Context, name string) (*FunctionInfo, error) {
	return a.impl.Get(ctx, GetFunctionRequest{
		Name: name,
	})
}

// List functions.
//
// List functions within the specified parent catalog and schema. If the user is
// a metastore admin, all functions are returned in the output list. Otherwise,
// the user must have the **USE_CATALOG** privilege on the catalog and the
// **USE_SCHEMA** privilege on the schema, and the output list contains only
// functions for which either the user has the **EXECUTE** privilege or the user
// is the owner. There is no guarantee of a specific ordering of the elements in
// the array.
func (a *FunctionsAPI) List(ctx context.Context, request ListFunctionsRequest) (*ListFunctionsResponse, error) {
	return a.impl.List(ctx, request)
}

// Update a function.
//
// Updates the function that matches the supplied name. Only the owner of the
// function can be updated. If the user is not a metastore admin, the user must
// be a member of the group that is the new function owner. - Is a metastore
// admin - Is the owner of the function's parent catalog - Is the owner of the
// function's parent schema and has the **USE_CATALOG** privilege on its parent
// catalog - Is the owner of the function itself and has the **USE_CATALOG**
// privilege on its parent catalog as well as the **USE_SCHEMA** privilege on
// the function's parent schema.
func (a *FunctionsAPI) Update(ctx context.Context, request UpdateFunction) (*FunctionInfo, error) {
	return a.impl.Update(ctx, request)
}

func NewGrants(client *client.DatabricksClient) *GrantsAPI {
	return &GrantsAPI{
		impl: &grantsImpl{
			client: client,
		},
	}
}

// In Unity Catalog, data is secure by default. Initially, users have no access
// to data in a metastore. Access can be granted by either a metastore admin,
// the owner of an object, or the owner of the catalog or schema that contains
// the object. Securable objects in Unity Catalog are hierarchical and
// privileges are inherited downward.
//
// Securable objects in Unity Catalog are hierarchical and privileges are
// inherited downward. This means that granting a privilege on the catalog
// automatically grants the privilege to all current and future objects within
// the catalog. Similarly, privileges granted on a schema are inherited by all
// current and future objects within that schema.
type GrantsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(GrantsService)
	impl GrantsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *GrantsAPI) WithImpl(impl GrantsService) *GrantsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Grants API implementation
func (a *GrantsAPI) Impl() GrantsService {
	return a.impl
}

// Get permissions.
//
// Gets the permissions for a securable.
func (a *GrantsAPI) Get(ctx context.Context, request GetGrantRequest) (*PermissionsList, error) {
	return a.impl.Get(ctx, request)
}

// Get permissions.
//
// Gets the permissions for a securable.
func (a *GrantsAPI) GetBySecurableTypeAndFullName(ctx context.Context, securableType SecurableType, fullName string) (*PermissionsList, error) {
	return a.impl.Get(ctx, GetGrantRequest{
		SecurableType: securableType,
		FullName:      fullName,
	})
}

// Get effective permissions.
//
// Gets the effective permissions for a securable.
func (a *GrantsAPI) GetEffective(ctx context.Context, request GetEffectiveRequest) (*EffectivePermissionsList, error) {
	return a.impl.GetEffective(ctx, request)
}

// Get effective permissions.
//
// Gets the effective permissions for a securable.
func (a *GrantsAPI) GetEffectiveBySecurableTypeAndFullName(ctx context.Context, securableType SecurableType, fullName string) (*EffectivePermissionsList, error) {
	return a.impl.GetEffective(ctx, GetEffectiveRequest{
		SecurableType: securableType,
		FullName:      fullName,
	})
}

// Update permissions.
//
// Updates the permissions for a securable.
func (a *GrantsAPI) Update(ctx context.Context, request UpdatePermissions) (*PermissionsList, error) {
	return a.impl.Update(ctx, request)
}

func NewMetastores(client *client.DatabricksClient) *MetastoresAPI {
	return &MetastoresAPI{
		impl: &metastoresImpl{
			client: client,
		},
	}
}

// A metastore is the top-level container of objects in Unity Catalog. It stores
// data assets (tables and views) and the permissions that govern access to
// them. Databricks account admins can create metastores and assign them to
// Databricks workspaces to control which workloads use each metastore. For a
// workspace to use Unity Catalog, it must have a Unity Catalog metastore
// attached.
//
// Each metastore is configured with a root storage location in a cloud storage
// account. This storage location is used for metadata and managed tables data.
//
// NOTE: This metastore is distinct from the metastore included in Databricks
// workspaces created before Unity Catalog was released. If your workspace
// includes a legacy Hive metastore, the data in that metastore is available in
// a catalog named hive_metastore.
type MetastoresAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(MetastoresService)
	impl MetastoresService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *MetastoresAPI) WithImpl(impl MetastoresService) *MetastoresAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Metastores API implementation
func (a *MetastoresAPI) Impl() MetastoresService {
	return a.impl
}

// Create an assignment.
//
// Creates a new metastore assignment. If an assignment for the same
// __workspace_id__ exists, it will be overwritten by the new __metastore_id__
// and __default_catalog_name__. The caller must be an account admin.
func (a *MetastoresAPI) Assign(ctx context.Context, request CreateMetastoreAssignment) error {
	return a.impl.Assign(ctx, request)
}

// Create a metastore.
//
// Creates a new metastore based on a provided name and storage root path.
func (a *MetastoresAPI) Create(ctx context.Context, request CreateMetastore) (*MetastoreInfo, error) {
	return a.impl.Create(ctx, request)
}

// Get metastore assignment for workspace.
//
// Gets the metastore assignment for the workspace being accessed.
func (a *MetastoresAPI) Current(ctx context.Context) (*MetastoreAssignment, error) {
	return a.impl.Current(ctx)
}

// Delete a metastore.
//
// Deletes a metastore. The caller must be a metastore admin.
func (a *MetastoresAPI) Delete(ctx context.Context, request DeleteMetastoreRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a metastore.
//
// Deletes a metastore. The caller must be a metastore admin.
func (a *MetastoresAPI) DeleteById(ctx context.Context, id string) error {
	return a.impl.Delete(ctx, DeleteMetastoreRequest{
		Id: id,
	})
}

// Get a metastore.
//
// Gets a metastore that matches the supplied ID. The caller must be a metastore
// admin to retrieve this info.
func (a *MetastoresAPI) Get(ctx context.Context, request GetMetastoreRequest) (*MetastoreInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a metastore.
//
// Gets a metastore that matches the supplied ID. The caller must be a metastore
// admin to retrieve this info.
func (a *MetastoresAPI) GetById(ctx context.Context, id string) (*MetastoreInfo, error) {
	return a.impl.Get(ctx, GetMetastoreRequest{
		Id: id,
	})
}

// List metastores.
//
// Gets an array of the available metastores (as __MetastoreInfo__ objects). The
// caller must be an admin to retrieve this info. There is no guarantee of a
// specific ordering of the elements in the array.
//
// This method is generated by Databricks SDK Code Generator.
func (a *MetastoresAPI) ListAll(ctx context.Context) ([]MetastoreInfo, error) {
	response, err := a.impl.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.Metastores, nil
}

// MetastoreInfoNameToMetastoreIdMap calls [MetastoresAPI.ListAll] and creates a map of results with [MetastoreInfo].Name as key and [MetastoreInfo].MetastoreId as value.
//
// Returns an error if there's more than one [MetastoreInfo] with the same .Name.
//
// Note: All [MetastoreInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *MetastoresAPI) MetastoreInfoNameToMetastoreIdMap(ctx context.Context) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.MetastoreId
	}
	return mapping, nil
}

// GetByName calls [MetastoresAPI.MetastoreInfoNameToMetastoreIdMap] and returns a single [MetastoreInfo].
//
// Returns an error if there's more than one [MetastoreInfo] with the same .Name.
//
// Note: All [MetastoreInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *MetastoresAPI) GetByName(ctx context.Context, name string) (*MetastoreInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	tmp := map[string][]MetastoreInfo{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("MetastoreInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of MetastoreInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Get a metastore summary.
//
// Gets information about a metastore. This summary includes the storage
// credential, the cloud vendor, the cloud region, and the global metastore ID.
func (a *MetastoresAPI) Summary(ctx context.Context) (*GetMetastoreSummaryResponse, error) {
	return a.impl.Summary(ctx)
}

// Delete an assignment.
//
// Deletes a metastore assignment. The caller must be an account administrator.
func (a *MetastoresAPI) Unassign(ctx context.Context, request UnassignRequest) error {
	return a.impl.Unassign(ctx, request)
}

// Delete an assignment.
//
// Deletes a metastore assignment. The caller must be an account administrator.
func (a *MetastoresAPI) UnassignByWorkspaceId(ctx context.Context, workspaceId int64) error {
	return a.impl.Unassign(ctx, UnassignRequest{
		WorkspaceId: workspaceId,
	})
}

// Update a metastore.
//
// Updates information for a specific metastore. The caller must be a metastore
// admin.
func (a *MetastoresAPI) Update(ctx context.Context, request UpdateMetastore) (*MetastoreInfo, error) {
	return a.impl.Update(ctx, request)
}

// Update an assignment.
//
// Updates a metastore assignment. This operation can be used to update
// __metastore_id__ or __default_catalog_name__ for a specified Workspace, if
// the Workspace is already assigned a metastore. The caller must be an account
// admin to update __metastore_id__; otherwise, the caller can be a Workspace
// admin.
func (a *MetastoresAPI) UpdateAssignment(ctx context.Context, request UpdateMetastoreAssignment) error {
	return a.impl.UpdateAssignment(ctx, request)
}

func NewProviders(client *client.DatabricksClient) *ProvidersAPI {
	return &ProvidersAPI{
		impl: &providersImpl{
			client: client,
		},
	}
}

// Databricks Delta Sharing: Providers REST API
type ProvidersAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ProvidersService)
	impl ProvidersService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ProvidersAPI) WithImpl(impl ProvidersService) *ProvidersAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Providers API implementation
func (a *ProvidersAPI) Impl() ProvidersService {
	return a.impl
}

// Create an auth provider.
//
// Creates a new authentication provider minimally based on a name and
// authentication type. The caller must be an admin on the metastore.
func (a *ProvidersAPI) Create(ctx context.Context, request CreateProvider) (*ProviderInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete a provider.
//
// Deletes an authentication provider, if the caller is a metastore admin or is
// the owner of the provider.
func (a *ProvidersAPI) Delete(ctx context.Context, request DeleteProviderRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a provider.
//
// Deletes an authentication provider, if the caller is a metastore admin or is
// the owner of the provider.
func (a *ProvidersAPI) DeleteByName(ctx context.Context, name string) error {
	return a.impl.Delete(ctx, DeleteProviderRequest{
		Name: name,
	})
}

// Get a provider.
//
// Gets a specific authentication provider. The caller must supply the name of
// the provider, and must either be a metastore admin or the owner of the
// provider.
func (a *ProvidersAPI) Get(ctx context.Context, request GetProviderRequest) (*ProviderInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a provider.
//
// Gets a specific authentication provider. The caller must supply the name of
// the provider, and must either be a metastore admin or the owner of the
// provider.
func (a *ProvidersAPI) GetByName(ctx context.Context, name string) (*ProviderInfo, error) {
	return a.impl.Get(ctx, GetProviderRequest{
		Name: name,
	})
}

// List providers.
//
// Gets an array of available authentication providers. The caller must either
// be a metastore admin or the owner of the providers. Providers not owned by
// the caller are not included in the response. There is no guarantee of a
// specific ordering of the elements in the array.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ProvidersAPI) ListAll(ctx context.Context, request ListProvidersRequest) ([]ProviderInfo, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Providers, nil
}

// ProviderInfoNameToMetastoreIdMap calls [ProvidersAPI.ListAll] and creates a map of results with [ProviderInfo].Name as key and [ProviderInfo].MetastoreId as value.
//
// Returns an error if there's more than one [ProviderInfo] with the same .Name.
//
// Note: All [ProviderInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ProvidersAPI) ProviderInfoNameToMetastoreIdMap(ctx context.Context, request ListProvidersRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.MetastoreId
	}
	return mapping, nil
}

// List shares by Provider.
//
// Gets an array of a specified provider's shares within the metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner.
func (a *ProvidersAPI) ListShares(ctx context.Context, request ListSharesRequest) (*ListProviderSharesResponse, error) {
	return a.impl.ListShares(ctx, request)
}

// List shares by Provider.
//
// Gets an array of a specified provider's shares within the metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner.
func (a *ProvidersAPI) ListSharesByName(ctx context.Context, name string) (*ListProviderSharesResponse, error) {
	return a.impl.ListShares(ctx, ListSharesRequest{
		Name: name,
	})
}

// Update a provider.
//
// Updates the information for an authentication provider, if the caller is a
// metastore admin or is the owner of the provider. If the update changes the
// provider name, the caller must be both a metastore admin and the owner of the
// provider.
func (a *ProvidersAPI) Update(ctx context.Context, request UpdateProvider) (*ProviderInfo, error) {
	return a.impl.Update(ctx, request)
}

func NewRecipientActivation(client *client.DatabricksClient) *RecipientActivationAPI {
	return &RecipientActivationAPI{
		impl: &recipientActivationImpl{
			client: client,
		},
	}
}

// Databricks Delta Sharing: Recipient Activation REST API
type RecipientActivationAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(RecipientActivationService)
	impl RecipientActivationService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *RecipientActivationAPI) WithImpl(impl RecipientActivationService) *RecipientActivationAPI {
	a.impl = impl
	return a
}

// Impl returns low-level RecipientActivation API implementation
func (a *RecipientActivationAPI) Impl() RecipientActivationService {
	return a.impl
}

// Get a share activation URL.
//
// Gets an activation URL for a share.
func (a *RecipientActivationAPI) GetActivationUrlInfo(ctx context.Context, request GetActivationUrlInfoRequest) error {
	return a.impl.GetActivationUrlInfo(ctx, request)
}

// Get a share activation URL.
//
// Gets an activation URL for a share.
func (a *RecipientActivationAPI) GetActivationUrlInfoByActivationUrl(ctx context.Context, activationUrl string) error {
	return a.impl.GetActivationUrlInfo(ctx, GetActivationUrlInfoRequest{
		ActivationUrl: activationUrl,
	})
}

// Get an access token.
//
// Retrieve access token with an activation url. This is a public API without
// any authentication.
func (a *RecipientActivationAPI) RetrieveToken(ctx context.Context, request RetrieveTokenRequest) (*RetrieveTokenResponse, error) {
	return a.impl.RetrieveToken(ctx, request)
}

// Get an access token.
//
// Retrieve access token with an activation url. This is a public API without
// any authentication.
func (a *RecipientActivationAPI) RetrieveTokenByActivationUrl(ctx context.Context, activationUrl string) (*RetrieveTokenResponse, error) {
	return a.impl.RetrieveToken(ctx, RetrieveTokenRequest{
		ActivationUrl: activationUrl,
	})
}

func NewRecipients(client *client.DatabricksClient) *RecipientsAPI {
	return &RecipientsAPI{
		impl: &recipientsImpl{
			client: client,
		},
	}
}

// Databricks Delta Sharing: Recipients REST API
type RecipientsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(RecipientsService)
	impl RecipientsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *RecipientsAPI) WithImpl(impl RecipientsService) *RecipientsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Recipients API implementation
func (a *RecipientsAPI) Impl() RecipientsService {
	return a.impl
}

// Create a share recipient.
//
// Creates a new recipient with the delta sharing authentication type in the
// metastore. The caller must be a metastore admin or has the
// **CREATE_RECIPIENT** privilege on the metastore.
func (a *RecipientsAPI) Create(ctx context.Context, request CreateRecipient) (*RecipientInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete a share recipient.
//
// Deletes the specified recipient from the metastore. The caller must be the
// owner of the recipient.
func (a *RecipientsAPI) Delete(ctx context.Context, request DeleteRecipientRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a share recipient.
//
// Deletes the specified recipient from the metastore. The caller must be the
// owner of the recipient.
func (a *RecipientsAPI) DeleteByName(ctx context.Context, name string) error {
	return a.impl.Delete(ctx, DeleteRecipientRequest{
		Name: name,
	})
}

// Get a share recipient.
//
// Gets a share recipient from the metastore if:
//
// * the caller is the owner of the share recipient, or: * is a metastore admin
func (a *RecipientsAPI) Get(ctx context.Context, request GetRecipientRequest) (*RecipientInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a share recipient.
//
// Gets a share recipient from the metastore if:
//
// * the caller is the owner of the share recipient, or: * is a metastore admin
func (a *RecipientsAPI) GetByName(ctx context.Context, name string) (*RecipientInfo, error) {
	return a.impl.Get(ctx, GetRecipientRequest{
		Name: name,
	})
}

// List share recipients.
//
// Gets an array of all share recipients within the current metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner. There is no
// guarantee of a specific ordering of the elements in the array.
//
// This method is generated by Databricks SDK Code Generator.
func (a *RecipientsAPI) ListAll(ctx context.Context, request ListRecipientsRequest) ([]RecipientInfo, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Recipients, nil
}

// RecipientInfoNameToMetastoreIdMap calls [RecipientsAPI.ListAll] and creates a map of results with [RecipientInfo].Name as key and [RecipientInfo].MetastoreId as value.
//
// Returns an error if there's more than one [RecipientInfo] with the same .Name.
//
// Note: All [RecipientInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *RecipientsAPI) RecipientInfoNameToMetastoreIdMap(ctx context.Context, request ListRecipientsRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.MetastoreId
	}
	return mapping, nil
}

// Rotate a token.
//
// Refreshes the specified recipient's delta sharing authentication token with
// the provided token info. The caller must be the owner of the recipient.
func (a *RecipientsAPI) RotateToken(ctx context.Context, request RotateRecipientToken) (*RecipientInfo, error) {
	return a.impl.RotateToken(ctx, request)
}

// Get recipient share permissions.
//
// Gets the share permissions for the specified Recipient. The caller must be a
// metastore admin or the owner of the Recipient.
func (a *RecipientsAPI) SharePermissions(ctx context.Context, request SharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error) {
	return a.impl.SharePermissions(ctx, request)
}

// Get recipient share permissions.
//
// Gets the share permissions for the specified Recipient. The caller must be a
// metastore admin or the owner of the Recipient.
func (a *RecipientsAPI) SharePermissionsByName(ctx context.Context, name string) (*GetRecipientSharePermissionsResponse, error) {
	return a.impl.SharePermissions(ctx, SharePermissionsRequest{
		Name: name,
	})
}

// Update a share recipient.
//
// Updates an existing recipient in the metastore. The caller must be a
// metastore admin or the owner of the recipient. If the recipient name will be
// updated, the user must be both a metastore admin and the owner of the
// recipient.
func (a *RecipientsAPI) Update(ctx context.Context, request UpdateRecipient) error {
	return a.impl.Update(ctx, request)
}

func NewSchemas(client *client.DatabricksClient) *SchemasAPI {
	return &SchemasAPI{
		impl: &schemasImpl{
			client: client,
		},
	}
}

// A schema (also called a database) is the second layer of Unity Catalog’s
// three-level namespace. A schema organizes tables, views and functions. To
// access (or list) a table or view in a schema, users must have the USE_SCHEMA
// data permission on the schema and its parent catalog, and they must have the
// SELECT permission on the table or view.
type SchemasAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(SchemasService)
	impl SchemasService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *SchemasAPI) WithImpl(impl SchemasService) *SchemasAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Schemas API implementation
func (a *SchemasAPI) Impl() SchemasService {
	return a.impl
}

// Create a schema.
//
// Creates a new schema for catalog in the Metatastore. The caller must be a
// metastore admin, or have the **CREATE_SCHEMA** privilege in the parent
// catalog.
func (a *SchemasAPI) Create(ctx context.Context, request CreateSchema) (*SchemaInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete a schema.
//
// Deletes the specified schema from the parent catalog. The caller must be the
// owner of the schema or an owner of the parent catalog.
func (a *SchemasAPI) Delete(ctx context.Context, request DeleteSchemaRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a schema.
//
// Deletes the specified schema from the parent catalog. The caller must be the
// owner of the schema or an owner of the parent catalog.
func (a *SchemasAPI) DeleteByFullName(ctx context.Context, fullName string) error {
	return a.impl.Delete(ctx, DeleteSchemaRequest{
		FullName: fullName,
	})
}

// Get a schema.
//
// Gets the specified schema within the metastore. The caller must be a
// metastore admin, the owner of the schema, or a user that has the
// **USE_SCHEMA** privilege on the schema.
func (a *SchemasAPI) Get(ctx context.Context, request GetSchemaRequest) (*SchemaInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a schema.
//
// Gets the specified schema within the metastore. The caller must be a
// metastore admin, the owner of the schema, or a user that has the
// **USE_SCHEMA** privilege on the schema.
func (a *SchemasAPI) GetByFullName(ctx context.Context, fullName string) (*SchemaInfo, error) {
	return a.impl.Get(ctx, GetSchemaRequest{
		FullName: fullName,
	})
}

// List schemas.
//
// Gets an array of schemas for a catalog in the metastore. If the caller is the
// metastore admin or the owner of the parent catalog, all schemas for the
// catalog will be retrieved. Otherwise, only schemas owned by the caller (or
// for which the caller has the **USE_SCHEMA** privilege) will be retrieved.
// There is no guarantee of a specific ordering of the elements in the array.
//
// This method is generated by Databricks SDK Code Generator.
func (a *SchemasAPI) ListAll(ctx context.Context, request ListSchemasRequest) ([]SchemaInfo, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Schemas, nil
}

// SchemaInfoNameToFullNameMap calls [SchemasAPI.ListAll] and creates a map of results with [SchemaInfo].Name as key and [SchemaInfo].FullName as value.
//
// Returns an error if there's more than one [SchemaInfo] with the same .Name.
//
// Note: All [SchemaInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *SchemasAPI) SchemaInfoNameToFullNameMap(ctx context.Context, request ListSchemasRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.FullName
	}
	return mapping, nil
}

// GetByName calls [SchemasAPI.SchemaInfoNameToFullNameMap] and returns a single [SchemaInfo].
//
// Returns an error if there's more than one [SchemaInfo] with the same .Name.
//
// Note: All [SchemaInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *SchemasAPI) GetByName(ctx context.Context, name string) (*SchemaInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListSchemasRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]SchemaInfo{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("SchemaInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of SchemaInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Update a schema.
//
// Updates a schema for a catalog. The caller must be the owner of the schema or
// a metastore admin. If the caller is a metastore admin, only the __owner__
// field can be changed in the update. If the __name__ field must be updated,
// the caller must be a metastore admin or have the **CREATE_SCHEMA** privilege
// on the parent catalog.
func (a *SchemasAPI) Update(ctx context.Context, request UpdateSchema) (*SchemaInfo, error) {
	return a.impl.Update(ctx, request)
}

func NewShares(client *client.DatabricksClient) *SharesAPI {
	return &SharesAPI{
		impl: &sharesImpl{
			client: client,
		},
	}
}

// Databricks Delta Sharing: Shares REST API
type SharesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(SharesService)
	impl SharesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *SharesAPI) WithImpl(impl SharesService) *SharesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Shares API implementation
func (a *SharesAPI) Impl() SharesService {
	return a.impl
}

// Create a share.
//
// Creates a new share for data objects. Data objects can be added at this time
// or after creation with **update**. The caller must be a metastore admin or
// have the **CREATE_SHARE** privilege on the metastore.
func (a *SharesAPI) Create(ctx context.Context, request CreateShare) (*ShareInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete a share.
//
// Deletes a data object share from the metastore. The caller must be an owner
// of the share.
func (a *SharesAPI) Delete(ctx context.Context, request DeleteShareRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a share.
//
// Deletes a data object share from the metastore. The caller must be an owner
// of the share.
func (a *SharesAPI) DeleteByName(ctx context.Context, name string) error {
	return a.impl.Delete(ctx, DeleteShareRequest{
		Name: name,
	})
}

// Get a share.
//
// Gets a data object share from the metastore. The caller must be a metastore
// admin or the owner of the share.
func (a *SharesAPI) Get(ctx context.Context, request GetShareRequest) (*ShareInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a share.
//
// Gets a data object share from the metastore. The caller must be a metastore
// admin or the owner of the share.
func (a *SharesAPI) GetByName(ctx context.Context, name string) (*ShareInfo, error) {
	return a.impl.Get(ctx, GetShareRequest{
		Name: name,
	})
}

// List shares.
//
// Gets an array of data object shares from the metastore. The caller must be a
// metastore admin or the owner of the share. There is no guarantee of a
// specific ordering of the elements in the array.
//
// This method is generated by Databricks SDK Code Generator.
func (a *SharesAPI) ListAll(ctx context.Context) ([]ShareInfo, error) {
	response, err := a.impl.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.Shares, nil
}

// Get permissions.
//
// Gets the permissions for a data share from the metastore. The caller must be
// a metastore admin or the owner of the share.
func (a *SharesAPI) SharePermissions(ctx context.Context, request SharePermissionsRequest) (*PermissionsList, error) {
	return a.impl.SharePermissions(ctx, request)
}

// Get permissions.
//
// Gets the permissions for a data share from the metastore. The caller must be
// a metastore admin or the owner of the share.
func (a *SharesAPI) SharePermissionsByName(ctx context.Context, name string) (*PermissionsList, error) {
	return a.impl.SharePermissions(ctx, SharePermissionsRequest{
		Name: name,
	})
}

// Update a share.
//
// Updates the share with the changes and data objects in the request. The
// caller must be the owner of the share or a metastore admin.
//
// When the caller is a metastore admin, only the __owner__ field can be
// updated.
//
// In the case that the share name is changed, **updateShare** requires that the
// caller is both the share owner and a metastore admin.
//
// For each table that is added through this method, the share owner must also
// have **SELECT** privilege on the table. This privilege must be maintained
// indefinitely for recipients to be able to access the table. Typically, you
// should use a group as the share owner.
//
// Table removals through **update** do not require additional privileges.
func (a *SharesAPI) Update(ctx context.Context, request UpdateShare) (*ShareInfo, error) {
	return a.impl.Update(ctx, request)
}

// Update permissions.
//
// Updates the permissions for a data share in the metastore. The caller must be
// a metastore admin or an owner of the share.
//
// For new recipient grants, the user must also be the owner of the recipients.
// recipient revocations do not require additional privileges.
func (a *SharesAPI) UpdatePermissions(ctx context.Context, request UpdateSharePermissions) error {
	return a.impl.UpdatePermissions(ctx, request)
}

func NewStorageCredentials(client *client.DatabricksClient) *StorageCredentialsAPI {
	return &StorageCredentialsAPI{
		impl: &storageCredentialsImpl{
			client: client,
		},
	}
}

// A storage credential represents an authentication and authorization mechanism
// for accessing data stored on your cloud tenant. Each storage credential is
// subject to Unity Catalog access-control policies that control which users and
// groups can access the credential. If a user does not have access to a storage
// credential in Unity Catalog, the request fails and Unity Catalog does not
// attempt to authenticate to your cloud tenant on the user’s behalf.
//
// Databricks recommends using external locations rather than using storage
// credentials directly.
//
// To create storage credentials, you must be a Databricks account admin. The
// account admin who creates the storage credential can delegate ownership to
// another user or group to manage permissions on it.
type StorageCredentialsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(StorageCredentialsService)
	impl StorageCredentialsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *StorageCredentialsAPI) WithImpl(impl StorageCredentialsService) *StorageCredentialsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level StorageCredentials API implementation
func (a *StorageCredentialsAPI) Impl() StorageCredentialsService {
	return a.impl
}

// Create a storage credential.
//
// Creates a new storage credential. The request object is specific to the
// cloud:
//
// * **AwsIamRole** for AWS credentials * **AzureServicePrincipal** for Azure
// credentials * **GcpServiceAcountKey** for GCP credentials.
//
// The caller must be a metastore admin and have the
// **CREATE_STORAGE_CREDENTIAL** privilege on the metastore.
func (a *StorageCredentialsAPI) Create(ctx context.Context, request CreateStorageCredential) (*StorageCredentialInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete a credential.
//
// Deletes a storage credential from the metastore. The caller must be an owner
// of the storage credential.
func (a *StorageCredentialsAPI) Delete(ctx context.Context, request DeleteStorageCredentialRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a credential.
//
// Deletes a storage credential from the metastore. The caller must be an owner
// of the storage credential.
func (a *StorageCredentialsAPI) DeleteByName(ctx context.Context, name string) error {
	return a.impl.Delete(ctx, DeleteStorageCredentialRequest{
		Name: name,
	})
}

// Get a credential.
//
// Gets a storage credential from the metastore. The caller must be a metastore
// admin, the owner of the storage credential, or have some permission on the
// storage credential.
func (a *StorageCredentialsAPI) Get(ctx context.Context, request GetStorageCredentialRequest) (*StorageCredentialInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a credential.
//
// Gets a storage credential from the metastore. The caller must be a metastore
// admin, the owner of the storage credential, or have some permission on the
// storage credential.
func (a *StorageCredentialsAPI) GetByName(ctx context.Context, name string) (*StorageCredentialInfo, error) {
	return a.impl.Get(ctx, GetStorageCredentialRequest{
		Name: name,
	})
}

// List credentials.
//
// Gets an array of storage credentials (as __StorageCredentialInfo__ objects).
// The array is limited to only those storage credentials the caller has
// permission to access. If the caller is a metastore admin, all storage
// credentials will be retrieved. There is no guarantee of a specific ordering
// of the elements in the array.
func (a *StorageCredentialsAPI) List(ctx context.Context) ([]StorageCredentialInfo, error) {
	return a.impl.List(ctx)
}

// StorageCredentialInfoNameToIdMap calls [StorageCredentialsAPI.List] and creates a map of results with [StorageCredentialInfo].Name as key and [StorageCredentialInfo].Id as value.
//
// Returns an error if there's more than one [StorageCredentialInfo] with the same .Name.
//
// Note: All [StorageCredentialInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *StorageCredentialsAPI) StorageCredentialInfoNameToIdMap(ctx context.Context) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.List(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// Update a credential.
//
// Updates a storage credential on the metastore. The caller must be the owner
// of the storage credential or a metastore admin. If the caller is a metastore
// admin, only the __owner__ credential can be changed.
func (a *StorageCredentialsAPI) Update(ctx context.Context, request UpdateStorageCredential) (*StorageCredentialInfo, error) {
	return a.impl.Update(ctx, request)
}

// Validate a storage credential.
//
// Validates a storage credential. At least one of __external_location_name__
// and __url__ need to be provided. If only one of them is provided, it will be
// used for validation. And if both are provided, the __url__ will be used for
// validation, and __external_location_name__ will be ignored when checking
// overlapping urls.
//
// Either the __storage_credential_name__ or the cloud-specific credential must
// be provided.
//
// The caller must be a metastore admin or the storage credential owner or have
// the **CREATE_EXTERNAL_LOCATION** privilege on the metastore and the storage
// credential.
func (a *StorageCredentialsAPI) Validate(ctx context.Context, request ValidateStorageCredential) (*ValidateStorageCredentialResponse, error) {
	return a.impl.Validate(ctx, request)
}

func NewTableConstraints(client *client.DatabricksClient) *TableConstraintsAPI {
	return &TableConstraintsAPI{
		impl: &tableConstraintsImpl{
			client: client,
		},
	}
}

// Primary key and foreign key constraints encode relationships between fields
// in tables.
//
// Primary and foreign keys are informational only and are not enforced. Foreign
// keys must reference a primary key in another table. This primary key is the
// parent constraint of the foreign key and the table this primary key is on is
// the parent table of the foreign key. Similarly, the foreign key is the child
// constraint of its referenced primary key; the table of the foreign key is the
// child table of the primary key.
//
// You can declare primary keys and foreign keys as part of the table
// specification during table creation. You can also add or drop constraints on
// existing tables.
type TableConstraintsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(TableConstraintsService)
	impl TableConstraintsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *TableConstraintsAPI) WithImpl(impl TableConstraintsService) *TableConstraintsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level TableConstraints API implementation
func (a *TableConstraintsAPI) Impl() TableConstraintsService {
	return a.impl
}

// Create a table constraint.
//
// Creates a new table constraint.
//
// For the table constraint creation to succeed, the user must satisfy both of
// these conditions: - the user must have the **USE_CATALOG** privilege on the
// table's parent catalog, the **USE_SCHEMA** privilege on the table's parent
// schema, and be the owner of the table. - if the new constraint is a
// __ForeignKeyConstraint__, the user must have the **USE_CATALOG** privilege on
// the referenced parent table's catalog, the **USE_SCHEMA** privilege on the
// referenced parent table's schema, and be the owner of the referenced parent
// table.
func (a *TableConstraintsAPI) Create(ctx context.Context, request CreateTableConstraint) (*TableConstraint, error) {
	return a.impl.Create(ctx, request)
}

// Delete a table constraint.
//
// Deletes a table constraint.
//
// For the table constraint deletion to succeed, the user must satisfy both of
// these conditions: - the user must have the **USE_CATALOG** privilege on the
// table's parent catalog, the **USE_SCHEMA** privilege on the table's parent
// schema, and be the owner of the table. - if __cascade__ argument is **true**,
// the user must have the following permissions on all of the child tables: the
// **USE_CATALOG** privilege on the table's catalog, the **USE_SCHEMA**
// privilege on the table's schema, and be the owner of the table.
func (a *TableConstraintsAPI) Delete(ctx context.Context, request DeleteTableConstraintRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a table constraint.
//
// Deletes a table constraint.
//
// For the table constraint deletion to succeed, the user must satisfy both of
// these conditions: - the user must have the **USE_CATALOG** privilege on the
// table's parent catalog, the **USE_SCHEMA** privilege on the table's parent
// schema, and be the owner of the table. - if __cascade__ argument is **true**,
// the user must have the following permissions on all of the child tables: the
// **USE_CATALOG** privilege on the table's catalog, the **USE_SCHEMA**
// privilege on the table's schema, and be the owner of the table.
func (a *TableConstraintsAPI) DeleteByFullName(ctx context.Context, fullName string) error {
	return a.impl.Delete(ctx, DeleteTableConstraintRequest{
		FullName: fullName,
	})
}

func NewTables(client *client.DatabricksClient) *TablesAPI {
	return &TablesAPI{
		impl: &tablesImpl{
			client: client,
		},
	}
}

// A table resides in the third layer of Unity Catalog’s three-level
// namespace. It contains rows of data. To create a table, users must have
// CREATE_TABLE and USE_SCHEMA permissions on the schema, and they must have the
// USE_CATALOG permission on its parent catalog. To query a table, users must
// have the SELECT permission on the table, and they must have the USE_CATALOG
// permission on its parent catalog and the USE_SCHEMA permission on its parent
// schema.
//
// A table can be managed or external. From an API perspective, a __VIEW__ is a
// particular kind of table (rather than a managed or external table).
type TablesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(TablesService)
	impl TablesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *TablesAPI) WithImpl(impl TablesService) *TablesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Tables API implementation
func (a *TablesAPI) Impl() TablesService {
	return a.impl
}

// Delete a table.
//
// Deletes a table from the specified parent catalog and schema. The caller must
// be the owner of the parent catalog, have the **USE_CATALOG** privilege on the
// parent catalog and be the owner of the parent schema, or be the owner of the
// table and have the **USE_CATALOG** privilege on the parent catalog and the
// **USE_SCHEMA** privilege on the parent schema.
func (a *TablesAPI) Delete(ctx context.Context, request DeleteTableRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a table.
//
// Deletes a table from the specified parent catalog and schema. The caller must
// be the owner of the parent catalog, have the **USE_CATALOG** privilege on the
// parent catalog and be the owner of the parent schema, or be the owner of the
// table and have the **USE_CATALOG** privilege on the parent catalog and the
// **USE_SCHEMA** privilege on the parent schema.
func (a *TablesAPI) DeleteByFullName(ctx context.Context, fullName string) error {
	return a.impl.Delete(ctx, DeleteTableRequest{
		FullName: fullName,
	})
}

// Get a table.
//
// Gets a table from the metastore for a specific catalog and schema. The caller
// must be a metastore admin, be the owner of the table and have the
// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
// privilege on the parent schema, or be the owner of the table and have the
// **SELECT** privilege on it as well.
func (a *TablesAPI) Get(ctx context.Context, request GetTableRequest) (*TableInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a table.
//
// Gets a table from the metastore for a specific catalog and schema. The caller
// must be a metastore admin, be the owner of the table and have the
// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
// privilege on the parent schema, or be the owner of the table and have the
// **SELECT** privilege on it as well.
func (a *TablesAPI) GetByFullName(ctx context.Context, fullName string) (*TableInfo, error) {
	return a.impl.Get(ctx, GetTableRequest{
		FullName: fullName,
	})
}

// List tables.
//
// Gets an array of all tables for the current metastore under the parent
// catalog and schema. The caller must be a metastore admin or an owner of (or
// have the **SELECT** privilege on) the table. For the latter case, the caller
// must also be the owner or have the **USE_CATALOG** privilege on the parent
// catalog and the **USE_SCHEMA** privilege on the parent schema. There is no
// guarantee of a specific ordering of the elements in the array.
//
// This method is generated by Databricks SDK Code Generator.
func (a *TablesAPI) ListAll(ctx context.Context, request ListTablesRequest) ([]TableInfo, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Tables, nil
}

// TableInfoNameToTableIdMap calls [TablesAPI.ListAll] and creates a map of results with [TableInfo].Name as key and [TableInfo].TableId as value.
//
// Returns an error if there's more than one [TableInfo] with the same .Name.
//
// Note: All [TableInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *TablesAPI) TableInfoNameToTableIdMap(ctx context.Context, request ListTablesRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.TableId
	}
	return mapping, nil
}

// GetByName calls [TablesAPI.TableInfoNameToTableIdMap] and returns a single [TableInfo].
//
// Returns an error if there's more than one [TableInfo] with the same .Name.
//
// Note: All [TableInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *TablesAPI) GetByName(ctx context.Context, name string) (*TableInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListTablesRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]TableInfo{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("TableInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of TableInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// List table summaries.
//
// Gets an array of summaries for tables for a schema and catalog within the
// metastore. The table summaries returned are either:
//
// * summaries for all tables (within the current metastore and parent catalog
// and schema), when the user is a metastore admin, or: * summaries for all
// tables and schemas (within the current metastore and parent catalog) for
// which the user has ownership or the **SELECT** privilege on the table and
// ownership or **USE_SCHEMA** privilege on the schema, provided that the user
// also has ownership or the **USE_CATALOG** privilege on the parent catalog.
//
// There is no guarantee of a specific ordering of the elements in the array.
func (a *TablesAPI) ListSummaries(ctx context.Context, request ListSummariesRequest) (*ListTableSummariesResponse, error) {
	return a.impl.ListSummaries(ctx, request)
}
