// Code generated by `ggen -ent SP -conf`; DO NOT EDIT.

package api

// Conf receives custom request config definition, e.g. custom headers, custom OData mod
func (sp *SP) Conf(config *RequestConfig) *SP {
	sp.config = config
	return sp
}
