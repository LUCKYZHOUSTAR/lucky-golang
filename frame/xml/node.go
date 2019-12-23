// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

const (
	NT_ROOT = iota
	NT_DIRECTIVE
	NT_PROCINST
	NT_COMMENT
	NT_TEXT
	NT_ELEMENT
)

// IndentPrefix holds the value for a single identation level, if one
// chooses to want indentation in the node.String() and node.Bytes() output.
// This would normally be set to a single tab, or a number of spaces.
var IndentPrefix = ""

type Attr struct {
	Name  xml.Name // Attribute namespace and name.
	Value string   // Attribute value.
}

type Node struct {
	Type       byte     // Node type.
	Name       xml.Name // Node namespace and name.
	Children   []*Node  // Child nodes.
	Attributes []*Attr  // Node attributes.
	Parent     *Node    // Parent node.
	Value      string   // Node value.
	Target     string   // procinst field.
}

func NewNode(tid byte) *Node {
	n := new(Node)
	n.Type = tid
	n.Children = make([]*Node, 0, 10)
	n.Attributes = make([]*Attr, 0, 10)
	return n
}

// This wraps the standard xml.Unmarshal function and supplies this particular
// node as the content to be unmarshalled.
func (n *Node) Unmarshal(obj interface{}) error {
	return xml.NewDecoder(bytes.NewBuffer(n.bytes())).Decode(obj)
}

func (n *Node) GetValue() string {
	res := ""
	for _, node := range n.Children {
		if node.Type == NT_TEXT {
			res += strings.TrimSpace(node.Value)
		}
	}

	return res
}

// Get node value as string
func (n *Node) S(namespace, name string) string {
	foundNode := rec_SelectNode(n, namespace, name)
	if foundNode == nil {
		return ""
	} else {
		return foundNode.GetValue()
	}
}

// Get node value as int
func (n *Node) I(namespace, name string) int {
	value := n.S(namespace, name)
	if value != "" {
		n, _ := strconv.ParseInt(value, 10, 0)
		return int(n)
	}
	return 0
}

// Get node value as int8
func (n *Node) I8(namespace, name string) int8 {
	value := n.S(namespace, name)
	if value != "" {
		n, _ := strconv.ParseInt(value, 10, 8)
		return int8(n)
	}
	return 0
}

// Get node value as int16
func (n *Node) I16(namespace, name string) int16 {
	value := n.S(namespace, name)
	if value != "" {
		n, _ := strconv.ParseInt(value, 10, 16)
		return int16(n)
	}
	return 0
}

// Get node value as int32
func (n *Node) I32(namespace, name string) int32 {
	value := n.S(namespace, name)
	if value != "" {
		n, _ := strconv.ParseInt(value, 10, 32)
		return int32(n)
	}
	return 0
}

// Get node value as int64
func (n *Node) I64(namespace, name string) int64 {
	value := n.S(namespace, name)
	if value != "" {
		n, _ := strconv.ParseInt(value, 10, 64)
		return n
	}
	return 0
}

// Get node value as uint
func (n *Node) U(namespace, name string) uint {
	value := n.S(namespace, name)
	if value != "" {
		n, _ := strconv.ParseUint(value, 10, 0)
		return uint(n)
	}
	return 0
}

// Get node value as uint8
func (n *Node) U8(namespace, name string) uint8 {
	value := n.S(namespace, name)
	if value != "" {
		n, _ := strconv.ParseUint(value, 10, 8)
		return uint8(n)
	}
	return 0
}

// Get node value as uint16
func (n *Node) U16(namespace, name string) uint16 {
	value := n.S(namespace, name)
	if value != "" {
		n, _ := strconv.ParseUint(value, 10, 16)
		return uint16(n)
	}
	return 0
}

// Get node value as uint32
func (n *Node) U32(namespace, name string) uint32 {
	value := n.S(namespace, name)
	if value != "" {
		n, _ := strconv.ParseUint(value, 10, 32)
		return uint32(n)
	}
	return 0
}

// Get node value as uint64
func (n *Node) U64(namespace, name string) uint64 {
	value := n.S(namespace, name)
	if value != "" {
		n, _ := strconv.ParseUint(value, 10, 64)
		return n
	}
	return 0
}

// Get node value as float32
func (n *Node) F32(namespace, name string) float32 {
	value := n.S(namespace, name)
	if value != "" {
		n, _ := strconv.ParseFloat(value, 32)
		return float32(n)
	}
	return 0
}

// Get node value as float64
func (n *Node) F64(namespace, name string) float64 {
	value := n.S(namespace, name)
	if value != "" {
		n, _ := strconv.ParseFloat(value, 64)
		return n
	}
	return 0
}

// Get node value as bool
func (n *Node) B(namespace, name string) bool {
	value := n.S(namespace, name)
	if value != "" {
		n, _ := strconv.ParseBool(value)
		return n
	}
	return false
}

// Get attribute value as string
func (n *Node) As(namespace, name string) string {
	for _, v := range n.Attributes {
		if (namespace == "*" || namespace == v.Name.Space) && name == v.Name.Local {
			return v.Value
		}
	}
	return ""
}

// Get attribute value as int
func (n *Node) Ai(namespace, name string) int {
	s := n.As(namespace, name)
	if s != "" {
		n, _ := strconv.ParseInt(s, 10, 0)
		return int(n)
	}
	return 0
}

// Get attribute value as int8
func (n *Node) Ai8(namespace, name string) int8 {
	s := n.As(namespace, name)
	if s != "" {
		n, _ := strconv.ParseInt(s, 10, 8)
		return int8(n)
	}
	return 0
}

// Get attribute value as int16
func (n *Node) Ai16(namespace, name string) int16 {
	s := n.As(namespace, name)
	if s != "" {
		n, _ := strconv.ParseInt(s, 10, 16)
		return int16(n)
	}
	return 0
}

// Get attribute value as int32
func (n *Node) Ai32(namespace, name string) int32 {
	s := n.As(namespace, name)
	if s != "" {
		n, _ := strconv.ParseInt(s, 10, 32)
		return int32(n)
	}
	return 0
}

// Get attribute value as int64
func (n *Node) Ai64(namespace, name string) int64 {
	s := n.As(namespace, name)
	if s != "" {
		n, _ := strconv.ParseInt(s, 10, 64)
		return n
	}
	return 0
}

// Get attribute value as uint
func (n *Node) Au(namespace, name string) uint {
	s := n.As(namespace, name)
	if s != "" {
		n, _ := strconv.ParseUint(s, 10, 0)
		return uint(n)
	}
	return 0
}

// Get attribute value as uint8
func (n *Node) Au8(namespace, name string) uint8 {
	s := n.As(namespace, name)
	if s != "" {
		n, _ := strconv.ParseUint(s, 10, 8)
		return uint8(n)
	}
	return 0
}

// Get attribute value as uint16
func (n *Node) Au16(namespace, name string) uint16 {
	s := n.As(namespace, name)
	if s != "" {
		n, _ := strconv.ParseUint(s, 10, 16)
		return uint16(n)
	}
	return 0
}

// Get attribute value as uint32
func (n *Node) Au32(namespace, name string) uint32 {
	s := n.As(namespace, name)
	if s != "" {
		n, _ := strconv.ParseUint(s, 10, 32)
		return uint32(n)
	}
	return 0
}

// Get attribute value as uint64
func (n *Node) Au64(namespace, name string) uint64 {
	s := n.As(namespace, name)
	if s != "" {
		n, _ := strconv.ParseUint(s, 10, 64)
		return n
	}
	return 0
}

// Get attribute value as float32
func (n *Node) Af32(namespace, name string) float32 {
	s := n.As(namespace, name)
	if s != "" {
		n, _ := strconv.ParseFloat(s, 32)
		return float32(n)
	}
	return 0
}

// Get attribute value as float64
func (n *Node) Af64(namespace, name string) float64 {
	s := n.As(namespace, name)
	if s != "" {
		n, _ := strconv.ParseFloat(s, 64)
		return n
	}
	return 0
}

// Get attribute value as bool
func (n *Node) Ab(namespace, name string) bool {
	s := n.As(namespace, name)
	if s != "" {
		n, _ := strconv.ParseBool(s)
		return n
	}
	return false
}

// Returns true if this node has the specified attribute. False otherwise.
func (n *Node) HasAttr(namespace, name string) bool {
	for _, v := range n.Attributes {
		if namespace != "*" && namespace != v.Name.Space {
			continue
		}

		if name == "*" || name == v.Name.Local {
			return true
		}
	}

	return false
}

// Select single node by name
func (n *Node) SelectNode(namespace, name string) *Node {
	return rec_SelectNode(n, namespace, name)
}

func rec_SelectNode(cn *Node, namespace, name string) *Node {
	if (namespace == "*" || cn.Name.Space == namespace) && (name == "*" || cn.Name.Local == name) {
		return cn
	}

	var tn *Node
	for _, v := range cn.Children {
		if tn = rec_SelectNode(v, namespace, name); tn != nil {
			return tn
		}
	}

	return nil
}

// Select multiple nodes by name
func (n *Node) SelectNodes(namespace, name string) []*Node {
	list := make([]*Node, 0, 16)
	rec_SelectNodes(n, namespace, name, &list, false)
	return list
}

// Select multiple nodes by name
func (n *Node) SelectNodesRecursive(namespace, name string) []*Node {
	list := make([]*Node, 0, 16)
	rec_SelectNodes(n, namespace, name, &list, true)
	return list
}

func rec_SelectNodes(cn *Node, namespace, name string, list *[]*Node, recurse bool) {
	if (namespace == "*" || cn.Name.Space == namespace) && (name == "*" || cn.Name.Local == name) {
		*list = append(*list, cn)
		if !recurse {
			return
		}
	}

	for _, v := range cn.Children {
		rec_SelectNodes(v, namespace, name, list, recurse)
	}
}

func (n *Node) RemoveNameSpace() {
	n.Name.Space = ""
	//	this.RemoveAttr("xmlns") //This is questionable

	for _, v := range n.Children {
		v.RemoveNameSpace()
	}
}

func (n *Node) RemoveAttr(name string) {
	for i, v := range n.Attributes {
		if name == v.Name.Local {
			//Delete it
			n.Attributes = append(n.Attributes[:i], n.Attributes[i+1:]...)
		}
	}
}

func (n *Node) SetAttr(name, value string) {
	for _, v := range n.Attributes {
		if name == v.Name.Local {
			v.Value = value
			return
		}
	}
	//Add
	attr := new(Attr)
	attr.Name.Local = name
	attr.Name.Space = ""
	attr.Value = value
	n.Attributes = append(n.Attributes, attr)
	return
}

// Convert node to appropriate []byte representation based on it's @Type.
// Note that NT_ROOT is a special-case empty node used as the root for a
// Document. This one has no representation by itself. It merely forwards the
// String() call to it's child nodes.
func (n *Node) Bytes() []byte { return n.bytes() }

func (n *Node) bytes() (b []byte) {
	switch n.Type {
	case NT_PROCINST:
		b = n.printProcInst()
	case NT_COMMENT:
		b = n.printComment()
	case NT_DIRECTIVE:
		b = n.printDirective()
	case NT_ELEMENT:
		b = n.printElement()
	case NT_TEXT:
		b = n.printText()
	case NT_ROOT:
		b = n.printRoot()
	}
	return
}

// Convert node to appropriate string representation based on it's @Type.
// Note that NT_ROOT is a special-case empty node used as the root for a
// Document. This one has no representation by itself. It merely forwards the
// String() call to it's child nodes.
func (n *Node) String() (s string) {
	return string(n.bytes())
}

func (n *Node) printRoot() []byte {
	var b bytes.Buffer
	for _, v := range n.Children {
		b.Write(v.bytes())
	}
	return b.Bytes()
}

func (n *Node) printProcInst() []byte {
	return []byte("<?" + n.Target + " " + n.Value + "?>")
}

func (n *Node) printComment() []byte {
	return []byte("<!-- " + n.Value + " -->")
}

func (n *Node) printDirective() []byte {
	return []byte("<!" + n.Value + "!>")
}

func (n *Node) printText() []byte {
	val := []byte(n.Value)
	if len(n.Parent.Children) > 1 {
		return val
	}
	var b bytes.Buffer
	xml.EscapeText(&b, val)
	return b.Bytes()
}

func (n *Node) printElement() []byte {
	var b bytes.Buffer

	if len(n.Name.Space) > 0 {
		b.WriteRune('<')
		b.WriteString(n.Name.Space)
		b.WriteRune(':')
		b.WriteString(n.Name.Local)
	} else {
		b.WriteRune('<')
		b.WriteString(n.Name.Local)
	}

	for _, v := range n.Attributes {
		if len(v.Name.Space) > 0 {
			prefix := n.spacePrefix(v.Name.Space)
			b.WriteString(fmt.Sprintf(` %s:%s="%s"`, prefix, v.Name.Local, v.Value))
		} else {
			b.WriteString(fmt.Sprintf(` %s="%s"`, v.Name.Local, v.Value))
		}
	}

	if len(n.Children) == 0 && len(n.Value) == 0 {
		b.WriteString(" />")
		return b.Bytes()
	}

	b.WriteRune('>')

	for _, v := range n.Children {
		b.Write(v.bytes())
	}

	xml.EscapeText(&b, []byte(n.Value))
	if len(n.Name.Space) > 0 {
		b.WriteString("</")
		b.WriteString(n.Name.Space)
		b.WriteRune(':')
		b.WriteString(n.Name.Local)
		b.WriteRune('>')
	} else {
		b.WriteString("</")
		b.WriteString(n.Name.Local)
		b.WriteRune('>')
	}

	return b.Bytes()
}

// spacePrefix resolves the given space (e.g. a url) to the prefix it was
// assigned by an attribute by the current node, or one of its parents.
func (n *Node) spacePrefix(space string) string {
	for _, attr := range n.Attributes {
		if attr.Name.Space == "xmlns" && attr.Value == space {
			return attr.Name.Local
		}
	}
	if n.Parent == nil {
		return space
	}
	return n.Parent.spacePrefix(space)
}

// Add a child node
func (n *Node) AddChild(t *Node) {
	if t.Parent != nil {
		t.Parent.RemoveChild(t)
	}
	t.Parent = n
	n.Children = append(n.Children, t)
}

// Remove a child node
func (n *Node) RemoveChild(t *Node) {
	p := -1
	for i, v := range n.Children {
		if v == t {
			p = i
			break
		}
	}

	if p == -1 {
		return
	}

	copy(n.Children[p:], n.Children[p+1:])
	n.Children = n.Children[0 : len(n.Children)-1]

	t.Parent = nil
}
