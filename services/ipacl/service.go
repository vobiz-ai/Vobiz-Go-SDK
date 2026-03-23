package ipacl
import "github.com/vobiz/all-vobiz-sdk/go-sdk/interfaces"
// Service exposes IP ACL CRUD operations.
type Service struct{ r interfaces.Requester }

// New creates a new ipacl service.
func New(r interfaces.Requester) *Service { return &Service{r: r} }