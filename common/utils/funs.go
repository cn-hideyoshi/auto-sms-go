package utils

import (
	"blog.hideyoshi.top/common/config"
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
)

func BuildAmqpUri(c *config.AmqpConfig) *string {
	return tea.String(fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Username, c.Password, c.Host, c.Port))
}
