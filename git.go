// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package charmrepo // import "gopkg.in/juju/charmrepo.v2"

import (
	"io/ioutil"
	"strings"

	"gopkg.in/errgo.v1"
	"gopkg.in/juju/charm.v6"
	"gopkg.in/src-d/go-git.v4"
)

// GitRepo is a repository Interface that describes
// a git remote used for checking out charm code
type GitRepo struct {
	remoteURI string
	reference string
}

var _ Interface = (*GitRepo)(nil)

// NewGitRepo holds parameters for instantiating a new GitRepo.
func NewGitRepo(ref *charm.URL) (Interface, error) {

	// Given the git revision won't always
	// match the charm revision, we need to
	// split the desired revision out from the
	// charm name
	tokens := strings.Split(ref.Name, "?")
	reference := "HEAD"

	if len(tokens) > 1 {
		reference = tokens[1]
	}

	return &GitRepo{
		remoteURI: tokens[0],
		reference: reference,
	}, nil
}

// Get implements Interface.Get.
func (g *GitRepo) Get(checkout *charm.URL) (charm.Charm, error) {

	if checkout.Series == "bundle" {
		return nil, errgo.Newf("expected a charm URL, got bundle URL %q", checkout)
	}
	path, err := g.archivePath(checkout)
	if err != nil {
		return nil, errgo.Mask(err, errgo.Any)
	}
	return charm.ReadCharmDir(path)
}

// GetBundle implements Interface.GetBundle.
func (g *GitRepo) GetBundle(checkout *charm.URL) (charm.Bundle, error) {
	if checkout.Series != "bundle" {
		return nil, errgo.Newf("expected a bundle URL, got charm URL %q", checkout)
	}
	path, err := g.archivePath(checkout)
	if err != nil {
		return nil, errgo.Mask(err, errgo.Any)
	}
	return charm.ReadBundleArchive(path)
}

// archivePath returns a local path to the checked out charm or bundle
func (g *GitRepo) archivePath(checkout *charm.URL) (string, error) {

	//get temporary directory for checkout
	tempDir, err := ioutil.TempDir("", "juju-clone")

	if err != nil {
		return "", errgo.Mask(err, errgo.Any)
	}

	//clone to a local directory
	_, err = git.PlainClone(tempDir, false, &git.CloneOptions{
		URL:               g.remoteURI,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	if err != nil {
		return "", errgo.Mask(err, errgo.Any)
	}
	// point at local repository, and resolve passed
	// reference, which will either be the default of HEAD
	// or whatever the user passed after the '?' in the
	// charm/bundle URI

	if err != nil {
		return "", errgo.Mask(err, errgo.Any)
	}

	return tempDir, nil
}

// Resolve implements Interface.Resolve.
func (g *GitRepo) Resolve(checkout *charm.URL) (*charm.URL, []string, error) {
	if checkout.Series == "" {
		return nil, nil, errgo.Newf("no series specified for %s", checkout)
	}
	if checkout.Revision != -1 {
		return checkout, nil, nil
	}
	if checkout.Series == "bundle" {
		// Bundles do not have revision files and the revision is not included
		// in metadata. For this reason, local bundles always have revision 0.
		return checkout.WithRevision(0), nil, nil
	}
	ch, err := g.Get(checkout)
	if err != nil {
		return nil, nil, err
	}
	// This is strictly speaking unnecessary, but just in case a bad charm is
	// used locally, we'll check the series.
	_, err = charm.SeriesForCharm(checkout.Series, ch.Meta().Series)
	if err != nil {
		return nil, nil, err
	}
	// We return the supported series read from the metadata
	return checkout.WithRevision(ch.Revision()), nil, nil
}
