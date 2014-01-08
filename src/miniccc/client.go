// Copyright (2014) Sandia Corporation.
// Under the terms of Contract DE-AC04-94AL85000 with Sandia Corporation,
// the U.S. Government retains certain rights in this software.

package main

func (r *ron) newClient() error {
	// start the periodic query to the parent
	go r.heartbeat()

	return nil
}