package odoo_api_wrapper

// Version describes odoo_api_wrapper instance version.
type Version struct {
	ServerVersion     *String     `xmlrpc:"server_version"`
	ServerVersionInfo interface{} `xmlrpc:"server_version_info"`
	ServerSerie       *String     `xmlrpc:"server_serie"`
	ProtocolVersion   *Int        `xmlrpc:"protocol_version"`
}
