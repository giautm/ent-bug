package confirmstatus

type Status string

const (
	StatusUnknown Status = "unknown"
	StatusConfirm Status = "confirm"
	StatusDecline Status = "decline"
)

func (Status) Values() (statuses []string) {
	for _, r := range []Status{StatusUnknown, StatusConfirm, StatusDecline} {
		statuses = append(statuses, string(r))
	}
	return
}
