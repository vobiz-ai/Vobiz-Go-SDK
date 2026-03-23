// Package compliance provides Vobiz Compliance API services:
// DocumentService, DocumentTypeService, RequirementService, ApplicationService.
package compliance

import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"

// DocumentService exposes compliance document API operations.
type DocumentService struct {
	r         interfaces.Requester
	authID    string
	authToken string
	baseURL   string
}

// NewDocument creates a new DocumentService.
func NewDocument(r interfaces.Requester) *DocumentService {
	return &DocumentService{r: r, authID: r.AuthID(), baseURL: "https://api.vobiz.ai"}
}

// DocumentTypeService exposes compliance document type API operations.
type DocumentTypeService struct {
	r interfaces.Requester
}

// NewDocumentType creates a new DocumentTypeService.
func NewDocumentType(r interfaces.Requester) *DocumentTypeService {
	return &DocumentTypeService{r: r}
}

// RequirementService exposes compliance requirement API operations.
type RequirementService struct {
	r interfaces.Requester
}

// NewRequirement creates a new RequirementService.
func NewRequirement(r interfaces.Requester) *RequirementService {
	return &RequirementService{r: r}
}

// ApplicationService exposes compliance application API operations.
type ApplicationService struct {
	r interfaces.Requester
}

// NewApplication creates a new ApplicationService.
func NewApplication(r interfaces.Requester) *ApplicationService {
	return &ApplicationService{r: r}
}
