package change

// Change types
const (
	Noop = iota
	Mirror_update
	Mirror_create
	Mirror_recreate
	Repo_create
	Gpg_add
)