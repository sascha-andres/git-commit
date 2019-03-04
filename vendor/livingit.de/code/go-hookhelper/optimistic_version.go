package hookhelper

import "errors"

// OptimisticVersion returns the calculated optimistic version
func OptimisticVersion(global, local []byte, globalVersion, localVersion string) (string, error) {
	if nil != global {
		if nil != local && localVersion != "" {
			if localVersion == globalVersion && globalVersion == "" {
				return "", errors.New("you have to provide versions for global and local config")
			}
			if localVersion != globalVersion {
				return "", errors.New("version mismatch for global and project version")
			}
		}
		return globalVersion, nil
	}

	if nil != local {
		return localVersion, nil
	}

	return "", errors.New("no suitable versioned configuration found")
}
