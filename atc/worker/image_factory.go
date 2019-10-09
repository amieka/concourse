package worker

import (
	"context"
	"io"
	"io/ioutil"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
)

//go:generate counterfeiter . ImageFactory

type ImageFactory interface {
	GetImage(
		logger lager.Logger,
		worker Worker,
		volumeClient VolumeClient,
		imageSpec ImageSpec,
		teamID int,
		delegate ImageFetchingDelegate,
		resourceTypes atc.VersionedResourceTypes,
	) (Image, error)
}

type FetchedImage struct {
	Version    atc.Version
	URL        string
	Privileged bool
}

//go:generate counterfeiter . Image

type Image interface {
	FetchForContainer(
		ctx context.Context,
		logger lager.Logger,
		container db.CreatingContainer,
	) (FetchedImage, error)
}

//go:generate counterfeiter . ImageFetchingDelegate

type ImageFetchingDelegate interface {
	Stdout() io.Writer
	Stderr() io.Writer
	ImageVersionDetermined(db.UsedResourceCache) error
}

type NoopImageFetchingDelegate struct{}

func (NoopImageFetchingDelegate) Stdout() io.Writer                                 { return ioutil.Discard }
func (NoopImageFetchingDelegate) Stderr() io.Writer                                 { return ioutil.Discard }
func (NoopImageFetchingDelegate) ImageVersionDetermined(db.UsedResourceCache) error { return nil }
