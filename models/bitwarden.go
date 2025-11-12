package models

type UnlockRequest struct {
	Password string `json:"password"`
}

type UnlockResponse struct {
	Success bool       `json:"success"`
	Data    UnlockData `json:"data"`
}

type UnlockData struct {
	NoColor bool   `json:"noColor"`
	Object  string `json:"object"`
	Title   string `json:"title"`
	Message string `json:"message"`
	Raw     string `json:"raw"`
}

type BitwardenItemResponse struct {
	Success bool          `json:"success"`
	Data    BitwardenItem `json:"data"`
}

type BitwardenItem struct {
	PasswordHistory []interface{}  `json:"passwordHistory"`
	RevisionDate    string         `json:"revisionDate"`
	CreationDate    string         `json:"creationDate"`
	Object          string         `json:"object"`
	ID              string         `json:"id"`
	OrganizationID  *string        `json:"organizationId"`
	FolderID        *string        `json:"folderId"`
	Type            int            `json:"type"`
	Reprompt        int            `json:"reprompt"`
	Name            string         `json:"name"`
	Notes           *string        `json:"notes"`
	Favorite        bool           `json:"favorite"`
	Fields          []interface{}  `json:"fields"`
	Login           BitwardenLogin `json:"login"`
	CollectionIDs   []string       `json:"collectionIds"`
	Attachments     []interface{}  `json:"attachments"`
}

type BitwardenLogin struct {
	URIs                 []BitwardenURI `json:"uris"`
	Username             string         `json:"username"`
	Password             string         `json:"password"`
	TOTP                 *string        `json:"totp"`
	PasswordRevisionDate *string        `json:"passwordRevisionDate"`
}

type BitwardenURI struct {
	Match *string `json:"match"`
	URI   string  `json:"uri"`
}

type BitwardenItemsResponse struct {
	Success bool            `json:"success"`
	Data    []BitwardenItem `json:"data"`
}

type BitwardenItemsListResponse struct {
	Success bool                  `json:"success"`
	Data    BitwardenItemsWrapper `json:"data"`
}

type BitwardenItemsWrapper struct {
	Object string          `json:"object"`
	Data   []BitwardenItem `json:"data"`
}
