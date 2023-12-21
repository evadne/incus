// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package ovsmodel

const ForwardingGroupTable = "Forwarding_Group"

// ForwardingGroup defines an object in Forwarding_Group table
type ForwardingGroup struct {
	UUID        string            `ovsdb:"_uuid"`
	ChildPort   []string          `ovsdb:"child_port"`
	ExternalIDs map[string]string `ovsdb:"external_ids"`
	Liveness    bool              `ovsdb:"liveness"`
	Name        string            `ovsdb:"name"`
	Vip         string            `ovsdb:"vip"`
	Vmac        string            `ovsdb:"vmac"`
}