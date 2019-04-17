package docker

import "time"

// Changed by best practices https://github.com/golang/go/wiki/CodeReviewComments#mixed-caps
const dockerFilterOptionName = "name"
const dockerRestartTimeout = 30
const ***REMOVED*** = "***REMOVED***"
const ***REMOVED*** = "***REMOVED***"
const ***REMOVED*** = "***REMOVED***"
const dockerDefaultAPIPort = ":2376"

// TLS Files
const certPath = "assets/certs/docker/cert.pem"
const keyPath = "assets/certs/docker/key.pem"
const CAPath = "assets/certs/docker/ca.pem"

// Default container properts
const neogridMountPath = "/data/neogrid:/data/neogrid"

// Errors handlers
const cannotCreateContainer = "Cannot create container."
const cannotStartContainer = "Cannot start container."
const cannotDeleteContainer = "Cannot delete container"

const containerNotFound = "Container not found."
const containerCreationTimeout = time.Second * 10
const containerCreationTimeExceeded = "Container creation timeout"

const missingImageName = "Image name is missing"
const missingImageTag = "Image tag is missing"
