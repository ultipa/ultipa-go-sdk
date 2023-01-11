package utils

//LeaderNotYetElectedError leader not yet elected error for cluster
type LeaderNotYetElectedError struct {
	Message string
}

func (err *LeaderNotYetElectedError) Error() string {
	if err.Message != "" {
		return err.Message
	}
	return "raft leader tot yet elected"
}

func NewLeaderNotYetElectedError(msg string) *LeaderNotYetElectedError {
	return &LeaderNotYetElectedError{
		Message: msg,
	}
}
