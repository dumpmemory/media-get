package meta

type IProcessor interface {
	FetchMetaAndResourceInfo() (mediaMeta *MediaMeta, err error)
}