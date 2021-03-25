package access

type PermissionLoader interface {
	PermissioLoad(id string) ([]string, error)
	Close() error
}
