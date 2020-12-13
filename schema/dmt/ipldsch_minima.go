package schemadmt

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	"fmt"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/schema"
)

const (
	midvalue  = schema.Maybe(4)
	allowNull = schema.Maybe(5)
)

type maState uint8

const (
	maState_initial maState = iota
	maState_midKey
	maState_expectValue
	maState_midValue
	maState_finished
)

type laState uint8

const (
	laState_initial laState = iota
	laState_midValue
	laState_finished
)

type _ErrorThunkAssembler struct {
	e error
}

func (ea _ErrorThunkAssembler) BeginMap(_ int) (ipld.MapAssembler, error)   { return nil, ea.e }
func (ea _ErrorThunkAssembler) BeginList(_ int) (ipld.ListAssembler, error) { return nil, ea.e }
func (ea _ErrorThunkAssembler) AssignNull() error                           { return ea.e }
func (ea _ErrorThunkAssembler) AssignBool(bool) error                       { return ea.e }
func (ea _ErrorThunkAssembler) AssignInt(int) error                         { return ea.e }
func (ea _ErrorThunkAssembler) AssignFloat(float64) error                   { return ea.e }
func (ea _ErrorThunkAssembler) AssignString(string) error                   { return ea.e }
func (ea _ErrorThunkAssembler) AssignBytes([]byte) error                    { return ea.e }
func (ea _ErrorThunkAssembler) AssignLink(ipld.Link) error                  { return ea.e }
func (ea _ErrorThunkAssembler) AssignNode(ipld.Node) error                  { return ea.e }
func (ea _ErrorThunkAssembler) Prototype() ipld.NodePrototype {
	panic(fmt.Errorf("cannot get prototype from error-carrying assembler: already derailed with error: %w", ea.e))
}
