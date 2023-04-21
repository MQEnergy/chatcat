package model

type FeedBack struct {
	Title       string   `json:"title,omitempty"`
	Body        string   `json:"body,omitempty"`
	Labels      []string `json:"labels,omitempty"`
	Assignee    string   `json:"assignee,omitempty"`
	State       string   `json:"state,omitempty"`
	StateReason string   `json:"state_reason,omitempty"`
	Milestone   int      `json:"milestone,omitempty"`
	Assignees   []string `json:"assignees,omitempty"`
}

type FeedbackReq struct {
	IssueType int    `json:"issue_type,omitempty"`
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
	Version   string `json:"version,omitempty"`
}
