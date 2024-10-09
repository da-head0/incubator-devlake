package migrationscripts

import (
	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/helpers/migrationhelper"
)

type JiraBoard20241009 struct {
	FilterId uint64 `gorm:"column:filter_id"`
}

func (JiraBoard20241009) TableName() string {
	return "_tool_jira_boards"
}

type addFilterIdToBoards struct{}

func (script *addFilterIdToBoards) Up(basicRes context.BasicRes) errors.Error {
	return migrationhelper.AutoMigrateTables(basicRes, &JiraBoard20241009{})
}

func (*addFilterIdToBoards) Version() uint64 {
	return 20241009000001
}

func (*addFilterIdToBoards) Name() string {
	return "add filter_id to _tool_jira_boards"
}
