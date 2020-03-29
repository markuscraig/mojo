package mojo

const (
	defaultSize          = float64(10.0)
	defaultDetail        = float64(10.0)
	defaultPoints        = Corner
	defaultWigglesPerSec = float64(2.0)
	defaultCorrelation   = float64(0.5)
	defaultRandomSeed    = uint32(0)
)

type PointsType int

const (
	Corner PointsType = iota
	Smooth
)

type WigglePath struct {
	Node
	Size          float64
	Detail        float64
	Points        PointsType
	WigglesPerSec float64
	Correlation   float64
	TemporalPhase Rotation
	SpatialPhase  Rotation
	RandomSeed    uint32
	Transform     *Transform
}

func NewWigglePath() *WigglePath {
	return &WigglePath{
		Size:          defaultSize,
		Detail:        defaultDetail,
		Points:        defaultPoints,
		WigglesPerSec: defaultWigglesPerSec,
		Correlation:   defaultCorrelation,
		TemporalPhase: defaultRotation,
		SpatialPhase:  defaultRotation,
		RandomSeed:    defaultRandomSeed,
		Transform:     defaultTransform,
	}
}
