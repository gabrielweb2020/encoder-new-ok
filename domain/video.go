package domain

type Video struct {
	ID string
	ResourceID string
	FilePath string
	CreatedAt time.Time
}

func main() {
	log.Info("Something noteworthy happened!")
}
