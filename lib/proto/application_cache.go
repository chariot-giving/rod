// This file is generated by "./lib/proto/generate"

package proto

/*

ApplicationCache

*/

// ApplicationCacheApplicationCacheResource Detailed application cache resource information.
type ApplicationCacheApplicationCacheResource struct {

	// URL Resource url.
	URL string `json:"url"`

	// Size Resource size.
	Size int `json:"size"`

	// Type Resource type.
	Type string `json:"type"`
}

// ApplicationCacheApplicationCache Detailed application cache information.
type ApplicationCacheApplicationCache struct {

	// ManifestURL Manifest URL.
	ManifestURL string `json:"manifestURL"`

	// Size Application cache size.
	Size float64 `json:"size"`

	// CreationTime Application cache creation time.
	CreationTime float64 `json:"creationTime"`

	// UpdateTime Application cache update time.
	UpdateTime float64 `json:"updateTime"`

	// Resources Application cache resources.
	Resources []*ApplicationCacheApplicationCacheResource `json:"resources"`
}

// ApplicationCacheFrameWithManifest Frame identifier - manifest URL pair.
type ApplicationCacheFrameWithManifest struct {

	// FrameID Frame identifier.
	FrameID PageFrameID `json:"frameId"`

	// ManifestURL Manifest URL.
	ManifestURL string `json:"manifestURL"`

	// Status Application cache status.
	Status int `json:"status"`
}

// ApplicationCacheEnable Enables application cache domain notifications.
type ApplicationCacheEnable struct {
}

// ProtoReq of the command
func (m ApplicationCacheEnable) ProtoReq() string { return "ApplicationCache.enable" }

// Call of the command, sessionID is optional.
func (m ApplicationCacheEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// ApplicationCacheGetApplicationCacheForFrame Returns relevant application cache data for the document in given frame.
type ApplicationCacheGetApplicationCacheForFrame struct {

	// FrameID Identifier of the frame containing document whose application cache is retrieved.
	FrameID PageFrameID `json:"frameId"`
}

// ProtoReq of the command
func (m ApplicationCacheGetApplicationCacheForFrame) ProtoReq() string {
	return "ApplicationCache.getApplicationCacheForFrame"
}

// Call of the command, sessionID is optional.
func (m ApplicationCacheGetApplicationCacheForFrame) Call(c Client) (*ApplicationCacheGetApplicationCacheForFrameResult, error) {
	var res ApplicationCacheGetApplicationCacheForFrameResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// ApplicationCacheGetApplicationCacheForFrameResult Returns relevant application cache data for the document in given frame.
type ApplicationCacheGetApplicationCacheForFrameResult struct {

	// ApplicationCache Relevant application cache data for the document in given frame.
	ApplicationCache *ApplicationCacheApplicationCache `json:"applicationCache"`
}

// ApplicationCacheGetFramesWithManifests Returns array of frame identifiers with manifest urls for each frame containing a document
// associated with some application cache.
type ApplicationCacheGetFramesWithManifests struct {
}

// ProtoReq of the command
func (m ApplicationCacheGetFramesWithManifests) ProtoReq() string {
	return "ApplicationCache.getFramesWithManifests"
}

// Call of the command, sessionID is optional.
func (m ApplicationCacheGetFramesWithManifests) Call(c Client) (*ApplicationCacheGetFramesWithManifestsResult, error) {
	var res ApplicationCacheGetFramesWithManifestsResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// ApplicationCacheGetFramesWithManifestsResult Returns array of frame identifiers with manifest urls for each frame containing a document
// associated with some application cache.
type ApplicationCacheGetFramesWithManifestsResult struct {

	// FrameIds Array of frame identifiers with manifest urls for each frame containing a document
	// associated with some application cache.
	FrameIds []*ApplicationCacheFrameWithManifest `json:"frameIds"`
}

// ApplicationCacheGetManifestForFrame Returns manifest URL for document in the given frame.
type ApplicationCacheGetManifestForFrame struct {

	// FrameID Identifier of the frame containing document whose manifest is retrieved.
	FrameID PageFrameID `json:"frameId"`
}

// ProtoReq of the command
func (m ApplicationCacheGetManifestForFrame) ProtoReq() string {
	return "ApplicationCache.getManifestForFrame"
}

// Call of the command, sessionID is optional.
func (m ApplicationCacheGetManifestForFrame) Call(c Client) (*ApplicationCacheGetManifestForFrameResult, error) {
	var res ApplicationCacheGetManifestForFrameResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// ApplicationCacheGetManifestForFrameResult Returns manifest URL for document in the given frame.
type ApplicationCacheGetManifestForFrameResult struct {

	// ManifestURL Manifest URL for document in the given frame.
	ManifestURL string `json:"manifestURL"`
}

// ApplicationCacheApplicationCacheStatusUpdated ...
type ApplicationCacheApplicationCacheStatusUpdated struct {

	// FrameID Identifier of the frame containing document whose application cache updated status.
	FrameID PageFrameID `json:"frameId"`

	// ManifestURL Manifest URL.
	ManifestURL string `json:"manifestURL"`

	// Status Updated application cache status.
	Status int `json:"status"`
}

// ProtoEvent interface
func (evt ApplicationCacheApplicationCacheStatusUpdated) ProtoEvent() string {
	return "ApplicationCache.applicationCacheStatusUpdated"
}

// ApplicationCacheNetworkStateUpdated ...
type ApplicationCacheNetworkStateUpdated struct {

	// IsNowOnline ...
	IsNowOnline bool `json:"isNowOnline"`
}

// ProtoEvent interface
func (evt ApplicationCacheNetworkStateUpdated) ProtoEvent() string {
	return "ApplicationCache.networkStateUpdated"
}