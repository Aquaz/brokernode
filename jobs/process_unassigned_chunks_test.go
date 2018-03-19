package jobs_test

import (
	"github.com/oysterprotocol/brokernode/models"
	"github.com/oysterprotocol/brokernode/jobs"
)

func (suite *JobsSuite) Test_ProcessUnassignedChunks() {
	genHash := "genHashTest"
	fileBytesCount := 8500

	models.SetChunkStatuses()

	vErr, err := models.BuildDataMaps(genHash, fileBytesCount)
	suite.Nil(err)
	suite.Equal(0, len(vErr.Errors))

	unassignedTimedOut := models.DataMap{}
	unassignedNotTimedOut := models.DataMap{}
	notUnassignedTimedOut := models.DataMap{}
	notUnassignedNotTimedOut := models.DataMap{}

	err = suite.DB.Find(&unassignedTimedOut, 1)
	unassignedTimedOut.Status = models.ChunkStatus["unassigned"]
	unassignedTimedOut.Message = "unassignedTimedOut"
	suite.DB.ValidateAndSave(&unassignedTimedOut)

	err = suite.DB.Find(&notUnassignedTimedOut, 3)
	notUnassignedTimedOut.Status = models.ChunkStatus["pending"]
	notUnassignedTimedOut.Message = "notUnassignedTimedOut"
	suite.DB.ValidateAndSave(&notUnassignedTimedOut)

	//currentTime := time.Now().Format("2006-01-02T15:04:05Z")

	err = suite.DB.Find(&unassignedNotTimedOut, 2)
	unassignedNotTimedOut.Status = models.ChunkStatus["unassigned"]
	unassignedNotTimedOut.Message = "unassignedNotTimedOut"
	suite.DB.ValidateAndSave(&unassignedNotTimedOut)

	err = suite.DB.Find(&notUnassignedNotTimedOut, 4)
	notUnassignedNotTimedOut.Status = models.ChunkStatus["pending"]
	notUnassignedNotTimedOut.Message = "notUnassignedNotTimedOut"
	suite.DB.ValidateAndSave(&notUnassignedNotTimedOut)

	result, err := jobs.GetUnassignedChunks()

	suite.Equal(2, len(result))

	for _, dMap := range result {
		suite.Equal(models.ChunkStatus["unassigned"], dMap.Status)
	}
}
