package models

type GasGiant struct {
	bodyType BodyType
}

func (gg *GasGiant) GetType() BodyType {
	return BodyType
}

func (gg *GasGiant) SetType(bt BodyType) Error {

	if bt == LargeGasGiant || bt == BodyType.SmallGasGiant {
		gg.bodyType = bt
		return nil
	}
	return Error("Incorrect BodyType")
}
