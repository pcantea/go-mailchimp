package mailchimp

// MemberResponse - see https://api.mailchimp.com/schema/3.0/Lists/Members/Instance.json?_ga=1.216961300.323879299.1464708316
type MemberResponse struct {
	ID              string                 `json:"id"` // The MD5 hash of the list member's email address.
	EmailAddress    string                 `json:"email_address"`
	UniqueEmailID   string                 `json:"unique_email_id"` // An identifier for the address across all of MailChimp.
	EmailType       string                 `json:"email_type"`      // Type of email this member asked to get ('html' or 'text').
	Status          string                 `json:"status"`
	VIP             bool                   `json:"vip"`
	IPSignup        string                 `json:"ip_signup"`        // IP address the subscriber signed up from.
	TimestampSignup string                 `json:"timestamp_signup"` // Date and time the subscriber signed up for the list.
	IPOpt           string                 `json:"ip_opt"`           // IP address the subscriber confirmed their opt-in status.
	TimestampOpt    string                 `json:"timestamp_opt"`    // Date and time the subscribe confirmed their opt-in status.
	MemberRating    uint                   `json:"member_rating"`    // Star rating for this member between 1 and 5.
	LastChanged     string                 `json:"last_changed"`     // Date and time the member's info was last changed.
	ListID          string                 `json:"list_id"`          // The id for the list.
	MergeFields     map[string]interface{} `json:"merge_fields"`     // merge fields
}

type BatchResponse struct {
	ID                 string `json:"id"`
	Status             string `json:"status"`
	TotalOperations    int    `json:"total_operations"`
	FinishedOperations int    `json:"finished_operations"`
	ErroredOperations  int    `json:"errored_operations"`
	SubmittedAt        string `json:"submitted_at"`
	CompletedAt        string `json:"completed_at"`
	ResponseURL        string `json:"response_body_url"`
}
