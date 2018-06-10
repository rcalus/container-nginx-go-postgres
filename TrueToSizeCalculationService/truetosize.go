package main

//TrueToSize data, created struct to keep signature if more fields were to be passed in
type TrueToSize struct {
	Size int `json:"size,string" validate:"min=1,max=5"`
}

//TrueToSizeResult a calculation of the average of the values in the system
type TrueToSizeResult struct {
	AverageTrueToSize float32 `json:"averageTrueToSize" validate:"required"`
}
