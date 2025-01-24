package postgresqlcnpgio

import (
	"time"
)

// RoleConfiguration is the representation, in Kubernetes, of a PostgreSQL role with the additional field Ensure specifying whether to ensure the presence or absence of the role in the database.
//
// The defaults of the CREATE ROLE command are applied
// Reference: https://www.postgresql.org/docs/current/sql-createrole.html
type ClusterSpecManagedRoles struct {
	// Name of the role.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether a role bypasses every row-level security (RLS) policy.
	//
	// Default is `false`.
	// Default: false`.
	//
	Bypassrls *bool `field:"optional" json:"bypassrls" yaml:"bypassrls"`
	// Description of the role.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// If the role can log in, this specifies how many concurrent connections the role can make.
	//
	// `-1` (the default) means no limit.
	ConnectionLimit *float64 `field:"optional" json:"connectionLimit" yaml:"connectionLimit"`
	// When set to `true`, the role being defined will be allowed to create new databases.
	//
	// Specifying `false` (default) will deny a role the
	// ability to create databases.
	Createdb *bool `field:"optional" json:"createdb" yaml:"createdb"`
	// Whether the role will be permitted to create, alter, drop, comment on, change the security label for, and grant or revoke membership in other roles.
	//
	// Default is `false`.
	// Default: false`.
	//
	Createrole *bool `field:"optional" json:"createrole" yaml:"createrole"`
	// DisablePassword indicates that a role's password should be set to NULL in Postgres.
	DisablePassword *bool `field:"optional" json:"disablePassword" yaml:"disablePassword"`
	// Ensure the role is `present` or `absent` - defaults to "present".
	Ensure ClusterSpecManagedRolesEnsure `field:"optional" json:"ensure" yaml:"ensure"`
	// Whether a role "inherits" the privileges of roles it is a member of.
	//
	// Defaults is `true`.
	// Default: true`.
	//
	Inherit *bool `field:"optional" json:"inherit" yaml:"inherit"`
	// List of one or more existing roles to which this role will be immediately added as a new member.
	//
	// Default empty.
	InRoles *[]*string `field:"optional" json:"inRoles" yaml:"inRoles"`
	// Whether the role is allowed to log in.
	//
	// A role having the `login`
	// attribute can be thought of as a user. Roles without this attribute
	// are useful for managing database privileges, but are not users in
	// the usual sense of the word. Default is `false`.
	// Default: false`.
	//
	Login *bool `field:"optional" json:"login" yaml:"login"`
	// Secret containing the password of the role (if present) If null, the password will be ignored unless DisablePassword is set.
	PasswordSecret *ClusterSpecManagedRolesPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	// Whether a role is a replication role.
	//
	// A role must have this
	// attribute (or be a superuser) in order to be able to connect to the
	// server in replication mode (physical or logical replication) and in
	// order to be able to create or drop replication slots. A role having
	// the `replication` attribute is a very highly privileged role, and
	// should only be used on roles actually used for replication. Default
	// is `false`.
	// Default: false`.
	//
	Replication *bool `field:"optional" json:"replication" yaml:"replication"`
	// Whether the role is a `superuser` who can override all access restrictions within the database - superuser status is dangerous and should be used only when really needed.
	//
	// You must yourself be a
	// superuser to create a new superuser. Defaults is `false`.
	// Default: false`.
	//
	Superuser *bool `field:"optional" json:"superuser" yaml:"superuser"`
	// Date and time after which the role's password is no longer valid.
	//
	// When omitted, the password will never expire (default).
	ValidUntil *time.Time `field:"optional" json:"validUntil" yaml:"validUntil"`
}

