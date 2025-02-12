package message


type ReplicationTopicMessage struct {
	Service string `json:"Service"`
	Action string `json:"Action"`
}
