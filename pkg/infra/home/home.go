package home

import (
	"fmt"

	gh "github.com/neverett8fr/go-home"
)

type HomeProvider struct {
	Home gh.Home
}

func InitialiseHomeProvider() (HomeProvider, error) {

	home, err := gh.NewHome()
	if err != nil {
		return HomeProvider{}, fmt.Errorf("error creating home provider, err %v", err)
	}

	return HomeProvider{
		Home: home,
	}, nil
}

func (hp *HomeProvider) AddRoute(name string, route string, method string) error {

	return hp.Home.RegisterHTTPEndpoint(name, route, method)
}

func (hp *HomeProvider) AddCondition(name string, condition string) error {

	return hp.Home.AddCondition(name, GetCondition(condition))
}

func (hp *HomeProvider) StartHome() error {

	go hp.Home.StartHandlers()

	return nil
}

func (hp *HomeProvider) StopHome() error {

	hp.Home.StopHandlers()

	return nil
}
