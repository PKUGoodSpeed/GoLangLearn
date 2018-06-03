package methods

import "fmt"

type NNmodel interface{
	train(epoch int, learning_rate float64) float64
	test(epoch int, learning_rate float64) float64
}

func GetModelSummary(nn NNmodel, epoch int, lr float64){
	fmt.Println("Training Accu: ", nn.train(epoch, lr))
	fmt.Println("Testing Accu: ", nn.test(epoch, lr))
}

type Resnet struct{
	Cnn int
	Fc int
	Dropout float64
}

func (res Resnet) train(ep int, lr float64) float64{
	return 1. - 1./(float64(res.Cnn) + float64(res.Fc)*res.Dropout + float64(ep))/lr
}

func (res Resnet) test(ep int, lr float64) float64{
	return 1. - 1./(float64(res.Cnn) + float64(res.Fc)*res.Dropout)
}

type Xception struct{
	Shaocong int
	Zebo float64
}

func (xcp Xception) train(ep int, lr float64) float64{
	return 1.
}

func (xcp Xception) test(ep int, lr float64) float64{
	return 1.
}