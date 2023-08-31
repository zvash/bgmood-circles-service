package worker

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

// MessagePublisher pushes the tasks to the redis
type MessagePublisher interface {
}
