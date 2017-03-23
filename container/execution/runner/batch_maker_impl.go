package runner

import (
	"github.com/streamsets/dataextractor/api"
)

type BatchMakerImpl struct {
	stagePipe   StagePipe
	stageOutput []api.Record
}

func (b *BatchMakerImpl) GetLanes() []string {
	return b.stagePipe.OutputLanes
}

func (b *BatchMakerImpl) AddRecord(record api.Record) {
	b.stageOutput = append(b.stageOutput, record)
}

func (b *BatchMakerImpl) getStageOutput() []api.Record {
	return b.stageOutput
}

func NewBatchMakerImpl(stagePipe StagePipe) *BatchMakerImpl {
	return &BatchMakerImpl{
		stagePipe:   stagePipe,
		stageOutput: make([]api.Record, 0),
	}
}
