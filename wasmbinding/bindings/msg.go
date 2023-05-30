package bindings

type JackalMsg struct {
	/// Contracts can create files
	/// will they be namespaced under the contract's address?
	/// A contract may create any number of independent files.
	PostFiles *PostFiles `json:"post_files,omitempty"`
	MakeRoot  *MakeRoot  `json:"make_root,omitempty"`
}

// / creator == broadcaster of the msg
type PostFiles struct {
	Creator        string `json:"creator"`
	Account        string `json:"account"`
	HashParent     string `json:"hashparent"`
	HashChild      string `json:"hashchild"`
	Contents       string `json:"contents"`
	Viewers        string `json:"viewers"`
	Editors        string `json:"editors"`
	TrackingNumber string `json:"trackingnumber"`
}

type MakeRoot struct {
	Creator        string `json:"creator"`
	Editors        string `json:"editors"`
	Viewers        string `json:"viewers"`
	TrackingNumber string `json:"trackingnumber"`
}