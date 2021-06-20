package gandiapi

import "time"

// DomainResponse struct stores response from zones endpoint.
type DomainResponse struct {
	Autorenew bool `json:"autorenew"`
	Dates     struct {
		RegistryCreatedAt   time.Time `json:"registry_created_at"`
		UpdatedAt           time.Time `json:"updated_at"`
		AuthinfoExpiresAt   time.Time `json:"authinfo_expires_at,omitempty"`
		CreatedAt           time.Time `json:"created_at,omitempty"`
		DeletesAt           time.Time `json:"deletes_at,omitempty"`
		HoldBeginsAt        time.Time `json:"hold_begins_at,omitempty"`
		HoldEndsAt          time.Time `json:"hold_ends_at,omitempty"`
		PendingDeleteEndsAt time.Time `json:"pending_delete_ends_at,omitempty"`
		RegistryEndsAt      time.Time `json:"registry_ends_at,omitempty"`
		RenewBeginsAt       time.Time `json:"renew_begins_at,omitempty"`
		RestoreEndsAt       time.Time `json:"restore_ends_at,omitempty"`
	} `json:"dates"`
	DomainOwner string `json:"domain_owner"`
	Fqdn        string `json:"fqdn"`
	FqdnUnicode string `json:"fqdn_unicode"`
	Href        string `json:"href"`
	ID          string `json:"id"`
	Nameserver  struct {
		Current string   `json:"current"`
		Hosts   []string `json:"hosts,omitempty"`
	} `json:"nameserver"`
	OrgaOwner string   `json:"orga_owner"`
	Owner     string   `json:"owner"`
	SharingID string   `json:"sharing_id,omitempty"`
	Status    []string `json:"status"`
	Tags      []string `json:"tags,omitempty"`
	Tld       string   `json:"tld"`
}

// DomainContact represents the contact details of a domain contact or owner.
type DomainContact struct {
	// A country code. See https://api.gandi.net/docs/domains/#appendix-Country-Codes
	// for possible values.
	Country    string `json:"country"`
	Email      string `json:"email"`
	Family     string `json:"family"`
	Given      string `json:"given"`
	Streetaddr string `json:"streetaddr"`
	// One of: "individual", "company", "association", "publicbody"
	Type        string `json:"type"`
	BrandNumber string `json:"brand_number,omitempty"`
	City        string `json:"city,omitempty"`
	// More info available at https://docs.gandi.net/en/domain_names/common_operations/whois_privacy.html
	Data_obfuscated bool `json:"data_obfuscated,omitempty"`

	// TODO: add support for this field
	// Extra parameters needed for some extensions. See https://api.gandi.net/docs/domains/#appendix-Contact-Extra-Parameters
	// for possible values
	// ExtraParameters struct {} `json:"extra_parameters,omitempty"`
	Fax               string `json:"fax,omitempty"`
	JoAnnounceNumber  string `json:"jo_announce_number,omitempty"`
	JoAnnouncePage    string `json:"jo_announce_page,omitempty"`
	JoDeclarationDate string `json:"jo_declaration_date,omitempty"`
	JoPublicationDate string `json:"jo_publication_date,omitempty"`
	// One of: "en", "es", "fr", "ja", "zh-hans", "zh-hant"
	Lang           string `json:"lang,omitempty"`
	MailObfuscated bool   `json:"mail_obfuscated,omitempty"`
	Mobile         string `json:"mobile,omitempty"`
	// Legal name of the company, association, or public body if the contact type is not "individual"
	Orgname string `json:"orgname,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Siren   string `json:"siren,omitempty"`
	// See https://docs.gandi.net/en/rest_api/domain_api/contacts_api.html for more information on state codes.
	State string `json:"state,omitempty"`
	// One of: "pending", "done", "failed", "deleted", "none"
	Validation string `json:"validation,omitempty"`
	Zip        string `json:"zip,omitempty"`
}

// DomainCreateRequest is the data sent with a CreateDomain request
type DomainCreateRequest struct {
	Fqdn   string        `json:"fqdn"`
	Owner  DomainContact `json:"owner"`
	Admin  DomainContact `json:"admin,omitempty"`
	Bill   DomainContact `json:"bill,omitempty"`
	Claims string        `json:"claims,omitempty"`
	// One of: "EUR", "USD", "GBP", "TWD", "CNY"
	Currency string `json:"currency,omitempty"`
	// Between 1 and 10
	Duration int `json:"duration,omitempty"`
	// Must be set to true if the domain is a premium domain
	EnforcePremium bool `json:"enforce_premium,omitempty"`
	// TODO: add support for this field
	// Extra parameters needed for some extensions. See https://api.gandi.net/docs/domains/#appendix-Contact-Extra-Parameters
	// for possible values
	// ExtraParameters struct {} `json:"extra_parameters,omitempty"`

	// ISO-639-2 language code of the domain. Required for some IDN domains.
	Lang string `json:"lang,omitempty"`
	// For glue records only. Dictionary associating a nameserver to a list of IP addresses.
	NameserverIPs map[string][]string `json:"nameserver_ips,omitempty"`
	// List of name servers. Gandi's LiveDNS nameservers are used if omitted.
	Nameservers []string `json:"nameservers,omitempty"`
	Price       int      `json:"price,omitempty"`
	ReselleeID  string   `json:"resellee_id,omitempty"`
	// Contents of a Signed Mark Data file (used for newgtld sunrises, `tld_period` must be sunrise).
	SMD  string        `json:"smd,omitempty"`
	Tech DomainContact `json:"tech,omitempty"`
	// Template applied when the process is done. It must be a template ID as you can retrieve it using the Template API:
	// https://api.gandi.net/docs/template/#get-v5-template-templates
	TemplateID string `json:"template_id,omitempty"`
	// One of: "sunrise", "landrush", "eap1", "eap2", "eap3", "eap4", "eap5", "eap6", "eap7", "eap8", "eap9", "golive"
	TLD_Period string `json:"tld_period,omitempty"`
}

// GetDomains returns a list of domains owned by the current user, as defined by the client API Key.
func (c *Client) GetDomains() (resp []DomainResponse, err error) {
	_, err = c.get("domain/domains", nil, &resp)
	return
}

// CreateDomain registers a new domain name. Note that this is *NOT* a free operation!
// Ensure your Gandi prepaid account has sufficient credit before performing this action.
func (c *Client) CreateDomain(data DomainCreateRequest) (err error) {
	_, err = c.post("domain/domains", data, nil)
	return
}
